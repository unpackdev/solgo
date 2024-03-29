package bitquery

import (
	"bytes"
	"context"
	"errors"
	"github.com/goccy/go-json"
	"net/http"
	"time"
)

// Provider represents a client for the blockchain data service,
// configured with options and capable of making requests.
type Provider struct {
	ctx    context.Context // The context for request cancellation and deadlines.
	opts   *Options        // Configuration options for the data service.
	client *http.Client    // The HTTP client used for making requests.
}

// NewProvider initializes and returns a new Provider instance.
// It validates the provided Options to ensure necessary configurations are set.
//
// Returns an error if the Options are nil, or if essential options like Endpoint or Key are not configured.
func NewProvider(ctx context.Context, opts *Options) (*Provider, error) {
	if opts == nil {
		return nil, errors.New("bitquery provider is not configured")
	}

	if opts.Endpoint == "" {
		return nil, errors.New("bitquery provider endpoint is not configured")
	}

	if opts.Key == "" {
		return nil, errors.New("bitquery provider key is not configured")
	}

	return &Provider{
		ctx:    ctx,
		opts:   opts,
		client: &http.Client{Timeout: time.Second * 30},
	}, nil
}

// GetContractCreationInfo sends a query to the blockchain data service to retrieve
// contract creation information such as the transaction hash and block height.
//
// This method accepts a context (for cancellation and deadlines) and a query map
// defining the parameters of the query. It returns a pointer to ContractCreationInfo
// containing the requested data, or an error if the request fails, the response status
// is not OK, or the response cannot be decoded.
func (b *Provider) GetContractCreationInfo(ctx context.Context, query map[string]string) (*ContractCreationInfo, error) {
	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, errors.New("failed to marshal query: " + err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", b.opts.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New("failed to create request: " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", b.opts.Key)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, errors.New("failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("received non-OK response: " + resp.Status)
	}

	var info ContractCreationInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}

	return &info, nil
}
