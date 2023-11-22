package exchanges

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type Exchange interface {
	GetNetwork() utils.Network
	GetRouterAddress() common.Address
	GetFactoryAddress() common.Address
	GetOptions() *ExchangeOptions
}
