package ir

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Event represents an event definition in the IR.
type Event struct {
	Unit       *ast.EventDefinition `json:"ast"`
	Id         int64                `json:"id"`
	NodeType   ast_pb.NodeType      `json:"nodeType"`
	Name       string               `json:"name"`
	Anonymous  bool                 `json:"anonymous"`
	Parameters []*Parameter         `json:"parameters"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the event definition.
func (e *Event) GetAST() *ast.EventDefinition {
	return e.Unit
}

// GetId returns the ID of the event definition.
func (e *Event) GetId() int64 {
	return e.Id
}

// GetNodeType returns the NodeType of the event definition.
func (e *Event) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetName returns the name of the event definition.
func (e *Event) GetName() string {
	return e.Name
}

// GetParameters returns the parameters of the event definition.
func (e *Event) GetParameters() []*Parameter {
	return e.Parameters
}

// IsAnonymous returns whether the event definition is anonymous.
func (e *Event) IsAnonymous() bool {
	return e.Anonymous
}

// GetSrc returns the source location of the event definition.
func (e *Event) GetSrc() ast.SrcNode {
	return e.Unit.GetSrc()
}

// GetSignature computes the Keccak-256 hash of the event signature to generate the 'topic0' hash.
// This method calls GetSignatureRaw to obtain the raw event signature string and then applies
// the Keccak-256 hash function to it. The resulting hash is commonly used in Ethereum as the
// identifier for the event in logs and is crucial for event tracking and decoding in smart contract
// interactions.
func (e *Event) GetSignature() common.Hash {
	signature := e.GetSignatureRaw()
	return crypto.Keccak256Hash([]byte(signature))
}

// GetSignatureRaw constructs the raw event signature string for the Event.
// It generates this signature by concatenating the event's name with a list of its parameters' types
// in their canonical form. The canonical form of each parameter type is obtained by using the
// canonicalizeType function. This method is particularly useful for creating a human-readable
// version of the event signature, which is essential for various Ethereum-related operations,
// such as logging and event filtering.
func (e *Event) GetSignatureRaw() string {
	paramTypes := make([]string, 0)
	for _, p := range e.Parameters {
		paramTypes = append(paramTypes, canonicalizeType(p.Type))
	}
	return fmt.Sprintf("%s(%s)", e.Name, strings.Join(paramTypes, ","))
}

// ToProto converts the Event to its protobuf representation.
func (e *Event) ToProto() *ir_pb.Event {
	proto := &ir_pb.Event{
		Id:         e.GetId(),
		NodeType:   e.GetNodeType(),
		Name:       e.GetName(),
		Anonymous:  e.IsAnonymous(),
		Parameters: make([]*ir_pb.Parameter, 0),
	}

	for _, parameter := range e.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	return proto
}

// processEvent processes the event definition unit and returns the Event.
func (b *Builder) processEvent(unit *ast.EventDefinition) *Event {
	toReturn := &Event{
		Unit:       unit,
		Id:         unit.GetId(),
		NodeType:   unit.GetType(),
		Name:       unit.GetName(),
		Anonymous:  unit.IsAnonymous(),
		Parameters: make([]*Parameter, 0),
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		toReturn.Parameters = append(toReturn.Parameters, &Parameter{
			Unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Indexed:         parameter.IsIndexed(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	return toReturn
}

// canonicalizeType converts a Solidity type into its canonical form as per Solidity's type system.
// This function handles various types, including basic types (uint, int, fixed, ufixed), bytes types,
// arrays (both fixed-size and dynamic), and tuples. The canonicalization is essential for ensuring
// consistency in how types are represented, particularly when generating event signatures. It
// transforms basic types to their full representation (e.g., 'uint' to 'uint256') and handles the
// formatting of array and tuple types. Note that complex or nested tuples might require additional
// parsing, which is not covered in this basic implementation.
func canonicalizeType(typ string) string {
	switch {
	case typ == "uint":
		return "uint256"
	case typ == "int":
		return "int256"
	case typ == "fixed":
		return "fixed128x18"
	case typ == "ufixed":
		return "ufixed128x18"
	case strings.HasPrefix(typ, "bytes") && len(typ) > 5:
		// bytes1 to bytes32 are unchanged
		return typ
	case strings.HasSuffix(typ, "[]"):
		// Dynamic array
		elementType := typ[:len(typ)-2]
		return canonicalizeType(elementType) + "[]"
	case strings.Contains(typ, "[") && strings.Contains(typ, "]"):
		// Fixed-size array
		elementType := typ[:strings.Index(typ, "[")]
		arraySize := typ[strings.Index(typ, "["):]
		return canonicalizeType(elementType) + arraySize
	default:
		// For all other types, return as-is or add specific handling
		return typ
	}
}
