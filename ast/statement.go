package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, parentStatement *ast_pb.Statement, statementCtx parser.IStatementContext) *ast_pb.Statement {
	if simpleStatement := statementCtx.SimpleStatement(); simpleStatement != nil {
		return b.parseSimpleStatement(node, bodyNode, simpleStatement.(*parser.SimpleStatementContext))
	}

	if returnStatement := statementCtx.ReturnStatement(); returnStatement != nil {
		return b.parseReturnStatement(node, bodyNode, parentStatement.Id, returnStatement.(*parser.ReturnStatementContext))
	}

	if ifStatement := statementCtx.IfStatement(); ifStatement != nil {
		return b.parseIfStatement(node, bodyNode, ifStatement.(*parser.IfStatementContext))
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
