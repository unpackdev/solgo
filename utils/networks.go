package utils

type Network string

const (
	Ethereum Network = "ethereum"
	Bsc      Network = "bsc"
	Polygon  Network = "polygon"
)

func (n Network) String() string {
	return string(n)
}
