package contracts

import (
	"context"
	"math/big"
)

// DiscoverToken retrieves token-related metadata for the contract at a specific block height.
func (c *Contract) DiscoverToken(ctx context.Context, atBlock *big.Int) error {
	descriptor, err := c.token.Unpack(ctx, atBlock)
	if err != nil {
		return err
	}
	c.descriptor.Token = descriptor
	return nil
}
