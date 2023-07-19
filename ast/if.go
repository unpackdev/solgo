package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseIfStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, ifCtx *parser.IfStatementContext) *ast_pb.Statement {
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

	condition := b.parseExpression(sourceUnit, node, bodyNode, nil, statement.Id, ifCtx.Expression())
	statement.Condition = condition

	if !ifCtx.IsEmpty() {
		if len(ifCtx.AllStatement()) > 0 {
			for _, statementCtx := range ifCtx.AllStatement() {
				if statementCtx.IsEmpty() {
					continue
				}

				if statementCtx.Block() != nil {
					blockCtx := statementCtx.Block()
					statement.TrueBody = &ast_pb.Statement{
						Id: atomic.AddInt64(&b.nextID, 1) - 1,
						Src: &ast_pb.Src{
							Line:        int64(blockCtx.GetStart().GetLine()),
							Start:       int64(blockCtx.GetStart().GetStart()),
							End:         int64(blockCtx.GetStop().GetStop()),
							Length:      int64(blockCtx.GetStop().GetStop() - blockCtx.GetStart().GetStart() + 1),
							ParentIndex: statement.Id,
						},
						NodeType: ast_pb.NodeType_BLOCK,
					}

					for _, stmtCtx := range statementCtx.Block().AllStatement() {
						statement.TrueBody.Statements = append(
							statement.TrueBody.Statements,
							b.parseStatement(
								sourceUnit, node, bodyNode, statement.TrueBody,
								stmtCtx,
							),
						)
					}
				}
			}
		}
	}

	return statement
}
