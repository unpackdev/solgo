package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulForStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Pre       Node[NodeType]  `json:"pre"`
	Post      Node[NodeType]  `json:"post"`
	Condition Node[NodeType]  `json:"condition"`
	Body      Node[NodeType]  `json:"body"`
}

func NewYulForStatement(b *ASTBuilder) *YulForStatement {
	return &YulForStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FOR,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulForStatement node.
func (y *YulForStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulForStatement) GetId() int64 {
	return y.Id
}

func (y *YulForStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulForStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulForStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Condition)
	toReturn = append(toReturn, y.Body)
	return toReturn
}

func (y *YulForStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulForStatement) GetPre() Node[NodeType] {
	return y.Pre
}

func (y *YulForStatement) GetPost() Node[NodeType] {
	return y.Post
}

func (y *YulForStatement) GetCondition() Node[NodeType] {
	return y.Condition
}

func (y *YulForStatement) GetBody() Node[NodeType] {
	return y.Body
}

func (y *YulForStatement) ToProto() NodeType {
	toReturn := ast_pb.YulForStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	if y.GetPre() != nil {
		toReturn.Pre = y.GetPre().ToProto().(*v3.TypedStruct)
	}

	if y.GetPost() != nil {
		toReturn.Post = y.GetPost().ToProto().(*v3.TypedStruct)
	}

	if y.GetCondition() != nil {
		toReturn.Condition = y.GetCondition().ToProto().(*v3.TypedStruct)
	}

	if y.GetBody() != nil {
		toReturn.Body = y.GetBody().ToProto().(*v3.TypedStruct)
	}

	return NewTypedStruct(&toReturn, "YulForStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulForStatement node.
func (f *YulForStatement) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulForStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulForStatementContext,
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

	if ctx.GetCond() != nil {
		y.Condition = ParseYulExpression(
			y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, nil,
			y, ctx.GetCond(),
		)
	}

	if ctx.GetInit() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Pre = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetInit().(*parser.YulBlockContext),
		)
	}

	if ctx.GetPost() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Post = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetPost().(*parser.YulBlockContext),
		)
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Body = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	return y
}
