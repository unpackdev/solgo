package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type FunctionCall struct {
	*ASTBuilder

	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"node_type"`
	Kind            ast_pb.NodeType  `json:"kind"`
	Src             SrcNode          `json:"src"`
	IsConstant      bool             `json:"is_constant"`
	IsLValue        bool             `json:"is_l_value"`
	IsPure          bool             `json:"is_pure"`
	LValueRequested bool             `json:"l_value_requested"`
	Arguments       []Node[NodeType] `json:"arguments"`
	TryCall         bool             `json:"try_call"`
	Expression      Node[NodeType]   `json:"expression"`
}

func NewFunctionCall(b *ASTBuilder) *FunctionCall {
	return &FunctionCall{
		ASTBuilder: b,
		Arguments:  make([]Node[NodeType], 0),
		NodeType:   ast_pb.NodeType_FUNCTION_CALL,
		Kind:       ast_pb.NodeType_FUNCTION_CALL,
	}
}

func (f *FunctionCall) GetId() int64 {
	return f.Id
}

func (f *FunctionCall) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *FunctionCall) GetSrc() SrcNode {
	return f.Src
}

func (f *FunctionCall) GetArguments() []Node[NodeType] {
	return f.Arguments
}

func (f *FunctionCall) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (f *FunctionCall) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *FunctionCall) GetExpression() Node[NodeType] {
	return f.Expression
}

func (f *FunctionCall) GetTryCall() bool {
	return f.TryCall
}

func (f *FunctionCall) GetTypeDescription() *TypeDescription {
	return nil
}

func (f *FunctionCall) GetNodes() []Node[NodeType] {
	return nil
}

func (f *FunctionCall) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.FunctionCallContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	expression := NewExpression(f.ASTBuilder)

	if ctx.CallArgumentList() != nil {
		for _, expressionCtx := range ctx.CallArgumentList().AllExpression() {
			f.Arguments = append(
				f.Arguments,
				expression.Parse(
					unit, contractNode, fnNode,
					bodyNode, nil, f, expressionCtx,
				),
			)
		}
	}

	if ctx.Expression() != nil {
		f.Expression = expression.Parse(
			unit, contractNode, fnNode, bodyNode, nil, f, ctx.Expression(),
		)
	}

	return f
}
