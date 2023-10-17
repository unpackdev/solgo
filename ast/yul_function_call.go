package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulFunctionCallStatement struct {
	*ASTBuilder

	Id           int64            `json:"id"`
	NodeType     ast_pb.NodeType  `json:"node_type"`
	Src          SrcNode          `json:"src"`
	FunctionName *YulIdentifier   `json:"function_name"`
	Arguments    []Node[NodeType] `json:"arguments"`
}

func NewYulFunctionCallStatement(b *ASTBuilder) *YulFunctionCallStatement {
	return &YulFunctionCallStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FUNCTION_CALL,
		Arguments:  make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulFunctionCallStatement node.
func (y *YulFunctionCallStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulFunctionCallStatement) GetId() int64 {
	return y.Id
}

func (y *YulFunctionCallStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulFunctionCallStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulFunctionCallStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Arguments...)
	return toReturn
}

func (y *YulFunctionCallStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulFunctionCallStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulFunctionCallStatement) Parse(
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

	if ctx.YulEVMBuiltin() != nil {
		builtin := ctx.YulEVMBuiltin()
		y.FunctionName = &YulIdentifier{
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
			if expression.YulPath() != nil {
				path := expression.YulPath()
				for _, identifier := range path.AllYulIdentifier() {
					y.Arguments = append(y.Arguments, &YulIdentifier{
						Id:       y.GetNextID(),
						NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
						Name:     identifier.GetText(),
						Src: SrcNode{
							Id:          y.GetNextID(),
							Line:        int64(path.GetStart().GetLine()),
							Column:      int64(path.GetStart().GetColumn()),
							Start:       int64(path.GetStart().GetStart()),
							End:         int64(path.GetStart().GetStop()),
							Length:      int64(path.GetStart().GetStop() - path.GetStart().GetStart() + 1),
							ParentIndex: y.GetId(),
						},
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
			}

			if expression.YulFunctionCall() != nil {
				fc := expression.YulFunctionCall()
				fcStatement := NewYulFunctionCallStatement(y.ASTBuilder)
				y.Arguments = append(y.Arguments, fcStatement.Parse(
					unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
					fc.(*parser.YulFunctionCallContext),
				))
			}

			if expression.YulLiteral() != nil {
				literal := expression.YulLiteral()
				literalStatement := NewYulLiteralStatement(y.ASTBuilder)
				y.Arguments = append(y.Arguments, literalStatement.Parse(
					unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
					literal.(*parser.YulLiteralContext),
				))
			}
		}
	}

	return y
}
