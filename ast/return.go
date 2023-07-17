package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseReturnStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, returnStatement *parser.ReturnStatementContext) *ast_pb.Statement {
	toReturn := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(returnStatement.GetStart().GetLine()),
			Column:      int64(returnStatement.GetStart().GetColumn()),
			Start:       int64(returnStatement.GetStart().GetStart()),
			End:         int64(returnStatement.GetStop().GetStop()),
			Length:      int64(returnStatement.GetStop().GetStop() - returnStatement.GetStart().GetStart() + 1),
			ParentIndex: bodyNode.Id,
		},
		NodeType: ast_pb.NodeType_RETURN_STATEMENT,
	}

	if expression := returnStatement.Expression(); expression != nil {
		toReturn.Expression = &ast_pb.Expression{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(expression.GetStart().GetLine()),
				Column:      int64(expression.GetStart().GetColumn()),
				Start:       int64(expression.GetStart().GetStart()),
				End:         int64(expression.GetStop().GetStop()),
				Length:      int64(expression.GetStop().GetStop() - expression.GetStart().GetStart() + 1),
				ParentIndex: toReturn.Id,
			},
			NodeType: ast_pb.NodeType_IDENTIFIER,
			Name:     expression.GetText(),
		}
	}

	// @TODO: Need to parse whole structure prior return types can be properly addressed.
	// It can be type that is in the body and not the type that is in arguments of the function.
	if node.ReturnParameters != nil {
		toReturn.FunctionReturnParameters = node.ReturnParameters.Id
	}

	return toReturn
}
