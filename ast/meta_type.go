package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type MetaTypeExpression struct {
	*ASTBuilder

	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"node_type"`
	Name            string           `json:"name"`
	Src             SrcNode          `json:"src"`
	TypeDescription *TypeDescription `json:"type_description"`
}

func NewMetaTypeExpression(b *ASTBuilder) *MetaTypeExpression {
	return &MetaTypeExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IDENTIFIER,
	}
}

func (m *MetaTypeExpression) GetId() int64 {
	return m.Id
}

func (m *MetaTypeExpression) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *MetaTypeExpression) GetSrc() SrcNode {
	return m.Src
}

func (m *MetaTypeExpression) GetName() string {
	return m.Name
}

func (m *MetaTypeExpression) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

func (m *MetaTypeExpression) GetNodes() []Node[NodeType] {
	return nil
}

func (m *MetaTypeExpression) ToProto() NodeType {
	return ast_pb.MetaTypeExpression{}
}

func (m *MetaTypeExpression) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx *parser.MetaTypeContext,
) Node[NodeType] {
	m.Src = SrcNode{
		Id:     m.GetNextID(),
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

	m.Name = ctx.Type().GetText()
	m.TypeDescription = &TypeDescription{
		TypeString: ctx.Type().GetText(),
	}

	return m
}
