package bindings

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

// BindingType is a custom type that defines various supported token standards. Currently, it includes ERC20 and
// ERC20Ownable types, but it can be extended to support more standards in the future.
const (
	Erc20        BindingType = "ERC20"
	Erc20Ownable BindingType = "ERC20Ownable"
)

// Token encapsulates data and functions for interacting with a blockchain token. It contains network information,
// execution context, and binding options to interact with the smart contract representing the token.
type Token struct {
	*Manager
	network utils.Network
	ctx     context.Context
	opts    []*BindOptions
}

// NewToken initializes a new Token instance. It requires a context, network, Manager instance, and a slice of
// BindOptions. The function validates the provided BindOptions and registers the bindings with the Manager.
// It returns a pointer to a Token instance or an error if the initialization fails.
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
		for _, oNetwork := range opt.Networks {
			if _, err := manager.RegisterBinding(oNetwork, opt.NetworkID, opt.Type, opt.Address, opt.ABI); err != nil {
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

// NewSimpleToken is a simplified version of NewToken that also initializes a new Token instance but does not perform
// binding registrations. It's intended for scenarios where the bindings are already registered or not needed.
func NewSimpleToken(ctx context.Context, network utils.Network, manager *Manager, opts []*BindOptions) (*Token, error) {
	if opts == nil {
		return nil, fmt.Errorf("no binding options provided for new token")
	}

	for _, opt := range opts {
		if err := opt.Validate(); err != nil {
			return nil, err
		}
	}

	return &Token{
		Manager: manager,
		network: network,
		ctx:     ctx,
		opts:    opts,
	}, nil
}

// GetAddress returns the primary address of the token's smart contract as specified in the first BindOptions entry.
func (t *Token) GetAddress() common.Address {
	return t.opts[0].Address
}

// GetBinding retrieves a specific binding based on the network and binding type. It utilizes the Manager's
// GetBinding method to find the requested binding and returns it.
func (t *Token) GetBinding(network utils.Network, bindingType BindingType) (*Binding, error) {
	return t.Manager.GetBinding(network, bindingType)
}

// Owner retrieves the owner address of an ERC20Ownable token by calling the owner() contract method. It returns
// the owner's address or an error if the call fails or the result cannot be asserted to common.Address.
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

// GetName queries the name of the token by calling the name() function in the ERC-20 contract. It returns the
// token name or an error if the call fails or the result cannot be asserted as a string.
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

// BalanceOf queries the balance of a specific address within an ERC-20 token contract. It calls the balanceOf
// method of the contract and returns the balance as a *big.Int. This function is essential for determining
// the token balance of a wallet or contract address.
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

// Allowance checks the amount of tokens that an owner allowed a spender to use on their behalf. It's a crucial
// component for enabling delegated token spending in ERC-20 tokens, supporting functionalities like token approvals
// for automated market makers or decentralized exchanges.
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

// GetOptionsByNetwork retrieves the BindOptions associated with a specific network. This function is useful
// for accessing network-specific configurations like contract addresses and ABIs, facilitating multi-network
// support within the same application.
func (t *Token) GetOptionsByNetwork(network utils.Network) *BindOptions {
	for _, opt := range t.opts {
		if t.network == network {
			return opt
		}
	}
	return nil
}

// DefaultTokenBindOptions generates a default set of BindOptions for ERC20 and ERC20Ownable tokens. It presets
// configurations such as networks, network IDs, types, contract addresses, and ABIs based on standard
// implementations. This function simplifies the setup process for common token types.
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
