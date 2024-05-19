package bindings

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

func (m *Manager) TraceCallMany(ctx context.Context, network utils.Network, sender common.Address, nonce int64) (*common.Hash, error) {
	client := m.clientPool.GetClientByGroup(network.String())
	if client == nil {
		return nil, fmt.Errorf("client not found for network %s", network)
	}

	rpcClient := client.GetRpcClient()
	var result *common.Hash

	if err := rpcClient.CallContext(ctx, &result, "trace_callMany", sender.Hex(), nonce); err != nil {
		return nil, errors.Wrap(err, "failed to execute trace_callMany")
	}

	return result, nil
}
