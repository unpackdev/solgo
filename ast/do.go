package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type DoWhiteStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      *BodyNode       `json:"body"`
}

func NewDoWhiteStatement(b *ASTBuilder) *DoWhiteStatement {
	return &DoWhiteStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_DO_WHILE_STATEMENT,
	}
}

func (d *DoWhiteStatement) GetId() int64 {
	return d.Id
}

func (d *DoWhiteStatement) GetType() ast_pb.NodeType {
	return d.NodeType
}

func (d *DoWhiteStatement) GetSrc() SrcNode {
	return d.Src
}

func (d *DoWhiteStatement) GetCondition() Node[NodeType] {
	return d.Condition
}

func (d *DoWhiteStatement) GetBody() *BodyNode {
	return d.Body
}

func (d *DoWhiteStatement) GetNodes() []Node[NodeType] {
	return d.Body.Statements
}

func (d *DoWhiteStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (d *DoWhiteStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (d *DoWhiteStatement) Parse(
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
