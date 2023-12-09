package tokens

import (
	"context"
	"fmt"
	"math/big"

	"github.com/unpackdev/solgo/accounts"
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

// @TODO: Figure out how to dynamically handle exchanges in the future...
func (t *Token) Buy(ctx context.Context, exchangeType utils.ExchangeType, simulatorType utils.SimulatorType, spender *accounts.Account, baseToken *entities.Token, amount *entities.CurrencyAmount, atBlock *big.Int) (any, error) {
	client, err := t.GetSimulatedClient(ctx, simulatorType, atBlock)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf("failed to get client: %s", err)
	}

	// let's ensure that spender itself is always connected to simulator client
	spender.SetClient(client)
	spender.SetAccountBalance(ctx, t.simulator.GetOptions().FaucetAccountDefaultBalance)

	block := t.descriptor.BlockNumber
	if atBlock != nil {
		block = atBlock
	}

	tokenBinding, err := t.GetTokenBind(ctx, simulatorType, t.bindManager, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get token bindings: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniswapBinding, err := t.GetUniswapV2Bind(ctx, simulatorType, t.bindManager, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get uniswap bindings: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniswapV2Exchange, err := exchanges.NewUniswapV2(ctx, t.GetClientPool(), t.simulator, uniswapBinding, exchanges.DefaultOptions().GetExchange(exchangeType))
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get uniswap exchange: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniV2 := exchanges.ToUniswapV2(uniswapV2Exchange)
	tradeResponse, err := uniV2.Buy(ctx, t.network, simulatorType, tokenBinding, spender, baseToken, t.GetEntity(), amount, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to buy token: %s, exchange: %s, base_addr: %s, quote_addr: %s", err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	return tradeResponse, nil
}

func (t *Token) Sell(ctx context.Context, exchangeType utils.ExchangeType, simulatorType utils.SimulatorType, spender *accounts.Account, baseToken *entities.Token, amount *entities.CurrencyAmount, atBlock *big.Int) (any, error) {
	client, err := t.GetSimulatedClient(ctx, simulatorType, atBlock)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf("failed to get client: %s", err)
	}

	// let's ensure that spender itself is always connected to simulator client
	spender.SetClient(client)
	spender.SetAccountBalance(ctx, t.simulator.GetOptions().FaucetAccountDefaultBalance)

	block := t.descriptor.BlockNumber
	if atBlock != nil {
		block = atBlock
	}

	tokenBinding, err := t.GetTokenBind(ctx, simulatorType, t.bindManager, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get token bindings: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniswapBinding, err := t.GetUniswapV2Bind(ctx, simulatorType, t.bindManager, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get uniswap bindings: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniswapV2Exchange, err := exchanges.NewUniswapV2(ctx, t.GetClientPool(), t.simulator, uniswapBinding, exchanges.DefaultOptions().GetExchange(exchangeType))
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to get uniswap exchange: %s, exchange: %s, base_addr: %s, quote_addr: %s",
			err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	uniV2 := exchanges.ToUniswapV2(uniswapV2Exchange)
	tradeResponse, err := uniV2.Sell(ctx, t.network, simulatorType, tokenBinding, spender, t.GetEntity(), baseToken, amount, block)
	if err != nil {
		return utils.ZeroAddress, fmt.Errorf(
			"failed to sell token: %s, exchange: %s, base_addr: %s, quote_addr: %s", err,
			exchangeType, baseToken.Address, t.descriptor.Address,
		)
	}

	return tradeResponse, nil
}
