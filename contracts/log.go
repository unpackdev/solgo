package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

// Log represents a structured version of an Ethereum log entry, including both
// the raw log information and its decoded form if applicable. It serves as a
// comprehensive data structure for handling log data within the system.
type Log struct {
	Log         *types.Log     // The original log entry from Ethereum.
	Address     common.Address // The address from which the log originated.
	Topics      []common.Hash  // Topics are essentially event signatures and optionally indexed event parameters.
	Data        []byte         // The data contains non-indexed event parameters.
	BlockNumber uint64         // The block number where this log was recorded.
	TxHash      common.Hash    // The hash of the transaction that generated this log.
	TxIndex     uint           // The index of the transaction in the block.
	BlockHash   common.Hash    // The hash of the block where this log was recorded.
	Index       uint           // The log's index position in the block.
	Removed     bool           // Flag indicating if the log was reverted due to a chain reorganization.
	DecodedLog  *bytecode.Log  // The decoded log information, providing a structured interpretation of the data.
}

// DecodeLog takes a raw Ethereum log entry and attempts to decode it into a structured
// Log instance using the contract's ABI. This method enables easier interpretation of
// log entries, which are crucial for monitoring contract events and state changes.
func (c *Contract) DecodeLog(ctx context.Context, log *types.Log) (*Log, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		toReturn := &Log{
			Log:         log,
			Address:     log.Address,
			Topics:      log.Topics,
			Data:        log.Data,
			BlockNumber: log.BlockNumber,
			TxHash:      log.TxHash,
			TxIndex:     log.TxIndex,
			BlockHash:   log.BlockHash,
			Index:       log.Index,
			Removed:     log.Removed,
		}

		if c.descriptor.Detector != nil && c.descriptor.Detector.GetIR() != nil && c.descriptor.Detector.GetIR().GetRoot() != nil {
			detector := c.descriptor.Detector
			irRoot := detector.GetIR().GetRoot()
			abiRoot := detector.GetABI().GetRoot()

			if irRoot.GetEntryContract() != nil {
				for _, contract := range abiRoot.GetContracts() {
					jsonData, err := utils.ToJSON(contract)
					if err != nil {
						return nil, fmt.Errorf("failed to convert contract to json: %s", err)
					}

					logData, err := bytecode.DecodeLogFromAbi(log, jsonData)
					if err != nil {
						zap.L().Debug(
							"failed to decode log from abi",
							zap.Error(err),
							zap.Any("network", c.network),
							zap.String("contract_address", c.addr.String()),
							zap.String("decode_type", "from_decoder"),
						)
						continue
					}
					toReturn.DecodedLog = logData
					return toReturn, nil
				}
			}
		} else if len(c.descriptor.ABI) > 0 {
			logData, err := bytecode.DecodeLogFromAbi(log, []byte(c.descriptor.ABI))
			if err != nil {
				zap.L().Debug(
					"failed to decode log from abi",
					zap.Error(err),
					zap.Any("network", c.network),
					zap.String("contract_address", c.addr.String()),
					zap.String("decode_type", "from_decoder"),
				)
				return nil, err
			}
			toReturn.DecodedLog = logData
			return toReturn, nil
		}

		return nil, fmt.Errorf("failed to decode log from abi: %s", "signature not found")
	}
}
