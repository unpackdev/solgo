package observers

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
)

// BlockSubscriberOptions defines the options for configuring a block subscriber.
type BlockSubscriberOptions struct {
	NetworkID        int64    `mapstructure:"network_id" yaml:"network_id" json:"network_id"`                         // The network ID of the blockchain.
	Group            string   `mapstructure:"group" yaml:"group" json:"group"`                                        // The group of the blockchain client.
	Type             string   `mapstructure:"type" yaml:"type" json:"type"`                                           // The type of the blockchain client.
	Head             bool     `mapstructure:"head" yaml:"head" json:"head"`                                           // Flag to indicate if subscribing to the latest block.
	StartBlockNumber *big.Int `mapstructure:"start_block_number" yaml:"start_block_number" json:"start_block_number"` // The starting block number for the subscription.
	EndBlockNumber   *big.Int `mapstructure:"end_block_number" yaml:"end_block_number" json:"end_block_number"`       // The ending block number for the subscription.
}

// BlockSubscriber provides methods to subscribe to blockchain block headers.
type BlockSubscriber struct {
	ctx    context.Context       // The context for the subscriber.
	client *clients.ClientPool   // The client pool for accessing blockchain clients.
	active atomic.Bool           // Flag to indicate if the subscriber is active.
	sub    ethereum.Subscription // The Ethereum subscription object.
}

// NewBlockSubscriber creates a new block subscriber with the given context and client pool.
func NewBlockSubscriber(ctx context.Context, client *clients.ClientPool) (*BlockSubscriber, error) {
	return &BlockSubscriber{
		ctx:    ctx,
		client: client,
	}, nil
}

// IsActive checks if the block subscriber is currently active.
func (b *BlockSubscriber) IsActive() bool {
	return b.active.Load()
}

// SubscribeHeader subscribes to block headers based on the provided options.
// It can either subscribe to the latest block headers or a range of block headers.
func (b *BlockSubscriber) SubscribeHeader(opts *BlockSubscriberOptions, blockCh chan *types.Header) error {
	if b.active.Load() {
		return errors.New("block subscriber is already active")
	}

	client := b.client.GetClientByGroupAndType(opts.Group, opts.Type)
	if client == nil {
		return errors.New("client not found")
	}

	if opts.Head {
		// Subscribe to block number
		sub, err := client.SubscribeNewHead(b.ctx, blockCh)
		if err != nil {
			return err
		}
		b.sub = sub

		// Set subscriber as active
		b.active.Store(true)

		for {
			select {
			case err := <-sub.Err():
				return err
			case <-b.ctx.Done():
				return nil
			}
		}
	} else {
		if opts.StartBlockNumber == nil && opts.EndBlockNumber == nil {
			return errors.New("start and end block numbers are not set")
		}

		if opts.EndBlockNumber.Int64() < opts.StartBlockNumber.Int64() {
			return errors.New("end block number is less than start block number")
		}

		b.active.Store(true)

		for i := opts.StartBlockNumber.Int64(); i <= opts.EndBlockNumber.Int64(); i++ {
			block, err := client.BlockByNumber(b.ctx, big.NewInt(i))
			if err != nil {
				return err
			}

			blockCh <- block.Header()
		}

		b.active.Store(false)
	}

	return nil
}

// SubscribeHeader subscribes to block based on the provided options.
// It can either subscribe to the latest blocks or a range of blocks.
func (b *BlockSubscriber) Subscribe(opts *BlockSubscriberOptions, blockCh chan *types.Block) error {
	if b.active.Load() {
		return errors.New("block subscriber is already active")
	}

	client := b.client.GetClientByGroupAndType(opts.Group, opts.Type)
	if client == nil {
		return errors.New("client not found")
	}

	if opts.Head {
		// Subscribe to block number
		headerCh := make(chan *types.Header)
		sub, err := client.SubscribeNewHead(b.ctx, headerCh)
		if err != nil {
			return err
		}
		b.sub = sub

		// Set subscriber as active
		b.active.Store(true)

		for {
			select {
			case header := <-headerCh:
				block, err := client.BlockByHash(b.ctx, header.Hash())
				if err != nil {
					return err
				}
				blockCh <- block
			case err := <-sub.Err():
				return err
			case <-b.ctx.Done():
				return nil
			}
		}
	} else {
		if opts.StartBlockNumber == nil && opts.EndBlockNumber == nil {
			return errors.New("start and end block numbers are not set")
		}

		if opts.EndBlockNumber.Int64() < opts.StartBlockNumber.Int64() {
			return errors.New("end block number is less than start block number")
		}

		b.active.Store(true)

		for i := opts.StartBlockNumber.Int64(); i <= opts.EndBlockNumber.Int64(); i++ {
			block, err := client.BlockByNumber(b.ctx, big.NewInt(i))
			if err != nil {
				return err
			}

			blockCh <- block
		}

		b.active.Store(false)
	}

	return nil
}

// Close terminates the block subscription and releases any associated resources.
func (b *BlockSubscriber) Close() error {
	if b.active.Load() {
		if b.sub != nil {
			b.sub.Unsubscribe()
		}
		b.active.Store(false)
	}

	return nil
}
