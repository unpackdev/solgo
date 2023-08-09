package eip

import (
	"strings"

	eip_pb "github.com/txpull/protos/dist/go/eip"
)

// Standard represents the type for Ethereum standards and EIPs.
type Standard string

// String returns the string representation of the Standard.
func (s Standard) String() string {
	return string(s)
}

// ToProto() converts a string representation of an Ethereum standard
// to its corresponding protobuf enum value. If the standard is not recognized,
// it returns unknown.
func (s Standard) ToProto() eip_pb.Standard {
	if standardValue, ok := eip_pb.Standard_value[strings.ToUpper(string(s))]; ok {
		return eip_pb.Standard(standardValue)
	}
	return eip_pb.Standard_UNKNOWN
}

// Constants representing various Ethereum standards and EIPs.
const (
	EIP20   Standard = "EIP20"   // ERC-20 Token Standard.
	EIP721  Standard = "EIP721"  // ERC-721 Non-Fungible Token Standard.
	EIP1822 Standard = "EIP1822" // EIP-1822 Universal Proxy Standard (UPS).
	EIP1820 Standard = "EIP1820" // EIP-1820 Pseudo-introspection Registry Contract.
	EIP777  Standard = "EIP777"  // ERC-777 Token Standard.
	EIP1155 Standard = "EIP1155" // ERC-1155 Multi Token Standard.
	EIP1337 Standard = "EIP1337" // ERC-1337 Subscription Standard.
	EIP1400 Standard = "EIP1400" // ERC-1400 Security Token Standard.
	EIP1410 Standard = "EIP1410" // ERC-1410 Partially Fungible Token Standard.
	EIP165  Standard = "EIP165"  // ERC-165 Standard Interface Detection.
	EIP820  Standard = "EIP820"  // ERC-820 Registry Standard.
	EIP1014 Standard = "EIP1014" // ERC-1014 Create2 Standard.
	EIP1948 Standard = "EIP1948" // ERC-1948 Non-Fungible Data Token Standard.
	EIP1967 Standard = "EIP1967" // EIP-1967 Proxy Storage Slots Standard.
	EIP2309 Standard = "EIP2309" // ERC-2309 Consecutive Transfer Standard.
	EIP2535 Standard = "EIP2535" // ERC-2535 Diamond Standard.
	EIP2771 Standard = "EIP2771" // ERC-2771 Meta Transactions Standard.
	EIP2917 Standard = "EIP2917" // ERC-2917 Interest-Bearing Tokens Standard.
	EIP3156 Standard = "EIP3156" // ERC-3156 Flash Loans Standard.
	EIP3664 Standard = "EIP3664" // ERC-3664 BitWords Standard.
)

// LoadStandards loads list of supported Ethereum EIPs into the registry.
func LoadStandards() error {
	if err := RegisterStandard(EIP20, NewEip20()); err != nil {
		return err
	}

	if err := RegisterStandard(EIP721, NewEip721()); err != nil {
		return err
	}

	if err := RegisterStandard(EIP1155, NewEip1155()); err != nil {
		return err
	}

	if err := RegisterStandard(EIP1820, NewEip1820()); err != nil {
		return err
	}

	if err := RegisterStandard(EIP1822, NewEip1822()); err != nil {
		return err
	}

	if err := RegisterStandard(EIP1967, NewEip1967()); err != nil {
		return err
	}

	return nil
}
