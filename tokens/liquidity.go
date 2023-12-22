package tokens

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

func (t *Token) PrepareBindings(ctx context.Context) error {

	return nil
}

func (t *Token) DiscoverLiquidityPairs(ctx context.Context) error {
	baseToken := entities.WETH9[uint(t.GetNetworkID().Uint64())]
	currentToken := t.GetEntity()

	// Ethereum based discovery pair process...
	if t.network == utils.Ethereum {
		if !t.descriptor.HasPairByExchange(utils.UniswapV2) {
			uniswapBinding, err := t.GetUniswapV2Bind(ctx, utils.NoSimulator, t.bindManager, nil)
			if err != nil {
				return fmt.Errorf(
					"failed to get uniswap bindings: %s, exchange: %s, base_addr: %s, quote_addr: %s",
					err,
					utils.UniswapV2, baseToken.Address, currentToken.Address,
				)
			}

			if uniswapPair, err := uniswapBinding.GetPair(ctx, baseToken.Address, t.descriptor.Address); err == nil {
				if uniswapPair != utils.ZeroAddress {
					t.descriptor.Pairs[utils.UniswapV2] = &Pair{
						BaseToken:   baseToken,
						QuoteToken:  currentToken,
						PairAddress: uniswapPair,
					}
				}
			}
		}

		if !t.descriptor.HasPairByExchange(utils.UniswapV3) {

		}
	}

	return nil
}
