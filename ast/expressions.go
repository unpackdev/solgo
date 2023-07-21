package ast

import (
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

func (b *ASTBuilder) parseExpressionStatement(sourceUnit *ast_pb.SourceUnit, fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, eCtx *parser.ExpressionStatementContext) *ast_pb.Statement {
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
}

func (b *ASTBuilder) parseAssignment(sourceUnit *ast_pb.SourceUnit, fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, assignmentCtx *parser.AssignmentContext) *ast_pb.Statement {
	statementNode.NodeType = ast_pb.NodeType_EXPRESSION_CONTEXT
	statementNode.Expression = &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(assignmentCtx.GetStart().GetLine()),
			Column:      int64(assignmentCtx.GetStart().GetColumn()),
			Start:       int64(assignmentCtx.GetStart().GetStart()),
			End:         int64(assignmentCtx.GetStop().GetStop()),
			ParentIndex: statementNode.Id,
		},
		NodeType: ast_pb.NodeType_ASSIGNMENT,
	}

	operator := assignmentCtx.AssignOp()
	if operator != nil {
		if operator.Assign() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_EQUAL
		} else if operator.AssignAdd() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_PLUS_EQUAL
		} else if operator.AssignSub() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_MINUS_EQUAL
		} else if operator.AssignMul() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_MUL_EQUAL
		} else if operator.AssignDiv() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_DIVISION
		} else if operator.AssignMod() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_MOD_EQUAL
		} else if operator.AssignBitAnd() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_AND_EQUAL
		} else if operator.AssignBitOr() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_OR_EQUAL
		} else if operator.AssignBitXor() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_XOR_EQUAL
		} else if operator.AssignShl() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_SHIFT_LEFT_EQUAL
		} else if operator.AssignShr() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_SHIFT_RIGHT_EQUAL
		} else if operator.AssignBitAnd() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_BIT_AND_EQUAL
		} else if operator.AssignBitOr() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_BIT_OR_EQUAL
		} else if operator.AssignBitXor() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_BIT_XOR_EQUAL
		} else if operator.AssignSar() != nil {
			statementNode.Expression.Operator = ast_pb.Operator_POW_EQUAL
		} else {
			zap.L().Warn(
				"Assignment operator not recognized",
				zap.String("type", reflect.TypeOf(operator).String()),
			)
		}
	}

	leftExpressionCtx := assignmentCtx.Expression(0)
	leftExpression := b.parseExpression(
		sourceUnit, fnNode, bodyNode, nil, statementNode.Id, leftExpressionCtx,
	)
	statementNode.Expression.LeftExpression = leftExpression

	rightExpressionCtx := assignmentCtx.Expression(1)
	rightExpression := b.parseExpression(
		sourceUnit, fnNode, bodyNode, nil, statementNode.Id, rightExpressionCtx,
	)
	statementNode.Expression.RightExpression = rightExpression

	return statementNode
}

