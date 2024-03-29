package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ModifierName represents the name of a modifier in the abstract syntax tree.
type ModifierName struct {
	Id       int64           `json:"id"`        // Unique identifier of the modifier name node.
	Name     string          `json:"name"`      // Name of the modifier.
	NodeType ast_pb.NodeType `json:"node_type"` // Type of the node.
	Src      SrcNode         `json:"src"`       // Source location information.
}

// ToProto converts the ModifierName node to its corresponding protobuf representation.
func (m *ModifierName) ToProto() *ast_pb.ModifierName {
	return &ast_pb.ModifierName{
		Id:       m.Id,
		Name:     m.Name,
		NodeType: m.NodeType,
		Src:      m.Src.ToProto(),
	}
}

// ModifierInvocation represents a modifier invocation node in the abstract syntax tree.
type ModifierInvocation struct {
	*ASTBuilder

	Id            int64              `json:"id"`                      // Unique identifier of the modifier invocation node.
	Name          string             `json:"name"`                    // Name of the modifier invocation.
	NodeType      ast_pb.NodeType    `json:"node_type"`               // Type of the node.
	Kind          ast_pb.NodeType    `json:"kind"`                    // Kind of the modifier invocation.
	Src           SrcNode            `json:"src"`                     // Source location information.
	ArgumentTypes []*TypeDescription `json:"argument_types"`          // Types of the arguments.
	Arguments     []Node[NodeType]   `json:"arguments"`               // Argument nodes.
	ModifierName  *ModifierName      `json:"modifier_name,omitempty"` // Modifier name node.
}

// NewModifierInvocation creates a new instance of ModifierInvocation with the provided ASTBuilder.
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

// SetReferenceDescriptor sets the reference descriptors of the ModifierInvocation node.
func (m *ModifierInvocation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the modifier invocation node.
func (m *ModifierInvocation) GetId() int64 {
	return m.Id
}

// GetType returns the type of the node.
func (m *ModifierInvocation) GetType() ast_pb.NodeType {
	return m.NodeType
}

// GetKind returns the kind of the modifier invocation.
func (m *ModifierInvocation) GetKind() ast_pb.NodeType {
	return m.Kind
}

// GetSrc returns the source location information of the modifier invocation node.
func (m *ModifierInvocation) GetSrc() SrcNode {
	return m.Src
}

// GetName returns the name of the modifier invocation.
func (m *ModifierInvocation) GetName() string {
	return m.Name
}

// GetTypeDescription returns the type description of the modifier invocation (returns nil).
func (m *ModifierInvocation) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetNodes returns a slice of nodes associated with the modifier invocation (arguments).
func (m *ModifierInvocation) GetNodes() []Node[NodeType] {
	return m.Arguments
}

// GetArguments returns a slice of argument nodes of the modifier invocation.
func (m *ModifierInvocation) GetArguments() []Node[NodeType] {
	return m.Arguments
}

// GetArgumentTypes returns a slice of argument types of the modifier invocation.
func (m *ModifierInvocation) GetArgumentTypes() []*TypeDescription {
	return m.ArgumentTypes
}

func (m *ModifierInvocation) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &m.Id); err != nil {
			return err
		}
	}

	if name, ok := tempMap["name"]; ok {
		if err := json.Unmarshal(name, &m.Name); err != nil {
			return err
		}
	}

	if kind, ok := tempMap["kind"]; ok {
		if err := json.Unmarshal(kind, &m.Kind); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &m.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &m.Src); err != nil {
			return err
		}
	}

	if modifierName, ok := tempMap["modifier_name"]; ok {
		if err := json.Unmarshal(modifierName, &m.ModifierName); err != nil {
			return err
		}
	}

	if arguments, ok := tempMap["arguments"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(arguments, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(tempNode, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(tempNode, tempNodeType)
			if err != nil {
				return err
			}
			m.Arguments = append(m.Arguments, node)
		}
	}

	if argumentTypes, ok := tempMap["argument_types"]; ok {
		if err := json.Unmarshal(argumentTypes, &m.ArgumentTypes); err != nil {
			return err
		}
	}

	return nil
}

// ToProto converts the ModifierInvocation node to its corresponding protobuf representation.
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

// Parse parses the modifier invocation context and populates the ModifierInvocation fields.
func (m *ModifierInvocation) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx parser.IModifierInvocationContext,
) {
	m.Src = SrcNode{
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
			expr := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, m, m.GetId(), expressionCtx)
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
