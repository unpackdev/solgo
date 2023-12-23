package exchanges

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

// exchangeFn is a function type that returns an Exchange instance and an error.
type exchangeFn func(ctx context.Context, clientsPool *clients.ClientPool, bindManager *bindings.Manager, opts *ExchangeOptions) (Exchange, error)

// exchanges is a map of exchange functions, keyed by a string.
var exchanges = make(map[utils.ExchangeType]exchangeFn)

// RegisterExchange stores the exchange function in the exchanges map.
func registerExchange(name utils.ExchangeType, exchange exchangeFn) error {
	if _, ok := exchanges[name]; ok {
		return fmt.Errorf("exchange %s already registered", name)
	}

	exchanges[name] = exchange
	return nil
}

// GetExchange retrieves an exchange function from the exchanges map.
func GetExchange(name utils.ExchangeType) (exchangeFn, bool) {
	if exchange, ok := exchanges[name]; ok {
		return exchange, true
	}

	return nil, false
}

// GetExchanges retrieves the exchanges map.
func GetExchanges() map[utils.ExchangeType]exchangeFn {
	return exchanges
}

/* func init() {
	registerExchange(utils.UniswapV2, func(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (Exchange, error) {
		uniswapBind.  := bindings.NewManager(ctx, clientsPool)
		return NewUniswapV2(ctx, clientsPool, opts)
	})
} */
