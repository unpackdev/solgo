package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ModifierDefinition struct {
	*ASTBuilder

	Id         int64             `json:"id"`
	Name       string            `json:"name"`
	NodeType   ast_pb.NodeType   `json:"node_type"`
	Src        SrcNode           `json:"src"`
	Visibility ast_pb.Visibility `json:"visibility"`
	Virtual    bool              `json:"virtual"`
	Parameters *ParameterList    `json:"parameters"`
	Body       *BodyNode         `json:"body"`
}

func NewModifierDefinition(b *ASTBuilder) *ModifierDefinition {
	return &ModifierDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_MODIFIER_DEFINITION,
		Visibility: ast_pb.Visibility_INTERNAL,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ModifierDefinition node.
func (m *ModifierDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (m *ModifierDefinition) GetId() int64 {
	return m.Id
}

func (m *ModifierDefinition) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *ModifierDefinition) GetSrc() SrcNode {
	return m.Src
}

func (m *ModifierDefinition) GetName() string {
	return m.Name
}

func (m *ModifierDefinition) GetTypeDescription() *TypeDescription {
	return nil
}

func (m *ModifierDefinition) GetNodes() []Node[NodeType] {
	return m.Body.GetNodes()
}

func (m *ModifierDefinition) IsVirtual() bool {
	return m.Virtual
}

func (m *ModifierDefinition) GetVisibility() ast_pb.Visibility {
	return m.Visibility
}

func (m *ModifierDefinition) GetParameters() *ParameterList {
	return m.Parameters
}

func (m *ModifierDefinition) GetBody() *BodyNode {
	return m.Body
}

func (m *ModifierDefinition) ToProto() NodeType {
	proto := ast_pb.Modifier{
		Id:         m.GetId(),
		Name:       m.GetName(),
		NodeType:   m.GetType(),
		Src:        m.GetSrc().ToProto(),
		Virtual:    m.IsVirtual(),
		Visibility: m.GetVisibility(),
		Parameters: m.GetParameters().ToProto(),
		Body:       m.GetBody().ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&proto, "Modifier")
}

func (m *ModifierDefinition) ParseDefinition(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.ModifierDefinitionContext,
) Node[NodeType] {
	m.Src = SrcNode{
		Id:          m.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	m.Name = ctx.Identifier().GetText()

	if ctx.AllVirtual() != nil {
		for _, virtualCtx := range ctx.AllVirtual() {
			if virtualCtx.GetText() == "virtual" {
				m.Virtual = true
			}
		}
	}

	parameters := NewParameterList(m.ASTBuilder)
	if ctx.ParameterList() != nil {
		parameters.Parse(unit, contractNode, ctx.ParameterList())
	} else {
		parameters.Src = m.Src
	}
	m.Parameters = parameters

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(m.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, m, ctx.Block())
		m.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(m.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, m, uncheckedCtx)
				m.Body.Statements = append(m.Body.Statements, bodyNode)
			}
		}
	}

	m.currentModifiers = append(m.currentModifiers, m)
	return m
}
