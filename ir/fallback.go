package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Fallback struct {
	unit             *ast.Fallback
	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"node_type"`
	Name             string            `json:"name"`
	Kind             ast_pb.NodeType   `json:"kind"`
	Implemented      bool              `json:"implemented"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	StateMutability  ast_pb.Mutability `json:"state_mutability"`
	Virtual          bool              `json:"virtual"`
	Modifiers        []*Modifier       `json:"modifiers"`
	Parameters       []*Parameter      `json:"parameters"`
	ReturnStatements []*Parameter      `json:"return"`
}

func (f *Fallback) GetAST() *ast.Fallback {
	return f.unit
}

func (f *Fallback) GetId() int64 {
	return f.Id
}

func (f *Fallback) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Fallback) GetName() string {
	return f.Name
}

func (f *Fallback) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Fallback) IsImplemented() bool {
	return f.Implemented
}

func (f *Fallback) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Fallback) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Fallback) IsVirtual() bool {
	return f.Virtual
}

func (f *Fallback) GetModifiers() []*Modifier {
	return f.Modifiers
}

func (f *Fallback) GetParameters() []*Parameter {
	return f.Parameters
}

func (b *Builder) processFallback(unit *ast.Fallback) *Fallback {
	toReturn := &Fallback{
		unit:             unit,
		Id:               unit.GetId(),
		NodeType:         unit.GetType(),
		Kind:             unit.GetKind(),
		Name:             "fallback",
		Implemented:      unit.IsImplemented(),
		Visibility:       unit.GetVisibility(),
		StateMutability:  unit.GetStateMutability(),
		Virtual:          unit.IsVirtual(),
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

	for _, parameter := range unit.GetReturnParameters().GetParameters() {
		toReturn.ReturnStatements = append(toReturn.ReturnStatements, &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	return toReturn
}
