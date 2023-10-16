package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

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
