package observers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type TransactionsProcessor struct {
	*Manager
	chTransactions chan *TransactionEntry
	hooks          map[HookType]TransactionHookFn
	contracts      *ContractsProcessor
	logs           *LogsProcessor
}

func NewTransactionsProcessor(manager *Manager) (*TransactionsProcessor, error) {
	hooks := map[HookType]TransactionHookFn{}
	for hook, fn := range manager.GetHooks(TransactionProcessor) {
		hooks[hook] = fn.(TransactionHookFn)
	}

	contracts, err := NewContractsProcessor(manager)
	if err != nil {
		return nil, err
	}

	logs, err := NewLogsProcessor(manager)
	if err != nil {
		return nil, err
	}

	return &TransactionsProcessor{
		Manager:        manager,
		chTransactions: make(chan *TransactionEntry, 100000),
		hooks:          hooks,
		contracts:      contracts,
		logs:           logs,
	}, nil
}

func (p *TransactionsProcessor) GetContractProcessor() *ContractsProcessor {
	return p.contracts
}

func (p *TransactionsProcessor) GetLogsProcessor() *LogsProcessor {
	return p.logs
}

func (p *TransactionsProcessor) Close() {
	p.contracts.Close()
}

func (p *TransactionsProcessor) Worker() error {
	for {
		select {
		case entry := <-p.chTransactions:
			entry, err := p.processTransaction(entry)
			if err != nil {
				continue
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *TransactionsProcessor) processTransaction(entry *TransactionEntry) (*TransactionEntry, error) {
	if hook, ok := p.hooks[PreHook]; ok {
		entry, err := hook(entry)
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
			return entry, fmt.Errorf("process transaction pre hook failed: %w", err)
		}
	}

	zap.L().Debug(
		"Received new inbound transaction",
		zap.Any("network_id", entry.NetworkID),
		zap.String("network", entry.Network.String()),
		zap.Any("strategy", entry.Strategy),
		zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
		zap.String("block_hash", entry.BlockHeader.Hash().String()),
		zap.String("tx_hash", entry.Transaction.Hash().String()),
	)

	client := p.clientsPool.GetClientByGroup(entry.Network.String())

	// In case that receipt is not available, we need to fetch one...
	// In the future we may have cases where receipt is already passed from some other source and if that is the case
	// we for sure do not want to fetch it again...
	if entry.Receipt == nil {
		if client == nil {
			zap.L().Error(
				"Client not found while processing transaction receipt",
				zap.Any("network_id", entry.NetworkID),
				zap.String("network", entry.Network.String()),
				zap.Any("strategy", entry.Strategy),
				zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
				zap.String("tx_hash", entry.Transaction.Hash().String()),
			)
			return entry, fmt.Errorf("client not found for network %s", entry.Network.String())
		}

		receipt, err := client.TransactionReceipt(p.ctx, entry.Transaction.Hash())
		if err != nil {
			zap.L().Error(
				"Failed to get transaction receipt",
				zap.Error(err),
				zap.String("tx_hash", entry.Transaction.Hash().String()),
			)
			return entry, err
		}
		entry.Receipt = receipt
	}

	// In case that sender address is not set, we need to compute one.
	if entry.Sender == (common.Address{}) {
		from, err := types.Sender(types.LatestSignerForChainID(entry.Transaction.ChainId()), entry.Transaction)
		if err != nil {
			zap.L().Error(
				"Failed to get transaction sender",
				zap.Error(err),
				zap.Any("network_id", entry.NetworkID),
				zap.String("network", entry.Network.String()),
				zap.Any("strategy", entry.Strategy),
				zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
				zap.String("tx_hash", entry.Transaction.Hash().String()),
			)
		}
		entry.Sender = from
		codeAt, _ := client.CodeAt(p.ctx, from, entry.BlockHeader.Number)
		if len(codeAt) > 10 {
			entry.SenderType = utils.ContractRecipient
		} else {
			entry.SenderType = utils.AddressRecipient
		}
	}

	// In case that contract address is not set, we need to fetch one.
	// This is under assumption that transaction is a contract deployment.
	if entry.Transaction.To() == nil {
		entry.TransactionType = utils.ContractCreationType
		entry.ContractAddress = entry.Receipt.ContractAddress
		entry.RecipientType = utils.ZeroAddressRecipient
		entry.Recipient = utils.ZeroAddress
	}

	// In case that recipient address is not set, we need to fetch one.
	if entry.Transaction.To() != nil && entry.Recipient == (common.Address{}) {
		entry.Recipient = *entry.Transaction.To()
		// @TODO: Handler errors properly....
		codeAt, _ := client.CodeAt(p.ctx, *entry.Transaction.To(), entry.BlockHeader.Number)
		if len(codeAt) > 10 {
			entry.RecipientType = utils.ContractRecipient
		} else {
			entry.RecipientType = utils.AddressRecipient
		}
	}

	group := errgroup.Group{}

	// Now we are going to do something quite clever (at least what I think for now).
	// Transaction destination or source can be a contract, if we don't have contract, it's pretty much useless to continue
	// with transaction processing as we're not having important pieces of information.
	// We're going to attempt sort out contract prior continuing anywhere else with this transaction....
	// This will slower down transaction processing but in the long run we'll have much more reliable information and
	// cache for future transactions...
	// MISSION: Get the transaction information (ABI) associated with the contract. If not possible will deal with
	// it later on...
	// WARN: This process is going to occur within contract processor as there we'll have potential source code
	// available... It's going to attempt decode transaction data.
	if entry.SenderType == utils.ContractRecipient {
		group.Go(func() error {
			contract, err := p.contracts.Unpack(p.ctx, entry.Network, entry.Sender, entry)
			if err != nil {

				return err
			}

			entry.SenderContract = contract
			return nil
		})
	}

	if entry.RecipientType == utils.ContractRecipient {
		group.Go(func() error {
			contract, err := p.contracts.Unpack(p.ctx, entry.Network, entry.Recipient, entry)
			if err != nil {

				return err
			}

			entry.RecipientContract = contract
			return nil
		})
	}

	if entry.TransactionType == utils.ContractCreationType {
		group.Go(func() error {
			contract, err := p.contracts.Unpack(p.ctx, entry.Network, entry.ContractAddress, entry)
			if err != nil {
				zap.L().Error(
					"Failed to unpack transaction contract",
					zap.Error(err),
					zap.Any("network_id", entry.NetworkID),
					zap.String("network", entry.Network.String()),
					zap.Any("strategy", entry.Strategy),
					zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
					zap.String("tx_hash", entry.Transaction.Hash().String()),
					zap.String("sender_address", entry.Sender.String()),
				)
				return err
			}
			entry.Contract = contract
			return nil
		})
	}

	if len(entry.Receipt.Logs) > 0 {
		group.Go(func() error {
			for _, logEntry := range entry.Receipt.Logs {
				contract, err := p.contracts.Unpack(p.ctx, entry.Network, logEntry.Address, entry)
				if err != nil {
					zap.L().Error(
						"Failed to unpack transaction log contract",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
						zap.String("sender_address", entry.Sender.String()),
						zap.String("log_address", logEntry.Address.String()),
					)
					continue
				}
				entry.LogContracts[logEntry.Address] = contract
			}

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		zap.L().Error(
			"failed to unpack transaction and associated information",
			zap.Error(err),
			zap.Any("network_id", entry.NetworkID),
			zap.String("network", entry.Network.String()),
			zap.Any("strategy", entry.Strategy),
			zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
			zap.String("tx_hash", entry.Transaction.Hash().String()),
		)
	}

	if err := p.decodeMethod(entry); err != nil {
		zap.L().Error(
			"failed to decode transaction method",
			zap.Error(err),
			zap.Any("network_id", entry.NetworkID),
			zap.String("network", entry.Network.String()),
			zap.Any("strategy", entry.Strategy),
			zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
			zap.String("tx_hash", entry.Transaction.Hash().String()),
		)
	}

	if err := p.decodeLogs(entry); err != nil {
		zap.L().Error(
			"failed to decode transaction logs",
			zap.Error(err),
			zap.Any("network_id", entry.NetworkID),
			zap.String("network", entry.Network.String()),
			zap.Any("strategy", entry.Strategy),
			zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
			zap.String("tx_hash", entry.Transaction.Hash().String()),
		)
	}

	// TODO: If we cannot discover contract in database we need to unpack contract and add it to the database...
	// Example is Tether or USDC....

	// @TODO: This is just basic implementation and is not fully correct as we are not checking for type but rather if
	// contract is found or not...

	// At this point we are going to try figuring out remaining transaction types...
	// Transaction type is basically a method that is being invoked on the contract.
	if entry.TransactionType != utils.ContractCreationType {
		if entry.SenderContract != nil && entry.RecipientContract != nil {
			// If we have both sender and recipient contract, we can assume that this is a contract to contract
			// transaction.
			zap.L().Info("Contract to contract transaction detected, this is not yet implemented")
		} else if entry.SenderContract == nil && entry.RecipientContract != nil {
			// If we have recipient contract but not sender contract, we can assume that this is contract method
			// invocation.
			transaction, err := entry.RecipientContract.DecodeTransaction(p.ctx, entry.Transaction.Data())
			if err != nil {
				if !strings.Contains(err.Error(), "no method with id") && !strings.Contains(err.Error(), "signature not found") {
					zap.L().Error(
						"failed to decode transaction",
						zap.Error(err),
						zap.Any("network_id", entry.NetworkID),
						zap.String("network", entry.Network.String()),
						zap.Any("strategy", entry.Strategy),
						zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
						zap.String("tx_hash", entry.Transaction.Hash().String()),
					)
				}
			} else {
				entry.TransactionType = transaction.Type
				entry.TransactionMethod = transaction
			}
		} else if entry.SenderContract != nil && entry.RecipientContract == nil {
			zap.L().Info("Address to address transaction detected, this is not yet implemented")
		} else {
			entry.TransactionType = utils.TransferMethodType
		}
	}

	if hook, ok := p.hooks[PostHook]; ok {
		entry, err := hook(entry)
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
			return entry, fmt.Errorf("process transaction post hook failed: %w", err)
		}
	}

	// Final thing is to queue contracts and logs in case that these functionalities are
	// enabled...
	p.QueueLog(entry)

	time.Sleep(1 * time.Second)
	utils.DumpNodeWithExit(entry)

	return entry, nil
}

func (p *TransactionsProcessor) decodeMethod(entry *TransactionEntry) error {
	return nil
}

func (p *TransactionsProcessor) decodeLogs(entry *TransactionEntry) error {
	if len(entry.Receipt.Logs) == 0 {
		return nil
	}

	for _, log := range entry.Receipt.Logs {
		if logContract, ok := entry.LogContracts[log.Address]; ok {
			logEntry, err := logContract.DecodeLog(p.ctx, log)
			if err != nil {
				zap.L().Error(
					"failed to decode log",
					zap.Error(err),
					zap.Any("network_id", entry.NetworkID),
					zap.String("network", entry.Network.String()),
					zap.Any("strategy", entry.Strategy),
					zap.Uint64("block_number", entry.BlockHeader.Number.Uint64()),
					zap.String("tx_hash", entry.Transaction.Hash().String()),
					zap.Uint("log_index", log.Index),
					zap.String("log_address", log.Address.String()),
				)
				continue
			}

			entry.Logs = append(entry.Logs, logEntry)
		}
	}

	return nil
}

func (p *TransactionsProcessor) ProcessTransaction(ctx context.Context, blockUUID uuid.UUID, network utils.Network, networkId utils.NetworkID, strategy utils.Strategy, block *types.Block, tx *types.Transaction) (*TransactionEntry, error) {
	entry := &TransactionEntry{
		BlockUUID:       blockUUID,
		NetworkID:       networkId,
		Network:         network,
		Strategy:        strategy,
		BlockHeader:     block.Header(),
		Transaction:     tx,
		TransactionType: utils.UnknownTransactionMethodType, // Default, will be changed later on if discovered...
		LogContracts:    make(map[common.Address]*contracts.Contract),
	}

	entry, err := p.processTransaction(entry)
	if err != nil {
		zap.L().Error(
			"Failed to process transaction",
			zap.Error(err),
		)
		return entry, err
	}

	return entry, nil
}

func (p *TransactionsProcessor) QueueLog(entry *TransactionEntry) error {
	for _, log := range entry.Logs {
		p.logs.QueueLog(&LogEntry{
			BlockUUID:         entry.BlockUUID,
			TransactionUUID:   entry.UUID,
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
			Log:               log,
			LogContract: func() *contracts.Contract {
				if contract, ok := entry.LogContracts[log.Address]; ok {
					return contract
				}

				return nil
			}(),
		})
	}
	return nil
}

func (p *TransactionsProcessor) QueueBlock(entry *BlockEntry) error {
	for _, tx := range entry.Block.Transactions() {
		p.chTransactions <- &TransactionEntry{
			BlockUUID:       entry.UUID,
			NetworkID:       entry.NetworkID,
			Network:         entry.Network,
			Strategy:        entry.Strategy,
			BlockHeader:     entry.Block.Header(),
			Transaction:     tx,
			TransactionType: utils.UnknownTransactionMethodType, // Default, will be changed later on if discovered...
			LogContracts:    make(map[common.Address]*contracts.Contract),
		}
	}

	return nil
}
