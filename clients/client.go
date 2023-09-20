package clients

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps the Ethereum client with additional context and options.
type Client struct {
	ctx  context.Context
	opts *Node
	*ethclient.Client
}

// NewClient initializes a new Ethereum client with the given options.
func NewClient(ctx context.Context, opts *Node) (*Client, error) {
	if opts.Endpoint == "" {
		return nil, errors.New("endpoint URL not set")
	}

	ethClient, err := ethclient.DialContext(ctx, opts.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client: %v", err)
	}

	if networkId, err := ethClient.NetworkID(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client: %v", err)
	} else {
		if networkId.Int64() != opts.GetNetworkID() {
			return nil, fmt.Errorf(
				"failed to initialize Ethereum client due to network ids missmatch %d->%d",
				opts.GetNetworkID(), networkId.Int64(),
			)
		}
	}

	return &Client{
		ctx:    ctx,
		opts:   opts,
		Client: ethClient,
	}, nil
}

// GetNetworkID retrieves the network ID for the client.
func (c *Client) GetNetworkID() int64 {
	return c.opts.NetworkId
}

// GetGroup retrieves the group for the client.
func (c *Client) GetGroup() string {
	return c.opts.Group
}

// GetType retrieves the type for the client.
func (c *Client) GetType() string {
	return c.opts.Type
}

// GetEndpoint retrieves the endpoint for the client.
func (c *Client) GetEndpoint() string {
	return c.opts.Endpoint
}

// Close closes the Ethereum client.
func (c *Client) Close() {
	c.Client.Close()
}
