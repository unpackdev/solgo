package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Parameter struct {
	unit            *ast.Parameter       `json:"-"`
	Id              int64                `json:"id"`
	NodeType        ast_pb.NodeType      `json:"node_type"`
	Name            string               `json:"name"`
	Type            string               `json:"type"`
	TypeDescription *ast.TypeDescription `json:"type_description"`
}

type Modifier struct {
	unit          *ast.ModifierInvocation `json:"-"`
	Id            int64                   `json:"id"`
	NodeType      ast_pb.NodeType         `json:"node_type"`
	Name          string                  `json:"name"`
	ArgumentTypes []*ast.TypeDescription  `json:"argument_types"`
}

type Function struct {
	unit *ast.Function

	Id                      int64             `json:"id"`
	NodeType                ast_pb.NodeType   `json:"node_type"`
	Kind                    ast_pb.NodeType   `json:"kind"`
	Name                    string            `json:"name"`
	Implemented             bool              `json:"implemented"`
	Visibility              ast_pb.Visibility `json:"visibility"`
	StateMutability         ast_pb.Mutability `json:"state_mutability"`
	Virtual                 bool              `json:"virtual"`
	ReferencedDeclarationId int64             `json:"referenced_declaration_id"`
	Modifiers               []*Modifier       `json:"modifiers"`
	Parameters              []*Parameter      `json:"parameters"`
	ReturnStatements        []*Parameter      `json:"return"`
}

func (f *Function) GetAST() *ast.Function {
	return f.unit
}

func (f *Function) GetId() int64 {
	return f.Id
}

func (f *Function) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Function) IsImplemented() bool {
	return f.Implemented
}

func (f *Function) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Function) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Function) IsVirtual() bool {
	return f.Virtual
}

func (f *Function) GetModifiers() []*Modifier {
	return f.Modifiers
}

func (f *Function) GetParameters() []*Parameter {
	return f.Parameters
}

func (f *Function) GetReturnStatements() []*Parameter {
	return f.ReturnStatements
}

func (b *Builder) processFunction(unit *ast.Function) *Function {
	toReturn := &Function{
		unit:                    unit,
		Id:                      unit.GetId(),
		NodeType:                unit.GetType(),
		Kind:                    unit.GetKind(),
		Name:                    unit.GetName(),
		Implemented:             unit.IsImplemented(),
		Visibility:              unit.GetVisibility(),
		StateMutability:         unit.GetStateMutability(),
		Virtual:                 unit.IsVirtual(),
		ReferencedDeclarationId: unit.GetReferencedDeclaration(),
		Modifiers:               make([]*Modifier, 0),
		Parameters:              make([]*Parameter, 0),
		ReturnStatements:        make([]*Parameter, 0),
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
