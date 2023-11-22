package observers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

type ContractsProcessor struct {
	*Manager
	chTransactions chan *TransactionEntry
	hooks          map[HookType]ContractHookFn
}

func NewContractsProcessor(manager *Manager) (*ContractsProcessor, error) {
	hooks := map[HookType]ContractHookFn{}
	for hook, fn := range manager.GetHooks(ContractProcessor) {
		hooks[hook] = fn.(ContractHookFn)
	}

	return &ContractsProcessor{
		Manager:        manager,
		chTransactions: make(chan *TransactionEntry, 100000),
		hooks:          hooks,
	}, nil
}

func (p *ContractsProcessor) Worker() error {
	for {
		select {
		case entry := <-p.chTransactions:
			contractEntry := &ContractEntry{
				NetworkID:         entry.NetworkID,
				Network:           entry.Network,
				Strategy:          entry.Strategy,
				TransactionType:   entry.TransactionType,
				ContractAddress:   entry.ContractAddress,
				Sender:            entry.Sender,
				SenderType:        entry.SenderType,
				SenderContract:    entry.SenderContract,
				Contract:          entry.Contract,
				Recipient:         entry.Recipient,
				RecipientType:     entry.RecipientType,
				RecipientContract: entry.RecipientContract,
				BlockHeader:       entry.BlockHeader,
				Transaction:       entry.Transaction,
				Receipt:           entry.Receipt,
			}

			if hook, ok := p.hooks[PreHook]; ok {
				entry, err := hook(contractEntry)
				if err != nil {
					zap.L().Error(
						"Pre hook failed",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
					)
					continue
				}
			}

			zap.L().Debug(
				"Received new inbound contract",
				zap.Any("network_id", entry.NetworkID),
				zap.String("network", entry.Network.String()),
				zap.Any("strategy", entry.Strategy),
				zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
				zap.String("block_hash", entry.BlockHeader.Hash().String()),
				zap.String("tx_hash", entry.Transaction.Hash().String()),
				zap.String("contract_address", entry.ContractAddress.String()),
			)

			contract, err := p.Unpack(p.ctx, entry.Network, entry.ContractAddress)
			if err != nil {
				zap.L().Error(
					"failed to create new contract instance",
					zap.Error(err),
					zap.Any("network_id", entry.NetworkID),
					zap.String("network", entry.Network.String()),
					zap.Any("strategy", entry.Strategy),
					zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
					zap.String("tx_hash", entry.Transaction.Hash().String()),
					zap.String("contract_address", entry.ContractAddress.String()),
				)
				continue
			}
			contractEntry.Contract = contract

			if hook, ok := p.hooks[PostHook]; ok {
				entry, err := hook(contractEntry)
				if err != nil {
					zap.L().Error(
						"Post hook failed",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
					)
					continue
				}
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *ContractsProcessor) Unpack(ctx context.Context, network utils.Network, addr common.Address) (*contracts.Contract, error) {
	contract, err := contracts.NewContract(ctx, network, p.clientsPool, nil, p.bqp, p.etherscan, p.compiler, p.bindings, addr)
	if err != nil {
		return nil, err
	}

	// This is critical error and we should not continue if we can't discover block, transaction and receipt information.
	if err := contract.DiscoverChainInfo(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover chain info: %s", err)
	}

	// Contract may have source code or may not. We should not treat this as critical error.
	contract.DiscoverSourceCode(ctx)

	// Contract may have token or may not. We should not treat this as critical error.
	contract.DiscoverToken(ctx)

	// Now we should attempt to parse contract's source code if we have it under our disposal.
	contract.Parse(ctx)

	// Lets see if we can discover contract constructor information
	contract.DiscoverConstructor(ctx)

	// Now lets try to audit the contract source code
	contract.Audit(ctx)

	// How about potential liquidity?
	contract.DiscoverLiquidity(ctx)

	return contract, nil
}

func (p *ContractsProcessor) QueueContract(entry *TransactionEntry) error {
	p.chTransactions <- entry
	return nil
}
