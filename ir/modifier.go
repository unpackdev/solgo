package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Modifier represents a Modifier in the Abstract Syntax Tree.
type Modifier struct {
	Unit          *ast.ModifierInvocation `json:"ast"`
	Id            int64                   `json:"id"`
	NodeType      ast_pb.NodeType         `json:"nodeType"`
	Name          string                  `json:"name"`
	ArgumentTypes []*ast.TypeDescription  `json:"argumentTypes"`
}

// GetAST returns the underlying AST node for the Modifier.
func (m *Modifier) GetAST() *ast.ModifierInvocation {
	return m.Unit
}

// GetId returns the ID of the Modifier.
func (m *Modifier) GetId() int64 {
	return m.Id
}

// GetName returns the name of the Modifier.
func (m *Modifier) GetName() string {
	return m.Name
}

// GetNodeType returns the AST node type of the Modifier.
func (m *Modifier) GetNodeType() ast_pb.NodeType {
	return m.NodeType
}

// GetArgumentTypes returns the argument types of the Modifier.
func (m *Modifier) GetArgumentTypes() []*ast.TypeDescription {
	return m.ArgumentTypes
}

// GetSrc returns the source code location for the Modifier.
func (m *Modifier) GetSrc() ast.SrcNode {
	return m.Unit.GetSrc()
}

// ToProto converts the Modifier to its corresponding protobuf representation.
func (m *Modifier) ToProto() *ir_pb.Modifier {
	proto := &ir_pb.Modifier{
		Id:       m.GetId(),
		NodeType: m.GetNodeType(),
		Name:     m.GetName(),
	}

	for _, typeArgument := range m.GetArgumentTypes() {
		proto.ArgumentTypes = append(proto.ArgumentTypes, typeArgument.ToProto())
	}

	return proto
}
