package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Eip1155 represents the EIP-1155 Multi Token Standard.
type Eip1155 struct {
	// Standard holds the details of the EIP-1155 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the EIP-1155 standard.
func (e Eip1155) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the EIP-1155 standard.
func (e Eip1155) GetType() Standard {
	return e.Standard.Type
}

// GetUrl returns the URL of the EIP-1155 standard.
func (e Eip1155) GetUrl() string {
	return e.Standard.Url
}

// GetFunctions returns the functions associated with the EIP-1155 standard.
func (e Eip1155) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the EIP-1155 standard.
func (e Eip1155) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete EIP-1155 contract standard.
func (e Eip1155) GetStandard() ContractStandard {
	return e.Standard
}

// IsStagnant returns a boolean indicating whether the EIP-1155 standard is stagnant.
func (e Eip1155) IsStagnant() bool {
	return e.Standard.Stagnant
}

// ConfidenceCheck checks the contract for the EIP-1155 standard compliance.
func (e Eip1155) ConfidenceCheck(contract *Contract) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// TokenCount returns the number of tokens associated with the EIP-1155 standard.
func (e Eip1155) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the EIP-1155 standard.
func (e Eip1155) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the EIP-1155 standard.
func (e Eip1155) String() string {
	return e.GetName()
}

// NewEip1155 initializes and returns an instance of the EIP-1155 standard.
func NewEip1155() EIP {
	return &Eip1155{
		Standard: ContractStandard{
			Name: "ERC-1155 Multi Token Standard",
			Url:  "https://eips.ethereum.org/EIPS/eip-1155",
			Type: EIP1155,
			Functions: []Function{
				newFunction("safeTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}, {Type: TypeUint256}, {Type: TypeBytes}}, nil),
				newFunction("safeBatchTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256Array}, {Type: TypeUint256Array}, {Type: TypeBytes}}, nil),
				newFunction("balanceOf", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []Output{{Type: TypeUint256}}),
				newFunction("balanceOfBatch", []Input{{Type: TypeAddressArray}, {Type: TypeUint256Array}}, []Output{{Type: TypeUint256Array}}),
				newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
				newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []Output{{Type: TypeBool}}),
			},
			Events: []Event{
				newEvent("TransferSingle", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}, {Type: TypeUint256}}, nil),
				newEvent("TransferBatch", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeAddressArray, Indexed: true}, {Type: TypeUint256Array}, {Type: TypeUint256Array}}, nil),
				newEvent("ApprovalForAll", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeBool}}, nil),
				newEvent("URI", []Input{{Type: TypeString, Indexed: false}, {Type: TypeUint256, Indexed: true}}, nil),
			},
		},
	}
}
