package exchanges

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Reserve struct {
	Token0             common.Address
	Reserve0           *big.Int
	Token1             common.Address
	Reserve1           *big.Int
	BlockNumber        *big.Int
	BlockTimestampLast uint32
}
