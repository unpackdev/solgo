package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

type YulLeaveStatement struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Src         SrcNode          `json:"src"`
	Identifiers []*YulIdentifier `json:"identifiers"`
	Expression  Node[NodeType]   `json:"expression"`
}

func NewYulLeaveStatement(b *ASTBuilder) *YulLeaveStatement {
	return &YulLeaveStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_ASSIGNMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulLeaveStatement node.
func (y *YulLeaveStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulLeaveStatement) GetId() int64 {
	return y.Id
}

func (y *YulLeaveStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulLeaveStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulLeaveStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

func (y *YulLeaveStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulLeaveStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulLeaveStatement) GetIdentifiers() []*YulIdentifier {
	return y.Identifiers
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulLeaveStatement node.
func (f *YulLeaveStatement) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulLeaveStatement) Parse(
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
	   				Start:       int64(identifier.GetSymbol().GetSymbol()),
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
