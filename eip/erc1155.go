package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Erc1155 represents the ERC-1155 Multi Token Standard.
type Erc1155 struct {
	// Standard holds the details of the ERC-1155 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the ERC-1155 standard.
func (e Erc1155) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the ERC-1155 standard.
func (e Erc1155) GetType() Standard {
	return e.Standard.Type
}

// GetFunctions returns the functions associated with the ERC-1155 standard.
func (e Erc1155) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the ERC-1155 standard.
func (e Erc1155) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete ERC-1155 contract standard.
func (e Erc1155) GetStandard() ContractStandard {
	return e.Standard
}

// TokenCount returns the number of tokens associated with the ERC-1155 standard.
func (e Erc1155) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the ERC-1155 standard.
func (e Erc1155) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the ERC-1155 standard.
func (e Erc1155) String() string {
	return e.GetName()
}

// NewErc1155 initializes and returns an instance of the ERC-1155 standard.
func NewErc1155() EIP {
	return &Erc1155{
		Standard: ContractStandard{
			Name: "ERC-1155 Multi Token Standard",
			Type: ERC1155,
			Functions: []Function{
				newFunction("safeTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}, {Type: TypeUint256}, {Type: TypeBytes}}, nil),
				newFunction("safeBatchTransferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256Array}, {Type: TypeUint256Array}, {Type: TypeBytes}}, nil),
				newFunction("balanceOf", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, []string{TypeUint256}),
				newFunction("balanceOfBatch", []Input{{Type: TypeAddressArray}, {Type: TypeUint256Array}}, []string{TypeUint256Array}),
				newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
				newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []string{TypeBool}),
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
