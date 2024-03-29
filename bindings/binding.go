package bindings

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// Binding represents a specific contract binding, including its network, type, address, and ABI. It serves as
// a central point for accessing contract-related data and functionalities, such as querying available methods
// and events defined in the contract ABI.
type Binding struct {
	network   utils.Network   // The network where the contract is deployed.
	networkID utils.NetworkID // The unique identifier of the network.
	Type      BindingType     // The type of the contract binding (e.g., ERC20, ERC721).
	Address   common.Address  // The blockchain address of the contract.
	RawABI    string          // The raw string representation of the contract's ABI.
	ABI       *abi.ABI        // The parsed contract ABI, providing programmatic access to its contents.
}

// GetNetwork returns the network where the contract is deployed.
func (b *Binding) GetNetwork() utils.Network {
	return b.network
}

// GetNetworkID returns the unique identifier of the network.
func (b *Binding) GetNetworkID() utils.NetworkID {
	return b.networkID
}

// GetType returns the type of the contract binding.
func (b *Binding) GetType() BindingType {
	return b.Type
}

// GetAddress returns the blockchain address of the contract.
func (b *Binding) GetAddress() common.Address {
	return b.Address
}

// GetRawABI returns the raw string representation of the contract's ABI.
func (b *Binding) GetRawABI() string {
	return b.RawABI
}

// GetABI returns the parsed contract ABI.
func (b *Binding) GetABI() *abi.ABI {
	return b.ABI
}

// GetAllMethods returns all the method names defined in the contract ABI.
func (b *Binding) GetAllMethods() []string {
	var methods []string
	for name := range b.ABI.Methods {
		methods = append(methods, name)
	}
	return methods
}

// MethodExist checks if a specified method exists in the contract ABI.
func (b *Binding) MethodExist(methodName string) bool {
	_, exists := b.ABI.Methods[methodName]
	return exists
}

// GetAllEvents returns all the event names defined in the contract ABI.
func (b *Binding) GetAllEvents() []string {
	var events []string
	for name := range b.ABI.Events {
		events = append(events, name)
	}
	return events
}

// EventExist checks if a specified event exists in the contract ABI.
func (b *Binding) EventExist(eventName string) bool {
	_, exists := b.ABI.Events[eventName]
	return exists
}
