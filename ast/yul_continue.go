package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

type YulContinueStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewYulContinueStatement(b *ASTBuilder) *YulContinueStatement {
	return &YulContinueStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_CONTINUE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulContinueStatement node.
func (y *YulContinueStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulContinueStatement) GetId() int64 {
	return y.Id
}

func (y *YulContinueStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulContinueStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulContinueStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	return toReturn
}

func (y *YulContinueStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulContinueStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulContinueStatement) Parse(
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
