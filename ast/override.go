package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type OverrideSpecifier struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewOverrideSpecifier(b *ASTBuilder) *OverrideSpecifier {
	return &OverrideSpecifier{
		ASTBuilder: b,
	}
}

func (o *OverrideSpecifier) GetId() int64 {
	return o.Id
}

func (o *OverrideSpecifier) GetType() ast_pb.NodeType {
	return o.NodeType
}

func (o *OverrideSpecifier) GetSrc() SrcNode {
	return o.Src
}

func (o *OverrideSpecifier) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], ctx parser.IOverrideSpecifierContext) {
	o.Id = o.GetNextID()
	o.Src = SrcNode{
		Id:          o.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}
	o.NodeType = ast_pb.NodeType_OVERRIDE_SPECIFIER
}
