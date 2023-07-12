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

type IpfsProvider struct {
	ctx    context.Context
	client *ipfs.Shell
}

// NewIpfsProvider returns a new instance of the IPFS provider.
// It requires a valid IPFS client to be passed as a parameter.
// If the client is nil, it will return an error otherwise it will return a new instance of the IPFS provider.
func NewIpfsProvider(ctx context.Context, client *ipfs.Shell) (Provider, error) {
	if client == nil {
		return nil, ErrInvalidIpfsClient
	}

	return Provider(&IpfsProvider{
		ctx:    ctx,
		client: client,
	}), nil
}

// extractHash extracts the cid hash from the string
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

// GetMetadataByCID returns the metadata of a contract by the CID (Content Identifier) of the contract
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
