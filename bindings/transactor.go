package bindings

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/utils"
)

var (
	// These values should be set based on current network conditions.
	// 2 Gwei is often a reasonable default for priority fee.
	defaultMaxPriorityFeePerGas = big.NewInt(2e9) // 2 Gwei

)

type TransactOpts struct {
	Network   utils.Network  // Network to send the transaction on
	From      common.Address // Ethereum account to send the transaction from
	Signer    bind.SignerFn  // Method to sign the transaction with
	GasPrice  *big.Int       // Gas price (nil if using dynamic fees)
	GasFeeCap *big.Int       // Max fee per gas (EIP-1559)
	GasTipCap *big.Int       // Max priority fee per gas (EIP-1559)
	Value     *big.Int       // Amount of ETH to send along with the transaction
	Nonce     uint64         // Nonce to use for the transaction (0 to use pending state)
	GasLimit  uint64         // Gas limit to set for the transaction
	Context   context.Context
	NoSend    bool // Set to true to sign but not send the transaction
}

func (m *Manager) SendTransaction(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	if opts.Context == nil || opts.GasPrice != nil && (opts.GasFeeCap != nil || opts.GasTipCap != nil) {
		return nil, errors.New("context and both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
	}

	client := m.clientPool.GetClientByGroup(opts.Network.String())
	if client == nil {
		return nil, errors.New("client not found for network")
	}

	// Determine the correct gas price to use
	var rawTx *types.Transaction
	var err error
	if opts.GasPrice != nil {
		rawTx, err = createLegacyTx(opts, contract, input)
	} else {
		head, errHead := client.HeaderByNumber(opts.Context, nil)
		if errHead != nil {
			return nil, errHead
		}
		if head.BaseFee != nil {
			rawTx, err = createDynamicTx(opts, contract, input, head)
		} else {
			rawTx, err = createLegacyTx(opts, contract, input)
		}
	}
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	// Send the transaction
	if !opts.NoSend {
		if err := client.SendTransaction(opts.Context, signedTx); err != nil {
			return nil, err
		}
	}

	return signedTx, nil
}

func createLegacyTx(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	// Create a legacy transaction
	tx := types.NewTransaction(opts.Nonce, *contract, opts.Value, opts.GasLimit, opts.GasPrice, input)
	return tx, nil
}

func createDynamicTx(opts *TransactOpts, contract *common.Address, input []byte, head *types.Header) (*types.Transaction, error) {
	// Calculate the effective gas fee cap and tip cap
	gasFeeCap := opts.GasFeeCap
	gasTipCap := opts.GasTipCap

	if gasFeeCap == nil {
		// Set default max fee per gas if not provided
		gasFeeCap = new(big.Int).Add(head.BaseFee, defaultMaxPriorityFeePerGas)
	}

	if gasTipCap == nil {
		// Set default priority fee if not provided
		gasTipCap = defaultMaxPriorityFeePerGas
	}

	// Create a dynamic fee transaction (EIP-1559)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(utils.EthereumNetworkID)),
		Nonce:     opts.Nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       opts.GasLimit,
		To:        contract,
		Value:     opts.Value,
		Data:      input,
	})
	return tx, nil
}
