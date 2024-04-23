package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	// ZeroAddress represents a zero Ethereum address
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// ZeroHash represents a hash value consisting of all zeros.
	ZeroHash = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
)

// NamedAddr encapsulates detailed information about an Ethereum address along with its metadata.
// It includes a human-readable name, an optional Ethereum Name Service (ENS) domain, a set of
// tags for categorization or annotation, the Ethereum address itself, and the type of address
// which provides additional context about its use or origin.
type NamedAddr struct {
	Name string         `json:"name"`
	Ens  string         `json:"ens"`
	Tags []string       `json:"tags"`
	Addr common.Address `json:"addr"`
	Type AddressType    `json:"type"`
}

// AppendTags appends unique tags to the named address
func (n NamedAddr) AppendTags(tags ...string) NamedAddr {
	for _, tag := range tags {
		if !StringInSlice(tag, n.Tags) {
			n.Tags = append(n.Tags, tag)
		}
	}
	return n
}

// IsZeroAddress checks whenever named address is zero address
func (n NamedAddr) IsZeroAddress() bool {
	return n.Addr == ZeroAddress
}
