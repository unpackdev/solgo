package clients

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client wraps the Ethereum client with additional context and options.
// It provides methods to retrieve client-specific configurations and to close the client connection.
type Client struct {
	ctx  context.Context
	opts *Node
	*ethclient.Client
}

// NewClient initializes a new Ethereum client with the given options.
// It returns an error if the endpoint URL is not set, if there's an issue initializing the Ethereum client,
// or if there's a mismatch between the provided network ID and the actual network ID.
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
				"failed to initialize Ethereum client due to network IDs mismatch %d->%d",
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
	return int64(c.opts.NetworkId)
}

// GetGroup retrieves the group associated with the client.
func (c *Client) GetGroup() string {
	return c.opts.Group
}

// GetType retrieves the type of the client.
func (c *Client) GetType() string {
	return c.opts.Type
}

// GetEndpoint retrieves the endpoint URL of the client.
func (c *Client) GetEndpoint() string {
	return c.opts.Endpoint
}

// GetFailoverGroup retrieves the failover group associated with the client.
func (c *Client) GetFailoverGroup() string {
	return c.opts.FailoverGroup
}

// GetFailoverType retrieves the type of failover for the client.
func (c *Client) GetFailoverType() string {
	return c.opts.FailoverType
}

// GetRpcClient retrieves the RPC client associated with the client.
func (c *Client) GetRpcClient() *rpc.Client {
	return c.Client.Client()
}

// Close gracefully closes the Ethereum client connection.
func (c *Client) Close() {
	c.Client.Close()
}
