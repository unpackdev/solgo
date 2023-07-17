package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseExpressionStatement(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, eCtx *parser.ExpressionStatementContext) *ast_pb.Statement {
	for _, child := range eCtx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.FunctionCallContext:
			statementNode = b.parseFunctionCall(
				fnNode, bodyNode, statementNode, childCtx,
			)
		}
	}

	return statementNode
}

func (b *ASTBuilder) parseExpression(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, arg *ast_pb.Argument, parentIndex int64, expressionCtx parser.IExpressionContext) *ast_pb.Expression {
	toReturn := &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(expressionCtx.GetStart().GetLine()),
			Column:      int64(expressionCtx.GetStart().GetColumn()),
			Start:       int64(expressionCtx.GetStart().GetStart()),
			End:         int64(expressionCtx.GetStop().GetStop()),
			Length:      int64(expressionCtx.GetStop().GetStop() - expressionCtx.GetStart().GetStart() + 1),
			ParentIndex: parentIndex,
		},
		Name:     expressionCtx.GetText(),
		NodeType: ast_pb.NodeType_IDENTIFIER,
		// TODO: Fix this...
		OverloadedDeclarations: []int64{},
	}

	referenceFound := false

	// Search for argument reference in statement declarations.
	for _, statement := range bodyNode.GetStatements() {
		for _, declaration := range statement.GetDeclarations() {
			if declaration.GetName() == expressionCtx.GetText() {
				referenceFound = true
				toReturn.ReferencedDeclaration = declaration.Id
				toReturn.TypeDescriptions = declaration.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	// If search for reference in statement declarations failed,
	// search for reference in function parameters.
	if !referenceFound {
		for _, parameter := range fnNode.GetParameters().Parameters {
			if parameter.GetName() == expressionCtx.GetText() {
				referenceFound = true
				toReturn.ReferencedDeclaration = parameter.Id
				toReturn.TypeDescriptions = parameter.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	// Let's see if there are any recursions that needs to be done to extract sub expressions.
	switch childCtx := expressionCtx.(type) {
	case *parser.MulDivModOperationContext:
		leftCtx := childCtx.Expression(0)
		rightCtx := childCtx.Expression(1)

		toReturn.LeftExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, leftCtx,
		)

		if childCtx.Div() != nil {
			toReturn.Operator = ast_pb.Operator_DIVISION
		} else if childCtx.Mul() != nil {
			toReturn.Operator = ast_pb.Operator_MULTIPLICATION
		} else if childCtx.Mod() != nil {
			toReturn.Operator = ast_pb.Operator_MODULO
		}

		toReturn.RightExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, rightCtx,
		)
	}

	return toReturn
}
