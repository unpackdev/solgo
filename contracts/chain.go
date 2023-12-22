package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

func (c *Contract) DiscoverChainInfo(ctx context.Context) error {
	// What we are going to do, as erigon node is used in this particular case, is to query etherscan only if
	// otterscan is not available.

	info, err := c.bindings.GetContractCreator(ctx, c.network, c.addr)
	if err != nil {
		return fmt.Errorf("failed to get contract creator: %w", err)
	}

	var txHash common.Hash

	if info == nil || info.CreationHash == utils.ZeroHash {
		// Prior we continue with the unpacking of the contract, we want to make sure that we can reach properly
		// contract transaction and associated creation block. If we can't, we're not going to unpack it.
		cInfo, err := c.etherscan.QueryContractCreationTx(ctx, c.addr)
		if err != nil {
			return fmt.Errorf("failed to query contract creation block and tx hash: %w", err)
		}
		txHash = cInfo.GetTransactionHash()
	} else {
		txHash = info.CreationHash
	}

	// Alright now lets extract block and transaction as well as receipt from the blockchain.
	// We're going to use archive node for this, as we want to be sure that we can get all the data.

	tx, _, err := c.client.TransactionByHash(ctx, txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction by hash: %s", err)
	}
	c.descriptor.Transaction = tx

	receipt, err := c.client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt by hash: %s", err)
	}
	c.descriptor.Receipt = receipt

	block, err := c.client.BlockByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		return fmt.Errorf("failed to get block by number: %s", err)
	}
	c.descriptor.Block = block

	return nil
}
