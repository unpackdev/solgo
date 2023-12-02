package clients

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// ClientPool manages a pool of Ethereum clients for different networks and types.
// It provides methods to retrieve clients based on various criteria and to close all clients in the pool.
type ClientPool struct {
	ctx           context.Context
	opts          *Options
	clients       map[string][]*Client
	next          uint32
	groupTypeNext map[string]uint32
}

// Len returns the number of clients in the pool.
func (c *ClientPool) Len() int {
	return len(c.clients)
}

// GetClient retrieves a client based on the group and type in a round-robin fashion.
func (c *ClientPool) GetClient(group, typ string) *Client {
	key := group + "_" + typ
	current := c.groupTypeNext[key]
	n := atomic.AddUint32(&current, 1)
	return c.clients[key][(int(n)-1)%len(c.clients[key])]
}

// GetClientByGroupAndType retrieves a client based on the group and type in a round-robin fashion.
// This method is functionally equivalent to GetClient and is provided for clarity.
func (c *ClientPool) GetClientByGroupAndType(group, typ string) *Client {
	return c.GetClient(group, typ)
}

// GetClientByGroup retrieves a client based on the group in a round-robin fashion.
// It aggregates all clients within the specified group and returns one of them.
func (c *ClientPool) GetClientByGroup(group string) *Client {
	var allClientsInGroup []*Client

	// Aggregate all clients within the specified group
	for key, clients := range c.clients {
		if strings.HasPrefix(key, group+"_") {
			allClientsInGroup = append(allClientsInGroup, clients...)
		}
	}

	// If no clients are found for the group, return nil
	if len(allClientsInGroup) == 0 {
		return nil
	}

	// Use atomic operation to get the next client index for the group in a round-robin fashion
	current := c.groupTypeNext[group]
	n := atomic.AddUint32(&current, 1)

	// Calculate the client index based on the current value of n
	clientIndex := (int(n) - 1) % len(allClientsInGroup)

	// Return the client at the calculated index
	return allClientsInGroup[clientIndex]
}

// GetClientDescriptionByNetworkId retrieves the group and type of a client based on the network ID.
// It returns an empty string for both group and type if no match is found.
func (c *ClientPool) GetClientDescriptionByNetworkId(networkId *big.Int) (string, string) {
	for _, node := range c.opts.GetNodes() {
		if big.NewInt(node.GetNetworkID()).Int64() == networkId.Int64() {
			return node.GetGroup(), node.GetType()
		}
	}
	return "", ""
}

// Close gracefully closes all the clients in the pool.
func (c *ClientPool) Close() {
	for _, clients := range c.clients {
		for _, client := range clients {
			client.Close()
		}
	}
}

// RegisterClient adds a new client to the pool. It takes the group and type of the client
// as well as the necessary parameters to create the client.
func (c *ClientPool) RegisterClient(ctx context.Context, networkId uint64, group, typ, endpoint string, concurrentClientsNumber int) error {
	if group == "" || typ == "" || endpoint == "" {
		return errors.New("invalid parameters: group, type, and endpoint are required")
	}
	if concurrentClientsNumber <= 0 {
		return errors.New("concurrentClientsNumber must be greater than 0")
	}

	mutex := sync.Mutex{}
	g, ctx := errgroup.WithContext(ctx)
	key := group + "_" + typ

	for i := 0; i < concurrentClientsNumber; i++ {
		g.Go(func() error {
			node := Node{Endpoint: endpoint, Group: group, Type: typ, NetworkId: int64(networkId), ConcurrentClientsNumber: concurrentClientsNumber}
			client, err := NewClient(ctx, &node)
			if err != nil {
				return err
			}

			// Additional checks or configurations for the client can be added here

			mutex.Lock()
			c.clients[key] = append(c.clients[key], client)
			mutex.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// NewClientPool initializes a new ClientPool with the given options.
// It returns an error if the options are not set, if there are no nodes specified in the options,
// or if there's an issue with any of the nodes' configurations.
func NewClientPool(ctx context.Context, opts *Options) (*ClientPool, error) {
	clients := make(map[string][]*Client)
	mutex := sync.Mutex{}
	g, ctx := errgroup.WithContext(ctx)

	if opts == nil {
		return nil, ErrOptionsNotSet
	}

	/* 	if len(opts.GetNodes()) == 0 {
		return nil, ErrNodesNotSet
	} */

	for _, node := range opts.GetNodes() {
		if node.GetEndpoint() == "" {
			return nil, ErrClientURLNotSet
		}

		if node.GetConcurrentClientsNumber() == 0 {
			return nil, ErrConcurrentClientsNotSet
		}

		for i := 0; i < node.GetConcurrentClientsNumber(); i++ {
			node := node // Shadow variable for goroutine
			g.Go(func() error {
				client, err := NewClient(ctx, &node)
				if err != nil {
					return err
				}

				networkId := client.GetNetworkID()
				if node.GetNetworkID() != networkId {
					client.Close()
					return errors.New("network ID mismatch")
				}

				key := node.GetGroup() + "_" + node.GetType()
				mutex.Lock()
				clients[key] = append(clients[key], client)
				mutex.Unlock()
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		zap.L().Error("failed to initialize ethereum client", zap.Error(err))
		return nil, err
	}

	return &ClientPool{
		ctx:           ctx,
		opts:          opts,
		clients:       clients,
		next:          0,
		groupTypeNext: make(map[string]uint32),
	}, nil
}
