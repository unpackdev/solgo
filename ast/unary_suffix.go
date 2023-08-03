package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type UnarySuffix struct {
	*ASTBuilder

	Id                    int64            `json:"id"`
	NodeType              ast_pb.NodeType  `json:"node_type"`
	Src                   SrcNode          `json:"src"`
	Operator              ast_pb.Operator  `json:"operator"`
	Expression            Node[NodeType]   `json:"expression"`
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription `json:"type_description"`
	Prefix                bool             `json:"prefix"`
	Constant              bool             `json:"is_constant"`
	LValue                bool             `json:"is_l_value"`
	Pure                  bool             `json:"is_pure"`
	LValueRequested       bool             `json:"l_value_requested"`
}

func NewUnarySuffixExpression(b *ASTBuilder) *UnarySuffix {
	return &UnarySuffix{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_UNARY_OPERATION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the UnarySuffix node.
func (u *UnarySuffix) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	u.ReferencedDeclaration = refId
	u.TypeDescription = refDesc
	return false
}

func (u *UnarySuffix) GetId() int64 {
	return u.Id
}

func (u *UnarySuffix) GetType() ast_pb.NodeType {
	return u.NodeType
}

func (u *UnarySuffix) GetSrc() SrcNode {
	return u.Src
}

func (u *UnarySuffix) GetOperator() ast_pb.Operator {
	return u.Operator
}

func (u *UnarySuffix) GetExpression() Node[NodeType] {
	return u.Expression
}

func (u *UnarySuffix) GetTypeDescription() *TypeDescription {
	return u.TypeDescription
}

func (u *UnarySuffix) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{u.Expression}
}

func (u *UnarySuffix) GetPrefix() bool {
	return u.Prefix
}

func (u *UnarySuffix) IsConstant() bool {
	return u.Constant
}

func (u *UnarySuffix) IsLValue() bool {
	return u.LValue
}

func (u *UnarySuffix) IsPure() bool {
	return u.Pure
}

func (u *UnarySuffix) IsLValueRequested() bool {
	return u.LValueRequested
}

func (u *UnarySuffix) GetReferencedDeclaration() int64 {
	return u.ReferencedDeclaration
}

func (u *UnarySuffix) ToProto() NodeType {
	proto := ast_pb.UnarySuffix{
		Id:                    u.GetId(),
		NodeType:              u.GetType(),
		Src:                   u.GetSrc().ToProto(),
		Operator:              u.GetOperator(),
		Prefix:                u.GetPrefix(),
		IsConstant:            u.IsConstant(),
		IsLValue:              u.IsLValue(),
		IsPure:                u.IsPure(),
		LValueRequested:       u.IsLValueRequested(),
		ReferencedDeclaration: u.GetReferencedDeclaration(),
		Expression:            u.GetExpression().ToProto().(*v3.TypedStruct),
		TypeDescription:       u.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "UnarySuffix")
}

func (u *UnarySuffix) Parse(
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
	u.TypeDescription = u.Expression.GetTypeDescription()
	return u
}
