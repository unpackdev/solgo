package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseDoWhileStatement(
	sourceUnit *ast_pb.SourceUnit,
	fnNode *ast_pb.Node,
	bodyNode *ast_pb.Body,
	ctx *parser.DoWhileStatementContext,
) *ast_pb.Statement {
	statement := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: fnNode.Id,
		},
		NodeType: ast_pb.NodeType_DO_WHILE_STATEMENT,
	}
	return statement
}
