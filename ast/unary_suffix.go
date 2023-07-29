package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type UnarySuffixExpression struct {
	*ASTBuilder

	Id                    int64            `json:"id"`
	NodeType              ast_pb.NodeType  `json:"node_type"`
	Src                   SrcNode          `json:"src"`
	Operator              ast_pb.Operator  `json:"operator"`
	Expression            Node[NodeType]   `json:"expression"`
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription `json:"type_description"`
	Prefix                bool             `json:"prefix"`
	IsConstant            bool             `json:"is_constant"`
	IsLValue              bool             `json:"is_l_value"`
	IsPure                bool             `json:"is_pure"`
	LValueRequested       bool             `json:"l_value_requested"`
}

func NewUnarySuffixExpression(b *ASTBuilder) *UnarySuffixExpression {
	return &UnarySuffixExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_UNARY_OPERATION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the UnarySuffixExpression node.
func (u *UnarySuffixExpression) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	u.ReferencedDeclaration = refId
	u.TypeDescription = refDesc
	return false
}

func (u *UnarySuffixExpression) GetId() int64 {
	return u.Id
}

func (u *UnarySuffixExpression) GetType() ast_pb.NodeType {
	return u.NodeType
}

func (u *UnarySuffixExpression) GetSrc() SrcNode {
	return u.Src
}

func (u *UnarySuffixExpression) GetOperator() ast_pb.Operator {
	return u.Operator
}

func (u *UnarySuffixExpression) GetExpression() Node[NodeType] {
	return u.Expression
}

func (u *UnarySuffixExpression) GetTypeDescription() *TypeDescription {
	return u.TypeDescription
}

func (u *UnarySuffixExpression) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{u.Expression}
}

func (u *UnarySuffixExpression) GetPrefix() bool {
	return u.Prefix
}

func (u *UnarySuffixExpression) GetIsConstant() bool {
	return u.IsConstant
}

func (u *UnarySuffixExpression) GetIsLValue() bool {
	return u.IsLValue
}

func (u *UnarySuffixExpression) GetIsPure() bool {
	return u.IsPure
}

func (u *UnarySuffixExpression) GetLValueRequested() bool {
	return u.LValueRequested
}

func (u *UnarySuffixExpression) ToProto() NodeType {
	return &ast_pb.UnarySuffixOperator{}
}

func (u *UnarySuffixExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.UnarySuffixOperationContext,
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
	}

	expression := NewExpression(u.ASTBuilder)
	u.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, u, ctx.Expression())
	return u
}
