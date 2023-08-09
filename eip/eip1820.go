package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Eip1820 represents the EIP-1820 Pseudo-introspection Registry Contract.
type Eip1820 struct {
	// Standard holds the details of the EIP-1820 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the EIP-1820 standard.
func (e Eip1820) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the EIP-1820 standard.
func (e Eip1820) GetType() Standard {
	return e.Standard.Type
}

// GetFunctions returns the functions associated with the EIP-1820 standard.
func (e Eip1820) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the EIP-1820 standard.
func (e Eip1820) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete EIP-1820 contract standard.
func (e Eip1820) GetStandard() ContractStandard {
	return e.Standard
}

// TokenCount returns the number of tokens associated with the EIP-1820 standard.
func (e Eip1820) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the EIP-1820 standard.
func (e Eip1820) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the EIP-1820 standard.
func (e Eip1820) String() string {
	return e.GetName()
}

// NewEip1820 initializes and returns an instance of the EIP-1820 standard.
func NewEip1820() EIP {
	return &Eip1820{
		Standard: ContractStandard{
			Name: "EIP-1820 Pseudo-introspection Registry Contract",
			Type: EIP1820,
			Functions: []Function{
				newFunction("setInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}, {Type: TypeAddress}}, nil),
				newFunction("getInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []string{TypeAddress}),
				newFunction("interfaceHash", []Input{{Type: TypeString}}, []string{TypeBytes32}),
				newFunction("updateERC165Cache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, nil),
				newFunction("implementsERC165InterfaceNoCache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []string{TypeBool}),
				newFunction("implementsERC165Interface", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []string{TypeBool}),
			},
			Events: []Event{
				newEvent("InterfaceImplementerSet", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeBytes32, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
				newEvent("ManagerChanged", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},
	}
}
