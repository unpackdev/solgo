package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseWhileStatement(sourceUnit *ast_pb.SourceUnit, fnNode *ast_pb.Node, bodyNode *ast_pb.Body, ctx *parser.WhileStatementContext) *ast_pb.Statement {
	statement := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: fnNode.Id,
		},
		NodeType: ast_pb.NodeType_WHILE_STATEMENT,
	}

	condition := b.parseExpression(sourceUnit, fnNode, bodyNode, nil, statement.Id, ctx.Expression())
	statement.Condition = condition

	if ctx.Statement() != nil && ctx.Statement().Block() != nil {
		blockCtx := ctx.Statement().Block()
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(blockCtx.GetStart().GetLine()),
				Column:      int64(blockCtx.GetStart().GetColumn()),
				Start:       int64(blockCtx.GetStart().GetStart()),
				End:         int64(blockCtx.GetStop().GetStop()),
				Length:      int64(blockCtx.GetStop().GetStop() - blockCtx.GetStart().GetStart() + 1),
				ParentIndex: statement.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range blockCtx.AllStatement() {
			if statement.IsEmpty() {
				continue
			}

			// Parent index statement in this case is used only to be able provide
			// index to the parent node. It is not used for anything else.
			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
				sourceUnit, fnNode, bodyNode, parentIndexStmt, statement,
			))
		}

		statement.Body = bodyNode
	}
	return statement
}
