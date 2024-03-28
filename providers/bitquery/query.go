package bitquery

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ContractCreationInfo holds data about smart contract creation events,
// including transaction hash and block height.
type ContractCreationInfo struct {
	// Data contains the nested structure of contract calls and their details.
	Data struct {
		// SmartContractCreation wraps the actual array of smart contract calls.
		SmartContractCreation struct {
			// SmartContractCalls is an array of smart contract call details.
			SmartContractCalls []struct {
				// Transaction contains the hash of the transaction.
				Transaction struct {
					Hash string `json:"hash"`
				} `json:"transaction"`
				// Block contains the height of the block.
				Block struct {
					Height int `json:"height"`
				} `json:"block"`
			} `json:"smartContractCalls"`
		} `json:"smartContractCreation"`
	} `json:"data"`
}

// GetTransactionHash returns the transaction hash of the first smart contract call
// if available, otherwise an empty string.
func (c *ContractCreationInfo) GetTransactionHash() string {
	if len(c.Data.SmartContractCreation.SmartContractCalls) == 0 {
		return ""
	}

	return c.Data.SmartContractCreation.SmartContractCalls[0].Transaction.Hash
}

// GetBlockHeight returns the block height of the first smart contract call
// if available, otherwise zero.
func (c *ContractCreationInfo) GetBlockHeight() int {
	if len(c.Data.SmartContractCreation.SmartContractCalls) == 0 {
		return 0
	}

	return c.Data.SmartContractCreation.SmartContractCalls[0].Block.Height
}

// QueryContractCreationBlockAndTxHash queries the blockchain for contract creation
// information by network group and address, returning the ContractCreationInfo.
func (b *Provider) QueryContractCreationBlockAndTxHash(ctx context.Context, networkGroup string, address common.Address) (*ContractCreationInfo, error) {
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
