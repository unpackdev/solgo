package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type MemberAccessExpression struct {
	*ASTBuilder

	Id                    int64              `json:"id"`
	IsConstant            bool               `json:"is_constant"`
	IsLValue              bool               `json:"is_l_value"`
	IsPure                bool               `json:"is_pure"`
	LValueRequested       bool               `json:"l_value_requested"`
	NodeType              ast_pb.NodeType    `json:"node_type"`
	Src                   SrcNode            `json:"src"`
	Expression            Node[NodeType]     `json:"expression"`
	MemberName            string             `json:"member_name"`
	ArgumentTypes         []*TypeDescription `json:"argument_types"`
	ReferencedDeclaration int64              `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription   `json:"type_description"`
}

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
	return false
}

func (m *MemberAccessExpression) GetId() int64 {
	return m.Id
}

func (m *MemberAccessExpression) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *MemberAccessExpression) GetSrc() SrcNode {
	return m.Src
}

func (m *MemberAccessExpression) GetExpression() Node[NodeType] {
	return m.Expression
}

func (m *MemberAccessExpression) GetMemberName() string {
	return m.MemberName
}

func (m *MemberAccessExpression) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

func (m *MemberAccessExpression) GetArgumentTypes() []*TypeDescription {
	return m.ArgumentTypes
}

func (m *MemberAccessExpression) GetNodes() []Node[NodeType] {
	return nil
}

func (m *MemberAccessExpression) ToProto() NodeType {
	return ast_pb.MemberAccessExpression{}
}

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

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}
	m.NodeType = ast_pb.NodeType_MEMBER_ACCESS
	m.MemberName = ctx.Identifier().GetText()

	if ctx.Expression() != nil {
		expression := NewExpression(m.ASTBuilder)
		m.Expression = expression.Parse(
			unit, contractNode, fnNode, bodyNode, vDeclar, m, ctx.Expression(),
		)
		m.TypeDescription = m.Expression.GetTypeDescription()

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
