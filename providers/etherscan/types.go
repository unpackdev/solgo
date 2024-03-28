package etherscan

// ProviderType defines a type for different blockchain explorer services supported by the etherscan package.
type ProviderType string

// String returns the string representation of the ProviderType.
func (p ProviderType) String() string {
	return string(p)
}

const (
	// EtherScan represents the Etherscan service for the Ethereum blockchain.
	EtherScan ProviderType = "etherscan"

	// BscScan represents the BSCScan service for the Binance Smart Chain.
	BscScan ProviderType = "bscscan"
)
