package eip

import eip_pb "github.com/txpull/protos/dist/go/eip"

// Erc721 represents the ERC-721 token standard (NFT standard).
type Erc721 struct {
	// Standard holds the details of the ERC-721 contract standard.
	Standard ContractStandard
}

// GetName returns the name of the ERC-721 standard.
func (e Erc721) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the ERC-721 standard.
func (e Erc721) GetType() Standard {
	return e.Standard.Type
}

// GetFunctions returns the functions associated with the ERC-721 standard.
func (e Erc721) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the ERC-721 standard.
func (e Erc721) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete ERC-721 contract standard.
func (e Erc721) GetStandard() ContractStandard {
	return e.Standard
}

// TokenCount returns the number of tokens associated with the ERC-721 standard.
func (e Erc721) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the ERC-721 standard.
func (e Erc721) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the ERC-721 standard.
func (e Erc721) String() string {
	return e.GetName()
}

// NewErc721 initializes and returns an instance of the ERC-721 standard.
// It sets up the standard with its name, type, associated functions, and events.
func NewErc721() EIP {
	return &Erc721{
		Standard: ContractStandard{
			Name: "ERC-721 Non-Fungible Token Standard",
			Type: ERC721,
			Functions: []Function{
				newFunction("name", nil, []string{TypeString}),
				newFunction("symbol", nil, []string{TypeString}),
				newFunction("totalSupply", nil, []string{TypeUint256}),
				newFunction("balanceOf", []Input{{Type: TypeAddress}}, []string{TypeUint256}),
				newFunction("ownerOf", []Input{{Type: TypeUint256}}, []string{TypeAddress}),
				newFunction("transferFrom", []Input{{Type: TypeAddress}, {Type: TypeAddress}, {Type: TypeUint256}}, nil),
				newFunction("approve", []Input{{Type: TypeAddress}, {Type: TypeUint256}}, nil),
				newFunction("setApprovalForAll", []Input{{Type: TypeAddress}, {Type: TypeBool}}, nil),
				newFunction("getApproved", []Input{{Type: TypeUint256}}, []string{TypeAddress}),
				newFunction("isApprovedForAll", []Input{{Type: TypeAddress}, {Type: TypeAddress}}, []string{TypeBool}),
			},
			Events: []Event{
				newEvent("Transfer", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("Approval", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeUint256}}, nil),
				newEvent("ApprovalForAll", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}, {Type: TypeBool}}, nil),
			},
		},
	}
}
