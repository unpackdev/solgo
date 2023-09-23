package observers

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
	"go.uber.org/zap"
)

// Contract represents a blockchain contract with its associated details.
type Contract struct {
	NetworkID       int64              `json:"network_id"`    // The network ID of the Ethereum blockchain.
	NetworkGroup    string             `json:"network_group"` // The group identifier for the blockchain client.
	NetworkType     string             `json:"network_type"`  // The type of the blockchain client.
	ContractAddress common.Address     `json:"contract_addr"` // Ethereum address of the contract.
	Block           *types.Block       `json:"block"`         // Block in which the contract transaction was included.
	Transaction     *types.Transaction `json:"transaction"`   // Transaction that created the contract.
	Receipt         *types.Receipt     `json:"receipt"`       // Receipt of the contract creation transaction.
}

// ContractSubscriberOptions defines the options for configuring a contract subscriber.
type ContractSubscriberOptions struct {
	NetworkID        int64    `mapstructure:"network_id" yaml:"network_id" json:"network_id"`                         // The network ID of the Ethereum blockchain.
	Group            string   `mapstructure:"group" yaml:"group" json:"group"`                                        // The group identifier for the blockchain client.
	Type             string   `mapstructure:"type" yaml:"type" json:"type"`                                           // The type of the blockchain client.
	Head             bool     `mapstructure:"head" yaml:"head" json:"head"`                                           // If true, subscribes to the latest block. Otherwise, subscribes to a range.
	StartBlockNumber *big.Int `mapstructure:"start_block_number" yaml:"start_block_number" json:"start_block_number"` // Starting block number for the subscription.
	EndBlockNumber   *big.Int `mapstructure:"end_block_number" yaml:"end_block_number" json:"end_block_number"`       // Ending block number for the subscription.
}

// ContractSubscriber provides methods to subscribe to and interact with Ethereum contracts.
type ContractSubscriber struct {
	ctx    context.Context       // The context for the subscriber operations.
	client *clients.ClientPool   // Client pool for accessing Ethereum clients.
	active atomic.Bool           // Indicates if the subscriber is currently active.
	sub    ethereum.Subscription // Ethereum subscription for real-time updates.
}

// NewContractSubscriber initializes a new contract subscriber with the provided context and client pool.
func NewContractSubscriber(ctx context.Context, client *clients.ClientPool) (*ContractSubscriber, error) {
	return &ContractSubscriber{
		ctx:    ctx,
		client: client,
	}, nil
}

// IsActive checks if the contract subscriber is currently active.
func (b *ContractSubscriber) IsActive() bool {
	return b.active.Load()
}

// Subscribe initiates a subscription based on the provided options.
// It can either subscribe to the latest blocks or a specific range of blocks.
func (b *ContractSubscriber) Subscribe(opts *ContractSubscriberOptions, contractsCh chan *Contract) error {
	if b.active.Load() {
		return errors.New("block subscriber is already active")
	}

	client := b.client.GetClientByGroupAndType(opts.Group, opts.Type)
	if client == nil {
		return errors.New("client not found")
	}

	if opts.Head {
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
					zap.L().Error(
						"failure while searching for block",
						zap.Error(err),
						zap.Int64("block_number", block.Number().Int64()),
					)
					continue
				}

				contracts, err := b.discoverContracts(block, opts)
				if err != nil {
					zap.L().Error(
						"failure while searching for block contracts",
						zap.Error(err),
						zap.Int64("block_number", block.Number().Int64()),
					)
					continue
				}

				zap.L().Debug(
					"Determined contracts in block",
					zap.Int64("block_number", block.Number().Int64()),
					zap.Int("contracts", len(contracts)),
				)

				if len(contracts) > 0 {
					for _, contract := range contracts {
						contractsCh <- contract
					}
				}
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

		if opts.EndBlockNumber == nil {
			header, err := client.HeaderByNumber(b.ctx, nil)
			if err != nil {
				return err
			}
			opts.EndBlockNumber = header.Number
		}

		if opts.EndBlockNumber.Int64() < opts.StartBlockNumber.Int64() {
			return errors.New("end block number is less than start block number")
		}

		b.active.Store(true)

		for i := opts.StartBlockNumber.Int64(); i <= opts.EndBlockNumber.Int64(); i++ {
			block, err := client.BlockByNumber(b.ctx, big.NewInt(i))
			if err != nil {
				zap.L().Error(
					"failure while searching for block",
					zap.Error(err),
					zap.Int64("block_number", i),
				)
				continue
			}

			contracts, err := b.discoverContracts(block, opts)
			if err != nil {
				zap.L().Error(
					"failure while searching for block contracts",
					zap.Error(err),
					zap.Int64("block_number", block.Number().Int64()),
				)
				continue
			}

			if len(contracts) > 0 {
				for _, contract := range contracts {
					contractsCh <- contract
				}
			}
		}

		b.active.Store(false)
	}

	return nil
}

// discoverContracts searches for contracts within a given block based on the provided options.
// It returns a list of discovered contracts.
func (b *ContractSubscriber) discoverContracts(block *types.Block, opts *ContractSubscriberOptions) ([]*Contract, error) {
	contracts := make([]*Contract, 0)

	errCh := make(chan error)

	var wg sync.WaitGroup

	for _, tx := range block.Transactions() {
		wg.Add(1)
		go func(tx *types.Transaction) {
			defer wg.Done()
			client := b.client.GetClientByGroupAndType(opts.Group, opts.Type)

			receipt, err := client.TransactionReceipt(b.ctx, tx.Hash())
			if err != nil {
				errCh <- err
				return
			}

			if receipt.Status == types.ReceiptStatusSuccessful && receipt.ContractAddress != (common.Address{}) {
				contracts = append(contracts, &Contract{
					NetworkID:       opts.NetworkID,
					NetworkGroup:    opts.Group,
					NetworkType:     opts.Type,
					ContractAddress: receipt.ContractAddress,
					Block:           block,
					Transaction:     tx,
					Receipt:         receipt,
				})
			}
		}(tx)
	}

	wg.Wait()

	close(errCh)

	if len(errCh) > 0 {
		return contracts, <-errCh
	}

	return contracts, nil
}

// Close terminates the contract subscription and releases any associated resources.
func (b *ContractSubscriber) Close() error {
	if b.active.Load() {
		if b.sub != nil {
			b.sub.Unsubscribe()
		}
		b.active.Store(false)
	}

	return nil
}
