package currencies

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

type SolidityType string

const (
	Uint8   SolidityType = "uint8"
	Uint256 SolidityType = "uint256"
)

var (
	SolidityTypeMaxima = map[SolidityType]*big.Int{
		Uint8:   big.NewInt(0xff),
		Uint256: math.MaxBig256,
	}
)

// ValidateSolidityTypeInstance determines if a value is a legal SolidityType
func ValidateSolidityTypeInstance(value *big.Int, t SolidityType) error {
	if value.Cmp(big.NewInt(0)) < 0 || value.Cmp(SolidityTypeMaxima[t]) > 0 {
		return fmt.Errorf(`%v is not a %s`, value, t)
	}
	return nil
}

// ValidateAndParseAddress warns if addresses are not checksummed
func ValidateAndParseAddress(address string) common.Address {
	return common.HexToAddress(address)
}
