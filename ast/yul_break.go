package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

type YulBreakStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewYulBreakStatement(b *ASTBuilder) *YulBreakStatement {
	return &YulBreakStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_BREAK,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulBreakStatement node.
func (y *YulBreakStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulBreakStatement) GetId() int64 {
	return y.Id
}

func (y *YulBreakStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulBreakStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulBreakStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	return toReturn
}

func (y *YulBreakStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulBreakStatement) ToProto() NodeType {
	toReturn := ast_pb.YulBreakStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	return NewTypedStruct(&toReturn, "YulBreakStatement")
}

func (y *YulBreakStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *antlr.TerminalNodeImpl,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetSymbol().GetLine()),
		Column:      int64(ctx.GetSymbol().GetColumn()),
		Start:       int64(ctx.GetSymbol().GetStart()),
		End:         int64(ctx.GetSymbol().GetStop()),
		Length:      int64(ctx.GetSymbol().GetStop() - ctx.GetSymbol().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	return y
}
