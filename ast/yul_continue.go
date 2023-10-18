package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// YulContinueStatement represents a YUL continue statement in the abstract syntax tree.
type YulContinueStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL continue statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL continue statement node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL continue statement.
	Src SrcNode `json:"src"`
}

// NewYulContinueStatement creates a new YulContinueStatement instance.
func NewYulContinueStatement(b *ASTBuilder) *YulContinueStatement {
	return &YulContinueStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_CONTINUE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulContinueStatement node.
func (y *YulContinueStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulContinueStatement.
func (y *YulContinueStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulContinueStatement.
func (y *YulContinueStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulContinueStatement.
func (y *YulContinueStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns an empty list of child nodes for the YulContinueStatement.
func (y *YulContinueStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetTypeDescription returns the type description of the YulContinueStatement.
func (y *YulContinueStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// ToProto converts the YulContinueStatement to its protocol buffer representation.
func (y *YulContinueStatement) ToProto() NodeType {
	toReturn := ast_pb.YulContinueStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	return NewTypedStruct(&toReturn, "YulContinueStatement")
}

// Parse parses a YulContinueStatement node.
func (y *YulContinueStatement) Parse(
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
