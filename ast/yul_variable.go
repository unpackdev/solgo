package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulVariable struct {
	*ASTBuilder

	Id        int64            `json:"id"`
	NodeType  ast_pb.NodeType  `json:"node_type"`
	Src       SrcNode          `json:"src"`
	Let       bool             `json:"let"`
	Value     Node[NodeType]   `json:"value"`
	Variables []*YulIdentifier `json:"variables"`
}

func NewYulVariable(b *ASTBuilder) *YulVariable {
	return &YulVariable{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_VARIABLE_DECLARATION,
		Variables:  make([]*YulIdentifier, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulVariable node.
func (y *YulVariable) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulVariable) GetId() int64 {
	return y.Id
}

func (y *YulVariable) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulVariable) GetSrc() SrcNode {
	return y.Src
}

func (y *YulVariable) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Value)
	return toReturn
}

func (y *YulVariable) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulVariable) ToProto() NodeType {
	return ast_pb.Statement{}
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulVariable node.
func (f *YulVariable) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulVariable) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulVariableDeclarationContext,
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

	y.Let = ctx.YulLet() != nil

	for _, variable := range ctx.GetVariables() {
		y.Variables = append(y.Variables, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     variable.GetText(),
			NodeType: ast_pb.NodeType_YUL_VARIABLE_NAME,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(variable.GetLine()),
				Column:      int64(variable.GetColumn()),
				Start:       int64(variable.GetStart()),
				End:         int64(variable.GetStop()),
				Length:      int64(variable.GetStop() - variable.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	if ctx.YulExpression() != nil {
		expression := ctx.YulExpression()
		yExpression := NewYulExpressionStatement(y.ASTBuilder)
		y.Value = yExpression.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
			expression.(*parser.YulExpressionContext),
		)
	}

	return y
}
