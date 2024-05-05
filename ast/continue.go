package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ContinueStatement represents a 'continue' statement in the abstract syntax tree.
type ContinueStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"nodeType"`
	Src      SrcNode         `json:"src"`
}

// NewContinueStatement creates a new instance of ContinueStatement.
func NewContinueStatement(b *ASTBuilder) *ContinueStatement {
	return &ContinueStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_CONTINUE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ContinueStatement node.
// This function always returns false for now.
func (b *ContinueStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the ContinueStatement.
func (b *ContinueStatement) GetId() int64 {
	return b.Id
}

// GetType returns the NodeType of the ContinueStatement.
func (b *ContinueStatement) GetType() ast_pb.NodeType {
	return b.NodeType
}

// GetSrc returns the source information of the ContinueStatement.
func (b *ContinueStatement) GetSrc() SrcNode {
	return b.Src
}

// GetTypeDescription returns the type description associated with the ContinueStatement.
func (b *ContinueStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "continue",
		TypeIdentifier: "$_t_continue",
	}
}

// GetNodes returns an empty list of child nodes for the ContinueStatement.
func (b *ContinueStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// ToProto converts the ContinueStatement to its corresponding protocol buffer representation.
func (b *ContinueStatement) ToProto() NodeType {
	proto := ast_pb.Continue{
		Id:       b.GetId(),
		NodeType: b.GetType(),
		Src:      b.GetSrc().ToProto(),
	}

	return NewTypedStruct(&proto, "Continue")
}

// Parse parses the ContinueStatement node from the parsing context and associates it with other nodes.
func (b *ContinueStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ContinueStatementContext,
) Node[NodeType] {
	b.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}
	return b
}
