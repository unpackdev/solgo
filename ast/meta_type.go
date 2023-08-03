package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type MetaType struct {
	*ASTBuilder

	Id                    int64            `json:"id"`
	NodeType              ast_pb.NodeType  `json:"node_type"`
	Name                  string           `json:"name"`
	Src                   SrcNode          `json:"src"`
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription `json:"type_description"`
}

func NewMetaTypeExpression(b *ASTBuilder) *MetaType {
	return &MetaType{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IDENTIFIER,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the MetaType node.
func (m *MetaType) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	m.ReferencedDeclaration = refId
	m.TypeDescription = refDesc
	return false
}

func (m *MetaType) GetId() int64 {
	return m.Id
}

func (m *MetaType) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *MetaType) GetSrc() SrcNode {
	return m.Src
}

func (m *MetaType) GetName() string {
	return m.Name
}

func (m *MetaType) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

func (m *MetaType) GetNodes() []Node[NodeType] {
	return nil
}

func (m *MetaType) GetReferencedDeclaration() int64 {
	return m.ReferencedDeclaration
}

func (m *MetaType) ToProto() NodeType {
	proto := ast_pb.MetaType{
		Id:                    m.GetId(),
		Name:                  m.GetName(),
		NodeType:              m.GetType(),
		Src:                   m.GetSrc().ToProto(),
		ReferencedDeclaration: m.GetReferencedDeclaration(),
		TypeDescription:       m.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "MetaType")
}

func (m *MetaType) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]],
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
