package exchanges

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/unpackdev/solgo/utils/currencies"
	"github.com/unpackdev/solgo/utils/number"
)

type Price struct {
	*number.Fraction
	BaseCurrency  *currencies.Currency // input i.e. denominator
	QuoteCurrency *currencies.Currency // output i.e. numerator
	Scalar        *number.Fraction     // used to adjust the raw fraction w/r/t the decimals of the {base,quote}Token
	BlockNumber   *big.Int
	Timestamp     time.Time
}

func NewPriceFromReserve(reserves *Reserve, baseCurrency *currencies.Currency, quoteCurrency *currencies.Currency) (*Price, error) {
	if reserves == nil || baseCurrency == nil || quoteCurrency == nil {
		return nil, fmt.Errorf("invalid input: nil Reserve or Currency provided")
	}

	if baseCurrency.AddressExists(reserves.Token0) {
		return NewPrice(
			baseCurrency,
			quoteCurrency,
			reserves.Reserve0,
			reserves.Reserve1,
			reserves.BlockNumber,
			reserves.BlockTimestampLast,
		), nil
	}

	// Base currency matches Reserve1, so invert the price
	return NewPrice(
		quoteCurrency,
		baseCurrency,
		reserves.Reserve1,
		reserves.Reserve0,
		reserves.BlockNumber,
		reserves.BlockTimestampLast,
	), nil
}

func NewPrice(baseCurrency, quoteCurrency *currencies.Currency, denominator, numerator *big.Int, blockNumber *big.Int, ts uint32) *Price {
	return &Price{
		Fraction:      number.NewFraction(denominator, numerator),
		BaseCurrency:  baseCurrency,
		QuoteCurrency: quoteCurrency,
		Scalar: number.NewFraction(
			math.BigPow(10, int64(baseCurrency.Decimals)),
			math.BigPow(10, int64(quoteCurrency.Decimals)),
		),
		BlockNumber: blockNumber,
		Timestamp:   time.Unix(int64(ts), 0),
	}
}

func (p *Price) Raw() *number.Fraction {
	return p.Fraction
}

func (p *Price) Adjusted() *number.Fraction {
	p.Fraction.Multiply(p.Scalar)
	return p.Fraction
}

func (p *Price) Invert() {
	p.BaseCurrency, p.QuoteCurrency = p.QuoteCurrency, p.BaseCurrency
}

func (p *Price) Multiply(other *Price) error {
	if !p.QuoteCurrency.Equals(other.BaseCurrency) {
		return fmt.Errorf("invalid currencies for price multiplication (%s, %s)", p.QuoteCurrency.Symbol, other.BaseCurrency.Symbol)
	}

	p.Fraction.Multiply(other.Fraction)
	p.QuoteCurrency = other.QuoteCurrency
	return nil
}

func (p *Price) Quote(currencyAmount *currencies.CurrencyAmount) (*currencies.CurrencyAmount, error) {
	if !p.BaseCurrency.Equals(currencyAmount.Currency) {
		return nil, fmt.Errorf("invalid currency for price quote (%s)", currencyAmount.Currency.Symbol)
	}

	return currencies.NewEther(p.Fraction.Multiply(number.NewFraction(currencyAmount.Raw(), nil)).Quotient())
}

func (p *Price) ToSignificant(significantDigits uint, opt ...number.Option) string {
	return p.Adjusted().ToSignificant(significantDigits, opt...)
}

func (p *Price) ToFixed(decimalPlaces uint, opt ...number.Option) string {
	return p.Adjusted().ToFixed(decimalPlaces, opt...)
}
