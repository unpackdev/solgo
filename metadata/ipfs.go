package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	cid "github.com/ipfs/go-cid"
	ipfs "github.com/ipfs/go-ipfs-api"
)

// IpfsProvider is a struct that holds the context and the client for IPFS operations.
type IpfsProvider struct {
	ctx    context.Context // The context to be used in IPFS operations.
	client *ipfs.Shell     // The IPFS client.
}

// NewIpfsProvider creates a new instance of IpfsProvider.
// It takes a context and an IPFS client as parameters.
// If the client is nil, it returns an error. Otherwise, it returns a new instance of IpfsProvider.
func NewIpfsProvider(ctx context.Context, client *ipfs.Shell) (Provider, error) {
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

	content, err := p.client.Cat(fmt.Sprintf("/ipfs/%s", hash))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}

	var toReturn ContractMetadata
	if err := json.Unmarshal(data, &toReturn); err != nil {
		return nil, err
	}

	// Inject as well raw data in case that we need to do some additional operations on top of it
	// eg. save it to the database...
	toReturn.Raw = string(data)

	return &toReturn, nil
}
