package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulForStatement struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Src         SrcNode          `json:"src"`
	Identifiers []*YulIdentifier `json:"identifiers"`
	Expression  Node[NodeType]   `json:"expression"`
	Init        Node[NodeType]   `json:"init"`
	Condition   Node[NodeType]   `json:"condition"`
	Body        Node[NodeType]   `json:"body"`
}

func NewYulForStatement(b *ASTBuilder) *YulForStatement {
	return &YulForStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_ASSIGNMENT,
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
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

func (y *YulForStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulForStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulForStatement) GetIdentifiers() []*YulIdentifier {
	return y.Identifiers
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

	/* 	ctx.YulFor()
	   	ctx.YulExpression()
	   	ctx.AllYulBlock()
	   	ctx.GetCond()
	   	ctx.GetBody()
	   	ctx.GetInit() */

	if ctx.GetInit() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Init = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil,
			ctx.GetInit().(*parser.YulBlockContext),
		)
	}

	if ctx.GetCond() != nil {
		y.Condition = ParseYulExpression(
			y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, nil,
			y, ctx.GetCond(),
		)
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Init = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	//utils.DumpNodeWithExit(ctx.GetText())

	/* 	if ctx.YulExpression() != nil {
	   		y.Expression = ParseYulExpression(
	   			y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, ctx,
	   			ctx.YulExpression(),
	   		)
	   	}

	   	for _, identifier := range ctx.AllYulIdentifier() {
	   		y.Identifiers = append(y.Identifiers, &YulIdentifier{
	   			Id:       y.GetNextID(),
	   			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
	   			Src: SrcNode{
	   				Id:          y.GetNextID(),
	   				Line:        int64(identifier.GetSymbol().GetLine()),
	   				Column:      int64(identifier.GetSymbol().GetColumn()),
	   				Start:       int64(identifier.GetSymbol().GetStart()),
	   				End:         int64(identifier.GetSymbol().GetStop()),
	   				Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
	   				ParentIndex: y.GetId(),
	   			},
	   			Name: identifier.GetText(),
	   			NameLocation: SrcNode{
	   				Id:          y.GetNextID(),
	   				Line:        int64(identifier.GetSymbol().GetLine()),
	   				Column:      int64(identifier.GetSymbol().GetColumn()),
	   				Start:       int64(identifier.GetSymbol().GetStart()),
	   				End:         int64(identifier.GetSymbol().GetStop()),
	   				Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
	   				ParentIndex: y.GetId(),
	   			},
	   		})
	   	} */

	return y
}
