package contracts

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/bytecode"
)

func (c *Contract) DiscoverMetadata(ctx context.Context) (*bytecode.Metadata, error) {
	if len(c.GetDeployedBytecode()) < 3 {
		return nil, fmt.Errorf("failed to discover metadata for contract %s due to invalid bytecode length", c.GetAddress())
	}

	bmetadata, err := bytecode.DecodeContractMetadata(c.GetDeployedBytecode())
	if err != nil {
		return nil, fmt.Errorf("failed to decode contract %s metadata: %s", c.GetAddress(), err)
	}
	c.descriptor.Metadata = bmetadata

	/* 	if len(bmetadata.GetIPFS()) > 10 {
		if len(bmetadata.GetCompilerVersion()) > 0 {
			utils.DumpNodeNoExit(bmetadata.GetCompilerVersion())
			c.descriptor.SetCompilerVersion(bmetadata.GetCompilerVersion())
		}

		cmetadata, err := c.ipfsProvider.GetMetadataByCID(bmetadata.GetIPFS())
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to get contract %s metadata from IPFS: %s", c.GetAddress(), err)
		}

		utils.DumpNodeNoExit(cmetadata)
		utils.DumpNodeWithExit(bmetadata)
	} */

	return bmetadata, nil
}
