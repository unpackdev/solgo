package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// MetaType represents a meta-type node in the abstract syntax tree.
type MetaType struct {
	*ASTBuilder

	Id                    int64            `json:"id"`                               // Unique identifier of the meta-type node.
	NodeType              ast_pb.NodeType  `json:"node_type"`                        // Type of the node.
	Name                  string           `json:"name"`                             // Name of the meta-type.
	Src                   SrcNode          `json:"src"`                              // Source location information.
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"` // Referenced declaration identifier.
	TypeDescription       *TypeDescription `json:"type_description"`                 // Type description of the meta-type.
}

// NewMetaTypeExpression creates a new instance of MetaType with the provided ASTBuilder.
func NewMetaTypeExpression(b *ASTBuilder) *MetaType {
	return &MetaType{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IDENTIFIER,
	}
}

// SetReferenceDescriptor sets the reference descriptors of the MetaType node.
func (m *MetaType) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	m.ReferencedDeclaration = refId
	m.TypeDescription = refDesc
	return false
}

// GetId returns the unique identifier of the meta-type node.
func (m *MetaType) GetId() int64 {
	return m.Id
}

// GetType returns the type of the node.
func (m *MetaType) GetType() ast_pb.NodeType {
	return m.NodeType
}

// GetSrc returns the source location information of the meta-type node.
func (m *MetaType) GetSrc() SrcNode {
	return m.Src
}

// GetName returns the name of the meta-type.
func (m *MetaType) GetName() string {
	return m.Name
}

// GetTypeDescription returns the type description of the meta-type.
func (m *MetaType) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

// GetNodes returns a slice of nodes associated with the meta-type.
func (m *MetaType) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetReferencedDeclaration returns the referenced declaration identifier of the meta-type.
func (m *MetaType) GetReferencedDeclaration() int64 {
	return m.ReferencedDeclaration
}

// ToProto converts the MetaType node to its corresponding protobuf representation.
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

// Parse parses the meta-type context and populates the MetaType fields.
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
