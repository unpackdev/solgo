package bindings

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

const (
	Erc20 BindingType = "ERC20"
)

type Token struct {
	*Manager
	ctx  context.Context
	opts []*BindOptions
}

func NewToken(ctx context.Context, manager *Manager, opts []*BindOptions) (*Token, error) {
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
		if _, err := manager.RegisterBinding(opt.Network, opt.NetworkID, opt.Type, opt.Address, opt.ABI); err != nil {
			return nil, err
		}
	}

	return &Token{
		Manager: manager,
		ctx:     ctx,
		opts:    opts,
	}, nil
}

// GetName calls the name() function in the ERC-20 contract.
func (t *Token) GetName() (string, error) {
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "name")
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
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "symbol")
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
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "decimals")
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
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "totalSupply")
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
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "balanceOf", address)
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
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "allowance", owner, spender)
	if err != nil {
		return nil, err
	}

	allowance, ok := result.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to assert result as *big.Int - allowance")
	}

	return allowance, nil
}

func (t *Token) Transfer(to common.Address, amount *big.Int) (bool, error) {
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "transfer", to, amount)
	if err != nil {
		return false, err
	}

	success, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("failed to assert result as bool - transfer")
	}

	return success, nil
}

func (t *Token) Approve(spender common.Address, amount *big.Int) (bool, error) {
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "approve", spender, amount)
	if err != nil {
		return false, err
	}

	success, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("failed to assert result as bool - approve")
	}

	return success, nil
}

func (t *Token) TransferFrom(from, to common.Address, amount *big.Int) (bool, error) {
	opts := t.GetOptionsByNetwork(utils.Ethereum)
	result, err := t.Manager.CallContractMethod(opts.Network, Erc20, "transferFrom", from, to, amount)
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
		if opt.Network == network {
			return opt
		}
	}
	return nil
}

func DefaultTokenBindOptions(address common.Address) []*BindOptions {
	eip, _ := standards.GetStandard(standards.ERC20)
	return []*BindOptions{
		{
			Network:   utils.Ethereum,
			NetworkID: utils.EthereumNetworkID,
			Type:      Erc20,
			Address:   address,
			ABI:       eip.GetStandard().ABI,
		},
	}
}
