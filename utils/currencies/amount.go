package currencies

import (
	"math/big"

	"github.com/unpackdev/solgo/utils/number"
)

// CurrencyAmount warps Fraction and Currency
type CurrencyAmount struct {
	*number.Fraction
	*Currency
}

// NewCurrencyAmount creates a CurrencyAmount
// amount _must_ be raw, i.e. in the native representation
func NewCurrencyAmount(currency *Currency, amount *big.Int) (*CurrencyAmount, error) {
	if err := ValidateSolidityTypeInstance(amount, Uint256); err != nil {
		return nil, err
	}

	fraction := number.NewFraction(amount, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(currency.Decimals)), nil))
	return &CurrencyAmount{
		Fraction: fraction,
		Currency: currency,
	}, nil
}

// Raw returns Fraction's Numerator
func (c *CurrencyAmount) Raw() *big.Int {
	return c.Numerator
}

// NewEther Helper that calls the constructor with the ETHER currency
// @param amount ether amount in wei
func NewEther(amount *big.Int) (*CurrencyAmount, error) {
	return NewCurrencyAmount(WETH, amount)
}
