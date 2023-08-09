// Package eip provides structures and functions related to Ethereum Improvement Proposals (EIPs).
package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Erc20 represents the ERC-20 token standard.
type Erc20 struct {
	// Standard holds the details of the ERC-20 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the ERC-20 standard.
func (e Erc20) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the ERC-20 standard.
func (e Erc20) GetType() Standard {
	return e.Standard.Type
}

// GetFunctions returns the functions associated with the ERC-20 standard.
func (e Erc20) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the ERC-20 standard.
func (e Erc20) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete ERC-20 contract standard.
func (e Erc20) GetStandard() ContractStandard {
	return e.Standard
}

// TokenCount returns the number of tokens associated with the ERC-20 standard.
func (e Erc20) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the ERC-20 standard.
func (e Erc20) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the ERC-20 standard.
func (e Erc20) String() string {
	return e.GetName()
}

// NewErc20 initializes and returns an instance of the ERC-20 standard.
// It sets up the standard with its name, type, associated functions, and events.
func NewErc20() EIP {
	return &Erc20{
		Standard: ContractStandard{
			Name: "ERC-20 Token Standard",
			Type: ERC20,
			Functions: []Function{
				newFunction("totalSupply", nil, []string{TypeUint256}),
				newFunction("balanceOf", []Input{{Type: TypeAddress}}, []string{TypeUint256}),
				newFunction("transfer", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []string{TypeBool}),
				newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, []string{TypeBool}),
				newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []string{TypeBool}),
				newFunction("allowance", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []string{TypeUint256}),
			},
			Events: []Event{
				newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
			},
		},
	}
}
