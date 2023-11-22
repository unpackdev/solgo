package contracts

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

func (c *Contract) DiscoverConstructor(ctx context.Context) error {
	if c.descriptor.Detector != nil && c.descriptor.Detector.GetIR() != nil && c.descriptor.Detector.GetIR().GetRoot() != nil {
		detector := c.descriptor.Detector
		irRoot := detector.GetIR().GetRoot()
		abiRoot := detector.GetABI().GetRoot()

		if irRoot.GetEntryContract() != nil && irRoot.GetEntryContract().GetConstructor() != nil &&
			abiRoot.GetEntryContract().GetMethodByType("constructor") != nil {
			cAbi, _ := utils.ToJSON(abiRoot.GetEntryContract().GetMethodByType("constructor"))
			constructorAbi := fmt.Sprintf("[%s]", string(cAbi))

			tx := c.descriptor.Transaction
			deployedBytecode := c.descriptor.DeployedBytecode
			position := bytes.Index(tx.Data(), deployedBytecode[:20])
			if position != -1 {
				adjustedData := tx.Data()[position:]

				constructor, err := bytecode.DecodeConstructorFromAbi(adjustedData[len(deployedBytecode):], constructorAbi)
				if err != nil {
					if !strings.Contains(err.Error(), "would go over slice boundary") {
						zap.L().Error(
							"failed to decode constructor from bytecode",
							zap.Error(err),
							zap.Any("network", c.network),
							zap.String("contract_address", c.addr.String()),
						)
					}
					return fmt.Errorf("failed to decode constructor from bytecode: %s", err)
				}
				c.descriptor.Constructor = constructor
			}
		}
	}

	return nil
}
