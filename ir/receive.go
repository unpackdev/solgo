package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Receive represents a receive function in the Intermediate Representation (IR) of Solidity contracts' Abstract Syntax Tree (AST).
type Receive struct {
	unit            *ast.Receive      `json:"-"`
	Id              int64             `json:"id"`               // Id is the unique identifier of the receive function.
	NodeType        ast_pb.NodeType   `json:"node_type"`        // NodeType is the type of the receive function node in the AST.
	Name            string            `json:"name"`             // Name is the name of the receive function (always "receive" for Solidity receive functions).
	Kind            ast_pb.NodeType   `json:"kind"`             // Kind is the kind of the receive function node (e.g., FunctionDefinition, FunctionType).
	Implemented     bool              `json:"implemented"`      // Implemented is true if the receive function is implemented in the contract, false otherwise.
	Visibility      ast_pb.Visibility `json:"visibility"`       // Visibility represents the visibility of the receive function (e.g., public, private, internal, external).
	StateMutability ast_pb.Mutability `json:"state_mutability"` // StateMutability represents the mutability of the receive function (e.g., pure, view, nonpayable, payable).
	Virtual         bool              `json:"virtual"`          // Virtual is true if the receive function is virtual, false otherwise.
	Modifiers       []*Modifier       `json:"modifiers"`        // Modifiers is a list of modifiers applied to the receive function.
	Overrides       []*Override       `json:"overrides"`        // Overrides is a list of functions overridden by the receive function.
	Parameters      []*Parameter      `json:"parameters"`       // Parameters is a list of parameters of the receive function.
}

// GetAST returns the underlying AST node of the receive function.
func (f *Receive) GetAST() *ast.Receive {
	return f.unit
}

// GetId returns the unique identifier of the receive function.
func (f *Receive) GetId() int64 {
	return f.Id
}

// GetNodeType returns the type of the receive function node in the AST.
func (f *Receive) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

// GetName returns the name of the receive function (always "receive" for Solidity receive functions).
func (f *Receive) GetName() string {
	return f.Name
}

// GetKind returns the kind of the receive function node (e.g., FunctionDefinition, FunctionType).
func (f *Receive) GetKind() ast_pb.NodeType {
	return f.Kind
}

// IsImplemented returns true if the receive function is implemented in the contract, false otherwise.
func (f *Receive) IsImplemented() bool {
	return f.Implemented
}

// GetVisibility returns the visibility of the receive function (e.g., public, private, internal, external).
func (f *Receive) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the mutability of the receive function (e.g., pure, view, nonpayable, payable).
func (f *Receive) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// IsVirtual returns true if the receive function is virtual, false otherwise.
func (f *Receive) IsVirtual() bool {
	return f.Virtual
}

// GetModifiers returns a list of modifiers applied to the receive function.
func (f *Receive) GetModifiers() []*Modifier {
	return f.Modifiers
}

// GetOverrides returns a list of functions overridden by the receive function.
func (f *Receive) GetOverrides() []*Override {
	return f.Overrides
}

// GetParameters returns a list of parameters of the receive function.
func (f *Receive) GetParameters() []*Parameter {
	return f.Parameters
}

// GetSrc returns the source code location of the receive function.
func (f *Receive) GetSrc() ast.SrcNode {
	return f.unit.GetSrc()
}

// ToProto is a function that converts the Receive to a protobuf message.
func (f *Receive) ToProto() *ir_pb.Receive {
	proto := &ir_pb.Receive{
		Id:              f.GetId(),
		NodeType:        f.GetNodeType(),
		Kind:            f.GetKind(),
		Name:            f.GetName(),
		Implemented:     f.IsImplemented(),
		Visibility:      f.GetVisibility(),
		StateMutability: f.GetStateMutability(),
		Virtual:         f.IsVirtual(),
		Modifiers:       make([]*ir_pb.Modifier, 0),
		Overrides:       make([]*ir_pb.Override, 0),
		Parameters:      make([]*ir_pb.Parameter, 0),
	}

	for _, modifier := range f.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, modifier.ToProto())
	}

	for _, overrides := range f.GetOverrides() {
		proto.Overrides = append(proto.Overrides, overrides.ToProto())
	}

	for _, parameter := range f.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	return proto
}

// processReceive is a function that processes the given receive function node and returns a Receive.
func (b *Builder) processReceive(unit *ast.Receive) *Receive {
	toReturn := &Receive{
		unit:            unit,
		Id:              unit.GetId(),
		NodeType:        unit.GetType(),
		Kind:            unit.GetKind(),
		Name:            "receive", // The name of the receive function is always "receive".
		Implemented:     unit.IsImplemented(),
		Visibility:      unit.GetVisibility(),
		StateMutability: unit.GetStateMutability(),
		Virtual:         unit.IsVirtual(),
		Modifiers:       make([]*Modifier, 0),
		Overrides:       make([]*Override, 0),
		Parameters:      make([]*Parameter, 0),
	}

	for _, modifier := range unit.GetModifiers() {
		toReturn.Modifiers = append(toReturn.Modifiers, &Modifier{
			unit:          modifier,
			Id:            modifier.GetId(),
			NodeType:      modifier.GetType(),
			Name:          modifier.GetName(),
			ArgumentTypes: modifier.GetArgumentTypes(),
		})
	}

	for _, oride := range unit.GetOverrides() {
		override := &Override{
			unit:                    oride,
			Id:                      oride.GetId(),
			NodeType:                oride.GetType(),
			Name:                    oride.GetName(),
			ReferencedDeclarationId: oride.GetReferencedDeclaration(),
			TypeDescription:         oride.GetTypeDescription(),
			Overrides:               make([]*Parameter, 0),
		}

		// @TODO: Fix this
		// for _, overrideParameter := range oride.GetOverrides() {
		// 	override.Overrides = append(override.Overrides, overrideParameter)
		// }

		toReturn.Overrides = append(toReturn.Overrides, override)
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		param := &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		}

		// If the type name is not set, but the path node is set, we use the name from the path node.
		if param.GetType() == "" && parameter.GetTypeName().GetPathNode() != nil {
			param.Type = parameter.GetTypeName().GetPathNode().Name
		}

		toReturn.Parameters = append(toReturn.Parameters, param)
	}

	return toReturn
}
