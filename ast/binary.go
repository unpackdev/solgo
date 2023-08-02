package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// BinaryOperationExpression represents a binary operation in a Solidity source file.
// A binary operation is an operation that operates on two operands like +, -, *, / etc.
type BinaryOperationExpression struct {
	*ASTBuilder // Embedding the ASTBuilder to provide common functionality across all AST nodes.

	// Id is the unique identifier of the binary operation.
	Id int64 `json:"id"`
	// IsConstant indicates whether the binary operation is a constant.
	IsConstant bool `json:"is_constant"`
	// IsPure indicates whether the binary operation is pure (i.e., it does not read or modify state).
	IsPure bool `json:"is_pure"`
	// NodeType is the type of the node.
	// For a BinaryOperationExpression, this is always NodeType_BINARY_OPERATION.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Src contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// Operator is the operator of the binary operation.
	Operator ast_pb.Operator `json:"operator"`
	// LeftExpression is the left operand of the binary operation.
	LeftExpression Node[NodeType] `json:"left_expression"`
	// RightExpression is the right operand of the binary operation.
	RightExpression Node[NodeType] `json:"right_expression"`
}

// NewBinaryOperationExpression is a constructor function that initializes a new BinaryOperationExpression with a unique ID and the NodeType set to NodeType_BINARY_OPERATION.
func NewBinaryOperationExpression(b *ASTBuilder) *BinaryOperationExpression {
	return &BinaryOperationExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_BINARY_OPERATION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the BinaryOperationExpression node.
func (a *BinaryOperationExpression) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId is a getter method that returns the unique identifier of the binary operation.
func (a *BinaryOperationExpression) GetId() int64 {
	return a.Id
}

// GetType is a getter method that returns the node type of the binary operation.
func (a *BinaryOperationExpression) GetType() ast_pb.NodeType {
	return a.NodeType
}

// GetSrc is a getter method that returns the source information of the binary operation.
func (a *BinaryOperationExpression) GetSrc() SrcNode {
	return a.Src
}

// GetOperator is a getter method that returns the operator of the binary operation.
func (a *BinaryOperationExpression) GetOperator() ast_pb.Operator {
	return a.Operator
}

// GetLeftExpression is a getter method that returns the left operand of the binary operation.
func (a *BinaryOperationExpression) GetLeftExpression() Node[NodeType] {
	return a.LeftExpression
}

// GetRightExpression is a getter method that returns the right operand of the binary operation.
func (a *BinaryOperationExpression) GetRightExpression() Node[NodeType] {
	return a.RightExpression
}

// GetTypeDescription is a getter method that returns the type description of the left operand of the binary operation.
func (a *BinaryOperationExpression) GetTypeDescription() *TypeDescription {
	return a.LeftExpression.GetTypeDescription()
}

// GetNodes is a getter method that returns a slice of the operands of the binary operation.
func (a *BinaryOperationExpression) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{a.LeftExpression, a.RightExpression}
}

// ToProto is a method that returns the protobuf representation of the binary operation.
func (a *BinaryOperationExpression) ToProto() NodeType {
	return ast_pb.BinaryOperationExpression{}
}

// ParseAddSub is a method that parses addition and subtraction operations.
func (a *BinaryOperationExpression) ParseAddSub(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.AddSubOperationContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if expNode != nil {
				return expNode.GetId()
			}

			if vDeclar != nil {
				return vDeclar.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	a.Operator = ast_pb.Operator_ADDITION
	if ctx.Sub() != nil {
		a.Operator = ast_pb.Operator_SUBTRACTION
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	return a
}

// ParseOrderComparison is a method that parses order comparison operations.
func (a *BinaryOperationExpression) ParseOrderComparison(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.OrderComparisonContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.GreaterThanOrEqual() != nil {
		a.Operator = ast_pb.Operator_GREATER_THAN_OR_EQUAL
	} else if ctx.LessThanOrEqual() != nil {
		a.Operator = ast_pb.Operator_LESS_THAN_OR_EQUAL
	} else if ctx.GreaterThan() != nil {
		a.Operator = ast_pb.Operator_GREATER_THAN
	} else if ctx.LessThan() != nil {
		a.Operator = ast_pb.Operator_LESS_THAN
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	return a
}

// ParseMulDivMod is a method that parses multiplication, division, and modulo operations.
func (a *BinaryOperationExpression) ParseMulDivMod(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.MulDivModOperationContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.Mul() != nil {
		a.Operator = ast_pb.Operator_MULTIPLICATION
	} else if ctx.Div() != nil {
		a.Operator = ast_pb.Operator_DIVISION
	} else if ctx.Mod() != nil {
		a.Operator = ast_pb.Operator_MODULO
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	return a
}

// ParseEqualityComparison is a method that parses equality comparison operations.
func (a *BinaryOperationExpression) ParseEqualityComparison(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.EqualityComparisonContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.Equal() != nil {
		a.Operator = ast_pb.Operator_EQUAL
	} else if ctx.NotEqual() != nil {
		a.Operator = ast_pb.Operator_NOT_EQUAL
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	return a
}

func (a *BinaryOperationExpression) ParseOr(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.OrOperationContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if expNode != nil {
				return expNode.GetId()
			}

			if vDeclar != nil {
				return vDeclar.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	a.Operator = ast_pb.Operator_OR

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	return a
}
