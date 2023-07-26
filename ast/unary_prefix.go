package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type UnaryPrefixExpression struct {
	*ASTBuilder

	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"node_type"`
	Src             SrcNode          `json:"src"`
	Operator        ast_pb.Operator  `json:"operator"`
	Expression      Node[NodeType]   `json:"expression"`
	TypeDescription *TypeDescription `json:"type_description"`
	Prefix          bool             `json:"prefix"`
	IsConstant      bool             `json:"is_constant"`
	IsLValue        bool             `json:"is_l_value"`
	IsPure          bool             `json:"is_pure"`
	LValueRequested bool             `json:"l_value_requested"`
}

func NewUnaryPrefixExpression(b *ASTBuilder) *UnaryPrefixExpression {
	return &UnaryPrefixExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_UNARY_OPERATION,
	}
}

func (u *UnaryPrefixExpression) GetId() int64 {
	return u.Id
}

func (u *UnaryPrefixExpression) GetType() ast_pb.NodeType {
	return u.NodeType
}

func (u *UnaryPrefixExpression) GetSrc() SrcNode {
	return u.Src
}

func (u *UnaryPrefixExpression) GetOperator() ast_pb.Operator {
	return u.Operator
}

func (u *UnaryPrefixExpression) GetExpression() Node[NodeType] {
	return u.Expression
}

func (u *UnaryPrefixExpression) GetTypeDescription() *TypeDescription {
	return u.TypeDescription
}

func (u *UnaryPrefixExpression) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{u.Expression}
}

func (u *UnaryPrefixExpression) GetPrefix() bool {
	return u.Prefix
}

func (u *UnaryPrefixExpression) GetIsConstant() bool {
	return u.IsConstant
}

func (u *UnaryPrefixExpression) GetIsLValue() bool {
	return u.IsLValue
}

func (u *UnaryPrefixExpression) GetIsPure() bool {
	return u.IsPure
}

func (u *UnaryPrefixExpression) GetLValueRequested() bool {
	return u.LValueRequested
}

func (u *UnaryPrefixExpression) ToProto() NodeType {
	return &ast_pb.UnaryPrefixOperator{}
}

func (u *UnaryPrefixExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.UnaryPrefixOperationContext,
) Node[NodeType] {
	u.Src = SrcNode{
		Id:     u.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if fnNode != nil {
				return fnNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	u.Operator = ast_pb.Operator_INCREMENT
	if ctx.Dec() != nil {
		u.Operator = ast_pb.Operator_DECREMENT
	} else if ctx.Not() != nil {
		u.Operator = ast_pb.Operator_NOT
	} else if ctx.BitNot() != nil {
		u.Operator = ast_pb.Operator_BIT_NOT
	} else if ctx.Sub() != nil {
		u.Operator = ast_pb.Operator_SUBTRACT
	}

	expression := NewExpression(u.ASTBuilder)
	u.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, u, ctx.Expression())
	return u
}
