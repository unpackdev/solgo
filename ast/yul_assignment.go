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

// SetReferenceDescriptor sets the reference descriptions of the YulAssignment node.
func (y *YulAssignment) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulAssignment) GetId() int64 {
	return y.Id
}

func (y *YulAssignment) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulAssignment) GetSrc() SrcNode {
	return y.Src
}

func (y *YulAssignment) GetNodes() []Node[NodeType] {
	return nil
}

func (y *YulAssignment) GetTypeDescription() *TypeDescription {
	return nil
}

func (y *YulAssignment) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulAssignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *AssemblyStatement,
	statementNode *YulStatement,
	ctx *parser.YulAssignmentContext,
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
	return y
}
