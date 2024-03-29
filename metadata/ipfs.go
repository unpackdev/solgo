package metadata

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"strings"
	"time"

	cid "github.com/ipfs/go-cid"
)

// IpfsProvider is a struct that holds the context and the client for IPFS operations.
type IpfsProvider struct {
	ctx    context.Context // The context to be used in IPFS operations.
	client Shell           // The IPFS client.
}

// NewIpfsProvider creates a new instance of IpfsProvider.
// It takes a context and an IPFS client as parameters.
// If the client is nil, it returns an error. Otherwise, it returns a new instance of IpfsProvider.
func NewIpfsProvider(ctx context.Context, client Shell) (Provider, error) {
	if client == nil {
		return nil, ErrInvalidIpfsClient
	}

	return Provider(&IpfsProvider{
		ctx:    ctx,
		client: client,
	}), nil
}

// extractHash is a helper method that extracts the CID hash from a string.
// It checks if the string starts with 'ipfs://', removes the prefix, and validates the remaining hash.
// If the string does not start with 'ipfs://' or the hash is invalid, it returns an error.
func (p *IpfsProvider) extractHash(s string) (string, error) {
	// Check if the string starts with 'ipfs://'
	if !strings.HasPrefix(s, "ipfs://") {
		return "", fmt.Errorf("invalid format: string does not start with 'ipfs://'")
	}

	// Remove 'ipfs://' from the string
	hash := strings.TrimPrefix(s, "ipfs://")

	// Validate the IPFS hash
	_, err := cid.Decode(hash)
	if err != nil {
		return "", fmt.Errorf("invalid IPFS hash: %w", err)
	}

	return hash, nil
}

// GetMetadataByCID retrieves the metadata of a contract by its CID (Content Identifier).
// It first extracts the hash from the CID, then uses the IPFS client to retrieve the content associated with the hash.
// It reads the content, unmarshals it into a ContractMetadata object, and returns it.
// If any of these operations fail, it returns an error.
func (p *IpfsProvider) GetMetadataByCID(cid string) (*ContractMetadata, error) {
	hash, err := p.extractHash(cid)
	if err != nil {
		return nil, err
	}

	done := time.After(10 * time.Second)
	result := make(chan *ContractMetadata)
	errs := make(chan error)

	go func() {
		content, err := p.client.Cat(fmt.Sprintf("/ipfs/%s", hash))
		if err != nil {
			errs <- err
			return
		}

		data, err := io.ReadAll(content)
		if err != nil {
			errs <- err
			return
		}

		var toReturn ContractMetadata
		if err := json.Unmarshal(data, &toReturn); err != nil {
			errs <- err
			return
		}

		// Inject as well raw data in case that we need to do some additional operations on top of it
		// eg. save it to the database...
		toReturn.Raw = string(data)

		// There are contracts that do contain sources but are not available immediately at first metadata lookup.
		// These sources usually contain bzz and ipfs locations. Sometimes reachable, sometimes not, never the less,
		// this is a way how to fetch those contracts.
		for sourceName, source := range toReturn.Sources {
			if len(source.Content) < 10 {
				for _, url := range source.Urls {
					if strings.HasPrefix(url, "dweb:/ipfs/") {
						url := strings.TrimPrefix(url, "dweb:/ipfs/")
						subContent, err := p.client.Cat(fmt.Sprintf("/ipfs/%s", url))
						if err != nil {
							errs <- err
							return
						}

						data, err := io.ReadAll(subContent)
						if err != nil {
							errs <- err
							return
						}
						source.Content = string(data)
						toReturn.Sources[sourceName] = source
					}
				}
			}

		}

		result <- &toReturn
	}()

	select {
	case <-p.ctx.Done():
		return nil, fmt.Errorf("context cancelled while fetching content from IPFS")
	case <-done:
		return nil, fmt.Errorf("timeout while fetching content from IPFS")
	case res := <-result:
		return res, nil
	case err := <-errs:
		return nil, err
	}
}