func (b *ASTBuilder) parseExpression(sourceUnit *ast_pb.SourceUnit, fnNode *ast_pb.Node, bodyNode *ast_pb.Body, parentExpression *ast_pb.Expression, parentIndex int64, expressionCtx parser.IExpressionContext) *ast_pb.Expression {
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
		Name:                   expressionCtx.GetText(),
		NodeType:               ast_pb.NodeType_IDENTIFIER,
		OverloadedDeclarations: []int64{},
	}

	referenceFound := false

	// Search for argument reference in state variable declarations.
	for _, node := range b.currentStateVariables {
		if node.GetName() == expressionCtx.GetText() {
			referenceFound = true
			toReturn.ReferencedDeclaration = node.Id
			toReturn.TypeDescriptions = node.GetTypeName().GetTypeDescriptions()
		}
	}

	// Search for argument reference in statement declarations.
	if !referenceFound {
		for _, statement := range bodyNode.GetStatements() {
			for _, declaration := range statement.GetDeclarations() {
				if declaration.GetName() == expressionCtx.GetText() {
					referenceFound = true
					toReturn.ReferencedDeclaration = declaration.Id
					toReturn.TypeDescriptions = declaration.GetTypeName().GetTypeDescriptions()
				}
			}

			for _, argument := range statement.GetArguments() {
				if argument.GetName() == expressionCtx.GetText() {
					referenceFound = true
					toReturn.ReferencedDeclaration = argument.Id
					toReturn.TypeDescriptions = argument.GetTypeDescriptions()
				}
			}
		}
	}

	// If search for reference in statement declarations failed,
	// search for reference in function parameters.
	if !referenceFound {
		if fnNode.GetParameters() != nil {
			for _, parameter := range fnNode.GetParameters().Parameters {
				if parameter.GetName() == expressionCtx.GetText() {
					referenceFound = true
					toReturn.ReferencedDeclaration = parameter.Id
					toReturn.TypeDescriptions = parameter.GetTypeName().GetTypeDescriptions()
				}
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
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, leftCtx,
		)

		if childCtx.Div() != nil {
			toReturn.Operator = ast_pb.Operator_DIVISION
		} else if childCtx.Mul() != nil {
			toReturn.Operator = ast_pb.Operator_MULTIPLICATION
		} else if childCtx.Mod() != nil {
			toReturn.Operator = ast_pb.Operator_MODULO
		}

		toReturn.RightExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, rightCtx,
		)

		// There's probably a better way to do this (not probably)
		// but for now let's go with this one and see if this will be a problem.
		// @TODO: Check in solc how they do this.
		toReturn.TypeDescriptions = toReturn.LeftExpression.TypeDescriptions

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
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(0),
		)

		toReturn.RightExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(1),
		)

	case *parser.EqualityComparisonContext:
		toReturn.NodeType = ast_pb.NodeType_BINARY_OPERATION

		if childCtx.Equal() != nil {
			toReturn.Operator = ast_pb.Operator_EQUAL
		} else if childCtx.NotEqual() != nil {
			toReturn.Operator = ast_pb.Operator_NOT_EQUAL
		}

		toReturn.LeftExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(0),
		)

		toReturn.RightExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(1),
		)

		// There's probably a better way to do this (not probably)
		// but for now let's go with this one and see if this will be a problem.
		// @TODO: Check in solc how they do this.
		toReturn.TypeDescriptions = toReturn.LeftExpression.TypeDescriptions

	case *parser.PrimaryExpressionContext:
		literalCtx := childCtx.Literal()

		if literalCtx != nil {
			toReturn.NodeType = ast_pb.NodeType_LITERAL
			toReturn.IsPure = true

			if literalCtx.BooleanLiteral() != nil {
				if toReturn.Name == "true" || toReturn.Name == "false" {
					toReturn.Name = ""
				}

				toReturn.Kind = ast_pb.NodeType_BOOLEAN
				toReturn.Value = strings.TrimSpace(
					// There can be hex 22 at beginning and end of literal.
					// We should drop it as that's ASCII for double quote.
					strings.ReplaceAll(literalCtx.BooleanLiteral().GetText(), "\"", ""),
				)
				toReturn.HexValue = hex.EncodeToString([]byte(toReturn.Value))

				toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
					TypeIdentifier: "t_bool",
					TypeString:     "bool",
				}
			}

			if literalCtx.StringLiteral() != nil {
				toReturn.Name = ""
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

			if literalCtx.HexStringLiteral() != nil {
				toReturn.Kind = ast_pb.NodeType_HEX_STRING

				toReturn.Value = strings.TrimSpace(
					// There can be hex 22 at beginning and end of literal.
					// We should drop it as that's ASCII for double quote.
					strings.ReplaceAll(literalCtx.StringLiteral().GetText(), "\"", ""),
				)
				toReturn.HexValue = hex.EncodeToString([]byte(toReturn.Value))

				toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
					TypeIdentifier: "t_string_hex_literal",
					TypeString: fmt.Sprintf(
						"literal_hex_string %s",
						literalCtx.StringLiteral().GetText(),
					),
				}

				return toReturn
			}
		}

		// Handle magic cases...
		if childCtx.GetText() == "msg" {
			toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeIdentifier: "t_magic_message",
				TypeString:     "msg",
			}
		}

		if toReturn.TypeDescriptions == nil {
			if parentExpression != nil {
				if parentExpression.TypeDescriptions != nil {
					toReturn.TypeDescriptions = parentExpression.TypeDescriptions
				} else {
					zap.L().Debug(
						"Unknown primary expression type description",
						zap.String("name", childCtx.GetText()),
					)
				}
			}
		}

	case *parser.MemberAccessContext:
		toReturn.NodeType = ast_pb.NodeType_MEMBER_ACCESS
		toReturn.MemberName = childCtx.Identifier().GetText()

		if childCtx.Expression() != nil {
			toReturn.Expression = b.parseExpression(
				sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(),
			)
		}

		if parentExpression != nil {
			for _, arguments := range parentExpression.Arguments {
				toReturn.ArgumentTypes = append(
					toReturn.ArgumentTypes,
					arguments.TypeDescriptions,
				)
			}
		}

		// Now we are going to search through all existing source units in hope
		// to discover reference declaration...
		for _, units := range b.sourceUnits {
			if units.GetRoot() != nil && len(units.GetRoot().Nodes) > 0 {
				for _, nodeCtx := range units.GetRoot().Nodes {
					for _, node := range nodeCtx.Nodes {
						if node.Name == toReturn.MemberName {
							toReturn.ReferencedDeclaration = node.Id
						}
					}
				}
			}
		}

		if toReturn.Expression != nil {
			if toReturn.Expression.TypeDescriptions != nil {
				if toReturn.Expression.TypeDescriptions.TypeIdentifier == "t_magic_message" {
					switch toReturn.MemberName {
					case "sender":
						toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
							TypeIdentifier: "t_address",
							TypeString:     "address",
						}
					case "data":
						toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
							TypeIdentifier: "t_bytes_calldata_ptr",
							TypeString:     "bytes calldata",
						}
					default:
						zap.L().Warn(
							"Unknown magic message member",
							zap.String("member", toReturn.MemberName),
						)
					}
				}
			}
		}

	case *parser.FunctionCallContext:
		toReturn.NodeType = ast_pb.NodeType_FUNCTION_CALL
		toReturn.Kind = ast_pb.NodeType_FUNCTION_CALL

		if childCtx.CallArgumentList() != nil {
			for _, argumentCtx := range childCtx.CallArgumentList().AllExpression() {
				toReturn.Arguments = append(
					toReturn.Arguments,
					b.parseExpression(
						sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, argumentCtx,
					),
				)
			}
		}

		typeString := "function("
		typeStringComponents := []string{}

		for _, component := range toReturn.Arguments {
			if !component.IsPure {
				toReturn.IsPure = false
			}

			// This is a problem...
			if component.TypeDescriptions == nil {
				zap.L().Warn(
					"Function call component type description is nil. Tuple type is corrupted.",
					zap.String("sourceUnit", sourceUnit.AbsolutePath),
					zap.String("function_name", fnNode.Name),
					zap.String("component_name", component.Name),
					zap.Int64("component_id", component.Id),
					zap.String("component_node_type", component.NodeType.String()),
				)
				continue
			}

			typeStringComponents = append(typeStringComponents, component.TypeDescriptions.TypeString)
		}

		if len(typeStringComponents) > 0 {
			typeString += strings.Join(typeStringComponents, ",")
		}

		typeString += ")"

		toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
			TypeIdentifier: "t_function_call",
			TypeString:     typeString,
		}

		if childCtx.Expression() != nil {
			toReturn.Expression = b.parseExpression(
				sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(),
			)
		}

	case *parser.AddSubOperationContext:
		toReturn.NodeType = ast_pb.NodeType_BINARY_OPERATION

		toReturn.Operator = ast_pb.Operator_ADDITION
		if childCtx.Sub() != nil {
			toReturn.Operator = ast_pb.Operator_SUBTRACTION
		}

		toReturn.LeftExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(0),
		)

		toReturn.RightExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(1),
		)

		// There's probably a better way to do this (not probably)
		// but for now let's go with this one and see if this will be a problem.
		// @TODO: Check in solc how they do this.
		toReturn.TypeDescriptions = toReturn.LeftExpression.TypeDescriptions

	case *parser.TupleContext:
		toReturn.NodeType = ast_pb.NodeType_TUPLE_EXPRESSION

		// In case that name starts with (, means we only have parameters inside
		// and showing it is useless...
		if strings.HasPrefix(childCtx.GetText(), "(") {
			toReturn.Name = ""
		}

		if childCtx.TupleExpression() != nil {
			for _, tupleExpressionCtx := range childCtx.TupleExpression().AllExpression() {
				toReturn.Components = append(
					toReturn.Components,
					b.parseExpression(
						sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, tupleExpressionCtx,
					),
				)
			}
		}

		toReturn.IsPure = true

		typeString := "tuple("
		typeStringComponents := []string{}

		for _, component := range toReturn.Components {
			if !component.IsPure {
				toReturn.IsPure = false
			}

			// This is a problem...
			if component.TypeDescriptions == nil {
				zap.L().Warn(
					"Tuple component type description is nil. Tuple type is corrupted.",
					zap.String("sourceUnit", sourceUnit.AbsolutePath),
					zap.String("function_name", fnNode.Name),
					zap.String("component_name", component.Name),
					zap.Int64("component_id", component.Id),
					zap.String("component_node_type", component.NodeType.String()),
				)
				continue
			}

			typeStringComponents = append(typeStringComponents, component.TypeDescriptions.TypeString)
		}

		if len(typeStringComponents) > 0 {
			typeString += strings.Join(typeStringComponents, ",")
		}

		typeString += ")"

		toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
			TypeIdentifier: "t_tuple",
			TypeString:     typeString,
		}

		if fnNode.ReturnParameters != nil {
			toReturn.FunctionReturnParameters = fnNode.ReturnParameters.Id
		}

	case *parser.MetaTypeContext: // @TODO: Type names could be improved, for now not...
		toReturn.Name = childCtx.Type().GetText()
		toReturn.TypeDescriptions = &ast_pb.TypeDescriptions{
			TypeString: childCtx.GetText(),
		}
	case *parser.IndexAccessContext:
		toReturn.NodeType = ast_pb.NodeType_INDEX_ACCESS

		toReturn.BaseExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(0),
		)
		toReturn.IndexExpression = b.parseExpression(
			sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(1),
		)

	default:
		zap.L().Warn(
			"Expression type not implemented...",
			zap.String("reflection_type", reflect.TypeOf(childCtx).String()),
		)
	}

	return toReturn
}
