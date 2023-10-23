package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
	"go.uber.org/zap"
)

// parseExpressionStatement is a utility function to parse an expression statement based on the provided context and parent node.
func parseExpressionStatement(
	b *ASTBuilder,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	parentNode Node[NodeType],
	parentNodeId int64,
	ctx *parser.ExpressionStatementContext,
) Node[NodeType] {
	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.FunctionCallContext:
			statementNode := NewFunctionCall(b)
			statementNode.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
			return statementNode
		case *parser.AssignmentContext:
			assignment := NewAssignment(b)
			assignment.ParseStatement(unit, contractNode, fnNode, bodyNode, parentNode, ctx, childCtx)
			return assignment
		case *parser.PrimaryExpressionContext:
			primaryExpression := NewPrimaryExpression(b)
			return primaryExpression.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.UnarySuffixOperationContext:
			unarySuffixOperation := NewUnarySuffixExpression(b)
			return unarySuffixOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.UnaryPrefixOperationContext:
			unaryPrefixOperation := NewUnaryPrefixExpression(b)
			return unaryPrefixOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.OrderComparisonContext:
			binaryExp := NewBinaryOperationExpression(b)
			return binaryExp.ParseOrderComparison(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.EqualityComparisonContext:
			binaryExp := NewBinaryOperationExpression(b)
			return binaryExp.ParseEqualityComparison(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.TupleContext:
			tupleExpr := NewTupleExpression(b)
			return tupleExpr.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.AndOperationContext:
			andOperation := NewAndOperationExpression(b)
			return andOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.AddSubOperationContext:
			binaryExp := NewBinaryOperationExpression(b)
			return binaryExp.ParseAddSub(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.MulDivModOperationContext:
			binaryExp := NewBinaryOperationExpression(b)
			return binaryExp.ParseMulDivMod(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.OrOperationContext:
			binaryExp := NewBinaryOperationExpression(b)
			return binaryExp.ParseOr(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.MemberAccessContext:
			memberAccess := NewMemberAccessExpression(b)
			return memberAccess.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.IndexAccessContext:
			indexAccess := NewIndexAccess(b)
			return indexAccess.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.MetaTypeContext:
			metaType := NewMetaTypeExpression(b)
			return metaType.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.PayableConversionContext:
			payableConversion := NewPayableConversionExpression(b)
			return payableConversion.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.NewExprContext:
			newExpr := NewExprExpression(b)
			return newExpr.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.FunctionCallOptionsContext:
			statementNode := NewFunctionCallOption(b)
			return statementNode.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.IndexRangeAccessContext:
			indexRangeAccess := NewIndexRangeAccessExpression(b)
			return indexRangeAccess.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.ExpOperationContext:
			expOperation := NewExprOperationExpression(b)
			return expOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.ConditionalContext:
			conditional := NewConditionalExpression(b)
			return conditional.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.BitAndOperationContext:
			bitAndOperation := NewBitAndOperationExpression(b)
			return bitAndOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.BitOrOperationContext:
			bitAndOperation := NewBitOrOperationExpression(b)
			return bitAndOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.BitXorOperationContext:
			bitXorOperation := NewBitXorOperationExpression(b)
			return bitXorOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *parser.ShiftOperationContext:
			shiftOperation := NewShiftOperationExpression(b)
			return shiftOperation.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, parentNodeId, childCtx)
		case *parser.InlineArrayContext:
			inlineArray := NewInlineArrayExpression(b)
			return inlineArray.Parse(unit, contractNode, fnNode, bodyNode, nil, parentNode, childCtx)
		case *antlr.TerminalNodeImpl:
			// @TODO: Not sure what to do with this... It's usually just a semicolon (;). Perhaps to
			// add to each expression statement semicolon_found?
			// Not important right now at all...
			continue
		default:
			zap.L().Warn(
				"Expression statement child not recognized @ ExpressionStatement.Parse",
				zap.String("child_type", fmt.Sprintf("%T", childCtx)),
			)
		}
	}

	return nil
}
