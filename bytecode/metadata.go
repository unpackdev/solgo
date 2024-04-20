package bytecode

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/mr-tron/base58"
	metadata_pb "github.com/unpackdev/protos/dist/go/metadata"
)

// Metadata represents the metadata contained in Ethereum contract creation bytecode.
// The structure and encoding of the metadata is defined by the Solidity compiler.
// More information can be found at https://docs.soliditylang.org/en/v0.8.20/metadata.html#encoding-of-the-metadata-hash-in-the-bytecode
type Metadata struct {
	executionBytecode []byte      // The execution bytecode of the contract
	cborLength        int16       // The length of the CBOR metadata
	auxbytes          []byte      // The raw CBOR metadata
	Ipfs              []byte      `cbor:"ipfs"`         // The IPFS hash of the metadata, if present
	Bzzr1             []byte      `cbor:"bzzr1"`        // The Swarm hash of the metadata, if present (version 1)
	Bzzr0             []byte      `cbor:"bzzr0"`        // The Swarm hash of the metadata, if present (version 0)
	Experimental      interface{} `cbor:"experimental"` // Experimental metadata, if present
	Solc              []byte      `cbor:"solc"`         // The version of the Solidity compiler used
}

// ToProto converts the Metadata instance into a protobuf representation, suitable for
// serialization and transmission across different systems or networks.
func (m *Metadata) ToProto() *metadata_pb.BytecodeMetadata {
	return &metadata_pb.BytecodeMetadata{
		ExecutionBytecode: m.executionBytecode,
		CborLength:        uint32(m.cborLength),
		Raw:               m.auxbytes,
		Ipfs:              m.GetIPFS(),
		Bzzr1:             m.GetBzzr1(),
		Bzzr0:             m.GetBzzr0(),
		Experimental:      m.GetExperimental(),
		Solc:              m.GetCompilerVersion(),
		Solgo:             "", // TODO: Solgo version should be appended...
	}
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
	if experimental, ok := m.Experimental.(string); ok {
		toReturn, err := strconv.ParseBool(experimental)
		if err != nil {
			return false
		}
		return toReturn
	}

	if experimental, ok := m.Experimental.(bool); ok {
		return experimental
	}

	return false
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

// GetExecutionBytecode returns the execution bytecode of the contract.
func (m *Metadata) GetExecutionBytecode() []byte {
	return m.executionBytecode
}

// GetAuxBytecode returns the raw CBOR metadata of the contract.
func (m *Metadata) GetAuxBytecode() []byte {
	return m.auxbytes
}

// GetCborLength returns the length of the CBOR metadata.
func (m *Metadata) GetCborLength() int16 {
	return m.cborLength
}

// AuxFound returns whether the CBOR metadata bytes are contained in the provided byte slice.
// This is used to verify if contract extracted cbor metadata can be found in the deployed/execution contract bytecode.
func (m *Metadata) AuxFound(b []byte) bool {
	return bytes.Contains(b, m.auxbytes)
}

// GetUrls returns the URLs of the contract's metadata.
func (m *Metadata) GetUrls() []string {
	urls := make([]string, 0)
	if len(m.GetIPFS()) > 7 {
		urls = append(urls, m.GetIPFS())
	}
	if len(m.GetBzzr0()) > 6 {
		urls = append(urls, m.GetBzzr1())
	}
	if len(m.GetBzzr1()) > 6 {
		urls = append(urls, m.GetBzzr0())
	}
	return urls
}

// IsPartial checks whenever decoded bytecode data is partial (some fields are missing, not known)
func (m *Metadata) IsPartial() bool {
	if m.cborLength == 0 {
		return true
	}

	if len(m.auxbytes) == 0 {
		return true
	}

	if len(m.Solc) == 0 {
		return true
	}

	if len(m.Ipfs) == 0 && len(m.Bzzr0) == 0 && len(m.Bzzr1) == 0 {
		return true
	}

	return false
}

// DecodeContractMetadata decodes the metadata from Ethereum contract creation bytecode.
// It returns a Metadata object and an error, if any occurred during decoding.
func DecodeContractMetadata(bytecode []byte) (*Metadata, error) {
	if len(bytecode) == 0 {
		return nil, errors.New("provided bytecode slice is empty")
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
	if len(bytecode) >= bytesLength+cborLength {
		toReturn.executionBytecode = bytecode[:len(bytecode)-bytesLength-cborLength]
		toReturn.auxbytes = bytecode[len(bytecode)-bytesLength-cborLength : len(bytecode)-bytesLength]
		if err := cbor.Unmarshal(toReturn.auxbytes, &toReturn); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("provided bytecode slice is smaller than the length of the cbor object")
	}

	return &toReturn, nil
}
