package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulFunctionCall struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Src         SrcNode          `json:"src"`
	Identifier  *YulIdentifier   `json:"identifier"`
	EVMBuiltin  *YulEVMBuiltin   `json:"evm_builtin"`
	Expressions []Node[NodeType] `json:"expressions"`
}

func NewYulFunctionCallStatement(b *ASTBuilder) *YulFunctionCall {
	return &YulFunctionCall{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FUNCTION_CALL,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulFunctionCall node.
func (y *YulFunctionCall) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulFunctionCall) GetId() int64 {
	return y.Id
}

func (y *YulFunctionCall) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulFunctionCall) GetSrc() SrcNode {
	return y.Src
}

func (y *YulFunctionCall) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Expressions...)
	return toReturn
}

func (y *YulFunctionCall) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulFunctionCall) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulFunctionCall) GetIdentifier() *YulIdentifier {
	return y.Identifier
}

func (y *YulFunctionCall) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulFunctionCallContext,
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

	ctx.AllYulExpression()

	if ctx.YulIdentifier() != nil {
		identifier := ctx.YulIdentifier()
		y.Identifier = &YulIdentifier{
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
		}
	}

	if ctx.YulEVMBuiltin() != nil {
		builtin := ctx.YulEVMBuiltin()
		y.EVMBuiltin = &YulEVMBuiltin{
			Id:       y.GetNextID(),
			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(builtin.GetSymbol().GetLine()),
				Column:      int64(builtin.GetSymbol().GetColumn()),
				Start:       int64(builtin.GetSymbol().GetStart()),
				End:         int64(builtin.GetSymbol().GetStop()),
				Length:      int64(builtin.GetSymbol().GetStop() - builtin.GetSymbol().GetStart() + 1),
				ParentIndex: y.GetId(),
			},
			Name: builtin.GetText(),
			NameLocation: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(builtin.GetSymbol().GetLine()),
				Column:      int64(builtin.GetSymbol().GetColumn()),
				Start:       int64(builtin.GetSymbol().GetStart()),
				End:         int64(builtin.GetSymbol().GetStop()),
				Length:      int64(builtin.GetSymbol().GetStop() - builtin.GetSymbol().GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		}
	}

	if ctx.AllYulExpression() != nil {
		for _, expression := range ctx.AllYulExpression() {
			y.Expressions = append(y.Expressions, ParseYulExpression(
				y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, nil,
				y, expression,
			))
		}
	}

	//utils.DumpNodeWithExit(y)

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
