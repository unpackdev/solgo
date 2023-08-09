package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Eip721 represents the EIP-721 token standard (NFT standard).
type Eip721 struct {
	// Standard holds the details of the EIP-721 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the EIP-721 standard.
func (e Eip721) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the EIP-721 standard.
func (e Eip721) GetType() Standard {
	return e.Standard.Type
}

// GetUrl returns the URL of the EIP-721 standard.
func (e Eip721) GetUrl() string {
	return e.Standard.Url
}

// GetFunctions returns the functions associated with the EIP-721 standard.
func (e Eip721) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the EIP-721 standard.
func (e Eip721) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete EIP-721 contract standard.
func (e Eip721) GetStandard() ContractStandard {
	return e.Standard
}

// IsStagnant returns a boolean indicating whether the EIP-721 standard is stagnant.
func (e Eip721) IsStagnant() bool {
	return e.Standard.Stagnant
}

// ConfidenceCheck checks the contract for the EIP-721 standard compliance.
func (e Eip721) ConfidenceCheck(contract *Contract) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// TokenCount returns the number of tokens associated with the EIP-721 standard.
func (e Eip721) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the EIP-721 standard.
func (e Eip721) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the EIP-721 standard.
func (e Eip721) String() string {
	return e.GetName()
}

// NewEip721 initializes and returns an instance of the EIP-721 standard.
// It sets up the standard with its name, type, associated functions, and events.
func NewEip721() EIP {
	return &Eip721{
		Standard: ContractStandard{
			Name: "EIP-721 Non-Fungible Token Standard",
			Url:  "https://eips.ethereum.org/EIPS/eip-721",
			Type: EIP721,
			Functions: []Function{
				newFunction("name", nil, []Output{{Type: TypeString}}),
				newFunction("symbol", nil, []Output{{Type: TypeString}}),
				newFunction("totalSupply", nil, []Output{{Type: TypeUint256}}),
				newFunction("balanceOf", []Input{{Type: TypeAddress}}, []Output{{Type: TypeUint256}}),
				newFunction("ownerOf", []Input{{Type: TypeUint256}}, []Output{{Type: TypeAddress}}),
				newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, nil),
				newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, nil),
				newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
				newFunction("getApproved", []Input{{Type: TypeUint256}}, []Output{{Type: TypeAddress}}),
				newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeBool}}),
			},
			Events: []Event{
				newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("ApprovalForAll", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeBool}}, nil),
			},
		},
	}
}
