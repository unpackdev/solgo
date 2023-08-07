package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type IfStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      Node[NodeType]  `json:"body"`
}

func NewIfStatement(b *ASTBuilder) *IfStatement {
	return &IfStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IF_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the IfStatement node.
func (i *IfStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (i *IfStatement) GetId() int64 {
	return i.Id
}

func (i *IfStatement) GetType() ast_pb.NodeType {
	return i.NodeType
}

func (i *IfStatement) GetSrc() SrcNode {
	return i.Src
}

func (i *IfStatement) GetCondition() Node[NodeType] {
	return i.Condition
}

func (i *IfStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (i *IfStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{i.Condition, i.Body}
}

func (i *IfStatement) GetBody() Node[NodeType] {
	return i.Body
}

func (i *IfStatement) ToProto() NodeType {
	proto := ast_pb.If{
		Id:        i.GetId(),
		NodeType:  i.GetType(),
		Src:       i.GetSrc().ToProto(),
		Condition: i.GetCondition().ToProto().(*v3.TypedStruct),
	}

	if i.GetBody() != nil {
		proto.Body = i.GetBody().ToProto().(*ast_pb.Body)
	}

	return NewTypedStruct(&proto, "If")
}

func (i *IfStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.IfStatementContext,
) Node[NodeType] {
	i.Src = SrcNode{
		Id:          i.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(i.ASTBuilder)

	i.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, i, ctx.Expression())

	// i.Body set is just ridicolous as there are so many different ways for parsed ast to show nil
	// instead of empty []. This was the way I've sorted it out. Future can decide if cleanup is necessary.
	body := NewBodyNode(i.ASTBuilder)
	if len(ctx.AllStatement()) > 0 {
		for _, statementCtx := range ctx.AllStatement() {
			if statementCtx.IsEmpty() {
				continue
			}

			if statementCtx.Block() != nil {
				body.ParseBlock(unit, contractNode, fnNode, statementCtx.Block())
				break
			}

			i.Body = body
		}

		i.Body = body
	} else {
		i.Body = body
	}

	return i
}
