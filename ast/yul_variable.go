package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulVariable struct {
	*ASTBuilder

	Id            int64               `json:"id"`
	NodeType      ast_pb.NodeType     `json:"node_type"`
	Src           SrcNode             `json:"src"`
	Identifiers   []*YulIdentifier    `json:"identifiers"`
	Let           bool                `json:"let"`
	Expression    Node[NodeType]      `json:"expression"`
	VariableNames []*YulVariableNames `json:"variable_names"`
}

func NewYulVariable(b *ASTBuilder) *YulVariable {
	return &YulVariable{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_YUL_VARIABLE_DECLARATION,
		VariableNames: make([]*YulVariableNames, 0),
		Identifiers:   make([]*YulIdentifier, 0),
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
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

func (y *YulVariable) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulVariable) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulVariable) GetIdentifiers() []*YulIdentifier {
	return y.Identifiers
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
		y.VariableNames = append(y.VariableNames, &YulVariableNames{
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

	if ctx.YulFunctionCall() != nil {
		fcStatement := NewYulFunctionCallStatement(y.ASTBuilder)
		y.Expression = fcStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
			ctx.YulFunctionCall().(*parser.YulFunctionCallContext),
		)
	}

	if ctx.YulExpression() != nil {
		y.Expression = ParseYulExpression(
			y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, ctx,
			y, ctx.YulExpression(),
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
	}

	return y
}
