package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type BinaryOperationExpression struct {
	*ASTBuilder

	Id              int64           `json:"id"`
	IsConstant      bool            `json:"is_constant"`
	IsLValue        bool            `json:"is_l_value"`
	IsPure          bool            `json:"is_pure"`
	LValueRequested bool            `json:"l_value_requested"`
	NodeType        ast_pb.NodeType `json:"node_type"`
	Src             SrcNode         `json:"src"`
	Operator        ast_pb.Operator `json:"operator"`
	LeftExpression  Node[NodeType]  `json:"left_expression"`
	RightExpression Node[NodeType]  `json:"right_expression"`
}

func NewBinaryOperationExpression(b *ASTBuilder) *BinaryOperationExpression {
	return &BinaryOperationExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_BINARY_OPERATION,
	}
}

func (a *BinaryOperationExpression) GetId() int64 {
	return a.Id
}

func (a *BinaryOperationExpression) GetType() ast_pb.NodeType {
	return a.NodeType
}

func (a *BinaryOperationExpression) GetSrc() SrcNode {
	return a.Src
}

func (a *BinaryOperationExpression) GetOperator() ast_pb.Operator {
	return a.Operator
}

func (a *BinaryOperationExpression) GetLeftExpression() Node[NodeType] {
	return a.LeftExpression
}

func (a *BinaryOperationExpression) GetRightExpression() Node[NodeType] {
	return a.RightExpression
}

func (a *BinaryOperationExpression) GetTypeDescription() *TypeDescription {
	return a.LeftExpression.GetTypeDescription()
}

func (a *BinaryOperationExpression) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{a.LeftExpression, a.RightExpression}
}

func (a *BinaryOperationExpression) ToProto() NodeType {
	return ast_pb.BinaryOperationExpression{}
}

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
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(1),
	)

	return a
}

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
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(1),
	)

	return a
}

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
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(1),
	)

	return a
}

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
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(1),
	)

	return a
}
