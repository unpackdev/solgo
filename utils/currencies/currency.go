package currencies

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type Currency struct {
	Name      string
	Symbol    string
	Decimals  uint8
	Addresses map[utils.Network]common.Address
}

func (c *Currency) String() string {
	return c.Symbol
}

func (c *Currency) NetworkExists(network utils.Network) bool {
	_, ok := c.Addresses[network]
	return ok
}

func (c *Currency) AddressForNetwork(network utils.Network) common.Address {
	return c.Addresses[network]
}

func (c *Currency) ToAmount(amount *big.Int) (*CurrencyAmount, error) {
	return NewCurrencyAmount(c, amount)
}

// Equals identifies whether A and B are equal
func (c *Currency) Equals(other *Currency) bool {
	return c == other ||
		(c.Decimals == other.Decimals && c.Symbol == other.Symbol && c.Name == other.Name)
}

func (c *Currency) AddressExists(addr common.Address) bool {
	for _, v := range c.Addresses {
		if v == addr {
			return true
		}
	}
	return false
}

var (
	WETH = &Currency{
		Name:     "Wrapped Ether",
		Symbol:   "WETH",
		Decimals: 18,
		Addresses: map[utils.Network]common.Address{
			utils.Ethereum: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		},
	}
	USDC = &Currency{
		Name:     "USD Coin",
		Symbol:   "USDC",
		Decimals: 6,
		Addresses: map[utils.Network]common.Address{
			utils.Ethereum: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		},
	}
)
