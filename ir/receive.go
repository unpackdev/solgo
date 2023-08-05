package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Receive struct {
	unit            *ast.Receive
	Id              int64             `json:"id"`
	NodeType        ast_pb.NodeType   `json:"node_type"`
	Name            string            `json:"name"`
	Kind            ast_pb.NodeType   `json:"kind"`
	Implemented     bool              `json:"implemented"`
	Visibility      ast_pb.Visibility `json:"visibility"`
	StateMutability ast_pb.Mutability `json:"state_mutability"`
	Virtual         bool              `json:"virtual"`
	Modifiers       []*Modifier       `json:"modifiers"`
	Parameters      []*Parameter      `json:"parameters"`
}

func (f *Receive) GetAST() *ast.Receive {
	return f.unit
}

func (f *Receive) GetId() int64 {
	return f.Id
}

func (f *Receive) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Receive) GetName() string {
	return f.Name
}

func (f *Receive) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Receive) IsImplemented() bool {
	return f.Implemented
}

func (f *Receive) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Receive) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Receive) IsVirtual() bool {
	return f.Virtual
}

func (f *Receive) GetModifiers() []*Modifier {
	return f.Modifiers
}

func (f *Receive) GetParameters() []*Parameter {
	return f.Parameters
}

func (b *Builder) processReceive(unit *ast.Receive) *Receive {
	toReturn := &Receive{
		unit:            unit,
		Id:              unit.GetId(),
		NodeType:        unit.GetType(),
		Kind:            unit.GetKind(),
		Name:            "receive",
		Implemented:     unit.IsImplemented(),
		Visibility:      unit.GetVisibility(),
		StateMutability: unit.GetStateMutability(),
		Virtual:         unit.IsVirtual(),
		Modifiers:       make([]*Modifier, 0),
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

	return toReturn
}
