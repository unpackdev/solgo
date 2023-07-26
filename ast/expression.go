package ast

import (
	"fmt"
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Expression struct {
	*ASTBuilder
}

func NewExpression(b *ASTBuilder) *Expression {
	return &Expression{
		ASTBuilder: b,
	}
}

func (e *Expression) GetTypeDescription() *TypeDescription {
	return nil
}

func (e *Expression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDecar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx parser.IExpressionContext,
) Node[NodeType] {
	switch ctxType := ctx.(type) {
	case *parser.AddSubOperationContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseAddSub(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.OrderComparisonContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseOrderComparison(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.MulDivModOperationContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseMulDivMod(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.EqualityComparisonContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseEqualityComparison(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.AssignmentContext:
		assignment := NewAssignment(e.ASTBuilder)
		assignment.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
		return assignment
	case *parser.FunctionCallContext:
		statementNode := NewFunctionCall(e.ASTBuilder)
		statementNode.Parse(unit, contractNode, fnNode, bodyNode, ctxType)
		return statementNode
	case *parser.MemberAccessContext:
		memberAccess := NewMemberAccessExpression(e.ASTBuilder)
		memberAccess.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
		return memberAccess
	case *parser.PrimaryExpressionContext:
		primaryExp := NewPrimaryExpression(e.ASTBuilder)
		primaryExp.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
		return primaryExp
	case *parser.IndexAccessContext:
		indexAccess := NewIndexAccess(e.ASTBuilder)
		indexAccess.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
		return indexAccess
	default:
		fmt.Println("Expression Type: ", reflect.TypeOf(ctx))
		panic("Expression type not supported @ Expression.Parse")
	}
}
