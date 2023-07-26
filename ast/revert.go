package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type RevertStatement struct {
	*ASTBuilder

	Id         int64            `json:"id"`
	NodeType   ast_pb.NodeType  `json:"node_type"`
	Src        SrcNode          `json:"src"`
	Arguments  []Node[NodeType] `json:"arguments"`
	Expression Node[NodeType]   `json:"expression"`
}

func NewRevertStatement(b *ASTBuilder) *RevertStatement {
	return &RevertStatement{
		ASTBuilder: b,

		Id:        b.GetNextID(),
		NodeType:  ast_pb.NodeType_REVERT_STATEMENT,
		Arguments: make([]Node[NodeType], 0),
	}
}

func (r *RevertStatement) GetId() int64 {
	return r.Id
}

func (r *RevertStatement) GetType() ast_pb.NodeType {
	return r.NodeType
}

func (r *RevertStatement) GetSrc() SrcNode {
	return r.Src
}

func (r *RevertStatement) GetArguments() []Node[NodeType] {
	return r.Arguments
}

func (r *RevertStatement) GetExpression() Node[NodeType] {
	return r.Expression
}

func (r *RevertStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (r *RevertStatement) ToProto() NodeType {
	return &ast_pb.Revert{}
}

func (r *RevertStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (r *RevertStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.RevertStatementContext,
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

	expression := NewExpression(r.ASTBuilder)

	if ctx.CallArgumentList() != nil {
		for _, expressionCtx := range ctx.CallArgumentList().AllExpression() {
			r.Arguments = append(
				r.Arguments,
				expression.Parse(
					unit, contractNode, fnNode,
					bodyNode, nil, r, expressionCtx,
				),
			)
		}
	}

	r.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())
	return r
}
