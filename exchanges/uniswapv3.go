package exchanges

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings/uniswapv3factory"
	"github.com/unpackdev/solgo/bindings/uniswapv3router"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

type UniswapV3Exchange struct {
	ctx        context.Context
	clientPool *clients.ClientPool
	opts       *ExchangeOptions
	router     *uniswapv3router.UniswapV3Router
	factory    *uniswapv3factory.UniswapV3Factory
}

func NewUniswapV3(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (*UniswapV3Exchange, error) {
	if clientsPool == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: clients pool is nil")
	}

	if opts == nil {
		return nil, fmt.Errorf("uniswapv2 exchange: options are nil")
	}

	router, err := uniswapv3router.NewUniswapV3Router(opts.RouterAddress, clientsPool.GetClientByGroup(opts.Network.String()))
	if err != nil {
		return nil, fmt.Errorf("uniswapv2 exchange: failed to create router contract: %w", err)
	}

	factory, err := uniswapv3factory.NewUniswapV3Factory(opts.FactoryAddress, clientsPool.GetClientByGroup(opts.Network.String()))
	if err != nil {
		return nil, fmt.Errorf("uniswapv2 exchange: failed to create factory contract: %w", err)
	}

	return &UniswapV3Exchange{
		ctx:        ctx,
		clientPool: clientsPool,
		opts:       opts,
		router:     router,
		factory:    factory,
	}, nil
}

// ToUniswapV3 converts an Exchange to a UniswapV3Exchange. This is a helper function that you can use to
// access interface methods that are not part of the Exchange interface.
func ToUniswapV3(exchange Exchange) *UniswapV3Exchange {
	return exchange.(*UniswapV3Exchange)
}

func (u *UniswapV3Exchange) GetRouter() *uniswapv3router.UniswapV3Router {
	return u.router
}

func (u *UniswapV3Exchange) GetFactory() *uniswapv3factory.UniswapV3Factory {
	return u.factory
}

func (u *UniswapV3Exchange) GetRouterAddress() common.Address {
	return u.opts.RouterAddress
}

func (u *UniswapV3Exchange) GetFactoryAddress() common.Address {
	return u.opts.FactoryAddress
}

func (u *UniswapV3Exchange) GetNetwork() utils.Network {
	return u.opts.Network
}

func (u *UniswapV3Exchange) GetOptions() *ExchangeOptions {
	return u.opts
}
