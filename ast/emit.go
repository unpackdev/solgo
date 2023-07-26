package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Emit struct {
	*ASTBuilder

	Id         int64            `json:"id"`
	NodeType   ast_pb.NodeType  `json:"node_type"`
	Src        SrcNode          `json:"src"`
	Arguments  []Node[NodeType] `json:"arguments"`
	Expression Node[NodeType]   `json:"expression"`
}

func NewEmitStatement(b *ASTBuilder) *Emit {
	return &Emit{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_EMIT_STATEMENT,
	}
}

func (e *Emit) GetId() int64 {
	return e.Id
}

func (e *Emit) GetType() ast_pb.NodeType {
	return e.NodeType
}

func (e *Emit) GetSrc() SrcNode {
	return e.Src
}

func (e *Emit) GetArguments() []Node[NodeType] {
	return e.Arguments
}

func (e *Emit) GetExpression() Node[NodeType] {
	return e.Expression
}

func (e *Emit) GetTypeDescription() *TypeDescription {
	return nil
}

func (e *Emit) GetNodes() []Node[NodeType] {
	return nil
}

func (e *Emit) ToProto() NodeType {
	return ast_pb.Emit{}
}

func (e *Emit) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.EmitStatementContext,
) Node[NodeType] {
	e.Src = SrcNode{
		Id:          e.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	expression := NewExpression(e.ASTBuilder)

	for _, argumentCtx := range ctx.CallArgumentList().AllExpression() {
		argument := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, argumentCtx)
		e.Arguments = append(e.Arguments, argument)
	}

	e.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())
	return e
}
