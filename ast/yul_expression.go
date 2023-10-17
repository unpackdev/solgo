package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulExpressionStatement struct {
	*ASTBuilder

	Id         int64           `json:"id"`
	NodeType   ast_pb.NodeType `json:"node_type"`
	Src        SrcNode         `json:"src"`
	Expression Node[NodeType]  `json:"expression"`
}

func NewYulExpressionStatement(b *ASTBuilder) *YulExpressionStatement {
	return &YulExpressionStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_EXPRESSION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulExpressionStatement node.
func (y *YulExpressionStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulExpressionStatement) GetId() int64 {
	return y.Id
}

func (y *YulExpressionStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulExpressionStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulExpressionStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

func (y *YulExpressionStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulExpressionStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulExpressionStatement node.
func (f *YulExpressionStatement) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulExpressionStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulExpressionContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStart().GetStop()),
		Length:      int64(ctx.GetStart().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(y.ASTBuilder)
		y.Expression = literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	if ctx.YulFunctionCall() != nil {
		fcStatement := NewYulFunctionCallStatement(y.ASTBuilder)
		y.Expression = fcStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulFunctionCall().(*parser.YulFunctionCallContext),
		)
	}

	return y
}

func ParseYulExpression(
	b *ASTBuilder,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	assignmentNode *parser.YulAssignmentContext,
	variableNode *parser.YulVariableDeclarationContext,
	parentNode Node[NodeType],
	ctx parser.IYulExpressionContext,
) Node[NodeType] {
	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(b)
		return literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	if ctx.YulFunctionCall() != nil {
		fcStatement := NewYulFunctionCallStatement(b)
		return fcStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulFunctionCall().(*parser.YulFunctionCallContext),
		)
	}

	/* 	zap.L().Warn(
		"ParseYulExpression: unimplemented child type",
		zap.Any("child_type", reflect.TypeOf(ctx).String()),
	) */

	return nil
}
