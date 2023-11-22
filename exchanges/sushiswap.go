package exchanges

import (
	"context"

	"github.com/unpackdev/solgo/clients"
)

type SushiswapExchange struct {
}

func NewSushiswap(ctx context.Context, clientsPool *clients.ClientPool, opts *ExchangeOptions) (*SushiswapExchange, error) {
	return &SushiswapExchange{}, nil
}
