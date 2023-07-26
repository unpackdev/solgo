package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type BreakStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewBreakStatement(b *ASTBuilder) *BreakStatement {
	return &BreakStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_BREAK,
	}
}

func (b *BreakStatement) GetId() int64 {
	return b.Id
}

func (b *BreakStatement) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *BreakStatement) GetSrc() SrcNode {
	return b.Src
}

func (b *BreakStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (b *BreakStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (b *BreakStatement) ToProto() NodeType {
	return ast_pb.Break{}
}

func (b *BreakStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.BreakStatementContext,
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
