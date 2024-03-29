package bindings

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// BindOptions defines the configuration parameters required to establish a binding to a smart contract.
// It includes network-specific settings, contract metadata, and the contract's ABI (Application Binary Interface).
// These options are used to initialize and manage contract bindings, allowing for interactions with the contract
// across different blockchain networks.
type BindOptions struct {
	Networks  []utils.Network // A list of networks on which the contract is deployed.
	NetworkID utils.NetworkID // The unique identifier for the blockchain network.
	Name      string          // The name of the binding, for identification purposes.
	Type      BindingType     // The type of binding, indicating the contract standard (e.g., ERC20, ERC721).
	Address   common.Address  // The blockchain address of the contract.
	ABI       string          // The JSON string representing the contract's Application Binary Interface.
}

// Validate checks the integrity and completeness of the BindOptions. It ensures that all necessary information
// is provided and valid, such as the presence of at least one network, a valid network ID, a specified binding
// type, a non-zero contract address, and a non-empty ABI definition. This validation step is crucial before
// establishing a binding to a contract to prevent runtime errors and ensure reliable contract interaction.
func (b *BindOptions) Validate() error {
	if len(b.Networks) == 0 {
		return fmt.Errorf("missing network")
	}
	if b.NetworkID == 0 {
		return fmt.Errorf("missing network id")
	}
	if b.Type == "" {
		return fmt.Errorf("missing binding type")
	}
	if b.Address == utils.ZeroAddress {
		return fmt.Errorf("missing address")
	}
	if b.ABI == "" {
		return fmt.Errorf("missing abi")
	}
	return nil
}
