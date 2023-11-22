package liquidity

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/clients"
)

type Liquidity struct {
	ctx    context.Context
	client *clients.ClientPool
}

func NewLiquidity(ctx context.Context, client *clients.ClientPool) (*Liquidity, error) {
	return &Liquidity{
		ctx:    ctx,
		client: client,
	}, nil
}

func (l *Liquidity) GetPool(ltype LiquidityType, addr common.Address) error {
	return nil
}
