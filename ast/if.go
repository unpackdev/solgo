package ast

import (
	"os"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseIfStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, ifCtx *parser.IfStatementContext) *ast_pb.Statement {
	statement := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ifCtx.GetStart().GetLine()),
			Start:       int64(ifCtx.GetStart().GetStart()),
			End:         int64(ifCtx.GetStop().GetStop()),
			Length:      int64(ifCtx.GetStop().GetStop() - ifCtx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType: ast_pb.NodeType_IF_STATEMENT,
	}

	condition := b.parseExpression(node, bodyNode, nil, statement.Id, ifCtx.Expression())
	statement.Condition = condition

	j, _ := b.NodeToPrettyJson(statement)
	println(string(j))
	os.Exit(1)

	//var statements []*ast_pb.Statement

	for _, _ = range ifCtx.AllStatement() {
		//statement := b.parseStatement(node, bodyNode, statementCtx.(*parser.StatementContext))
		//statements = append(statements, statement)
	}

	return statement
}
