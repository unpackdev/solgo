package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ReturnStatement struct {
	*ASTBuilder

	Id                       int64           `json:"id"`
	NodeType                 ast_pb.NodeType `json:"node_type"`
	Src                      SrcNode         `json:"src"`
	FunctionReturnParameters int64           `json:"function_return_parameters"`
	Expression               Node[NodeType]  `json:"expression"`
}

func NewReturnStatement(b *ASTBuilder) *ReturnStatement {
	return &ReturnStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_RETURN_STATEMENT,
	}
}

func (r *ReturnStatement) GetId() int64 {
	return r.Id
}

func (r *ReturnStatement) GetType() ast_pb.NodeType {
	return r.NodeType
}

func (r *ReturnStatement) GetSrc() SrcNode {
	return r.Src
}

func (r *ReturnStatement) GetExpression() Node[NodeType] {
	return r.Expression
}

func (r *ReturnStatement) GetFunctionReturnParameters() int64 {
	return r.FunctionReturnParameters
}

func (r *ReturnStatement) GetTypeDescription() *TypeDescription {
	if r.Expression != nil {
		return r.Expression.GetTypeDescription()
	}
	return nil
}

func (r *ReturnStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (r *ReturnStatement) ToProto() NodeType {
	return ast_pb.Return{}
}

func (r *ReturnStatement) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ReturnStatementContext,
) Node[NodeType] {
	r.Src = SrcNode{
		Id:          r.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	fnCtx := fnNode.(FunctionNode[ast_pb.Function])
	if fnCtx.GetReturnParameters() != nil {
		r.FunctionReturnParameters = fnCtx.GetId()
	}

	expression := NewExpression(r.ASTBuilder)
	r.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())
	return r
}
