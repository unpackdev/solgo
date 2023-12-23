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

type Log struct {
	Log         *types.Log
	Address     common.Address
	Topics      []common.Hash
	Data        []byte
	BlockNumber uint64
	TxHash      common.Hash
	TxIndex     uint
	BlockHash   common.Hash
	Index       uint
	Removed     bool
	DecodedLog  *bytecode.Log
}

func (c *Contract) DecodeLog(ctx context.Context, log *types.Log) (*Log, error) {
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

	// There are contracts that just reuse different contracts a lot so we'll try to decode from storage
	// This is a very expensive operation so we'll only do it if we have no other choice.
	/* 	if storages := storage.GetStorages(); len(storages) > 0 {
		for _, storage := range storages {
			logData, err := bytecode.DecodeLogFromAbi(log, []byte(storage.ABI))
			if err != nil {
				zap.L().Debug(
					"failed to decode log from storage abi",
					zap.Error(err),
					zap.Any("network", c.network),
					zap.String("contract_address", c.addr.String()),
					zap.String("decode_type", "from_storage"),
				)
				continue
			}
			toReturn.DecodedLog = logData
			return toReturn, nil
		}
	} */

	// Last attempt as we have no ABI and no IR to decode from :(
	// What we have is bindings, quite a lot of them so let's try to decode from them.
	for _, binding := range c.bindings.GetBindings(c.network) {
		logData, err := bytecode.DecodeLogFromAbi(log, []byte(binding.GetRawABI()))
		if err != nil {
			zap.L().Debug(
				"failed to decode log from binding abi",
				zap.Error(err),
				zap.Any("network", c.network),
				zap.String("contract_address", c.addr.String()),
				zap.String("decode_type", "from_bindings"),
			)
			continue
		}
		toReturn.DecodedLog = logData
		return toReturn, nil
	}

	return nil, fmt.Errorf("failed to decode log from abi: %s", "signature not found")
}
