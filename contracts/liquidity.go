package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings/uniswapv2factory"
	"github.com/unpackdev/solgo/bindings/uniswapv2router"
	"github.com/unpackdev/solgo/exchanges"
)

func (c *Contract) DiscoverLiquidity(ctx context.Context) error {
	routerAddr := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") // Uniswap V2 Router

	router, err := uniswapv2router.NewUniswapV2Router(routerAddr, c.client.Client)
	if err != nil {
		return fmt.Errorf("failed to create uniswap router: %w", err)
	}

	weth, err := router.WETH(nil)
	if err != nil {
		return fmt.Errorf("failed to get WETH address: %w", err)
	}

	factory, err := uniswapv2factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), c.client.Client)
	if err != nil {
		return fmt.Errorf("failed to create uniswap factory: %w", err)
	}

	pair, err := factory.GetPair(nil, weth, c.addr)
	if err != nil {
		return fmt.Errorf("failed to get pair: %w", err)
	}

	c.descriptor.LiquidityPairs[exchanges.UniswapV2] = pair

	//fmt.Println("Pair: ", pair.Hex())

	return nil
}
