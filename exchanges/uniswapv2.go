package exchanges

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings/uniswapv2factory"
	"github.com/unpackdev/solgo/bindings/uniswapv2pool"
	"github.com/unpackdev/solgo/bindings/uniswapv2router"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/currencies"
	"go.uber.org/zap"
)

var (
	InitCodeHash = common.FromHex("0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f")
)

type UniswapV2Exchange struct {
	ctx        context.Context
	clientPool *clients.ClientPool
	opts       *ExchangeOptions
	router     *uniswapv2router.UniswapV2Router
	factory    *uniswapv2factory.UniswapV2Factory
}

func NewUniswapV2(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (*UniswapV2Exchange, error) {
	if clientsPool == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: clients pool is nil")
	}

	if opts == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: options are nil")
	}

	router, err := uniswapv2router.NewUniswapV2Router(opts.RouterAddress, clientsPool.GetClientByGroup(opts.Network.String()))
	if err != nil {
		return nil, fmt.Errorf("uniswapv2 exchange: failed to create router contract: %w", err)
	}

	factory, err := uniswapv2factory.NewUniswapV2Factory(opts.FactoryAddress, clientsPool.GetClientByGroup(opts.Network.String()))
	if err != nil {
		return nil, fmt.Errorf("uniswapv2 exchange: failed to create factory contract: %w", err)
	}

	return &UniswapV2Exchange{
		ctx:        ctx,
		clientPool: clientsPool,
		opts:       opts,
		router:     router,
		factory:    factory,
	}, nil
}

// ToUniswapV2 converts an Exchange to a UniswapV2Exchange. This is a helper function that you can use to
// access interface methods that are not part of the Exchange interface.
func ToUniswapV2(exchange Exchange) *UniswapV2Exchange {
	return exchange.(*UniswapV2Exchange)
}

func (u *UniswapV2Exchange) GetRouter() *uniswapv2router.UniswapV2Router {
	return u.router
}

func (u *UniswapV2Exchange) GetFactory() *uniswapv2factory.UniswapV2Factory {
	return u.factory
}

func (u *UniswapV2Exchange) GetRouterAddress() common.Address {
	return u.opts.RouterAddress
}

func (u *UniswapV2Exchange) GetFactoryAddress() common.Address {
	return u.opts.FactoryAddress
}

func (u *UniswapV2Exchange) GetNetwork() utils.Network {
	return u.opts.Network
}

func (u *UniswapV2Exchange) GetOptions() *ExchangeOptions {
	return u.opts
}

func (u *UniswapV2Exchange) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	return u.factory.GetPair(nil, tokenA, tokenB)
}

func (u *UniswapV2Exchange) GetPair(pairAddr common.Address) (*uniswapv2pool.UniswapV2Pool, error) {
	return uniswapv2pool.NewUniswapV2Pool(pairAddr, u.clientPool.GetClientByGroup(u.opts.Network.String()))
}

func (u *UniswapV2Exchange) GetPriceByAmount(ctx context.Context, amount *currencies.CurrencyAmount, baseCurrency *currencies.Currency, quoteCurrency *currencies.Currency, atBlock *big.Int) (*Price, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("uniswapv2 exchange: context is done")
	default:
		baseAddr := baseCurrency.AddressForNetwork(u.opts.Network)
		quoteAddr := quoteCurrency.AddressForNetwork(u.opts.Network)
		withLog := zap.L().With(
			zap.Uint64("wei", amount.Raw().Uint64()),
			zap.String("exchange", string(UniswapV2)),
			zap.String("base_currency", baseAddr.Hex()),
			zap.String("quote_currency", quoteAddr.Hex()),
		)

		withLog.Debug("Looking up for the token price")

		pairAddress, err := u.GetPairAddress(baseAddr, quoteAddr)
		if err != nil {
			return nil, err
		}

		if pairAddress == utils.ZeroAddress {
			return nil, fmt.Errorf("could not discover pair address for token pair [%s, %s]", baseAddr.Hex(), quoteAddr.Hex())
		}

		pair, err := u.GetPair(pairAddress)
		if err != nil {
			return nil, err
		}

		a, _ := pair.TotalSupply(&bind.CallOpts{BlockNumber: atBlock})
		withLog.Debug("Total supply", zap.Uint64("total_supply", a.Uint64()))
		withLog.Debug("Found pair address", zap.String("pair_address", pairAddress.Hex()))

		pairReserves, err := pair.GetReserves(&bind.CallOpts{BlockNumber: atBlock})
		if err != nil {
			return nil, err
		}

		baseToken, err := pair.Token0(&bind.CallOpts{BlockNumber: atBlock})
		if err != nil {
			return nil, err
		}

		quoteToken, err := pair.Token1(&bind.CallOpts{BlockNumber: atBlock})
		if err != nil {
			return nil, err
		}

		withLog.Debug(
			"Pair reserves information",
			zap.String("pair_address", pairAddress.Hex()),
			zap.Uint64("base_reserve", pairReserves.Reserve0.Uint64()),
			zap.Uint64("quote_reserve", pairReserves.Reserve1.Uint64()),
			zap.Uint32("latest_reserve_timestamp", pairReserves.BlockTimestampLast),
			zap.String("base_reserve_token", baseToken.Hex()),
			zap.String("quote_reserve_token", quoteToken.Hex()),
		)

		reserves := &Reserve{
			Token0:             baseToken,
			Reserve0:           pairReserves.Reserve0,
			Token1:             quoteToken,
			Reserve1:           pairReserves.Reserve1,
			BlockNumber:        atBlock,
			BlockTimestampLast: pairReserves.BlockTimestampLast,
		}

		am := big.NewInt(0).Mul(amount.Raw(), pairReserves.Reserve0)
		den := big.NewInt(0).Div(am, pairReserves.Reserve1)

		baseCurrencyAmount, _ := baseCurrency.ToAmount(den)
		withLog.Debug(
			"Base currency amount",
			zap.String("amount", baseCurrencyAmount.ToFixed(6)),
		)

		price, err := NewPriceFromReserve(reserves, baseCurrency, quoteCurrency)
		if err != nil {
			return nil, err
		}

		quote := big.NewInt(0).Mul(amount.Raw(), pairReserves.Reserve1)
		quote = big.NewInt(0).Div(quote, pairReserves.Reserve0)

		withLog.Info(
			"Calculated price",
			zap.String("base_currency", price.BaseCurrency.Symbol),
			zap.String("base_reserve", reserves.Reserve0.String()),
			zap.String("quote_currency", price.QuoteCurrency.Symbol),
			zap.String("quote_reserve", reserves.Reserve1.String()),
			zap.String("significant_6", price.Adjusted().ToSignificant(8)),
			zap.String("quote", quote.String()),
		)

		return price, nil
	}
}
