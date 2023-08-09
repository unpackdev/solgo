package eip

// Standard represents the type for Ethereum standards and EIPs.
type Standard string

// String returns the string representation of the Standard.
func (s Standard) String() string {
	return string(s)
}

// Constants representing various Ethereum standards and EIPs.
const (
	ERC20   Standard = "ERC20"   // ERC-20 Token Standard.
	ERC721  Standard = "ERC721"  // ERC-721 Non-Fungible Token Standard.
	EIP1822 Standard = "EIP1822" // EIP-1822 Universal Proxy Standard (UPS).
	EIP1820 Standard = "EIP1820" // EIP-1820 Pseudo-introspection Registry Contract.
	ERC777  Standard = "ERC777"  // ERC-777 Token Standard.
	ERC1155 Standard = "ERC1155" // ERC-1155 Multi Token Standard.
	ERC1337 Standard = "ERC1337" // ERC-1337 Subscription Standard.
	ERC1400 Standard = "ERC1400" // ERC-1400 Security Token Standard.
	ERC1410 Standard = "ERC1410" // ERC-1410 Partially Fungible Token Standard.
	ERC165  Standard = "ERC165"  // ERC-165 Standard Interface Detection.
	ERC820  Standard = "ERC820"  // ERC-820 Registry Standard.
	ERC1014 Standard = "ERC1014" // ERC-1014 Create2 Standard.
	ERC1948 Standard = "ERC1948" // ERC-1948 Non-Fungible Data Token Standard.
	ERC1967 Standard = "ERC1967" // ERC-1967 Proxy Storage Slots.
	ERC2309 Standard = "ERC2309" // ERC-2309 Consecutive Transfer Standard.
	ERC2535 Standard = "ERC2535" // ERC-2535 Diamond Standard.
	ERC2771 Standard = "ERC2771" // ERC-2771 Meta Transactions Standard.
	ERC2917 Standard = "ERC2917" // ERC-2917 Interest-Bearing Tokens Standard.
	ERC3156 Standard = "ERC3156" // ERC-3156 Flash Loans Standard.
	ERC3664 Standard = "ERC3664" // ERC-3664 BitWords Standard.
	EIP1967 Standard = "EIP1967" // EIP-1967 Proxy Storage Slots Standard.
)
