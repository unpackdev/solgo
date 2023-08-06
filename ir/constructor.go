package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

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

func (f *Constructor) GetAST() *ast.Constructor {
	return f.unit
}

func (f *Constructor) GetId() int64 {
	return f.Id
}

func (f *Constructor) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Constructor) GetName() string {
	return f.Name
}

func (f *Constructor) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Constructor) IsImplemented() bool {
	return f.Implemented
}

func (f *Constructor) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Constructor) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Constructor) IsVirtual() bool {
	return f.Virtual
}

func (f *Constructor) GetModifiers() []*Modifier {
	return f.Modifiers
}

func (f *Constructor) GetParameters() []*Parameter {
	return f.Parameters
}

func (f *Constructor) GetReturnParameters() []*Parameter {
	return f.ReturnStatements
}

func (f *Constructor) ToProto() *ir_pb.Constructor {
	proto := &ir_pb.Constructor{
		Id: f.GetId(),
	}

	return proto
}

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
