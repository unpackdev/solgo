package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type DoWhileStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      *BodyNode       `json:"body"`
}

func NewDoWhileStatement(b *ASTBuilder) *DoWhileStatement {
	return &DoWhileStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_DO_WHILE_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the DoWhileStatement node.
func (d *DoWhileStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (d *DoWhileStatement) GetId() int64 {
	return d.Id
}

func (d *DoWhileStatement) GetType() ast_pb.NodeType {
	return d.NodeType
}

func (d *DoWhileStatement) GetSrc() SrcNode {
	return d.Src
}

func (d *DoWhileStatement) GetCondition() Node[NodeType] {
	return d.Condition
}

func (d *DoWhileStatement) GetBody() *BodyNode {
	return d.Body
}

func (d *DoWhileStatement) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{d.Condition}
	toReturn = append(toReturn, d.Body.GetNodes()...)
	return toReturn
}

func (d *DoWhileStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (d *DoWhileStatement) ToProto() NodeType {
	protos := ast_pb.Do{
		Id:        d.GetId(),
		NodeType:  d.GetType(),
		Src:       d.GetSrc().ToProto(),
		Condition: d.GetCondition().ToProto().(*v3.TypedStruct),
		Body:      d.GetBody().ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&protos, "Do")
}

func (d *DoWhileStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.DoWhileStatementContext,
) Node[NodeType] {
	d.Src = SrcNode{
		Id:          d.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(d.ASTBuilder)
	d.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())

	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(d.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, d, ctx.Statement().Block())
		d.Body = bodyNode

		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(d.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, d, uncheckedCtx)
				d.Body.Statements = append(d.Body.Statements, bodyNode)
			}
		}
	}

	return d
}
