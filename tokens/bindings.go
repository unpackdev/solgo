package tokens

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/accounts"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

// GetTokenBind connects to a blockchain simulator or live network to create a token binding.
// It uses bindings.Manager and can target a specific block number.
func (t *Token) GetTokenBind(ctx context.Context, simulatorType utils.SimulatorType, bindManager *bindings.Manager, atBlock *big.Int) (*bindings.Token, error) {
	if t.IsInSimulation() && t.simulator != nil {
		if _, node, err := t.simulator.GetClient(ctx, simulatorType, atBlock); err != nil {
			return nil, fmt.Errorf("failed to get client from simulator: %w", err)
		} else {
			t.descriptor.BlockNumber = node.BlockNumber
		}
	}

	return bindings.NewToken(t.ctx, t.network, bindManager, bindings.DefaultTokenBindOptions(t.descriptor.Address))
}

func (t *Token) GetBinding() *bindings.Token {
	return t.tokenBind
}

// GetTokenBind connects to a blockchain simulator or live network to create a token binding.
// It uses bindings.Manager and can target a specific block number.
func (t *Token) GetUniswapV2Bind(ctx context.Context, simulatorType utils.SimulatorType, bindManager *bindings.Manager, atBlock *big.Int) (*bindings.Uniswap, error) {
	if t.IsInSimulation() && t.simulator != nil {
		if _, node, err := t.simulator.GetClient(ctx, simulatorType, atBlock); err != nil {
			return nil, fmt.Errorf("failed to get client from simulator: %w", err)
		} else {
			t.descriptor.BlockNumber = node.BlockNumber
		}
	}

	return bindings.NewUniswapV2(t.ctx, t.network, bindManager, bindings.DefaultUniswapBindOptions())
}

// ResolveName fetches the name property of the token from the blockchain.
func (t *Token) ResolveName(ctx context.Context, from common.Address, bind *bindings.Token) (string, error) {
	return bind.GetName(ctx, from)
}

// ResolveSymbol fetches the symbol property of the token from the blockchain.
func (t *Token) ResolveSymbol(ctx context.Context, from common.Address, bind *bindings.Token) (string, error) {
	return bind.GetSymbol(ctx, from)
}

// ResolveDecimals fetches the decimals property of the token from the blockchain.
func (t *Token) ResolveDecimals(ctx context.Context, from common.Address, bind *bindings.Token) (uint8, error) {
	return bind.GetDecimals(ctx, from)
}

// ResolveTotalSupply fetches the total supply of the token from the blockchain.
func (t *Token) ResolveTotalSupply(ctx context.Context, from common.Address, bind *bindings.Token) (*big.Int, error) {
	return bind.GetTotalSupply(ctx, from)
}

// ResolveBalance retrieves the balance of a specific address for the token.
func (t *Token) ResolveBalance(ctx context.Context, from common.Address, bind *bindings.Token, address common.Address) (*big.Int, error) {
	return bind.BalanceOf(ctx, from, address)
}

// Approve facilitates the approval of a spender to spend a specified amount of tokens.
func (t *Token) Approve(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, bind *bindings.Token, spender *accounts.Account, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
	if spender == nil {
		return nil, nil, errors.New("spender address is nil")
	}

	if amount == nil {
		return nil, nil, errors.New("amount is nil")
	}

	client, err := t.GetClient(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get client: %w", err)
	}

	opts, err := spender.TransactOpts(client, nil, t.IsInSimulation())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transact opts: %w", err)
	}

	return bind.Approve(ctx, network, simulatorType, client, opts, spender.GetAddress(), amount, atBlock)
}

// Transfer facilitates the transfer of tokens to a to address.
func (t *Token) Transfer(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, bind *bindings.Token, spender *accounts.Account, to *accounts.Account, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
	if spender == nil {
		return nil, nil, errors.New("spender address is nil")
	}

	if amount == nil {
		return nil, nil, errors.New("amount is nil")
	}

	client, err := t.GetClient(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get client: %w", err)
	}

	opts, err := to.TransactOpts(client, nil, t.IsInSimulation())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transact opts: %w", err)
	}

	return bind.Transfer(ctx, network, simulatorType, client, opts, to.GetAddress(), amount, atBlock)
}

// TransferFrom facilitates the transfer of tokens from one address to another.
func (t *Token) TransferFrom(ctx context.Context, network utils.Network, simulatorType utils.SimulatorType, client *clients.Client, bind *bindings.Token, spender *accounts.Account, to common.Address, amount *big.Int, atBlock *big.Int) (*types.Transaction, *types.Receipt, error) {
	if spender == nil {
		return nil, nil, errors.New("spender address is nil")
	}

	if amount == nil {
		return nil, nil, errors.New("amount is nil")
	}

	client, err := t.GetClient(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get client: %w", err)
	}

	opts, err := spender.TransactOpts(client, nil, t.IsInSimulation())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transact opts: %w", err)
	}

	return bind.TransferFrom(ctx, network, simulatorType, client, opts, spender.GetAddress(), to, amount, atBlock)
}
