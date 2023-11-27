package storage

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/contracts"
)

// DecodeContract retrieves and decodes contract information from the blockchain.
// It takes a contract reference and a context, then performs several operations
// to enrich the contract with additional data such as chain info, source code, and parsing.
//
// The function attempts to fetch the contract from a local registry based on its network and address.
// If the contract is not found in the registry, it proceeds to discover chain information
// (like block, transaction, and receipt details) and source code associated with the contract.
// The contract's source code is then parsed for further processing.
//
// If any step of discovering chain info, source code, or parsing fails, the function
// returns an error, detailing the failure. Successful operations result in updating
// the contract object with the discovered data.
//
// Finally, the function registers the contract in the registry for future quick access
// and returns the updated contract.
//
// Parameters:
// ctx      - The context used for network calls and potentially for cancellation.
// contract - The contract object that needs to be decoded. It should have at least the contract address set.
//
// Returns:
// A pointer to the updated contract object and nil error on success, or nil and an error on failure.
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
