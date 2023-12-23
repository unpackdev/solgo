package contracts

import (
	"context"
	"fmt"
	"strings"

	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

func (c *Contract) DecodeTransaction(ctx context.Context, data []byte) (*bytecode.Transaction, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("invalid transaction data length: %d", len(data))
	}

	// The first 4 bytes of the t represent the ID of the method in the ABI.
	methodSigData := data[:4]

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

				transaction, err := bytecode.DecodeTransactionFromAbi(data, jsonData)
				if err != nil {
					if !strings.Contains(err.Error(), "failed to get method by id") {
						zap.L().Error(
							"failed to decode transaction from abi",
							zap.Error(err),
							zap.Any("network", c.network),
							zap.String("contract_address", c.addr.String()),
							zap.Binary("method_signature_data", methodSigData),
							zap.String("decode_type", "from_contract"),
						)
					}
					continue
				}

				return transaction, nil
			}
		}
	} else if len(c.descriptor.ABI) > 0 { // We have ABI thankfully...
		transaction, err := bytecode.DecodeTransactionFromAbi(data, []byte(c.descriptor.ABI))
		if err != nil {
			if !strings.Contains(err.Error(), "failed to get method by id") {
				zap.L().Error(
					"failed to decode transaction from abi",
					zap.Error(err),
					zap.Any("network", c.network),
					zap.String("contract_address", c.addr.String()),
					zap.Binary("method_signature_data", methodSigData),
					zap.String("decode_type", "from_raw_abi"),
				)
			}
			return nil, fmt.Errorf("failed to decode transaction from abi: %s", err)
		}

		return transaction, nil
	}

	// Signature itself is not found due to any of the above reasons...
	// What you could do is have post hook to figure out based on your own signature logic.

	return nil, fmt.Errorf("failed to decode transaction from abi: %s", "signature not found")
}
