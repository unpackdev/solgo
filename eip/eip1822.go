package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Eip1822 represents the EIP-1822 Universal Proxy Standard (UPS).
type Eip1822 struct {
	// Standard holds the details of the EIP-1822 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the EIP-1822 standard.
func (e Eip1822) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the EIP-1822 standard.
func (e Eip1822) GetType() Standard {
	return e.Standard.Type
}

// GetUrl returns the URL of the EIP-1822 standard.
func (e Eip1822) GetUrl() string {
	return e.Standard.Url
}

// GetFunctions returns the functions associated with the EIP-1822 standard.
func (e Eip1822) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the EIP-1822 standard.
func (e Eip1822) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete EIP-1822 contract standard.
func (e Eip1822) GetStandard() ContractStandard {
	return e.Standard
}

// IsStagnant returns a boolean indicating whether the EIP-1822 standard is stagnant.
func (e Eip1822) IsStagnant() bool {
	return e.Standard.Stagnant
}

// ConfidenceCheck checks the contract for the EIP-1822 standard compliance.
func (e Eip1822) ConfidenceCheck(contract *Contract) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// TokenCount returns the number of tokens associated with the EIP-1822 standard.
func (e Eip1822) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the EIP-1822 standard.
func (e Eip1822) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the EIP-1822 standard.
func (e Eip1822) String() string {
	return e.GetName()
}

// NewEip1822 initializes and returns an instance of the EIP-1822 standard.
func NewEip1822() EIP {
	return &Eip1822{
		Standard: ContractStandard{
			Name:     "EIP-1822 Universal Proxy Standard (UPS)",
			Url:      "https://eips.ethereum.org/EIPS/eip-1822",
			Stagnant: true,
			Type:     EIP1822,
			Functions: []Function{
				newFunction("getImplementation", nil, []Output{{Type: TypeAddress}}),
				newFunction("upgradeTo", []Input{{Type: TypeAddress}}, nil),
				newFunction("upgradeToAndCall", []Input{{Type: TypeAddress, Indexed: false}, {Type: TypeString, Indexed: false}}, nil),
				newFunction("setProxyOwner", []Input{{Type: TypeAddress}}, nil),
			},
			Events: []Event{
				newEvent("Upgraded", []Input{{Type: TypeAddress, Indexed: true}}, nil),
				newEvent("ProxyOwnershipTransferred", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			},
		},
	}
}
