package exchanges

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/clients"
)

// exchangeFn is a function type that returns an Exchange instance and an error.
type exchangeFn func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error)

// exchanges is a map of exchange functions, keyed by a string.
var exchanges map[ExchangeType]exchangeFn

// RegisterExchange stores the exchange function in the exchanges map.
func registerExchange(name ExchangeType, exchange exchangeFn) error {
	if _, ok := exchanges[name]; ok {
		return fmt.Errorf("exchange %s already registered", name)
	}

	exchanges[name] = exchange
	return nil
}

// GetExchange retrieves an exchange function from the exchanges map.
func GetExchange(name ExchangeType) (exchangeFn, bool) {
	if exchange, ok := exchanges[name]; ok {
		return exchange, true
	}

	return nil, false
}

// GetExchanges retrieves the exchanges map.
func GetExchanges() map[ExchangeType]exchangeFn {
	return exchanges
}

func init() {
	exchanges = make(map[ExchangeType]exchangeFn)

	registerExchange(UniswapV2, func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error) {
		return NewUniswapV2(ctx, clientsPool, opts)
	})

	registerExchange(UniswapV3, func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error) {
		return NewUniswapV3(ctx, clientsPool, opts)
	})

	/*
		 	RegisterExchange(Sushiswap, func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error) {
				return NewSushiswap(ctx, clientsPool, opts)
			})

			RegisterExchange(PancakeswapV2, func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error) {
				return NewPancakeswapV2(ctx, clientsPool, opts)
			})
	*/
}
