package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Contract) DiscoverOwner(ctx context.Context) error {
	if c.token == nil || c.tokenBind == nil {
		return fmt.Errorf(
			"failed to discover owner of contract %s: token is not set. Forgot to call discover token?",
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

func (c *Contract) IsRenounounced() bool {
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
