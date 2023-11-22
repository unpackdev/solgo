package observers

import (
	"context"
	"fmt"

	"github.com/0x19/solc-switch"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/providers/bitquery"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

type Manager struct {
	ctx         context.Context
	opts        Options
	clientsPool *clients.ClientPool
	blockCh     chan *BlockEntry
	hooks       map[ProcessorType]map[HookType]interface{}
	bqp         *bitquery.BitQueryProvider
	etherscan   *etherscan.EtherScanProvider
	compiler    *solc.Solc
	bindings    *bindings.Manager
}

func NewManager(ctx context.Context, clientsPool *clients.ClientPool, bqp *bitquery.BitQueryProvider, etherscan *etherscan.EtherScanProvider, compiler *solc.Solc, bindings *bindings.Manager, opts Options, blockCh chan *BlockEntry) (*Manager, error) {
	return &Manager{
		ctx:         ctx,
		opts:        opts,
		clientsPool: clientsPool,
		bqp:         bqp,
		etherscan:   etherscan,
		blockCh:     blockCh,
		hooks:       make(map[ProcessorType]map[HookType]interface{}),
		compiler:    compiler,
		bindings:    bindings,
	}, nil
}

func (m *Manager) Run() error {
	zap.L().Info(
		"Starting observers manager",
		zap.Any("strategies", m.opts.Strategies),
		zap.Any("start_block", m.opts.StartBlock),
		zap.Any("end_block", m.opts.EndBlock),
	)

	errCh := make(chan error, 1)

	subscriber, err := NewBlocksProcessor(m)
	if err != nil {
		return fmt.Errorf("failed to create block subscriber: %w", err)
	}
	defer subscriber.Close()

	for _, strategy := range m.opts.Strategies {
		switch strategy {
		case utils.HeadStrategy:
			go func() {
				if err := subscriber.SubscribeHeader(&SubscriberOptions{
					NetworkID: m.opts.NetworkID,
					Network:   m.opts.Network,
					Head:      true,
				}, m.blockCh); err != nil {
					errCh <- err
					return
				}
			}()
		case utils.ArchiveStrategy:
			go func() {
				if err := subscriber.Subscribe(&SubscriberOptions{
					NetworkID:        m.opts.NetworkID,
					Network:          m.opts.Network,
					StartBlockNumber: m.opts.StartBlock,
					EndBlockNumber:   m.opts.EndBlock,
				}, m.blockCh); err != nil {
					errCh <- err
					return
				}
			}()
		}
	}

	select {
	case err := <-errCh:
		return err
	case <-m.ctx.Done():
		return nil
	}
}

func (m *Manager) Process() error {
	errCh := make(chan error, 1)

	txsProcessor, err := NewTransactionsProcessor(m)
	if err != nil {
		return fmt.Errorf("failed to create transaction processor: %w", err)
	}
	defer txsProcessor.Close()

	contractProcessor := txsProcessor.GetContractProcessor()
	defer contractProcessor.Close()

	for i := 1; i <= 2; i++ {
		go func() {
			if err := txsProcessor.Worker(); err != nil {
				errCh <- err
				return
			}
		}()
	}

	for i := 1; i <= 2; i++ {
		go func() {
			if err := contractProcessor.Worker(); err != nil {
				errCh <- err
				return
			}
		}()
	}

	for {
		select {
		case entry := <-m.blockCh:
			zap.L().Debug(
				"New inbound block received",
				zap.Any("network_id", entry.NetworkID),
				zap.String("network", entry.Network.String()),
				zap.Any("strategy", entry.Strategy),
				zap.Uint64("block_number", entry.Block.NumberU64()),
				zap.String("block_hash", entry.Block.Hash().String()),
			)

			if err := txsProcessor.QueueBlock(entry); err != nil {
				errCh <- err
				return nil
			}

		case err := <-errCh:
			return err
		case <-m.ctx.Done():
			return nil
		}
	}
}

func (m *Manager) Close() {}
