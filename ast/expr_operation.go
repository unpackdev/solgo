package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ExprOperation struct {
	*ASTBuilder

	Id               int64              `json:"id"`
	NodeType         ast_pb.NodeType    `json:"node_type"`
	Src              SrcNode            `json:"src"`
	LeftExpression   Node[NodeType]     `json:"left_expression"`
	RightExpression  Node[NodeType]     `json:"right_expression"`
	TypeDescriptions []*TypeDescription `json:"type_descriptions"`
}

func NewExprOperationExpression(b *ASTBuilder) *ExprOperation {
	return &ExprOperation{
		ASTBuilder:       b,
		Id:               b.GetNextID(),
		NodeType:         ast_pb.NodeType_EXPRESSION_OPERATION,
		TypeDescriptions: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ExprOperation node.
func (b *ExprOperation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (f *ExprOperation) GetId() int64 {
	return f.Id
}

func (f *ExprOperation) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *ExprOperation) GetSrc() SrcNode {
	return f.Src
}

func (f *ExprOperation) GetTypeDescription() *TypeDescription {
	return f.TypeDescriptions[0]
}

func (f *ExprOperation) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{f.LeftExpression, f.RightExpression}
}

func (f *ExprOperation) GetLeftExpression() Node[NodeType] {
	return f.LeftExpression
}

func (f *ExprOperation) GetRightExpression() Node[NodeType] {
	return f.RightExpression
}

func (f *ExprOperation) ToProto() NodeType {
	proto := ast_pb.ExprOperation{
		Id:              f.GetId(),
		NodeType:        f.GetType(),
		Src:             f.GetSrc().ToProto(),
		LeftExpression:  f.GetLeftExpression().ToProto().(*v3.TypedStruct),
		RightExpression: f.GetRightExpression().ToProto().(*v3.TypedStruct),
		TypeDescription: f.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "ExprOperation")
}

func (f *ExprOperation) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.ExpOperationContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Src = SrcNode{
		Id:     f.GetNextID(),
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

			if bodyNode != nil {
				return bodyNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return contractNode.GetId()
		}(),
	}

	expression := NewExpression(f.ASTBuilder)

	f.LeftExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, ctx.Expression(0))
	f.RightExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, ctx.Expression(1))

	f.TypeDescriptions = append(f.TypeDescriptions, f.LeftExpression.GetTypeDescription())
	f.TypeDescriptions = append(f.TypeDescriptions, f.RightExpression.GetTypeDescription())

	return f
}
