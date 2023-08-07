package ast

import (
	"fmt"
	"strings"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type PayableConversion struct {
	*ASTBuilder

	Id                    int64              `json:"id"`
	NodeType              ast_pb.NodeType    `json:"node_type"`
	Src                   SrcNode            `json:"src"`
	Arguments             []Node[NodeType]   `json:"arguments"`
	ArgumentTypes         []*TypeDescription `json:"argument_types"`
	ReferencedDeclaration int64              `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription   `json:"type_description"`
	Payable               bool               `json:"payable"`
}

func NewPayableConversionExpression(b *ASTBuilder) *PayableConversion {
	return &PayableConversion{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_PAYABLE_CONVERSION,
		Arguments:     make([]Node[NodeType], 0),
		ArgumentTypes: []*TypeDescription{},
	}
}

// SetReferenceDescriptor sets the reference descriptions of the PayableConversion node.
func (p *PayableConversion) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	p.ReferencedDeclaration = refId
	p.TypeDescription = refDesc
	return false
}

func (p *PayableConversion) GetId() int64 {
	return p.Id
}

func (p *PayableConversion) GetType() ast_pb.NodeType {
	return p.NodeType
}

func (p *PayableConversion) GetSrc() SrcNode {
	return p.Src
}

func (p *PayableConversion) GetTypeDescription() *TypeDescription {
	return p.TypeDescription
}

func (p *PayableConversion) GetArgumentTypes() []*TypeDescription {
	return p.ArgumentTypes
}

func (p *PayableConversion) GetArguments() []Node[NodeType] {
	return p.Arguments
}

func (p *PayableConversion) IsPayable() bool {
	return p.Payable
}

func (p *PayableConversion) GetNodes() []Node[NodeType] {
	return p.Arguments
}

func (p *PayableConversion) GetReferencedDeclaration() int64 {
	return p.ReferencedDeclaration
}

func (p *PayableConversion) ToProto() NodeType {
	proto := ast_pb.PayableConversion{
		Id:                    p.GetId(),
		Src:                   p.GetSrc().ToProto(),
		NodeType:              p.GetType(),
		Payable:               p.IsPayable(),
		ReferencedDeclaration: p.GetReferencedDeclaration(),
		ArgumentTypes:         make([]*ast_pb.TypeDescription, 0),
		Arguments:             make([]*v3.TypedStruct, 0),
	}

	for _, arg := range p.GetArgumentTypes() {
		proto.ArgumentTypes = append(proto.ArgumentTypes, arg.ToProto())
	}

	for _, arg := range p.GetArguments() {
		proto.Arguments = append(proto.Arguments, arg.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "PayableConversion")
}

func (p *PayableConversion) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx *parser.PayableConversionContext,
) Node[NodeType] {
	p.Src = SrcNode{
		Id:     p.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if exprNode != nil {
				return exprNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}
	p.Payable = ctx.Payable() != nil

	expression := NewExpression(p.ASTBuilder)

	typeStrings := []string{}
	typeIdentifiers := []string{}

	if ctx.CallArgumentList() != nil {
		for _, expressionCtx := range ctx.CallArgumentList().AllExpression() {
			expr := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, p, expressionCtx)
			p.Arguments = append(
				p.Arguments,
				expr,
			)

			typeStrings = append(typeStrings, expr.GetTypeDescription().TypeString)
			typeIdentifiers = append(typeIdentifiers, expr.GetTypeDescription().TypeIdentifier)

			p.ArgumentTypes = append(
				p.ArgumentTypes,
				expr.GetTypeDescription(),
			)
		}
	}

	p.TypeDescription = &TypeDescription{
		TypeString: func() string {
			return fmt.Sprintf(
				"function(%s) payable",
				strings.Join(typeStrings, ","),
			)
		}(),
		TypeIdentifier: func() string {
			return fmt.Sprintf(
				"t_function_payable$_%s$",
				strings.Join(typeIdentifiers, "$_"),
			)
		}(),
	}

	return p
}
