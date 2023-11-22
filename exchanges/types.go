package exchanges

import "math/big"

type ExchangeType string

const (
	UniswapV2     ExchangeType = "UniswapV2"
	UniswapV3     ExchangeType = "UniswapV3"
	Sushiswap     ExchangeType = "Sushiswap"
	PancakeswapV2 ExchangeType = "PancakeswapV2"
	PancakeswapV3 ExchangeType = "PancakeswapV3"
)

var (
	MinimumLiquidity = big.NewInt(1000)

	Zero  = big.NewInt(0)
	One   = big.NewInt(1)
	Two   = big.NewInt(2)
	Three = big.NewInt(3)
	Five  = big.NewInt(5)
	Ten   = big.NewInt(10)

	B100  = big.NewInt(100)
	B997  = big.NewInt(997)
	B1000 = big.NewInt(1000)
)
