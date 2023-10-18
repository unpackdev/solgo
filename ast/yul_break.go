package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// YulBreakStatement represents a YUL break statement in the abstract syntax tree.
type YulBreakStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL break statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL break statement node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL break statement.
	Src SrcNode `json:"src"`
}

// NewYulBreakStatement creates a new YulBreakStatement instance.
func NewYulBreakStatement(b *ASTBuilder) *YulBreakStatement {
	return &YulBreakStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_BREAK,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulBreakStatement node.
func (y *YulBreakStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulBreakStatement.
func (y *YulBreakStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulBreakStatement.
func (y *YulBreakStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulBreakStatement.
func (y *YulBreakStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns an empty list of child nodes for the YulBreakStatement.
func (y *YulBreakStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetTypeDescription returns the type description of the YulBreakStatement.
func (y *YulBreakStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// ToProto converts the YulBreakStatement to its protocol buffer representation.
func (y *YulBreakStatement) ToProto() NodeType {
	toReturn := ast_pb.YulBreakStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	return NewTypedStruct(&toReturn, "YulBreakStatement")
}

// Parse parses a YulBreakStatement node.
func (y *YulBreakStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *antlr.TerminalNodeImpl,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetSymbol().GetLine()),
		Column:      int64(ctx.GetSymbol().GetColumn()),
		Start:       int64(ctx.GetSymbol().GetStart()),
		End:         int64(ctx.GetSymbol().GetStop()),
		Length:      int64(ctx.GetSymbol().GetStop() - ctx.GetSymbol().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	return y
}
