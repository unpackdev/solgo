package etherscan

type ProviderType string

func (p ProviderType) String() string {
	return string(p)
}

const (
	EtherScan ProviderType = "etherscan"
	BscScan   ProviderType = "bscscan"
)
