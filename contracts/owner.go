package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// DiscoverOwner queries the blockchain to find the current owner of the contract and updates the
// contract's descriptor with the owner's address.
func (c *Contract) DiscoverOwner(ctx context.Context) error {
	if c.token == nil {
		return fmt.Errorf(
			"failed to discover owner of contract %s: token is not set",
			c.GetAddress().Hex(),
		)
	}

	if c.tokenBind == nil {
		return fmt.Errorf(
			"failed to discover owner of contract %s: token bind is not set",
			c.GetAddress().Hex(),
		)
	}

	owner, err := c.tokenBind.Owner(ctx, c.GetAddress())
	if err != nil {
		return fmt.Errorf(
			"failed to discover owner of contract %s: %w",
			c.GetAddress().Hex(),
			err,
		)
	}
	c.descriptor.Owner = owner

	return nil
}

// IsRenounced checks if the contract's ownership has been renounced by comparing the owner's address
// against a predefined list of "zero" addresses, which represent renounced ownership in various contexts.
// This method helps in identifying contracts that are deliberately made ownerless, usually as a measure
// of decentralization or to prevent further administrative intervention.
func (c *Contract) IsRenounced() bool {
	zeroAddresses := []common.Address{
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		common.HexToAddress("0x000000000000000000000000000000000000dead"),
		common.HexToAddress("0x000000000000000000000000000000000000dEaD"),
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		common.HexToAddress("0x0000000000000000000000000000000000000003"),
		common.HexToAddress("0x0000000000000000000000000000000000000004"),
		common.HexToAddress("0x0000000000000000000000000000000000000005"),
		common.HexToAddress("0x0000000000000000000000000000000000000006"),
		common.HexToAddress("0x0000000000000000000000000000000000000007"),
		common.HexToAddress("0x0000000000000000000000000000000000000008"),
		common.HexToAddress("0x0000000000000000000000000000000000000009"),
	}

	for _, address := range zeroAddresses {
		if c.descriptor.Owner == address {
			return true
		}
	}

	return false
}
