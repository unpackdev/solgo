package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulFunctionDefinition struct {
	*ASTBuilder

	Id               int64            `json:"id"`
	NodeType         ast_pb.NodeType  `json:"node_type"`
	Src              SrcNode          `json:"src"`
	Arguments        []*YulIdentifier `json:"arguments"`
	Body             Node[NodeType]   `json:"body"`
	ReturnParameters []*YulIdentifier `json:"return_parameters"`
}

func NewYulFunctionDefinition(b *ASTBuilder) *YulFunctionDefinition {
	return &YulFunctionDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FUNCTION_DEFINITION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulFunctionDefinition node.
func (y *YulFunctionDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulFunctionDefinition) GetId() int64 {
	return y.Id
}

func (y *YulFunctionDefinition) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulFunctionDefinition) GetSrc() SrcNode {
	return y.Src
}

func (y *YulFunctionDefinition) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Body)
	return toReturn
}

func (y *YulFunctionDefinition) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulFunctionDefinition) ToProto() NodeType {
	return ast_pb.Statement{}
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulFunctionDefinition node.
func (f *YulFunctionDefinition) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulFunctionDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulFunctionDefinitionContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStart().GetStop()),
		Length:      int64(ctx.GetStart().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	for _, argument := range ctx.GetArguments() {
		y.Arguments = append(y.Arguments, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     argument.GetText(),
			NodeType: ast_pb.NodeType_YUL_VARIABLE_NAME,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(argument.GetLine()),
				Column:      int64(argument.GetColumn()),
				Start:       int64(argument.GetStart()),
				End:         int64(argument.GetStop()),
				Length:      int64(argument.GetStop() - argument.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Body = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	for _, argument := range ctx.GetReturnParameters() {
		y.ReturnParameters = append(y.ReturnParameters, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     argument.GetText(),
			NodeType: ast_pb.NodeType_YUL_VARIABLE_NAME,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(argument.GetLine()),
				Column:      int64(argument.GetColumn()),
				Start:       int64(argument.GetStart()),
				End:         int64(argument.GetStop()),
				Length:      int64(argument.GetStop() - argument.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	return y
}
