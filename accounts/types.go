package accounts

// Network defines a type for representing various Ethereum-compatible networks.
// It is used to specify and differentiate between different blockchain networks.
type Network string

const (
	// Ethereum represents the Ethereum mainnet.
	Ethereum Network = "ethereum"

	// Bsc represents the Binance Smart Chain network.
	Bsc Network = "bsc"

	// Polygon represents the Polygon (formerly Matic Network) network.
	Polygon Network = "polygon"
)
