package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
	"go.uber.org/zap"
)

type YulStatement struct {
	*ASTBuilder

	Id         int64            `json:"id"`
	NodeType   ast_pb.NodeType  `json:"node_type"`
	Src        SrcNode          `json:"src"`
	Statements []Node[NodeType] `json:"body"`
}

func NewYulStatement(b *ASTBuilder) *YulStatement {
	return &YulStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_STATEMENT,
		Statements: make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulStatement node.
func (y *YulStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulStatement) GetId() int64 {
	return y.Id
}

func (y *YulStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulStatement) GetNodes() []Node[NodeType] {
	return y.Statements
}

func (y *YulStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	ctx *parser.YulStatementContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: assemblyNode.GetId(),
	}

	for _, childCtx := range ctx.GetChildren() {
		switch child := childCtx.(type) {
		case *parser.YulBlockContext:
			block := NewYulBlockStatement(y.ASTBuilder)
			y.Statements = append(y.Statements,
				block.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, y, nil, child),
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
