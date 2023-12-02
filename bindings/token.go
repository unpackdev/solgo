package bindings

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

const (
	Erc20 BindingType = "ERC20"
)

type Token struct {
	*Manager
	network utils.Network
	ctx     context.Context
	opts    []*BindOptions
}

func NewToken(ctx context.Context, network utils.Network, manager *Manager, opts []*BindOptions) (*Token, error) {
	if opts == nil {
		return nil, fmt.Errorf("no binding options provided for new token")
	}

	for _, opt := range opts {
		if err := opt.Validate(); err != nil {
			return nil, err
		}
	}

	// Now lets register all the bindings with the manager
	for _, opt := range opts {
		for _, network := range opt.Networks {
			if _, err := manager.RegisterBinding(network, opt.NetworkID, opt.Type, opt.Address, opt.ABI); err != nil {
				return nil, err
			}
		}
	}

	return &Token{
		Manager: manager,
		network: network,
		ctx:     ctx,
		opts:    opts,
	}, nil
}

func (t *Token) GetBinding(network utils.Network, bindingType BindingType) (*Binding, error) {
	return t.Manager.GetBinding(network, bindingType)
}

// GetName calls the name() function in the ERC-20 contract.
func (t *Token) GetName() (string, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "name")
	if err != nil {
		return "", err
	}

	name, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("failed to assert result as string - name")
	}

	return name, nil
}

// GetSymbol calls the symbol() function in the ERC-20 contract.
func (t *Token) GetSymbol() (string, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "symbol")
	if err != nil {
		return "", err
	}

	symbol, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("failed to assert result as string - symbol")
	}

	return symbol, nil
}

// GetDecimals calls the decimals() function in the ERC-20 contract.
func (t *Token) GetDecimals() (uint8, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "decimals")
	if err != nil {
		return 0, err
	}

	decimals, ok := result.(uint8)
	if !ok {
		return 0, fmt.Errorf("failed to assert result as uint8 - decimals")
	}

	return decimals, nil
}

// GetTotalSupply calls the totalSupply() function in the ERC-20 contract.
func (t *Token) GetTotalSupply() (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "totalSupply")
	if err != nil {
		return nil, err
	}

	totalSupply, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - totalSupply")
	}

	return totalSupply, nil
}

func (t *Token) BalanceOf(address common.Address) (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "balanceOf", address)
	if err != nil {
		return nil, err
	}

	balance, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - balanceOf")
	}

	return balance, nil
}

func (t *Token) Allowance(owner, spender common.Address) (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "allowance", owner, spender)
	if err != nil {
		return nil, err
	}

	allowance, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - allowance")
	}

	return allowance, nil
}

func (t *Token) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	binding, err := t.GetBinding(utils.Ethereum, Erc20)
	if err != nil {
		return nil, nil, err
	}
	bindingAbi := binding.GetABI()

	method, exists := bindingAbi.Methods["transfer"]
	if !exists {
		return nil, nil, errors.New("transfer method not found")
	}

	input, err := method.Inputs.Pack(method.Name, to, amount)
	if err != nil {
		return nil, nil, err
	}

	txHash, err := t.Manager.SendSimulatedTransaction(opts, t.network, &binding.Address, method, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send transfer transaction: %w", err)
	}

	// Wait for the receipt
	receipt, err := t.Manager.WaitForReceipt(t.ctx, t.network, *txHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transfer transaction receipt: %w", err)
	}

	tx, _, err := t.Manager.GetTransactionByHash(t.ctx, t.network, receipt.TxHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transfer transaction by hash: %w", err)
	}

	return tx, receipt, nil
}

func (t *Token) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int, simulate bool) (*types.Transaction, *types.Receipt, error) {
	binding, err := t.GetBinding(utils.Ethereum, Erc20)
	if err != nil {
		return nil, nil, err
	}
	bindingAbi := binding.GetABI()

	method, exists := bindingAbi.Methods["approve"]
	if !exists {
		return nil, nil, errors.New("approve method not found")
	}

	input, err := bindingAbi.Pack(method.Name, spender, amount)
	if err != nil {
		return nil, nil, err
	}

	if simulate {
		txHash, err := t.Manager.SendSimulatedTransaction(opts, t.network, &binding.Address, method, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send approve transaction: %w", err)
		}

		// Wait for the receipt
		receipt, err := t.Manager.WaitForReceipt(t.ctx, t.network, *txHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get approve transaction receipt: %w", err)
		}

		tx, _, err := t.Manager.GetTransactionByHash(t.ctx, t.network, receipt.TxHash)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get approve transaction by hash: %w", err)
		}

		return tx, receipt, nil
	}

	tx, err := t.Manager.SendTransaction(opts, t.network, &binding.Address, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send approve transaction: %w", err)
	}

	// Wait for the receipt
	receipt, err := t.Manager.WaitForReceipt(t.ctx, t.network, tx.Hash())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get approve transaction receipt: %w", err)
	}

	return tx, receipt, nil
}

func (t *Token) TransferFrom(from, to common.Address, amount *big.Int) (bool, error) {
	result, err := t.Manager.CallContractMethod(t.network, Erc20, "transferFrom", from, to, amount)
	if err != nil {
		return false, err
	}

	success, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("failed to assert result as bool - transferFrom")
	}

	return success, nil
}

func (t *Token) GetOptionsByNetwork(network utils.Network) *BindOptions {
	for _, opt := range t.opts {
		if t.network == network {
			return opt
		}
	}
	return nil
}

func DefaultTokenBindOptions(address common.Address) []*BindOptions {
	eip, _ := standards.GetStandard(standards.ERC20)
	return []*BindOptions{
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Type:      Erc20,
			Address:   address,
			ABI:       eip.GetStandard().ABI,
		},
	}
}
