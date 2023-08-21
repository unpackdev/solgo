package standards

var standards = map[Standard]ContractStandard{
	EIP20: {
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
	EIP721: {
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
	EIP1155: {
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
	EIP1820: {
		Name: "EIP-1820 Pseudo-introspection Registry Contract",
		Url:  "https://eips.ethereum.org/EIPS/eip-1820",
		Type: EIP1820,
		Functions: []Function{
			newFunction("setInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}, {Type: TypeAddress}}, nil),
			newFunction("getInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeAddress}}),
			newFunction("interfaceHash", []Input{{Type: TypeString}}, []Output{{Type: TypeBytes32}}),
			newFunction("updateERC165Cache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, nil),
			newFunction("implementsERC165InterfaceNoCache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
			newFunction("implementsERC165Interface", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
		},
		Events: []Event{
			newEvent("InterfaceImplementerSet", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeBytes32, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			newEvent("ManagerChanged", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
		},
	},
	EIP1822: {
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
	EIP1967: {
		Name: "EIP-1967 Proxy Storage Slots",
		Url:  "https://eips.ethereum.org/EIPS/eip-1967",
		Type: EIP1967,
		Functions: []Function{
			newFunction("setInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}, {Type: TypeAddress}}, nil),
			newFunction("getInterfaceImplementer", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeAddress}}),
			newFunction("interfaceHash", []Input{{Type: TypeString}}, []Output{{Type: TypeBytes32}}),
			newFunction("updateERC165Cache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, nil),
			newFunction("implementsERC165InterfaceNoCache", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
			newFunction("implementsERC165Interface", []Input{{Type: TypeAddress}, {Type: TypeBytes32}}, []Output{{Type: TypeBool}}),
		},
		Events: []Event{
			newEvent("InterfaceImplementerSet", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeBytes32, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
			newEvent("AdminChanged", []Input{{Type: TypeAddress, Indexed: true}, {Type: TypeAddress, Indexed: true}}, nil),
		},
	},
}
