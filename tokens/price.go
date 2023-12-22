package tokens

import (
	"context"

	"github.com/unpackdev/solgo/utils/entities"
)

func (t *Token) DiscoverPrice(ctx context.Context) error {
	wethToken := entities.WETH9[uint(t.networkID.Uint64())]
	usdtToken := entities.USDT[uint(t.networkID.Uint64())]
	_, _ = wethToken, usdtToken

	return nil
}

func (t *Token) GetPrice(ctx context.Context) (*entities.Price, error) {
	if t.descriptor.Price == nil {
		if err := t.DiscoverPrice(ctx); err != nil {
			return nil, err
		}

		return t.descriptor.Price, nil
	}

	return t.descriptor.Price, nil
}
