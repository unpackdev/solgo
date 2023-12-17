package bindings

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

const (
	Erc20        BindingType = "ERC20"
	Erc20Ownable BindingType = "ERC20Ownable"
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

func (t *Token) Owner(ctx context.Context, from common.Address) (common.Address, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20Ownable, from, "owner")
	if err != nil {
		return utils.ZeroAddress, err
	}

	addr, ok := result.(common.Address)
	if !ok {
		return utils.ZeroAddress, fmt.Errorf("failed to assert result as common.address - owner")
	}

	return addr, nil
}

// GetName calls the name() function in the ERC-20 contract.
func (t *Token) GetName(ctx context.Context, from common.Address) (string, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "name")
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
func (t *Token) GetSymbol(ctx context.Context, from common.Address) (string, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "symbol")
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
func (t *Token) GetDecimals(ctx context.Context, from common.Address) (uint8, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "decimals")
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
func (t *Token) GetTotalSupply(ctx context.Context, from common.Address) (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "totalSupply")
	if err != nil {
		return nil, err
	}

	totalSupply, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - totalSupply")
	}

	return totalSupply, nil
}

func (t *Token) BalanceOf(ctx context.Context, from common.Address, address common.Address) (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "balanceOf", address)
	if err != nil {
		return nil, err
	}

	balance, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - balanceOf")
	}

	return balance, nil
}

func (t *Token) Allowance(ctx context.Context, owner, from common.Address, spender common.Address) (*big.Int, error) {
	result, err := t.Manager.CallContractMethod(ctx, t.network, Erc20, from, "allowance", owner, spender)
	if err != nil {
		return nil, err
	}

	allowance, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - allowance")
	}

	return allowance, nil
}

func (t *Token) Transfer(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, opts *bind.TransactOpts, to common.Address, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
	binding, err := t.GetBinding(utils.Ethereum, Erc20)
	if err != nil {
		return nil, nil, err
	}
	bindingAbi := binding.GetABI()

	method, exists := bindingAbi.Methods["transfer"]
	if !exists {
		return nil, nil, errors.New("transfer method not found")
	}

	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	default:
		input, err := bindingAbi.Pack(method.Name, to, amount)
		if err != nil {
			return nil, nil, err
		}

		tx, err := t.Manager.SendTransaction(opts, t.network, simulatorType, client, &binding.Address, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send transfer transaction: %w", err)
		}

		receipt, err := t.Manager.WaitForReceipt(t.ctx, network, simulatorType, client, tx.Hash())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get transfer transaction receipt: %w", err)
		}

		return tx, receipt, nil
	}
}

func (t *Token) Approve(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
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

	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	default:
		tx, err := t.Manager.SendTransaction(opts, t.network, simulatorType, client, &binding.Address, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send approve transaction: %w", err)
		}

		receipt, err := t.Manager.WaitForReceipt(t.ctx, network, simulatorType, client, tx.Hash())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get approve transaction receipt: %w", err)
		}

		zap.L().Debug(
			"Approve transaction sent and receipt received",
			zap.String("tx_hash", tx.Hash().Hex()),
			zap.String("tx_from", spender.Hex()),
			zap.String("tx_to", tx.To().Hex()),
			zap.String("tx_nonce", fmt.Sprintf("%d", tx.Nonce())),
			zap.String("tx_gas_price", tx.GasPrice().String()),
			zap.String("tx_gas", fmt.Sprintf("%d", tx.Gas())),
		)

		return tx, receipt, nil
	}
}

func (t *Token) TransferFrom(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, opts *bind.TransactOpts, from, to common.Address, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
	binding, err := t.GetBinding(utils.Ethereum, Erc20)
	if err != nil {
		return nil, nil, err
	}
	bindingAbi := binding.GetABI()

	method, exists := bindingAbi.Methods["transferFrom"]
	if !exists {
		return nil, nil, errors.New("transfer method not found")
	}

	input, err := bindingAbi.Pack(method.Name, from, to, amount)
	if err != nil {
		return nil, nil, err
	}

	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	default:
		tx, err := t.Manager.SendTransaction(opts, t.network, simulatorType, client, &binding.Address, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to send transfer transaction: %w", err)
		}

		receipt, err := t.Manager.WaitForReceipt(t.ctx, network, simulatorType, client, tx.Hash())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get transfer transaction receipt: %w", err)
		}

		return tx, receipt, nil
	}
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
	eipOwnable, _ := standards.GetStandard(standards.OZOWNABLE)
	return []*BindOptions{
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Type:      Erc20,
			Address:   address,
			ABI:       eip.GetStandard().ABI,
		},
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Type:      Erc20Ownable,
			Address:   address,
			ABI:       eipOwnable.GetStandard().ABI,
		},
	}
}
