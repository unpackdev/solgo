package number

import (
	"math/big"
)

var (
	// Percent100 percent 100
	Percent100 = NewFraction(big.NewInt(100), big.NewInt(1))
)

// Percent warps Fraction
type Percent struct {
	*Fraction
}

// NewPercent creates Percent
func NewPercent(num, deno *big.Int) *Percent {
	return &Percent{
		Fraction: NewFraction(num, deno),
	}
}

// ToSignificant format output
func (p *Percent) ToSignificant(significantDigits uint, opt ...Option) string {
	return p.Multiply(Percent100).ToSignificant(significantDigits, opt...)
}

// ToFixed format output
func (p *Percent) ToFixed(decimalPlaces uint, opt ...Option) string {
	return p.Multiply(Percent100).ToFixed(decimalPlaces, opt...)
}
