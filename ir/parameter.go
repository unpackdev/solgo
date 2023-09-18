package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Parameter represents a Parameter in the Abstract Syntax Tree.
type Parameter struct {
	unit            *ast.Parameter       `json:"-"`
	Id              int64                `json:"id"`
	NodeType        ast_pb.NodeType      `json:"node_type"`
	Name            string               `json:"name"`
	Type            string               `json:"type"`
	TypeDescription *ast.TypeDescription `json:"type_description"`
	Indexed         bool                 `json:"indexed"`
}

// GetAST returns the underlying AST node for the Parameter.
func (p *Parameter) GetAST() *ast.Parameter {
	return p.unit
}

// GetId returns the ID of the Parameter.
func (p *Parameter) GetId() int64 {
	return p.Id
}

// GetName returns the name of the Parameter.
func (p *Parameter) GetName() string {
	return p.Name
}

// GetNodeType returns the AST node type of the Parameter.
func (p *Parameter) GetNodeType() ast_pb.NodeType {
	return p.NodeType
}

// IsIndexed returns whether the Parameter is indexed.
func (p *Parameter) IsIndexed() bool {
	return p.Indexed
}

// GetType returns the type of the Parameter.
func (p *Parameter) GetType() string {
	return p.Type
}

// GetTypeDescription returns the type description of the Parameter.
func (p *Parameter) GetTypeDescription() *ast.TypeDescription {
	return p.TypeDescription
}

// GetSrc returns the source code location for the Parameter.
func (p *Parameter) GetSrc() ast.SrcNode {
	return p.unit.GetSrc()
}

// ToProto converts the Parameter to its corresponding protobuf representation.
func (p *Parameter) ToProto() *ir_pb.Parameter {
	proto := &ir_pb.Parameter{
		Id:              p.GetId(),
		NodeType:        p.GetNodeType(),
		Name:            p.GetName(),
		Type:            p.GetType(),
		TypeDescription: p.GetTypeDescription().ToProto(),
	}

	return proto
}
