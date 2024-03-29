package contracts

import (
	"fmt"
)

// DiscoverDeployedBytecode retrieves the deployed bytecode of the contract deployed at the specified address.
// It queries the blockchain using the provided client to fetch the bytecode associated with the contract address.
// The fetched bytecode is then stored in the contract descriptor for further processing.
func (c *Contract) DiscoverDeployedBytecode() error {
	code, err := c.client.CodeAt(c.ctx, c.addr, nil)
	if err != nil {
		return fmt.Errorf("failed to get code at address %s: %s", c.addr.Hex(), err)
	}
	c.descriptor.DeployedBytecode = code

	return nil
}
