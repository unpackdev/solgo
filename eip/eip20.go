// Package eip provides structures and functions related to Ethereum Improvement Proposals (EIPs).
package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Eip20 represents the ERC-20 token standard.
type Eip20 struct {
	// Standard holds the details of the ERC-20 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the ERC-20 standard.
func (e Eip20) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the ERC-20 standard.
func (e Eip20) GetType() Standard {
	return e.Standard.Type
}

// GetUrl returns the URL of the ERC-20 standard.
func (e Eip20) GetUrl() string {
	return e.Standard.Url
}

// GetFunctions returns the functions associated with the ERC-20 standard.
func (e Eip20) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the ERC-20 standard.
func (e Eip20) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete ERC-20 contract standard.
func (e Eip20) GetStandard() ContractStandard {
	return e.Standard
}

// IsStagnant returns a boolean indicating whether the ERC-20 standard is stagnant.
func (e Eip20) IsStagnant() bool {
	return e.Standard.Stagnant
}

func (e Eip20) ConfidenceCheck(contract *Contract) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// TokenCount returns the number of tokens associated with the ERC-20 standard.
func (e Eip20) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the ERC-20 standard.
func (e Eip20) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the ERC-20 standard.
func (e Eip20) String() string {
	return e.GetName()
}

// NewEip20 initializes and returns an instance of the ERC-20 standard.
// It sets up the standard with its name, type, associated functions, and events.
func NewEip20() EIP {
	return &Eip20{
		Standard: ContractStandard{
			Name: "ERC-20 Token Standard",
			Url:  "https://eips.ethereum.org/EIPS/eip-20",
			Type: EIP20,
			Functions: []Function{
				newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
				newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
				newFunction("transfer", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
				newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
				newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeBool}}),
				newFunction("allowance", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
			},
			Events: []Event{
				newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
			},
		},
	}
}
