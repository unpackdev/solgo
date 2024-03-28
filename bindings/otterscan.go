package bindings

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type CreatorInformation struct {
	ContractCreator common.Address `json:"creator"`
	CreationHash    common.Hash    `json:"hash"`
}

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
