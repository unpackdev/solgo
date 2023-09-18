package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// UnarySuffix represents a unary operation applied as a suffix to an expression.
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

// NewUnarySuffixExpression creates a new UnarySuffix instance with the given ASTBuilder.
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

// GetId returns the unique identifier of the UnarySuffix.
func (u *UnarySuffix) GetId() int64 {
	return u.Id
}

// GetType returns the node type of the UnarySuffix.
func (u *UnarySuffix) GetType() ast_pb.NodeType {
	return u.NodeType
}

// GetSrc returns the source location information of the UnarySuffix.
func (u *UnarySuffix) GetSrc() SrcNode {
	return u.Src
}

// GetOperator returns the unary operator applied to the expression.
func (u *UnarySuffix) GetOperator() ast_pb.Operator {
	return u.Operator
}

// GetExpression returns the expression to which the unary operation is applied.
func (u *UnarySuffix) GetExpression() Node[NodeType] {
	return u.Expression
}

// GetTypeDescription returns the type description associated with the UnarySuffix.
func (u *UnarySuffix) GetTypeDescription() *TypeDescription {
	return u.TypeDescription
}

// GetNodes returns a list of child nodes for traversal within the UnarySuffix.
func (u *UnarySuffix) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{u.Expression}
}

// GetPrefix returns true if the unary operation is a prefix operation.
func (u *UnarySuffix) GetPrefix() bool {
	return u.Prefix
}

// IsConstant returns true if the operation's result is a constant.
func (u *UnarySuffix) IsConstant() bool {
	return u.Constant
}

// IsLValue returns true if the expression is an l-value.
func (u *UnarySuffix) IsLValue() bool {
	return u.LValue
}

// IsPure returns true if the operation is pure, i.e., it doesn't modify state.
func (u *UnarySuffix) IsPure() bool {
	return u.Pure
}

// IsLValueRequested returns true if an l-value is requested from the operation.
func (u *UnarySuffix) IsLValueRequested() bool {
	return u.LValueRequested
}

// GetReferencedDeclaration returns the referenced declaration of the UnarySuffix.
func (u *UnarySuffix) GetReferencedDeclaration() int64 {
	return u.ReferencedDeclaration
}

// ToProto converts the UnarySuffix instance to its corresponding protocol buffer representation.
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

// Parse populates the UnarySuffix instance with information parsed from the provided contexts.
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
