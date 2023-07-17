package ast

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseNamedArgument(statementNode *ast_pb.Statement, argumentCtx *parser.NamedArgumentContext) *ast_pb.Argument {
	argument := &ast_pb.Argument{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(argumentCtx.GetStart().GetLine()),
			Column:      int64(argumentCtx.GetStart().GetColumn()),
			Start:       int64(argumentCtx.GetStart().GetStart()),
			End:         int64(argumentCtx.GetStop().GetStop()),
			Length:      int64(argumentCtx.GetStop().GetStop() - argumentCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		IsConstant:      false, // @TODO
		IsLValue:        false, // @TODO
		IsPure:          false, // @TODO
		LValueRequested: false, // @TODO
	}

	return argument
}

func (b *ASTBuilder) parseArgumentFromOrderComparasion(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, expressionCtx *parser.OrderComparisonContext) *ast_pb.Argument {
	argument := &ast_pb.Argument{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(expressionCtx.GetStart().GetLine()),
			Column:      int64(expressionCtx.GetStart().GetColumn()),
			Start:       int64(expressionCtx.GetStart().GetStart()),
			End:         int64(expressionCtx.GetStop().GetStop()),
			Length:      int64(expressionCtx.GetStop().GetStop() - expressionCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		// Comparison operators can end up only with boolean type.
		TypeDescriptions: &ast_pb.TypeDescriptions{
			TypeIdentifier: "t_bool",
			TypeString:     "bool",
		},
		IsConstant:      false, // @TODO
		IsLValue:        false, // @TODO
		IsPure:          false, // @TODO
		LValueRequested: false, // @TODO
		NodeType:        ast_pb.NodeType_BINARY_OPERATION,
	}

	if expressionCtx.GreaterThanOrEqual() != nil {
		argument.Operator = ast_pb.Operator_GREATER_THAN_OR_EQUAL
	} else if expressionCtx.LessThanOrEqual() != nil {
		argument.Operator = ast_pb.Operator_LESS_THAN_OR_EQUAL
	} else if expressionCtx.GreaterThan() != nil {
		argument.Operator = ast_pb.Operator_GREATER_THAN
	} else if expressionCtx.LessThan() != nil {
		argument.Operator = ast_pb.Operator_LESS_THAN
	} else {
		panic("Not implemented order comparison operator in parseArgumentFromOrderComparasion")
	}

	allExpressions := expressionCtx.AllExpression()

	if len(allExpressions) != 2 {
		panic("Not implemented order comparison with more than 2 expressions in parseArgumentFromOrderComparasion")
	}

	lE := allExpressions[0].(*parser.PrimaryExpressionContext)
	rE := allExpressions[1].(*parser.PrimaryExpressionContext)

	leftExpression := &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(lE.GetStart().GetLine()),
			Column:      int64(lE.GetStart().GetColumn()),
			Start:       int64(lE.GetStart().GetStart()),
			End:         int64(lE.GetStop().GetStop()),
			Length:      int64(lE.GetStop().GetStop() - lE.GetStart().GetStart() + 1),
			ParentIndex: argument.Id,
		},
		Name:     lE.GetText(),
		NodeType: ast_pb.NodeType_IDENTIFIER,
		// TODO: Fix this...
		OverloadedDeclarations: []int64{},
	}

	leftReferenceFound := false
	for _, statement := range bodyNode.GetStatements() {
		for _, declaration := range statement.GetDeclarations() {
			if declaration.GetName() == lE.GetText() {
				leftReferenceFound = true
				leftExpression.ReferencedDeclaration = declaration.Id
				leftExpression.TypeDescriptions = declaration.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	if !leftReferenceFound {
		for _, parameter := range fnNode.GetParameters().Parameters {
			if parameter.GetName() == rE.GetText() {
				leftReferenceFound = true
				leftExpression.ReferencedDeclaration = parameter.Id
				leftExpression.TypeDescriptions = parameter.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	argument.LeftExpression = leftExpression

	rightExpression := &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(rE.GetStart().GetLine()),
			Column:      int64(rE.GetStart().GetColumn()),
			Start:       int64(rE.GetStart().GetStart()),
			End:         int64(rE.GetStop().GetStop()),
			Length:      int64(rE.GetStop().GetStop() - rE.GetStart().GetStart() + 1),
			ParentIndex: argument.Id,
		},
		Name:     rE.GetText(),
		NodeType: ast_pb.NodeType_IDENTIFIER,
		// TODO: Fix this...
		OverloadedDeclarations: []int64{},
	}

	rightReferenceFound := false
	for _, statement := range bodyNode.GetStatements() {
		for _, declaration := range statement.GetDeclarations() {
			if declaration.GetName() == rE.GetText() {
				rightReferenceFound = true
				rightExpression.ReferencedDeclaration = declaration.Id
				rightExpression.TypeDescriptions = declaration.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	if !rightReferenceFound {
		for _, parameter := range fnNode.GetParameters().Parameters {
			if parameter.GetName() == rE.GetText() {
				rightReferenceFound = true
				rightExpression.ReferencedDeclaration = parameter.Id
				rightExpression.TypeDescriptions = parameter.GetTypeName().GetTypeDescriptions()
			}
		}
	}

	argument.RightExpression = rightExpression

	return argument
}

func (b *ASTBuilder) parseArgumentFromPrimaryExpression(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, expressionCtx *parser.PrimaryExpressionContext) *ast_pb.Argument {
	argument := &ast_pb.Argument{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(expressionCtx.GetStart().GetLine()),
			Column:      int64(expressionCtx.GetStart().GetColumn()),
			Start:       int64(expressionCtx.GetStart().GetStart()),
			End:         int64(expressionCtx.GetStop().GetStop()),
			Length:      int64(expressionCtx.GetStop().GetStop() - expressionCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		NodeType: ast_pb.NodeType_BINARY_OPERATION,
	}

	if expressionCtx.Literal() != nil {
		argument.IsPure = true
		argument.NodeType = ast_pb.NodeType_LITERAL
		argument.Value = strings.TrimSpace(
			// There can be hex 22 at beginning and end of literal.
			// We should drop it as that's ASCII for double quote.
			strings.ReplaceAll(expressionCtx.Literal().GetText(), "\"", ""),
		)
		argument.HexValue = hex.EncodeToString([]byte(argument.Value))

		argument.TypeDescriptions = &ast_pb.TypeDescriptions{
			TypeIdentifier: "t_string_literal",
			TypeString: fmt.Sprintf(
				"literal_string %s",
				expressionCtx.Literal().GetText(),
			),
		}

		return argument
	}

	panic("Not implemented primary expression in parseArgumentFromPrimaryExpression")

	return argument
}
