package exchanges

import (
	"context"

	"github.com/unpackdev/solgo/clients"
)

type PancakeswapV2Exchange struct {
}

func NewPancakeswapV2(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (*UniswapV2Exchange, error) {
	return &UniswapV2Exchange{}, nil
}
