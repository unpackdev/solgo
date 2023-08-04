package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ModifierName struct {
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func (m *ModifierName) ToProto() *ast_pb.ModifierName {
	return &ast_pb.ModifierName{
		Id:       m.Id,
		Name:     m.Name,
		NodeType: m.NodeType,
		Src:      m.Src.ToProto(),
	}
}

type ModifierInvocation struct {
	*ASTBuilder

	Id            int64              `json:"id"`
	Name          string             `json:"name"`
	NodeType      ast_pb.NodeType    `json:"node_type"`
	Kind          ast_pb.NodeType    `json:"kind"`
	Src           SrcNode            `json:"src"`
	ArgumentTypes []*TypeDescription `json:"argument_types"`
	Arguments     []Node[NodeType]   `json:"arguments"`
	ModifierName  *ModifierName      `json:"modifier_name,omitempty"`
}

func NewModifierInvocation(b *ASTBuilder) *ModifierInvocation {
	return &ModifierInvocation{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_MODIFIER_INVOCATION,
		Kind:          ast_pb.NodeType_MODIFIER_INVOCATION,
		Arguments:     make([]Node[NodeType], 0),
		ArgumentTypes: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ModifierInvocation node.
func (m *ModifierInvocation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (m *ModifierInvocation) GetId() int64 {
	return m.Id
}

func (m *ModifierInvocation) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *ModifierInvocation) GetKind() ast_pb.NodeType {
	return m.Kind
}

func (m *ModifierInvocation) GetSrc() SrcNode {
	return m.Src
}

func (m *ModifierInvocation) GetName() string {
	return m.Name
}

func (m *ModifierInvocation) GetTypeDescription() *TypeDescription {
	return nil
}

func (m *ModifierInvocation) GetNodes() []Node[NodeType] {
	return m.Arguments
}

func (m *ModifierInvocation) GetArguments() []Node[NodeType] {
	return m.Arguments
}

func (m *ModifierInvocation) GetArgumentTypes() []*TypeDescription {
	return m.ArgumentTypes
}

func (m *ModifierInvocation) ToProto() NodeType {
	toReturn := &ast_pb.ModifierInvocation{
		Id:            m.GetId(),
		Name:          m.GetName(),
		NodeType:      m.GetType(),
		Kind:          m.GetKind(),
		Src:           m.Src.ToProto(),
		ModifierName:  m.ModifierName.ToProto(),
		Arguments:     make([]*v3.TypedStruct, 0),
		ArgumentTypes: []*ast_pb.TypeDescription{},
	}

	for _, arg := range m.GetArguments() {
		toReturn.Arguments = append(
			toReturn.Arguments,
			arg.ToProto().(*v3.TypedStruct),
		)

		toReturn.ArgumentTypes = append(
			toReturn.ArgumentTypes,
			arg.GetTypeDescription().ToProto(),
		)
	}

	return toReturn
}

func (m *ModifierInvocation) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx parser.IModifierInvocationContext,
) {
	m.Src = SrcNode{
		Id:          m.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}
	m.NodeType = ast_pb.NodeType_MODIFIER_INVOCATION
	m.Arguments = make([]Node[NodeType], 0)

	if ctx.IdentifierPath() != nil {
		iCtx := ctx.IdentifierPath()
		m.ModifierName = &ModifierName{
			Id:   m.GetNextID(),
			Name: iCtx.GetText(),
			Src: SrcNode{
				Id:          m.GetNextID(),
				Line:        int64(iCtx.GetStart().GetLine()),
				Column:      int64(iCtx.GetStart().GetColumn()),
				Start:       int64(iCtx.GetStart().GetStart()),
				End:         int64(iCtx.GetStop().GetStop()),
				Length:      int64(iCtx.GetStop().GetStop() - iCtx.GetStart().GetStart() + 1),
				ParentIndex: m.GetId(),
			},
		}
		m.Name = m.ModifierName.Name
	}

	expression := NewExpression(m.ASTBuilder)
	if ctx.CallArgumentList() != nil {
		for _, expressionCtx := range ctx.CallArgumentList().AllExpression() {
			expr := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, m, expressionCtx)
			m.Arguments = append(
				m.Arguments,
				expr,
			)

			m.ArgumentTypes = append(
				m.ArgumentTypes,
				expr.GetTypeDescription(),
			)
		}
	}
}