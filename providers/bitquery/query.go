package bitquery

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type ContractCreationInfo struct {
	Data struct {
		SmartContractCreation struct {
			SmartContractCalls []struct {
				Transaction struct {
					Hash string `json:"hash"`
				} `json:"transaction"`
				Block struct {
					Height int `json:"height"`
				} `json:"block"`
			} `json:"smartContractCalls"`
		} `json:"smartContractCreation"`
	} `json:"data"`
}

func (c *ContractCreationInfo) GetTransactionHash() string {
	if len(c.Data.SmartContractCreation.SmartContractCalls) == 0 {
		return ""
	}

	return c.Data.SmartContractCreation.SmartContractCalls[0].Transaction.Hash
}

func (c *ContractCreationInfo) GetBlockHeight() int {
	if len(c.Data.SmartContractCreation.SmartContractCalls) == 0 {
		return 0
	}

	return c.Data.SmartContractCreation.SmartContractCalls[0].Block.Height
}

func (b *BitQueryProvider) QueryContractCreationBlockAndTxHash(ctx context.Context, networkGroup string, address common.Address) (*ContractCreationInfo, error) {
	queryData := map[string]string{
		"query": fmt.Sprintf(`{
		smartContractCreation: ethereum(network: %s) {
		  smartContractCalls(
			smartContractAddress: {is: "%s"}
			smartContractMethod: {is: "Contract Creation"}
		  ) {
			transaction {
			  hash
			}
			block {
			  height
			}
		  }
		}
	  }`, networkGroup, address.Hex()),
	}

	return b.GetContractCreationInfo(ctx, queryData)
}
