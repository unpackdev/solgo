package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
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
	}
}

func (a *YulStatement) GetId() int64 {
	return a.Id
}

func (a *YulStatement) GetType() ast_pb.NodeType {
	return a.NodeType
}

func (a *YulStatement) GetSrc() SrcNode {
	return a.Src
}

func (a *YulStatement) GetNodes() []Node[NodeType] {
	return a.Statements
}

func (a *YulStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (a *YulStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (a *YulStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *AssemblyStatement,
	ctx *parser.YulStatementContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:          a.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: assemblyNode.GetId(),
	}

	for _, childCtx := range ctx.GetChildren() {
		switch child := childCtx.(type) {
		case *parser.YulAssignmentContext:
			assignment := NewYulAssignment(a.ASTBuilder)
			a.Statements = append(a.Statements,
				assignment.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, a, child),
			)
		default:
			panic(fmt.Sprintf("Unimplemented YulStatementContext @ YulStatement.Parse(): %T", child))
		}
	}

	return a
}
