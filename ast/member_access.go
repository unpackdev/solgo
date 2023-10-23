package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// MemberAccessExpression represents a member access expression node in the AST.
// It contains information about the accessed member, expression, type description, and related metadata.
type MemberAccessExpression struct {
	*ASTBuilder

	Id                    int64              `json:"id"`
	Constant              bool               `json:"is_constant"`
	LValue                bool               `json:"is_l_value"`
	Pure                  bool               `json:"is_pure"`
	LValueRequested       bool               `json:"l_value_requested"`
	NodeType              ast_pb.NodeType    `json:"node_type"`
	Src                   SrcNode            `json:"src"`
	MemberLocation        SrcNode            `json:"member_location"`
	Expression            Node[NodeType]     `json:"expression"`
	MemberName            string             `json:"member_name"`
	ArgumentTypes         []*TypeDescription `json:"argument_types"`
	ReferencedDeclaration int64              `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription   `json:"type_description"`
}

// NewMemberAccessExpression creates a new MemberAccessExpression instance with initial values.
func NewMemberAccessExpression(b *ASTBuilder) *MemberAccessExpression {
	return &MemberAccessExpression{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_MEMBER_ACCESS,
		ArgumentTypes: []*TypeDescription{},
	}
}

// SetReferenceDescriptor sets the reference descriptions of the MemberAccessExpression node.
func (m *MemberAccessExpression) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	m.ReferencedDeclaration = refId
	m.TypeDescription = refDesc
	return true
}

// GetId returns the ID of the MemberAccessExpression node.
func (m *MemberAccessExpression) GetId() int64 {
	return m.Id
}

// GetType returns the NodeType of the MemberAccessExpression node.
func (m *MemberAccessExpression) GetType() ast_pb.NodeType {
	return m.NodeType
}

// GetSrc returns the source information of the MemberAccessExpression node.
func (m *MemberAccessExpression) GetSrc() SrcNode {
	return m.Src
}

// GetMemberLocation returns the source information of the accessed member.
func (m *MemberAccessExpression) GetMemberLocation() SrcNode {
	return m.MemberLocation
}

// GetExpression returns the expression being accessed in the member access.
func (m *MemberAccessExpression) GetExpression() Node[NodeType] {
	return m.Expression
}

// GetMemberName returns the name of the accessed member.
func (m *MemberAccessExpression) GetMemberName() string {
	return m.MemberName
}

// GetTypeDescription returns the type description associated with the member access.
func (m *MemberAccessExpression) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

// GetArgumentTypes returns the type descriptions of arguments in case of function call member access.
func (m *MemberAccessExpression) GetArgumentTypes() []*TypeDescription {
	return m.ArgumentTypes
}

// GetNodes returns the list of child nodes of the MemberAccessExpression node.
func (m *MemberAccessExpression) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{m.Expression}
}

// GetReferencedDeclaration returns the ID of the referenced declaration in the context of member access.
func (m *MemberAccessExpression) GetReferencedDeclaration() int64 {
	return m.ReferencedDeclaration
}

// IsConstant returns whether the member access is constant.
func (m *MemberAccessExpression) IsConstant() bool {
	return m.Constant
}

// IsLValue returns whether the member access is an l-value.
func (m *MemberAccessExpression) IsLValue() bool {
	return m.LValue
}

// IsPure returns whether the member access is pure.
func (m *MemberAccessExpression) IsPure() bool {
	return m.Pure
}

// IsLValueRequested returns whether an l-value is requested in the context of member access.
func (m *MemberAccessExpression) IsLValueRequested() bool {
	return m.LValueRequested
}

func (m *MemberAccessExpression) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &m.Id); err != nil {
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

	if memberLocation, ok := tempMap["member_location"]; ok {
		if err := json.Unmarshal(memberLocation, &m.MemberLocation); err != nil {
			return err
		}
	}

	if referencedDeclaration, ok := tempMap["referenced_declaration"]; ok {
		if err := json.Unmarshal(referencedDeclaration, &m.ReferencedDeclaration); err != nil {
			return err
		}
	}

	if typeDescription, ok := tempMap["type_description"]; ok {
		if err := json.Unmarshal(typeDescription, &m.TypeDescription); err != nil {
			return err
		}
	}

	if memberName, ok := tempMap["member_name"]; ok {
		if err := json.Unmarshal(memberName, &m.MemberName); err != nil {
			return err
		}
	}

	if argumentTypes, ok := tempMap["argument_types"]; ok {
		if err := json.Unmarshal(argumentTypes, &m.ArgumentTypes); err != nil {
			return err
		}
	}

	if constant, ok := tempMap["is_constant"]; ok {
		if err := json.Unmarshal(constant, &m.Constant); err != nil {
			return err
		}
	}

	if lValue, ok := tempMap["is_l_value"]; ok {
		if err := json.Unmarshal(lValue, &m.LValue); err != nil {
			return err
		}
	}

	if pure, ok := tempMap["is_pure"]; ok {
		if err := json.Unmarshal(pure, &m.Pure); err != nil {
			return err
		}
	}

	if lValueRequested, ok := tempMap["l_value_requested"]; ok {
		if err := json.Unmarshal(lValueRequested, &m.LValueRequested); err != nil {
			return err
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &m.Expression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(expression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(expression, tempNodeType)
			if err != nil {
				return err
			}
			m.Expression = node
		}
	}

	return nil
}

// ToProto converts the MemberAccessExpression node to its corresponding protobuf representation.
func (m *MemberAccessExpression) ToProto() NodeType {
	proto := ast_pb.MemberAccess{
		Id:                    m.GetId(),
		MemberName:            m.GetMemberName(),
		NodeType:              m.GetType(),
		Src:                   m.GetSrc().ToProto(),
		MemberLocation:        m.GetMemberLocation().ToProto(),
		ReferencedDeclaration: m.GetReferencedDeclaration(),
		IsConstant:            m.IsConstant(),
		IsLValue:              m.IsLValue(),
		IsPure:                m.IsPure(),
		LValueRequested:       m.IsLValueRequested(),
		Expression:            m.GetExpression().ToProto().(*v3.TypedStruct),
		ArgumentTypes:         make([]*ast_pb.TypeDescription, 0),
	}

	for _, arg := range m.GetArgumentTypes() {
		proto.ArgumentTypes = append(proto.ArgumentTypes, arg.ToProto())
	}

	if m.GetTypeDescription() != nil {
		proto.TypeDescription = m.GetTypeDescription().ToProto()
	}

	return NewTypedStruct(&proto, "MemberAccess")
}

// Parse populates the MemberAccessExpression node based on the provided context and other information.
func (m *MemberAccessExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.MemberAccessContext,
) Node[NodeType] {
	m.Src = SrcNode{
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

			return bodyNode.GetId()
		}(),
	}
	m.NodeType = ast_pb.NodeType_MEMBER_ACCESS
	m.MemberName = ctx.Identifier().GetText()
	m.MemberLocation = SrcNode{
		Line:        int64(ctx.Identifier().GetStart().GetLine()),
		Column:      int64(ctx.Identifier().GetStart().GetColumn()),
		Start:       int64(ctx.Identifier().GetStart().GetStart()),
		End:         int64(ctx.Identifier().GetStop().GetStop()),
		Length:      int64(ctx.Identifier().GetStop().GetStop() - ctx.Identifier().GetStart().GetStart() + 1),
		ParentIndex: m.Id,
	}

	// Parsing the expression in the member access.
	if ctx.Expression() != nil {
		expression := NewExpression(m.ASTBuilder)
		m.Expression = expression.Parse(
			unit, contractNode, fnNode, bodyNode, vDeclar, m, m.GetId(), ctx.Expression(),
		)

		m.TypeDescription = m.Expression.GetTypeDescription()

		// Handling edge case in type discovery.
		if m.Expression != nil && m.Expression.GetTypeDescription() == nil {
			if refId, refTypeDescription := m.GetResolver().ResolveByNode(m, m.MemberName); refTypeDescription != nil {
				m.ReferencedDeclaration = refId
				m.TypeDescription = refTypeDescription
			} else {
				if primary, ok := m.Expression.(*PrimaryExpression); ok {
					if refId, refTypeDescription := m.GetResolver().ResolveByNode(primary, primary.GetName()); refTypeDescription != nil {
						m.ReferencedDeclaration = refId
						m.TypeDescription = refTypeDescription
					}
				}
			}
		}

		// Forward type declaration for specific cases.
		if m.TypeDescription != nil {
			if m.TypeDescription.TypeIdentifier == "t_magic_message" {
				switch m.MemberName {
				case "sender":
					m.TypeDescription = &TypeDescription{
						TypeIdentifier: "t_address",
						TypeString:     "address",
					}
				case "data":
					m.TypeDescription = &TypeDescription{
						TypeIdentifier: "t_bytes_calldata_ptr",
						TypeString:     "bytes calldata",
					}
				case "value":
					m.TypeDescription = &TypeDescription{
						TypeIdentifier: "t_uint256",
						TypeString:     "uint256",
					}
				case "timestamp":
					m.TypeDescription = &TypeDescription{
						TypeIdentifier: "t_uint256",
						TypeString:     "uint256",
					}
				}
			}

			if m.TypeDescription.TypeIdentifier == "t_magic_block" {
				switch m.MemberName {
				case "timestamp":
					m.TypeDescription = &TypeDescription{
						TypeIdentifier: "t_uint256",
						TypeString:     "uint256",
					}
				}
			}
		}
	}

	// Handling function call argument types.
	if expNode != nil {
		if expNode.GetType() == ast_pb.NodeType_FUNCTION_CALL {
			fcNode := expNode.(*FunctionCall)
			for _, arguments := range fcNode.GetArguments() {
				m.ArgumentTypes = append(
					m.ArgumentTypes,
					arguments.GetTypeDescription(),
				)
			}
		}
	}

	return m
}
