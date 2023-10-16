package ast

import (
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulSwitchStatement struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Src         SrcNode          `json:"src"`
	Identifiers []*YulIdentifier `json:"identifiers"`
	Expression  Node[NodeType]   `json:"expression"`
}

func NewYulSwitchStatement(b *ASTBuilder) *YulSwitchStatement {
	return &YulSwitchStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_ASSIGNMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulSwitchStatement node.
func (y *YulSwitchStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulSwitchStatement) GetId() int64 {
	return y.Id
}

func (y *YulSwitchStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulSwitchStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulSwitchStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

func (y *YulSwitchStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulSwitchStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulSwitchStatement) GetIdentifiers() []*YulIdentifier {
	return y.Identifiers
}

func (y *YulSwitchStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulSwitchStatementContext,
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

	if ctx.AllYulSwitchCase() != nil {
		for _, switchCase := range ctx.AllYulSwitchCase() {
			fmt.Println(switchCase)
		}
	}

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
