package bytecode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/mr-tron/base58"
)

// Metadata represents the metadata contained in Ethereum contract creation bytecode.
// The structure and encoding of the metadata is defined by the Solidity compiler.
// More information can be found at https://docs.soliditylang.org/en/v0.8.20/metadata.html#encoding-of-the-metadata-hash-in-the-bytecode
type Metadata struct {
	executionBytecode []byte // The execution bytecode of the contract
	cborLength        int16  // The length of the CBOR metadata
	raw               []byte // The raw CBOR metadata
	Ipfs              []byte `cbor:"ipfs"`         // The IPFS hash of the metadata, if present
	Bzzr1             []byte `cbor:"bzzr1"`        // The Swarm hash of the metadata, if present (version 1)
	Bzzr0             []byte `cbor:"bzzr0"`        // The Swarm hash of the metadata, if present (version 0)
	Experimental      []byte `cbor:"experimental"` // Experimental metadata, if present
	Solc              []byte `cbor:"solc"`         // The version of the Solidity compiler used
}

// GetCompilerVersion returns the version of the Solidity compiler used to compile the contract.
func (m *Metadata) GetCompilerVersion() string {
	s := make([]string, 0, len(m.Solc))
	for _, i := range m.Solc {
		s = append(s, strconv.Itoa(int(i)))
	}
	return strings.Join(s, ".")
}

// GetExperimental returns whether the contract includes experimental metadata.
func (m *Metadata) GetExperimental() bool {
	toReturn, err := strconv.ParseBool(string(m.Experimental))
	if err != nil {
		return false
	}
	return toReturn
}

// GetIPFS returns the IPFS hash of the contract's metadata, if present.
func (m *Metadata) GetIPFS() string {
	return fmt.Sprintf("ipfs://%s", base58.Encode(m.Ipfs))
}

// GetBzzr0 returns the Swarm (version 0) hash of the contract's metadata, if present.
func (m *Metadata) GetBzzr0() string {
	return fmt.Sprintf("bzz://%s", base58.Encode(m.Bzzr0))
}

// GetBzzr1 returns the Swarm (version 1) hash of the contract's metadata, if present.
func (m *Metadata) GetBzzr1() string {
	return fmt.Sprintf("bzz://%s", base58.Encode(m.Bzzr1))
}

// GetRawMetadata returns the raw CBOR metadata of the contract.
func (m *Metadata) GetRawMetadata() []byte {
	return m.raw
}

// GetCborLength returns the length of the CBOR metadata.
func (m *Metadata) GetCborLength() int16 {
	return m.cborLength
}

// DecodeContractCreationMetadata decodes the metadata from Ethereum contract creation bytecode.
// It returns a Metadata object and an error, if any occurred during decoding.
func DecodeContractCreationMetadata(bytecode []byte) (*Metadata, error) {
	if len(bytecode) == 0 {
		return nil, errors.New("provided bytecode slice is empty")
	}

	if bytecode[0] != 0x60 {
		return nil, errors.New("provided bytecode slice is not a contract")
	}

	toReturn := Metadata{}

	// Per solidity docs, last two bytes of the bytecode are the length of the cbor object
	bytesLength := 2

	// Take latest 2 bytes of the bytecode (length of the cbor object)
	cborLength := int(bytecode[len(bytecode)-2])<<8 | int(bytecode[len(bytecode)-1])
	toReturn.cborLength = int16(cborLength)

	// If the length of the cbor is more or equal to the length of the execution bytecode, it means there is no cbor
	if len(bytecode)-bytesLength <= 0 {
		return nil, errors.New("provided bytecode slice does not contain cbor metadata")
	}

	// Split the bytecode into execution bytecode and auxdata
	toReturn.executionBytecode = bytecode[:len(bytecode)-bytesLength-cborLength]
	toReturn.raw = bytecode[len(bytecode)-bytesLength-cborLength : len(bytecode)-bytesLength]

	if err := cbor.Unmarshal(toReturn.raw, &toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
