package tokens

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils/entities"
)

// Descriptor contains detailed information about a specific Ethereum token.
// It includes metadata such as the token's address, name, symbol, decimals,
// total supply, total burned supply, and the latest block number for context.
type Descriptor struct {
	BlockNumber       *big.Int        `json:"block_number"`        // Block number at which the token info was queried.
	Address           common.Address  `json:"address"`             // Ethereum address of the token contract.
	Name              string          `json:"name"`                // Name of the token.
	Symbol            string          `json:"symbol"`              // Symbol of the token.
	Decimals          uint8           `json:"decimals"`            // Decimal precision of the token.
	TotalSupply       *big.Int        `json:"total_supply"`        // Total supply of the token.
	TotalBurnedSupply *big.Int        `json:"total_burned_supply"` // Supply of tokens that have been burned.
	LatestBlockNumber *big.Int        `json:"latest_block_number"` // Latest block number processed by the application.
	Entity            *entities.Token `json:"-"`                   // Associated token entity, not serialized to JSON.
}

// GetAddress returns the Ethereum address of the token contract.
func (d *Descriptor) GetAddress() common.Address {
	return d.Address
}

// GetName returns the name of the token.
func (d *Descriptor) GetName() string {
	return d.Name
}

// GetSymbol returns the symbol of the token.
func (d *Descriptor) GetSymbol() string {
	return d.Symbol
}

// GetDecimals returns the decimal precision of the token.
func (d *Descriptor) GetDecimals() uint8 {
	return d.Decimals
}

// GetTotalSupply returns the total supply of the token.
func (d *Descriptor) GetTotalSupply() *big.Int {
	return d.TotalSupply
}

// GetTotalBurnedSupply returns the total supply of tokens that have been burned.
func (d *Descriptor) GetTotalBurnedSupply() *big.Int {
	return d.TotalBurnedSupply
}

// GetEntity returns the associated token entity.
func (d *Descriptor) GetEntity() *entities.Token {
	return d.Entity
}

// GetTotalCirculatingSupply calculates and returns the total circulating supply of the token,
// which is the total supply minus the total burned supply.
func (d *Descriptor) GetTotalCirculatingSupply() *big.Int {
	if d.TotalBurnedSupply == nil {
		return d.TotalSupply
	}

	return new(big.Int).Sub(d.TotalSupply, d.TotalBurnedSupply)
}
