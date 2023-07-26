package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

type MemberAccessExpression struct {
	*ASTBuilder

	Id              int64             `json:"id"`
	IsConstant      bool              `json:"is_constant"`
	IsLValue        bool              `json:"is_l_value"`
	IsPure          bool              `json:"is_pure"`
	LValueRequested bool              `json:"l_value_requested"`
	NodeType        ast_pb.NodeType   `json:"node_type"`
	Src             SrcNode           `json:"src"`
	Expression      Node[NodeType]    `json:"expression"`
	MemberName      string            `json:"member_name"`
	ArgumentTypes   []TypeDescription `json:"argument_types"`
	TypeDescription TypeDescription   `json:"type_descriptions"`
}

func NewMemberAccessExpression(b *ASTBuilder) *MemberAccessExpression {
	return &MemberAccessExpression{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_MEMBER_ACCESS,
		ArgumentTypes: []TypeDescription{},
	}
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

func (m *MemberAccessExpression) GetTypeDescription() TypeDescription {
	return m.TypeDescription
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
) {

	m.NodeType = ast_pb.NodeType_MEMBER_ACCESS
	m.MemberName = ctx.Identifier().GetText()

	if ctx.Expression() != nil {
		expression := NewExpression(m.ASTBuilder)
		m.Expression = expression.Parse(
			unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(),
		)
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

	// Now we are going to search through all existing source units in hope
	// to discover reference declaration...
	/** for _, units := range m.sourceUnits {
		if units.GetNodes() != nil && len(units.GetNodes()) > 0 {
			for _, nodeCtx := range units.GetNodes() {
				fmt.Println("Node Type: ", nodeCtx.GetType())
				if nodeCtx.GetType() == ast_pb.NodeType_CONTRACT_DEFINITION {
					contractDef := nodeCtx.(ContractNode[*ast_pb.Contract])
					for _, node := range contractDef.Nodes {
						fmt.Println("Node Type AAA: ", node.GetType())
					}
				}
				 				for _, node := range nodeCtx.Nodes {
					if node.Name == toReturn.MemberName {
						toReturn.ReferencedDeclaration = node.Id
					}
				}
			}
		}
	} **/

	/* 		for _, enum := range b.currentEnums {
	   			if enum.Name == toReturn.Expression.Name {
	   				toReturn.ReferencedDeclaration = enum.Id
	   				toReturn.TypeDescriptions = enum.TypeDescriptions

	   				if toReturn.Expression.TypeDescriptions == nil {
	   					toReturn.Expression.TypeDescriptions = enum.TypeDescriptions
	   				}
	   			}
	   		}
	*/

	if m.Expression.GetTypeDescription().TypeString == "t_magic_message" {
		switch m.MemberName {
		case "sender":
			m.TypeDescription = TypeDescription{
				TypeIdentifier: "t_address",
				TypeString:     "address",
			}
		case "data":
			m.TypeDescription = TypeDescription{
				TypeIdentifier: "t_bytes_calldata_ptr",
				TypeString:     "bytes calldata",
			}
		case "value":
			m.TypeDescription = TypeDescription{
				TypeIdentifier: "t_uint256",
				TypeString:     "uint256",
			}
		default:
			zap.L().Warn(
				"Unknown magic message member",
				zap.String("member", m.MemberName),
			)
		}
	}
}
