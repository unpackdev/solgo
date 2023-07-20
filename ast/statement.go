package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, parentStatement *ast_pb.Statement, statementCtx parser.IStatementContext) *ast_pb.Statement {
	if simpleStatement := statementCtx.SimpleStatement(); simpleStatement != nil {
		return b.parseSimpleStatement(sourceUnit, node, bodyNode, simpleStatement.(*parser.SimpleStatementContext))
	}

	if returnStatement := statementCtx.ReturnStatement(); returnStatement != nil {
		return b.parseReturnStatement(sourceUnit, node, bodyNode, parentStatement.Id, returnStatement.(*parser.ReturnStatementContext))
	}

	if ifStatement := statementCtx.IfStatement(); ifStatement != nil {
		return b.parseIfStatement(sourceUnit, node, bodyNode, ifStatement.(*parser.IfStatementContext))
	}

	if revertStatement := statementCtx.RevertStatement(); revertStatement != nil {
		panic("It's revert statement...")
	}

	if forStatement := statementCtx.ForStatement(); forStatement != nil {
		panic("It's for statement...")
	}

	if whileStatement := statementCtx.WhileStatement(); whileStatement != nil {
		panic("It's while statement...")
	}

	if doWhileStatement := statementCtx.DoWhileStatement(); doWhileStatement != nil {
		panic("It's do while statement...")
	}

	if continueStatement := statementCtx.ContinueStatement(); continueStatement != nil {
		panic("It's continue statement...")
	}

	if breakStatement := statementCtx.BreakStatement(); breakStatement != nil {
		panic("It's break statement...")
	}

	if emitStatement := statementCtx.EmitStatement(); emitStatement != nil {
		panic("It's emit statement...")
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
