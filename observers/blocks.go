package observers

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

// SubscriberOptions defines the options for configuring a block subscriber.
type SubscriberOptions struct {
	NetworkID        utils.NetworkID `mapstructure:"network_id" yaml:"network_id" json:"network_id"`                         // The network ID of the blockchain.
	Network          utils.Network   `mapstructure:"network" yaml:"network" json:"network"`                                  // The group of the blockchain client.                                      // The type of the blockchain client.
	Head             bool            `mapstructure:"head" yaml:"head" json:"head"`                                           // Flag to indicate if subscribing to the latest block.
	StartBlockNumber *big.Int        `mapstructure:"start_block_number" yaml:"start_block_number" json:"start_block_number"` // The starting block number for the subscription.
	EndBlockNumber   *big.Int        `mapstructure:"end_block_number" yaml:"end_block_number" json:"end_block_number"`       // The ending block number for the subscription.
}

// BlocksProcessor provides methods to subscribe to blockchain block headers.
type BlocksProcessor struct {
	*Manager
	active atomic.Bool           // Flag to indicate if the subscriber is active.
	sub    ethereum.Subscription // The Ethereum subscription object.
	hooks  map[HookType]BlockHookFn
}

// NewBlocksProcessor creates a new block subscriber with the given context and client pool.
func NewBlocksProcessor(manager *Manager) (*BlocksProcessor, error) {
	hooks := map[HookType]BlockHookFn{}
	for hook, fn := range manager.GetHooks(BlockProcessor) {
		hooks[hook] = fn.(BlockHookFn)
	}

	return &BlocksProcessor{
		Manager: manager,
		hooks:   hooks,
	}, nil
}

func (b *BlocksProcessor) ReloadHooks() bool {
	hooks := map[HookType]BlockHookFn{}
	for hook, fn := range b.GetHooks(BlockProcessor) {
		hooks[hook] = fn.(BlockHookFn)
	}

	b.hooks = hooks
	return true
}

// IsActive checks if the block subscriber is currently active.
func (b *BlocksProcessor) IsActive() bool {
	return b.active.Load()
}

// SubscribeHeader subscribes to block headers based on the provided options.
// It can either subscribe to the latest block headers or a range of block headers.
func (b *BlocksProcessor) SubscribeHeader(opts *SubscriberOptions, blockCh chan *BlockEntry) error {
	zap.L().Info(
		"Subscribing to block headers",
		zap.Any("network_id", opts.NetworkID),
		zap.String("network", opts.Network.String()),
	)

	client := b.clientsPool.GetClientByGroup(opts.Network.String())
	if client == nil {
		return errors.New("client not found")
	}

	headCh := make(chan *types.Header, 100)

	if opts.Head {
		sub, err := client.SubscribeNewHead(b.ctx, headCh)
		if err != nil {
			return err
		}
		b.sub = sub

		b.ReloadHooks()

		// Set subscriber as active
		b.active.Store(true)

		for {
			select {
			case err := <-sub.Err():
				return err
			case header := <-headCh:
				block, err := client.BlockByHash(b.ctx, header.Hash())
				if err != nil {
					zap.L().Error(
						"Failed to get block by hash",
						zap.Error(err),
						zap.Uint64("header_number", header.Number.Uint64()),
						zap.String("header_hash", header.Hash().String()),
					)
					continue
				}

				entry := &BlockEntry{
					NetworkID:   opts.NetworkID,
					Network:     opts.Network,
					Strategy:    utils.HeadStrategy,
					BlockHash:   block.Hash(),
					BlockNumber: block.Number(),
					Block:       block,
				}

				if hook, ok := b.hooks[PostHook]; ok {
					entry, err = hook(entry)
					if err != nil {
						zap.L().Error(
							"Post hook failed",
							zap.Error(err),
							zap.Uint64("block_number", block.Number().Uint64()),
							zap.String("block_hash", block.Hash().String()),
						)
						continue
					}
				}

				blockCh <- entry
			case <-b.ctx.Done():
				return nil
			}
		}
	}

	return nil
}

// Subscribe subscribes to block based on the provided options.
// It can either subscribe to the latest blocks or a range of blocks.
func (b *BlocksProcessor) Subscribe(opts *SubscriberOptions, blockCh chan *BlockEntry) error {
	zap.L().Info("Subscribing to blocks", zap.Any("options", opts))

	client := b.clientsPool.GetClientByGroup(opts.Network.String())
	if client == nil {
		return errors.New("client not found")
	}

	if opts.StartBlockNumber == nil && opts.EndBlockNumber == nil {
		return errors.New("start and end block numbers are not set")
	}

	if opts.EndBlockNumber.Int64() < opts.StartBlockNumber.Int64() {
		return errors.New("end block number is less than start block number")
	}

	b.ReloadHooks()

	b.active.Store(true)

	for i := opts.StartBlockNumber.Int64(); i <= opts.EndBlockNumber.Int64(); i++ {
		block, err := client.BlockByNumber(b.ctx, big.NewInt(i))
		if err != nil {
			zap.L().Error(
				"Failed to get block by number",
				zap.Error(err),
				zap.Int64("block_number", i),
			)
			continue
		}

		entry := &BlockEntry{
			NetworkID:   opts.NetworkID,
			Network:     opts.Network,
			Strategy:    utils.ArchiveStrategy,
			BlockHash:   block.Hash(),
			BlockNumber: block.Number(),
			Block:       block,
		}

		if hook, ok := b.hooks[PostHook]; ok {
			entry, err = hook(entry)
			if err != nil {
				zap.L().Error(
					"Post hook failed",
					zap.Error(err),
					zap.Uint64("block_number", block.Number().Uint64()),
					zap.String("block_hash", block.Hash().String()),
				)
				continue
			}
		}
		blockCh <- entry
	}

	b.active.Store(false)

	return nil
}

func (b *BlocksProcessor) ProcessBlock(ctx context.Context, network utils.Network, networkId utils.NetworkID, strategy utils.Strategy, block *types.Block) (*BlockEntry, error) {
	entry := &BlockEntry{
		NetworkID:   networkId,
		Network:     network,
		Strategy:    strategy,
		BlockHash:   block.Hash(),
		BlockNumber: block.Number(),
		Block:       block,
	}

	if hook, ok := b.hooks[PostHook]; ok {
		entry, err := hook(entry)
		if err != nil {
			zap.L().Error(
				"Post process block hook failed",
				zap.Error(err),
				zap.Uint64("block_number", block.Number().Uint64()),
				zap.String("block_hash", block.Hash().String()),
			)
			return nil, err
		}
		return entry, nil
	}

	return entry, nil
}

// Close terminates the block subscription and releases any associated resources.
func (b *BlocksProcessor) Close() error {
	if b.active.Load() {
		if b.sub != nil {
			b.sub.Unsubscribe()
		}
		b.active.Store(false)
	}

	return nil
}
