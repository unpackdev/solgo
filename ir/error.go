package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Error represents an error definition in the IR.
type Error struct {
	Unit            *ast.ErrorDefinition `json:"ast"`
	Id              int64                `json:"id"`
	NodeType        ast_pb.NodeType      `json:"nodeType"`
	Name            string               `json:"name"`
	Parameters      []*Parameter         `json:"parameters"`
	TypeDescription *ast.TypeDescription `json:"type_description"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the error definition.
func (e *Error) GetAST() *ast.ErrorDefinition {
	return e.Unit
}

// GetId returns the ID of the error definition.
func (e *Error) GetId() int64 {
	return e.Id
}

// GetName returns the name of the error definition.
func (e *Error) GetName() string {
	return e.Name
}

// GetParameters returns the parameters of the error definition.
func (e *Error) GetParameters() []*Parameter {
	return e.Parameters
}

// GetNodeType returns the NodeType of the error definition.
func (e *Error) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetTypeDescription returns the type description of the error definition.
func (e *Error) GetTypeDescription() *ast.TypeDescription {
	return e.TypeDescription
}

// GetSrc returns the source location of the error definition.
func (e *Error) GetSrc() ast.SrcNode {
	return e.Unit.GetSrc()
}

// ToProto converts the Error to its protobuf representation.
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

// processError processes the error definition unit and returns the Error.
func (b *Builder) processError(unit *ast.ErrorDefinition) *Error {
	toReturn := &Error{
		Unit:            unit,
		Id:              unit.GetId(),
		NodeType:        unit.GetType(),
		Name:            unit.GetName(),
		Parameters:      make([]*Parameter, 0),
		TypeDescription: unit.GetTypeDescription(),
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		toReturn.Parameters = append(toReturn.Parameters, &Parameter{
			Unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	return toReturn
}
