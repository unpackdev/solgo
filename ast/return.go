package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseReturnStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, parentIndex int64, returnStatement *parser.ReturnStatementContext) *ast_pb.Statement {
	id := atomic.AddInt64(&b.nextID, 1) - 1

	return &ast_pb.Statement{
		Id: id,
		Src: &ast_pb.Src{
			Line:        int64(returnStatement.GetStart().GetLine()),
			Column:      int64(returnStatement.GetStart().GetColumn()),
			Start:       int64(returnStatement.GetStart().GetStart()),
			End:         int64(returnStatement.GetStop().GetStop()),
			Length:      int64(returnStatement.GetStop().GetStop() - returnStatement.GetStart().GetStart() + 1),
			ParentIndex: parentIndex,
		},
		NodeType: ast_pb.NodeType_RETURN_STATEMENT,
		FunctionReturnParameters: func() int64 {
			if node.ReturnParameters != nil {
				return node.ReturnParameters.Id
			}

			return 0
		}(),
		Expression: b.parseExpression(
			sourceUnit, node, bodyNode, nil, id, returnStatement.Expression(),
		),
	}
}
