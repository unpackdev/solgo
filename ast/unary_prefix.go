package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type UnaryPrefix struct {
	*ASTBuilder

	Id                    int64            `json:"id"`
	NodeType              ast_pb.NodeType  `json:"node_type"`
	Src                   SrcNode          `json:"src"`
	Operator              ast_pb.Operator  `json:"operator"`
	Prefix                bool             `json:"prefix"`
	Constant              bool             `json:"is_constant"`
	LValue                bool             `json:"is_l_value"`
	Pure                  bool             `json:"is_pure"`
	LValueRequested       bool             `json:"l_value_requested"`
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"`
	Expression            Node[NodeType]   `json:"expression"`
	TypeDescription       *TypeDescription `json:"type_description"`
}

func NewUnaryPrefixExpression(b *ASTBuilder) *UnaryPrefix {
	return &UnaryPrefix{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_UNARY_OPERATION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the UnaryPrefix node.
func (u *UnaryPrefix) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	u.ReferencedDeclaration = refId
	u.TypeDescription = refDesc
	return false
}

func (u *UnaryPrefix) GetId() int64 {
	return u.Id
}

func (u *UnaryPrefix) GetType() ast_pb.NodeType {
	return u.NodeType
}

func (u *UnaryPrefix) GetSrc() SrcNode {
	return u.Src
}

func (u *UnaryPrefix) GetOperator() ast_pb.Operator {
	return u.Operator
}

func (u *UnaryPrefix) GetExpression() Node[NodeType] {
	return u.Expression
}

func (u *UnaryPrefix) GetTypeDescription() *TypeDescription {
	return u.TypeDescription
}

func (u *UnaryPrefix) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{u.Expression}
}

func (u *UnaryPrefix) GetPrefix() bool {
	return u.Prefix
}

func (u *UnaryPrefix) IsConstant() bool {
	return u.Constant
}

func (u *UnaryPrefix) IsLValue() bool {
	return u.LValue
}

func (u *UnaryPrefix) IsPure() bool {
	return u.Pure
}

func (u *UnaryPrefix) IsLValueRequested() bool {
	return u.LValueRequested
}

func (u *UnaryPrefix) GetReferencedDeclaration() int64 {
	return u.ReferencedDeclaration
}

func (u *UnaryPrefix) ToProto() NodeType {
	if u.GetTypeDescription() == nil {
		u.dumpNode(u)
	}
	proto := ast_pb.UnaryPrefix{
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

	return NewTypedStruct(&proto, "UnaryPrefix")
}

func (u *UnaryPrefix) Parse(
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
	u.TypeDescription = u.Expression.GetTypeDescription()
	return u
}
