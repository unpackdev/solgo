package ast

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ExpressionStatement struct {
	*ASTBuilder
}

func NewExpressionStatement(b *ASTBuilder) *ExpressionStatement {
	return &ExpressionStatement{
		ASTBuilder: b,
	}
}

func (p *ExpressionStatement) GetTypeDescription() TypeDescription {
	return TypeDescription{}
}

func (e *ExpressionStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ExpressionStatementContext,
) Node[NodeType] {
	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.FunctionCallContext:
			statementNode := NewFunctionCall(e.ASTBuilder)
			statementNode.Parse(unit, contractNode, fnNode, bodyNode, childCtx)
			return statementNode
		case *parser.AssignmentContext:
			assignment := NewAssignment(e.ASTBuilder)
			assignment.ParseStatement(unit, contractNode, fnNode, bodyNode, ctx, childCtx)
			return assignment
		case *antlr.TerminalNodeImpl:
			// @TODO: Not sure what to do with this... It's usually just a semicolon. Perhaps to
			// add to each expression statement semicolon_found?
			// Not important right now at all...
			continue
		case *parser.PrimaryExpressionContext:
			primaryExpression := NewPrimaryExpression(e.ASTBuilder)
			primaryExpression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, childCtx)
			return primaryExpression
		default:
			fmt.Println("Expression Type: ", reflect.TypeOf(childCtx).String())
			panic("Expression statement child not recognized @ ExpressionStatement.Parse")
		}
	}

	return nil
}

/**
	for _, child := range eCtx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.FunctionCallContext:
			statementNode.NodeType = ast_pb.NodeType_FUNCTION_CALL
			statementNode = b.parseFunctionCall(
				sourceUnit, fnNode, bodyNode, statementNode, childCtx,
			)
		case *parser.AssignmentContext:
			statementNode = b.parseAssignment(
				sourceUnit, fnNode, bodyNode, statementNode, childCtx,
			)
		case *antlr.TerminalNodeImpl:
			continue
		default:
			zap.L().Warn(
				"Expression statement child not recognized",
				zap.String("type", reflect.TypeOf(childCtx).String()),
			)
		}
	}

	return statementNode
**/
