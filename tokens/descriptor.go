package tokens

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils/entities"
)

type Descriptor struct {
	BlockNumber       *big.Int        `json:"block_number"`
	Address           common.Address  `json:"address"`
	Name              string          `json:"name"`
	Symbol            string          `json:"symbol"`
	Decimals          uint8           `json:"decimals"`
	TotalSupply       *big.Int        `json:"total_supply"`
	TotalBurnedSupply *big.Int        `json:"total_burned_supply"`
	LatestBlockNumber *big.Int        `json:"latest_block_number"`
	Entity            *entities.Token `json:"-"`
}

func (d *Descriptor) GetAddress() common.Address {
	return d.Address
}

func (d *Descriptor) GetName() string {
	return d.Name
}

func (d *Descriptor) GetSymbol() string {
	return d.Symbol
}

func (d *Descriptor) GetDecimals() uint8 {
	return d.Decimals
}

func (d *Descriptor) GetTotalSupply() *big.Int {
	return d.TotalSupply
}

func (d *Descriptor) GetTotalBurnedSupply() *big.Int {
	return d.TotalBurnedSupply
}

func (d *Descriptor) GetEntity() *entities.Token {
	return d.Entity
}

func (d *Descriptor) GetTotalCirculatingSupply() *big.Int {
	if d.TotalBurnedSupply == nil {
		return d.TotalSupply
	}

	return new(big.Int).Sub(d.TotalSupply, d.TotalBurnedSupply)
}
