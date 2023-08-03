package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// BreakStatement represents a 'break' statement in Solidity.
type BreakStatement struct {
	*ASTBuilder // Embedding ASTBuilder for building the AST.

	Id       int64           `json:"id"`        // Unique identifier for the break statement.
	NodeType ast_pb.NodeType `json:"node_type"` // Type of the node, which is 'BREAK' for a break statement.
	Src      SrcNode         `json:"src"`       // Source information about the node, such as its line and column numbers in the source file.
}

// NewBreakStatement creates a new BreakStatement instance.
func NewBreakStatement(b *ASTBuilder) *BreakStatement {
	return &BreakStatement{
		ASTBuilder: b,                     // ASTBuilder instance for building the AST.
		Id:         b.GetNextID(),         // Generating a new unique identifier for the break statement.
		NodeType:   ast_pb.NodeType_BREAK, // Setting the node type to 'BREAK'.
	}
}

// SetReferenceDescriptor sets the reference descriptions of the BreakStatement node.
func (b *BreakStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the break statement.
func (b *BreakStatement) GetId() int64 {
	return b.Id
}

// GetType returns the type of the node, which is 'BREAK' for a break statement.
func (b *BreakStatement) GetType() ast_pb.NodeType {
	return b.NodeType
}

// GetSrc returns the source information about the node, such as its line and column numbers in the source file.
func (b *BreakStatement) GetSrc() SrcNode {
	return b.Src
}

// GetTypeDescription returns the type description of the break statement.
// As the break statement doesn't have a type description, it returns nil.
func (b *BreakStatement) GetTypeDescription() *TypeDescription {
	return nil
}

// GetNodes returns the child nodes of the break statement.
// As the break statement doesn't have any child nodes, it returns nil.
func (b *BreakStatement) GetNodes() []Node[NodeType] {
	return nil
}

// ToProto returns the protobuf representation of the break statement.
func (b *BreakStatement) ToProto() NodeType {
	proto := ast_pb.Break{
		Id:       b.GetId(),
		NodeType: b.GetType(),
		Src:      b.GetSrc().ToProto(),
	}

	return NewTypedStruct(&proto, "Break")
}

// Parse populates the BreakStatement fields using the provided parser context.
func (b *BreakStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.BreakStatementContext,
) Node[NodeType] {
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}
	return b
}
