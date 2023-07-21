package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, parentStatement *ast_pb.Statement, statementCtx parser.IStatementContext) *ast_pb.Statement {
	if simpleStatement := statementCtx.SimpleStatement(); simpleStatement != nil {
		return b.parseSimpleStatement(
			sourceUnit,
			node, bodyNode,
			simpleStatement.(*parser.SimpleStatementContext),
		)
	}

	if returnStatement := statementCtx.ReturnStatement(); returnStatement != nil {
		return b.parseReturnStatement(
			sourceUnit, node, bodyNode, parentStatement.Id,
			returnStatement.(*parser.ReturnStatementContext),
		)
	}

	if ifStatement := statementCtx.IfStatement(); ifStatement != nil {
		return b.parseIfStatement(
			sourceUnit,
			node, bodyNode,
			ifStatement.(*parser.IfStatementContext),
		)
	}

	if emitStatement := statementCtx.EmitStatement(); emitStatement != nil {
		return b.parseEmitStatement(
			sourceUnit,
			node, bodyNode,
			emitStatement.(*parser.EmitStatementContext),
		)
	}

	if whileStatement := statementCtx.WhileStatement(); whileStatement != nil {
		return b.parseWhileStatement(
			sourceUnit,
			node, bodyNode,
			whileStatement.(*parser.WhileStatementContext),
		)
	}

	if breakStatement := statementCtx.BreakStatement(); breakStatement != nil {
		return b.parseBreakStatement(
			sourceUnit,
			node, bodyNode,
			breakStatement.(*parser.BreakStatementContext),
		)
	}

	if continueStatement := statementCtx.ContinueStatement(); continueStatement != nil {
		return b.parseContinueStatement(
			sourceUnit,
			node, bodyNode,
			continueStatement.(*parser.ContinueStatementContext),
		)
	}

	if revertStatement := statementCtx.RevertStatement(); revertStatement != nil {
		return b.parseRevertStatement(
			sourceUnit,
			node, bodyNode,
			revertStatement.(*parser.RevertStatementContext),
		)
	}

	if forStatement := statementCtx.ForStatement(); forStatement != nil {
		return b.parseForLoopStatement(
			sourceUnit,
			node, bodyNode,
			forStatement.(*parser.ForStatementContext),
		)
	}

	if doWhileStatement := statementCtx.DoWhileStatement(); doWhileStatement != nil {
		return b.parseDoWhileStatement(
			sourceUnit,
			node, bodyNode,
			doWhileStatement.(*parser.DoWhileStatementContext),
		)
	}

	return nil
}

func (b *ASTBuilder) parseSimpleStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, statement *parser.SimpleStatementContext) *ast_pb.Statement {
	toReturn := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(statement.GetStart().GetLine()),
			Column:      int64(statement.GetStart().GetColumn()),
			Start:       int64(statement.GetStart().GetStart()),
			End:         int64(statement.GetStop().GetStop()),
			Length:      int64(statement.GetStop().GetStop() - statement.GetStart().GetStart() + 1),
			ParentIndex: bodyNode.Id,
		},
	}

	if variableDeclaration := statement.VariableDeclarationStatement(); variableDeclaration != nil {
		toReturn.NodeType = ast_pb.NodeType_VARIABLE_DECLARATION_STATEMENT
		toReturn = b.parseVariableDeclaration(
			sourceUnit, node, bodyNode, toReturn,
			variableDeclaration.(*parser.VariableDeclarationStatementContext),
		)
	} else if expressionStatement := statement.ExpressionStatement(); expressionStatement != nil {
		toReturn = b.parseExpressionStatement(
			sourceUnit, node, bodyNode, toReturn,
			expressionStatement.(*parser.ExpressionStatementContext),
		)
	}

	return toReturn
}
