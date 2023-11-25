package storage

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/contracts"
)

func (s *Storage) DecodeContract(ctx context.Context, contract *contracts.Contract) (*contracts.Contract, error) {
	if contract := contracts.GetContract(s.network, contract.GetAddress()); contract != nil {
		return contract, nil
	}

	// This is critical error and we should not continue if we can't discover block, transaction and receipt information.
	if err := contract.DiscoverChainInfo(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover contract chain info: %s", err)
	}

	if err := contract.DiscoverSourceCode(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover contract source code: %s", err)
	}

	// Now we should attempt to parse contract's source code if we have it under our disposal.
	if err := contract.Parse(ctx); err != nil {
		return nil, fmt.Errorf("failed to parse contract: %s", err)
	}

	// Register contract in our registry for faster access in the future
	contracts.RegisterContract(s.network, contract)
	return contract, nil
}
