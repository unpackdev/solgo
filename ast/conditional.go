package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Conditional struct {
	*ASTBuilder

	Id               int64              `json:"id"`
	NodeType         ast_pb.NodeType    `json:"node_type"`
	Src              SrcNode            `json:"src"`
	Expressions      []Node[NodeType]   `json:"right_expression"`
	TypeDescriptions []*TypeDescription `json:"type_descriptions"`
}

func NewConditionalExpression(b *ASTBuilder) *Conditional {
	return &Conditional{
		ASTBuilder:       b,
		Id:               b.GetNextID(),
		NodeType:         ast_pb.NodeType_CONDITIONAL_EXPRESSION,
		TypeDescriptions: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Conditional node.
func (b *Conditional) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (f *Conditional) GetId() int64 {
	return f.Id
}

func (f *Conditional) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Conditional) GetSrc() SrcNode {
	return f.Src
}

func (f *Conditional) GetTypeDescription() *TypeDescription {
	return f.TypeDescriptions[0]
}

func (f *Conditional) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	for _, exp := range f.Expressions {
		toReturn = append(toReturn, exp)
	}
	return toReturn
}

func (f *Conditional) GetExpressions() []Node[NodeType] {
	return f.Expressions
}

func (f *Conditional) ToProto() NodeType {
	/* 	proto := ast_pb.Conditional{
	   		Id:              f.GetId(),
	   		NodeType:        f.GetType(),
	   		Src:             f.GetSrc().ToProto(),
	   		LeftExpression:  f.GetLeftExpression().ToProto().(*v3.TypedStruct),
	   		RightExpression: f.GetRightExpression().ToProto().(*v3.TypedStruct),
	   		TypeDescription: f.GetTypeDescription().ToProto(),
	   	}
	*/
	return NewTypedStruct(nil, "Conditional")
}

func (f *Conditional) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.ConditionalContext,
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

	for _, expr := range ctx.AllExpression() {
		parsedExp := expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, expr)
		f.Expressions = append(
			f.Expressions,
			parsedExp,
		)
		f.TypeDescriptions = append(f.TypeDescriptions, parsedExp.GetTypeDescription())
	}

	return f
}
