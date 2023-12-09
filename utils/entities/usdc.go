package entities

import "github.com/ethereum/go-ethereum/common"

// Known USDC implementation addresses
var USDC = map[uint]*Token{
	1: NewToken(1, common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), 6, "USDC", "USDC Coin"),
}

// Known USDC implementation addresses
var USDT = map[uint]*Token{
	1: NewToken(1, common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), 6, "USDT", "USDT Coin"),
}
