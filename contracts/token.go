package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	ctx      context.Context
	contract *Contract

	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"total_supply"`
}

func (c *Contract) IsToken() bool {
	return c.descriptor.Token != nil && c.descriptor.Token.Name != ""
}

func (c *Contract) DiscoverToken(ctx context.Context) (*Token, error) {
	toReturn := &Token{
		ctx:      ctx,
		contract: c,
	}

	if err := toReturn.Discover(); err != nil {
		return nil, fmt.Errorf("failed to discover token metadata: %s", err)
	}

	c.token = toReturn
	c.descriptor.Token = &TokenDescriptor{
		Name:        toReturn.Name,
		Symbol:      toReturn.Symbol,
		Decimals:    toReturn.Decimals,
		TotalSupply: toReturn.TotalSupply,
	}

	return toReturn, nil
}

func (t *Token) Discover() error {
	name, err := t.contract.tokenBind.GetName()
	if err != nil {
		return fmt.Errorf("failed to get token name: %s", err)
	}
	t.Name = name

	symbol, err := t.contract.tokenBind.GetSymbol()
	if err != nil {
		return fmt.Errorf("failed to get token symbol: %s", err)
	}
	t.Symbol = symbol

	decimals, err := t.contract.tokenBind.GetDecimals()
	if err != nil {
		return fmt.Errorf("failed to get token decimals: %s", err)
	}
	t.Decimals = decimals

	totalSupply, err := t.contract.tokenBind.GetTotalSupply()
	if err != nil {
		return fmt.Errorf("failed to get token total supply: %s", err)
	}
	t.TotalSupply = totalSupply

	return nil
}

func (t *Token) BalanceOf(addr common.Address) (*big.Int, error) {
	return t.contract.tokenBind.BalanceOf(addr)
}

func (t *Token) Allowance(owner, spender common.Address) (*big.Int, error) {
	return t.contract.tokenBind.Allowance(owner, spender)
}

/* func (t *Token) Transfer(opts *bindings.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	input, err := t.contract.bindings.GetABi(bindings.Erc20, "transfer", to, amount)
	if err != nil {
		return nil, err
	}

	return t.contract.bindings.SendTransaction(opts, &t.contract.addr, input)
}

func (t *Token) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.tokenBind.Approve(nil, spender, value)
}

func (t *Token) TransferFrom(from, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.tokenBind.TransferFrom(nil, from, to, value)
}
*/
