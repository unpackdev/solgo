package ast

import (
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
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

func (b *ASTBuilder) parseExpression(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, arg *ast_pb.Expression, parentIndex int64, expressionCtx parser.IExpressionContext) *ast_pb.Expression {
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
		toReturn.NodeType = ast_pb.NodeType_BINARY_OPERATION

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
	case *parser.OrderComparisonContext:
		toReturn.NodeType = ast_pb.NodeType_BINARY_OPERATION

		if toReturn.TypeDescriptions == nil {
			toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeIdentifier: "t_bool",
				TypeString:     "bool",
			}
		}

		if childCtx.GreaterThanOrEqual() != nil {
			toReturn.Operator = ast_pb.Operator_GREATER_THAN_OR_EQUAL
		} else if childCtx.LessThanOrEqual() != nil {
			toReturn.Operator = ast_pb.Operator_LESS_THAN_OR_EQUAL
		} else if childCtx.GreaterThan() != nil {
			toReturn.Operator = ast_pb.Operator_GREATER_THAN
		} else if childCtx.LessThan() != nil {
			toReturn.Operator = ast_pb.Operator_LESS_THAN
		}

		toReturn.LeftExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, childCtx.Expression(0),
		)

		toReturn.RightExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, childCtx.Expression(1),
		)

	case *parser.EqualityComparisonContext:
		toReturn.NodeType = ast_pb.NodeType_BINARY_OPERATION

		if childCtx.Equal() != nil {
			toReturn.Operator = ast_pb.Operator_EQUAL
		} else if childCtx.NotEqual() != nil {
			toReturn.Operator = ast_pb.Operator_NOT_EQUAL
		}

		toReturn.LeftExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, childCtx.Expression(0),
		)

		toReturn.RightExpression = b.parseExpression(
			fnNode, bodyNode, arg, toReturn.Id, childCtx.Expression(1),
		)
	case *parser.PrimaryExpressionContext:
		if childCtx.Literal() != nil {
			toReturn.NodeType = ast_pb.NodeType_LITERAL
			literalCtx := childCtx.Literal()
			toReturn.IsPure = true

			if literalCtx.StringLiteral() != nil {
				toReturn.Kind = ast_pb.NodeType_STRING

				toReturn.Value = strings.TrimSpace(
					// There can be hex 22 at beginning and end of literal.
					// We should drop it as that's ASCII for double quote.
					strings.ReplaceAll(literalCtx.StringLiteral().GetText(), "\"", ""),
				)
				toReturn.HexValue = hex.EncodeToString([]byte(toReturn.Value))

				toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
					TypeIdentifier: "t_string_literal",
					TypeString: fmt.Sprintf(
						"literal_string %s",
						literalCtx.StringLiteral().GetText(),
					),
				}

				return toReturn
			}

			if literalCtx.NumberLiteral() != nil {
				toReturn.Kind = ast_pb.NodeType_NUMBER

				toReturn.Value = strings.TrimSpace(
					// There can be hex 22 at beginning and end of literal.
					// We should drop it as that's ASCII for double quote.
					strings.ReplaceAll(literalCtx.NumberLiteral().GetText(), "\"", ""),
				)
				toReturn.HexValue = hex.EncodeToString([]byte(toReturn.Value))

				// Check if the number is a floating-point number
				if strings.Contains(toReturn.Value, ".") {
					parts := strings.Split(toReturn.Value, ".")

					// The numerator is the number without the decimal point
					numerator, _ := strconv.Atoi(parts[0] + parts[1])

					// The denominator is a power of 10 equal to the number of digits in the fractional part
					denominator := int(math.Pow(10, float64(len(parts[1]))))

					toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
						TypeIdentifier: fmt.Sprintf("t_rational_%d_by_%d", numerator, denominator),
						TypeString: fmt.Sprintf(
							"fixed_const %s",
							literalCtx.NumberLiteral().GetText(),
						),
					}
				} else {
					numerator, _ := strconv.Atoi(toReturn.Value)

					// The denominator for an integer is 1
					denominator := 1

					toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
						TypeIdentifier: fmt.Sprintf("t_rational_%d_by_%d", numerator, denominator),
						TypeString: fmt.Sprintf(
							"int_const %s",
							literalCtx.NumberLiteral().GetText(),
						),
					}
				}

				return toReturn
			}

		}

	default:
		panic(fmt.Sprintf("Expression Reflect Unimplemented: %s \n", reflect.TypeOf(childCtx)))
	}

	return toReturn
}
