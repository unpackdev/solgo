package utils

import (
	"fmt"
	"math/big"
)

var Ether = big.NewInt(1e18)
var GWei = big.NewInt(1e9)

// FromWei converts a balance in wei to Ether.
func FromWei(wei *big.Int, unit *big.Int) *big.Float {
	if wei == nil {
		return big.NewFloat(0)
	}

	// Convert wei to a big.Float for division.
	weiFloat := new(big.Float).SetInt(wei)

	eUnit := Ether

	if unit == nil {
		eUnit = Ether
	}

	// Divide by 1e18 to convert wei to Ether.
	ether := new(big.Float).Quo(weiFloat, new(big.Float).SetInt(eUnit))

	return ether
}

// ToWei converts an Ether value (as a decimal) to Wei.
func ToWei(etherValueStr string, unit *big.Int) (*big.Int, error) {
	etherValue, ok := new(big.Float).SetString(etherValueStr)
	if !ok {
		return nil, fmt.Errorf("invalid ether value: %s", etherValueStr)
	}

	eUnit := Ether

	if unit == nil {
		eUnit = Ether
	}

	// Multiply the Ether value by 1e18 to convert it to Wei.
	wei := new(big.Float).Mul(etherValue, new(big.Float).SetInt(eUnit))

	// Convert the result to *big.Int.
	result := new(big.Int)
	wei.Int(result) // Note: This truncates the fractional part.

	return result, nil
}
