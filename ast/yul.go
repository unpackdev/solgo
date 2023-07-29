package ast

import (
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
	return nil
}

func (y *YulStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *AssemblyStatement,
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

	/*
		NEXT RELEASE PRIORITY
		for _, childCtx := range ctx.GetChildren() {
			switch child := childCtx.(type) {
			case *parser.YulAssignmentContext:
				assignment := NewYulAssignment(y.ASTBuilder)
				y.Statements = append(y.Statements,
					assignment.Parse(unit, contractNode, fnNode, bodyNode, assemblyNode, a, child),
				)
			default:
				panic(fmt.Sprintf("Unimplemented YulStatementContext @ YulStatement.Parse(): %T", child))
			}
		} */

	return y
}
