package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type AssemblyStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
	Body     *BodyNode       `json:"body"`
}

func NewAssemblyStatement(b *ASTBuilder) *AssemblyStatement {
	return &AssemblyStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_ASSEMBLY_STATEMENT,
	}
}

func (a *AssemblyStatement) GetId() int64 {
	return a.Id
}

func (a *AssemblyStatement) GetType() ast_pb.NodeType {
	return a.NodeType
}

func (a *AssemblyStatement) GetSrc() SrcNode {
	return a.Src
}

func (a *AssemblyStatement) GetBody() *BodyNode {
	return a.Body
}

func (a *AssemblyStatement) GetNodes() []Node[NodeType] {
	return a.Body.Statements
}

func (a *AssemblyStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (a *AssemblyStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (a *AssemblyStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.AssemblyStatementContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:          a.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	yulStatement := NewYulStatement(a.ASTBuilder)

	a.Body = NewBodyNode(a.ASTBuilder)
	a.Body.Src = a.Src
	a.Body.Src.ParentIndex = a.Id
	a.Body.NodeType = ast_pb.NodeType_AST

	for _, yulCtx := range ctx.AllYulStatement() {
		a.Body.Statements = append(a.Body.Statements,
			yulStatement.Parse(
				unit, contractNode, fnNode, a.Body, a, yulCtx.(*parser.YulStatementContext),
			),
		)
	}

	if ctx.AssemblyDialect() != nil {
		fmt.Println("Assembly Dialect: ", ctx.AssemblyDialect().GetText())
	}

	return a
}
