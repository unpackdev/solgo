package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ContinueStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewContinueStatement(b *ASTBuilder) *ContinueStatement {
	return &ContinueStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_CONTINUE,
	}
}

func (b *ContinueStatement) GetId() int64 {
	return b.Id
}

func (b *ContinueStatement) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *ContinueStatement) GetSrc() SrcNode {
	return b.Src
}

func (b *ContinueStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (b *ContinueStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (b *ContinueStatement) ToProto() NodeType {
	return ast_pb.Continue{}
}

func (b *ContinueStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ContinueStatementContext,
) Node[NodeType] {
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}
	return b
}
