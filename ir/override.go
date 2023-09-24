package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Override represents an Override in the Abstract Syntax Tree.
type Override struct {
	unit                    *ast.OverrideSpecifier `json:"-"`
	Id                      int64                  `json:"id"`
	NodeType                ast_pb.NodeType        `json:"node_type"`
	Name                    string                 `json:"name"`
	ReferencedDeclarationId int64                  `json:"referenced_declaration_id"`
	TypeDescription         *ast.TypeDescription   `json:"type_description"`
}

// GetAST returns the underlying AST node for the Override.
func (m *Override) GetAST() *ast.OverrideSpecifier {
	return m.unit
}

// GetId returns the ID of the Override.
func (m *Override) GetId() int64 {
	return m.Id
}

// GetName returns the name of the Override.
func (m *Override) GetName() string {
	return m.Name
}

// GetNodeType returns the AST node type of the Override.
func (m *Override) GetNodeType() ast_pb.NodeType {
	return m.NodeType
}

// GetReferencedDeclarationId returns the ID of the referenced declaration for the Override.
func (m *Override) GetReferencedDeclarationId() int64 {
	return m.ReferencedDeclarationId
}

// GetTypeDescription returns the type description of the Override.
func (m *Override) GetTypeDescription() *ast.TypeDescription {
	return m.TypeDescription
}

// GetSrc returns the source node of the Override.
func (m *Override) GetSrc() ast.SrcNode {
	return m.unit.GetSrc()
}

// ToProto converts the Override to its corresponding protobuf representation.
func (m *Override) ToProto() *ir_pb.Override {
	proto := &ir_pb.Override{
		Id:                      m.GetId(),
		NodeType:                m.GetNodeType(),
		Name:                    m.GetName(),
		ReferencedDeclarationId: m.GetReferencedDeclarationId(),
		TypeDescription:         m.GetTypeDescription().ToProto(),
	}
	return proto
}
