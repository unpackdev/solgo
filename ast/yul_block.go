package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulBlockStatement struct {
	*ASTBuilder

	Id         int64            `json:"id"`
	NodeType   ast_pb.NodeType  `json:"node_type"`
	Src        SrcNode          `json:"src"`
	Statements []Node[NodeType] `json:"statements"`
}

func NewYulBlockStatement(b *ASTBuilder) *YulBlockStatement {
	return &YulBlockStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_BLOCK,
		Statements: make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulBlockStatement node.
func (y *YulBlockStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulBlockStatement) GetId() int64 {
	return y.Id
}

func (y *YulBlockStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulBlockStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulBlockStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Statements...)
	return toReturn
}

func (y *YulBlockStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulBlockStatement) GetStatements() []Node[NodeType] {
	return y.Statements
}

func (y *YulBlockStatement) ToProto() NodeType {
	toReturn := ast_pb.YulBlockStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	for _, statement := range y.GetStatements() {
		toReturn.Statements = append(
			toReturn.Statements,
			statement.ToProto().(*v3.TypedStruct),
		)
	}

	return NewTypedStruct(&toReturn, "YulBlockStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulBlockStatement node.
func (f *YulBlockStatement) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulBlockStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ifStatement *parser.YulIfStatementContext,
	parentNode Node[NodeType],
	ctx *parser.YulBlockContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	if ctx.AllYulStatement() != nil {
		for _, statement := range ctx.AllYulStatement() {
			yStatement := NewYulStatement(y.ASTBuilder)

			y.Statements = append(y.Statements, yStatement.Parse(
				unit, contractNode, fnNode, bodyNode, assemblyNode, y, statement.(*parser.YulStatementContext),
			))
		}
	}

	return y
}
