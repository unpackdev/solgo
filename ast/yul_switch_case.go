package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulSwitchCaseStatement struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Src         SrcNode          `json:"src"`
	Identifiers []*YulIdentifier `json:"identifiers"`
	Literal     Node[NodeType]   `json:"literal"`
	Block       Node[NodeType]   `json:"block"`
}

func NewYulSwitchCaseStatement(b *ASTBuilder) *YulSwitchCaseStatement {
	return &YulSwitchCaseStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_SWITCH_CASE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulSwitchCaseStatement node.
func (y *YulSwitchCaseStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulSwitchCaseStatement) GetId() int64 {
	return y.Id
}

func (y *YulSwitchCaseStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulSwitchCaseStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulSwitchCaseStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Literal)
	return toReturn
}

func (y *YulSwitchCaseStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulSwitchCaseStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulSwitchCaseStatement) GetIdentifiers() []*YulIdentifier {
	return y.Identifiers
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulSwitchCaseStatement node.
func (f *YulSwitchCaseStatement) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulSwitchCaseStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulSwitchCaseContext,
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

	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(y.ASTBuilder)
		y.Literal = literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	if ctx.YulBlock() != nil {
		block := NewYulBlockStatement(y.ASTBuilder)
		y.Block = block.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.YulBlock().(*parser.YulBlockContext),
		)
	}

	return y
}
