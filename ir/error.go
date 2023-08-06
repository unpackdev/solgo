package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Error struct {
	unit *ast.ErrorDefinition

	Id              int64                `json:"id"`
	NodeType        ast_pb.NodeType      `json:"node_type"`
	Name            string               `json:"name"`
	Parameters      []*Parameter         `json:"parameters"`
	TypeDescription *ast.TypeDescription `json:"type_description"`
}

func (e *Error) GetAST() *ast.ErrorDefinition {
	return e.unit
}

// GetId returns the unique identifier of the node.
func (e *Error) GetId() int64 {
	return e.Id
}

// GetName returns the name of the node.
func (e *Error) GetName() string {
	return e.Name
}

// GetParameters returns the parameters of the error.
func (e *Error) GetParameters() []*Parameter {
	return e.Parameters
}

// GetNodeType returns the type of the node in the AST.
func (e *Error) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetTypeDescription returns the type description of the node.
func (e *Error) GetTypeDescription() *ast.TypeDescription {
	return e.TypeDescription
}

// GetSrc returns the source location of the node.
func (e *Error) GetSrc() ast.SrcNode {
	return e.unit.GetSrc()
}

func (e *Error) ToProto() *ir_pb.Error {
	proto := &ir_pb.Error{
		Id:              e.GetId(),
		NodeType:        e.GetNodeType(),
		Name:            e.GetName(),
		Parameters:      make([]*ir_pb.Parameter, 0),
		TypeDescription: e.GetTypeDescription().ToProto(),
	}

	for _, parameter := range e.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	return proto
}

func (b *Builder) processError(unit *ast.ErrorDefinition) *Error {
	toReturn := &Error{
		unit:            unit,
		Id:              unit.GetId(),
		NodeType:        unit.GetType(),
		Name:            unit.GetName(),
		Parameters:      make([]*Parameter, 0),
		TypeDescription: unit.GetTypeDescription(),
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
