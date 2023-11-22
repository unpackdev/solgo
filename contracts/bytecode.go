package contracts

import (
	"fmt"
)

func (c *Contract) discoverDeployedBytecode() error {
	code, err := c.client.CodeAt(c.ctx, c.addr, nil)
	if err != nil {
		return fmt.Errorf("failed to get code at address %s: %s", c.addr.Hex(), err)
	}
	c.descriptor.DeployedBytecode = code

	return nil
}
