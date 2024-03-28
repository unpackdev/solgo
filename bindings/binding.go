package bindings

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type Binding struct {
	network   utils.Network
	networkID utils.NetworkID
	Type      BindingType
	Address   common.Address
	RawABI    string
	ABI       *abi.ABI
}

func (b *Binding) GetNetwork() utils.Network {
	return b.network
}

func (b *Binding) GetNetworkID() utils.NetworkID {
	return b.networkID
}

func (b *Binding) GetType() BindingType {
	return b.Type
}

func (b *Binding) GetAddress() common.Address {
	return b.Address
}

func (b *Binding) GetRawABI() string {
	return b.RawABI
}

func (b *Binding) GetABI() *abi.ABI {
	return b.ABI
}

// GetAllMethods returns all the method names in the contract ABI.
func (b *Binding) GetAllMethods() []string {
	var methods []string
	for name := range b.ABI.Methods {
		methods = append(methods, name)
	}
	return methods
}

// MethodExist checks if a method exists in the contract ABI.
func (b *Binding) MethodExist(methodName string) bool {
	_, exists := b.ABI.Methods[methodName]
	return exists
}

// GetAllEvents returns all the event names in the contract ABI.
func (b *Binding) GetAllEvents() []string {
	var events []string
	for name := range b.ABI.Events {
		events = append(events, name)
	}
	return events
}

// EventExist checks if an event exists in the contract ABI.
func (b *Binding) EventExist(eventName string) bool {
	_, exists := b.ABI.Events[eventName]
	return exists
}
