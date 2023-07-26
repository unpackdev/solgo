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
		return assignment.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.FunctionCallContext:
		statementNode := NewFunctionCall(e.ASTBuilder)
		return statementNode.Parse(unit, contractNode, fnNode, bodyNode, ctxType)
	case *parser.MemberAccessContext:
		memberAccess := NewMemberAccessExpression(e.ASTBuilder)
		return memberAccess.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.PrimaryExpressionContext:
		primaryExp := NewPrimaryExpression(e.ASTBuilder)
		return primaryExp.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.IndexAccessContext:
		indexAccess := NewIndexAccess(e.ASTBuilder)
		return indexAccess.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.MetaTypeContext:
		metaType := NewMetaTypeExpression(e.ASTBuilder)
		return metaType.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.PayableConversionContext:
		payableConversion := NewPayableConversionExpression(e.ASTBuilder)
		return payableConversion.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.UnarySuffixOperationContext:
		unarySuffixOperation := NewUnarySuffixExpression(e.ASTBuilder)
		return unarySuffixOperation.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	case *parser.UnaryPrefixOperationContext:
		unaryPrefixOperation := NewUnaryPrefixExpression(e.ASTBuilder)
		return unaryPrefixOperation.Parse(unit, contractNode, fnNode, bodyNode, vDecar, exprNode, ctxType)
	default:
		fmt.Println("Expression Type: ", reflect.TypeOf(ctx))
		panic("Expression type not supported @ Expression.Parse")
	}
}
