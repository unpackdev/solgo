package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseEmitStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, ctx *parser.EmitStatementContext) *ast_pb.Statement {
	statement := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType: ast_pb.NodeType_EMIT_STATEMENT,
	}

	for _, argumentCtx := range ctx.CallArgumentList().AllExpression() {
		argument := b.parseExpression(sourceUnit, node, bodyNode, nil, statement.Id, argumentCtx)
		statement.Arguments = append(statement.Arguments, argument)
	}

	if ctx.Expression() != nil {
		statement.Expression = b.parseExpression(sourceUnit, node, bodyNode, nil, statement.Id, ctx.Expression())
	}

	return statement
}
