package contracts

import (
	"context"
	"math/big"

	"github.com/unpackdev/solgo/utils"
)

func (c *Contract) DiscoverToken(ctx context.Context, atBlock *big.Int, simulate bool, simulatorType utils.SimulatorType) error {
	descriptor, err := c.token.Unpack(ctx, atBlock, simulate, simulatorType)
	if err != nil {
		return err
	}
	c.descriptor.Token = descriptor
	return nil
}
