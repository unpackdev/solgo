package bindings

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// CreatorInformation holds data related to the creation of a smart contract on the Ethereum blockchain. It includes
// the address of the contract creator and the transaction hash of the contract creation transaction. This struct
// is typically used to retrieve and store information about how and by whom a contract was deployed to the blockchain.
type CreatorInformation struct {
	ContractCreator common.Address `json:"creator"` // The Ethereum address of the contract creator.
	CreationHash    common.Hash    `json:"hash"`    // The hash of the transaction used to create the contract.
}

// GetContractCreator queries the Ethereum blockchain to find the creator of a specified smart contract. This method
// utilizes the Ethereum JSON-RPC API to request creator information, which includes both the creator's address and
// the transaction hash of the contract's creation. It's a valuable tool for auditing and tracking the origins of
// contracts on the network. WORKS ONLY WITH ERIGON NODE - OR NODES THAT SUPPORT OTTERSCAN!
func (m *Manager) GetContractCreator(ctx context.Context, network utils.Network, contract common.Address) (*CreatorInformation, error) {
	client := m.clientPool.GetClientByGroup(string(network))
	if client == nil {
		return nil, fmt.Errorf("client not found for network %s", network)
	}

	rpcClient := client.GetRpcClient()
	var result *CreatorInformation

	if err := rpcClient.CallContext(ctx, &result, "ots_getContractCreator", contract.Hex()); err != nil {
		return nil, fmt.Errorf("failed to fetch otterscan creator information: %v", err)
	}

	return result, nil
}
