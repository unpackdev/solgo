package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
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

// SetReferenceDescriptor sets the reference descriptions of the RevertStatement node.
func (r *RevertStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
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
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, r.Arguments...)
	toReturn = append(toReturn, r.Expression)
	return toReturn
}

func (r *RevertStatement) ToProto() NodeType {
	proto := ast_pb.Revert{
		Id:         r.Id,
		NodeType:   r.NodeType,
		Src:        r.Src.ToProto(),
		Arguments:  make([]*v3.TypedStruct, 0),
		Expression: r.Expression.ToProto().(*v3.TypedStruct),
	}

	for _, arg := range r.Arguments {
		proto.Arguments = append(proto.Arguments, arg.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Revert")
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
