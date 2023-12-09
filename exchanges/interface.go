package exchanges

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// Exchange is an interface defining methods to interact with cryptocurrency exchanges.
// It allows fetching network details, router and factory addresses, and specific exchange options.
type Exchange interface {
	// GetRouterAddress returns the Ethereum address of the router contract for the exchange.
	// This address is used to interact with the router contract, which handles token swaps and other operations.
	GetRouterAddress() common.Address

	// GetFactoryAddress returns the Ethereum address of the factory contract for the exchange.
	// The factory contract is responsible for creating new liquidity pools and other administrative tasks.
	GetFactoryAddress() common.Address

	// GetOptions returns a pointer to ExchangeOptions which includes optional settings and configurations
	// for interacting with the exchange. This could include fees, slippage tolerance, and other parameters.
	GetOptions() *ExchangeOptions

	// GetType returns the type of exchange.
	// This is used to determine which exchange to use when interacting with the blockchain.
	GetType() utils.ExchangeType
}
