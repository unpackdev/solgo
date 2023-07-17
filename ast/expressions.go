package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseExpressionStatement(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, eCtx *parser.ExpressionStatementContext) *ast_pb.Statement {
	for _, child := range eCtx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.FunctionCallContext:
			statementNode = b.parseFunctionCall(
				fnNode, bodyNode, statementNode, childCtx,
			)
		}
	}

	return statementNode
}
