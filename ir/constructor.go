// Package ir provides an intermediate representation of the AST.
package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Constructor represents a contract constructor.
type Constructor struct {
	unit *ast.Constructor

	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"node_type"`
	Kind             ast_pb.NodeType   `json:"kind"`
	Name             string            `json:"name"`
	Implemented      bool              `json:"implemented"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	StateMutability  ast_pb.Mutability `json:"state_mutability"`
	Virtual          bool              `json:"virtual"`
	Modifiers        []*Modifier       `json:"modifiers"`
	Parameters       []*Parameter      `json:"parameters"`
	ReturnStatements []*Parameter      `json:"return"`
}

// GetAST returns the underlying ast.Constructor.
func (f *Constructor) GetAST() *ast.Constructor {
	return f.unit
}

// GetId returns the ID of the constructor.
func (f *Constructor) GetId() int64 {
	return f.Id
}

// GetNodeType returns the node type of the constructor.
func (f *Constructor) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

// GetName returns the name of the constructor.
func (f *Constructor) GetName() string {
	return f.Name
}

// GetKind returns the kind of the constructor.
func (f *Constructor) GetKind() ast_pb.NodeType {
	return f.Kind
}

// IsImplemented checks if the constructor is implemented.
func (f *Constructor) IsImplemented() bool {
	return f.Implemented
}

// GetVisibility returns the visibility of the constructor.
func (f *Constructor) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the state mutability of the constructor.
func (f *Constructor) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// IsVirtual checks if the constructor is virtual.
func (f *Constructor) IsVirtual() bool {
	return f.Virtual
}

// GetModifiers returns the modifiers of the constructor.
func (f *Constructor) GetModifiers() []*Modifier {
	return f.Modifiers
}

// GetParameters returns the parameters of the constructor.
func (f *Constructor) GetParameters() []*Parameter {
	return f.Parameters
}

// GetReturnStatements returns the return statements of the constructor.
func (f *Constructor) GetReturnStatements() []*Parameter {
	return f.ReturnStatements
}

// GetSrc returns the source code location of the constructor.
func (f *Constructor) GetSrc() ast.SrcNode {
	return f.unit.GetSrc()
}

// ToProto converts the constructor to its protobuf representation.
func (f *Constructor) ToProto() *ir_pb.Constructor {
	proto := &ir_pb.Constructor{
		Id:              f.GetId(),
		NodeType:        f.GetNodeType(),
		Kind:            f.GetKind(),
		Name:            f.GetName(),
		Implemented:     f.IsImplemented(),
		Visibility:      f.GetVisibility(),
		StateMutability: f.GetStateMutability(),
		Virtual:         f.IsVirtual(),
		Modifiers:       make([]*ir_pb.Modifier, 0),
		Parameters:      make([]*ir_pb.Parameter, 0),
		Return:          make([]*ir_pb.Parameter, 0),
	}

	for _, modifier := range f.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, modifier.ToProto())
	}

	for _, parameter := range f.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	for _, returnStatement := range f.GetReturnStatements() {
		proto.Return = append(proto.Return, returnStatement.ToProto())
	}

	return proto
}

// processConstructor processes the given ast.Constructor and returns its IR representation.
func (b *Builder) processConstructor(unit *ast.Constructor) *Constructor {
	toReturn := &Constructor{
		unit:             unit,
		Id:               unit.GetId(),
		NodeType:         unit.GetType(),
		Kind:             unit.GetKind(),
		Name:             "constructor",
		Implemented:      unit.IsImplemented(),
		Visibility:       unit.GetVisibility(),
		StateMutability:  unit.GetStateMutability(),
		Modifiers:        make([]*Modifier, 0),
		Parameters:       make([]*Parameter, 0),
		ReturnStatements: make([]*Parameter, 0),
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

	for _, parameter := range unit.GetParameters().GetParameters() {
		toReturn.Parameters = append(toReturn.Parameters, &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	for _, returnStatement := range unit.GetReturnParameters().GetParameters() {
		toReturn.ReturnStatements = append(toReturn.ReturnStatements, &Parameter{
			unit:            returnStatement,
			Id:              returnStatement.GetId(),
			NodeType:        returnStatement.GetType(),
			Name:            returnStatement.GetName(),
			Type:            returnStatement.GetTypeName().GetName(),
			TypeDescription: returnStatement.GetTypeDescription(),
		})
	}

	return toReturn
}
