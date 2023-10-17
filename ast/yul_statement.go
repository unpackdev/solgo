package ast

import (
	"encoding/json"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
	"go.uber.org/zap"
)

// YulStatement represents a statement in the Yul language.
type YulStatement struct {
	*ASTBuilder

	Id         int64            `json:"id"`         // Unique identifier for the statement node.
	NodeType   ast_pb.NodeType  `json:"node_type"`  // The type of the node.
	Src        SrcNode          `json:"src"`        // Source information about the node.
	Statements []Node[NodeType] `json:"statements"` // Statements within this Yul statement.
}

// NewYulStatement creates a new YulStatement node and initializes its fields.
func NewYulStatement(b *ASTBuilder) *YulStatement {
	return &YulStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_STATEMENT,
		Statements: make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulStatement node.
// It always returns false for this node type.
func (y *YulStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the YulStatement node.
func (y *YulStatement) GetId() int64 {
	return y.Id
}

// GetType returns the type of the YulStatement node.
func (y *YulStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulStatement node.
func (y *YulStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns the statements within the YulStatement node.
func (y *YulStatement) GetNodes() []Node[NodeType] {
	return y.Statements
}

// GetTypeDescription returns an empty TypeDescription for the YulStatement node.
func (y *YulStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulStatement) GetStatements() []Node[NodeType] {
	return y.Statements
}

// ToProto converts the YulStatement node into its protocol buffer representation.
func (y *YulStatement) ToProto() NodeType {
	proto := ast_pb.YulStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	for _, statement := range y.GetStatements() {
		proto.Statements = append(
			proto.Statements,
			statement.ToProto().(*v3.TypedStruct),
		)
	}

	return NewTypedStruct(&proto, "YulStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulStatement node.
// Currently, this function does not perform any unmarshalling and always returns nil.
func (f *YulStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &f.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &f.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if statements, ok := tempMap["statements"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(statements, &nodes); err != nil {
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
			f.Statements = append(f.Statements, node)
		}
	}

	return nil
}

// Parse processes the provided YulStatementContext and populates the YulStatement node's fields based on its content.
func (y *YulStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	parentNode Node[NodeType],
	ctx *parser.YulStatementContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	for _, childCtx := range ctx.GetChildren() {
		switch child := childCtx.(type) {
		case *parser.YulBlockContext:
			block := NewYulBlockStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				block.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, nil, y, child),
			)
		case *parser.YulAssignmentContext:
			assignment := NewYulAssignment(y.ASTBuilder)
			y.Statements = append(y.Statements,
				assignment.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *parser.YulVariableDeclarationContext:
			variable := NewYulVariable(y.ASTBuilder)
			y.Statements = append(y.Statements,
				variable.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *parser.YulFunctionCallContext:
			functionCall := NewYulFunctionCallStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				functionCall.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, nil, child),
			)
		case *parser.YulSwitchStatementContext:
			switchStatement := NewYulSwitchStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				switchStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *parser.YulIfStatementContext:
			ifStatement := NewYulIfStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				ifStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *parser.YulForStatementContext:
			forStatement := NewYulForStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				forStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *parser.YulFunctionDefinitionContext:
			functionDefinition := NewYulFunctionDefinition(y.ASTBuilder)
			y.Statements = append(y.Statements,
				functionDefinition.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
			)
		case *antlr.TerminalNodeImpl:
			switch child.GetText() {
			case "break":
				breakStatement := NewYulBreakStatement(y.ASTBuilder)
				y.Statements = append(y.Statements,
					breakStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
				)
			case "continue":
				breakStatement := NewYulContinueStatement(y.ASTBuilder)
				y.Statements = append(y.Statements,
					breakStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
				)
			case "leave":
				breakStatement := NewYulLeaveStatement(y.ASTBuilder)
				y.Statements = append(y.Statements,
					breakStatement.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, child),
				)
			default:
				y.NodeType = ast_pb.NodeType_YUL_TERMINAL_NODE
			}
		default:
			zap.L().Warn(fmt.Sprintf("Unimplemented YulStatementContext @ YulStatement.Parse(): %T", child))
		}
	}

	return y
}
