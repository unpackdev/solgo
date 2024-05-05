package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ExpressionContext represents an AST node for an expression context in Solidity.
type ExpressionContext struct {
	*ASTBuilder

	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"nodeType"`
	Src             SrcNode          `json:"src"`
	Value           string           `json:"value"`
	TypeDescription *TypeDescription `json:"typeDescription"`
}

// NewExpressionContext creates a new ExpressionContext instance with the provided ASTBuilder.
// The ASTBuilder is used to facilitate the construction of the AST.
func NewExpressionContext(b *ASTBuilder) *ExpressionContext {
	return &ExpressionContext{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_EXPRESSION_CONTEXT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ExpressionContext node.
// This function always returns false for now.
func (f *ExpressionContext) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return true
}

// GetId returns the ID of the ExpressionContext.
func (f *ExpressionContext) GetId() int64 {
	return f.Id
}

// GetType returns the NodeType of the ExpressionContext.
func (f *ExpressionContext) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the source information of the ExpressionContext.
func (f *ExpressionContext) GetSrc() SrcNode {
	return f.Src
}

// GetTypeDescription returns the type description associated with the ExpressionContext.
func (f *ExpressionContext) GetTypeDescription() *TypeDescription {
	return f.TypeDescription
}

// GetNodes returns the child nodes of the ExpressionContext.
func (f *ExpressionContext) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	return toReturn
}

// ToProto converts the ExpressionContext to its corresponding protocol buffer representation.
func (f *ExpressionContext) ToProto() NodeType {
	proto := ast_pb.Expression{
		Id:              f.GetId(),
		NodeType:        f.GetType(),
		Src:             f.GetSrc().ToProto(),
		TypeDescription: f.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "ExpressionContext")
}

// Parse parses the ExpressionContext node from the parsing context and associates it with other nodes.
func (f *ExpressionContext) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	parentNodeId int64,
	ctx *parser.ExpressionContext,
) Node[NodeType] {
	f.Id = f.GetNextID()

	f.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNodeId,
	}
	f.Value = ctx.GetText()
	f.TypeDescription = &TypeDescription{
		TypeString:     "string",
		TypeIdentifier: "$_t_string",
	}
	return f
}
