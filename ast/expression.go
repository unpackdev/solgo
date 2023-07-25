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

func (e *Expression) GetTypeDescription() TypeDescription {
	return TypeDescription{}
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
	fmt.Println("Expression: ", ctx.GetText())
	fmt.Println("Expression Type: ", reflect.TypeOf(ctx))

	switch ctxType := ctx.(type) {
	case *parser.AddSubOperationContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseAddSub(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.OrderComparisonContext:
		binaryExp := NewBinaryOperationExpression(e.ASTBuilder)
		return binaryExp.ParseOrderComparison(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.PrimaryExpressionContext:
		primaryExp := NewPrimaryExpression(e.ASTBuilder)
		primaryExp.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
		return primaryExp
	default:
		panic("Expression type not supported @ Expression.Parse")
	}
}
