package contracts

import (
	"context"
	"fmt"
	"github.com/unpackdev/solgo/bytecode"
)

// DiscoverMetadata attempts to extract metadata from the contract's deployed bytecode.
// This method decodes the bytecode to retrieve metadata information, such as the compiler version
// used for deployment and the content of the deployed code. It is critical in understanding
// the capabilities and design of the contract on the blockchain.
func (c *Contract) DiscoverMetadata(ctx context.Context) (*bytecode.Metadata, error) {
	if len(c.GetDeployedBytecode()) < 3 {
		return nil, fmt.Errorf("failed to discover metadata for contract %s due to invalid bytecode length", c.GetAddress())
	}

	select {
	case <-ctx.Done():
		return nil, nil
	default:
		bMetadata, err := bytecode.DecodeContractMetadata(c.GetDeployedBytecode())
		if err != nil {
			return nil, fmt.Errorf("failed to decode contract %s metadata: %s", c.GetAddress(), err)
		}
		c.descriptor.Metadata = bMetadata

		return bMetadata, nil
	}
}
