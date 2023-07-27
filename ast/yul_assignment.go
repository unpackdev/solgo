package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type YulVariableNames struct {
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

type YulAssignment struct {
	*ASTBuilder

	Id            int64               `json:"id"`
	NodeType      ast_pb.NodeType     `json:"node_type"`
	Src           SrcNode             `json:"src"`
	VariableNames []*YulVariableNames `json:"variable_names"`
}

func NewYulAssignment(b *ASTBuilder) *YulAssignment {
	return &YulAssignment{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_YUL_STATEMENT,
		VariableNames: make([]*YulVariableNames, 0),
	}
}

func (a *YulAssignment) GetId() int64 {
	return a.Id
}

func (a *YulAssignment) GetType() ast_pb.NodeType {
	return a.NodeType
}

func (a *YulAssignment) GetSrc() SrcNode {
	return a.Src
}

func (a *YulAssignment) GetNodes() []Node[NodeType] {
	return nil
}

func (a *YulAssignment) GetTypeDescription() *TypeDescription {
	return nil
}

func (a *YulAssignment) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (a *YulAssignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *AssemblyStatement,
	statementNode *YulStatement,
	ctx *parser.YulAssignmentContext,
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
	return a
}
