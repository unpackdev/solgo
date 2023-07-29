package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type WhileStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      *BodyNode       `json:"body"`
	Kind      ast_pb.NodeType `json:"kind"`
}

func NewWhileStatement(b *ASTBuilder) *WhileStatement {
	return &WhileStatement{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_WHILE_STATEMENT,
		Kind:       ast_pb.NodeType_WHILE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the WhileStatement node.
func (w *WhileStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (w *WhileStatement) GetId() int64 {
	return w.Id
}

func (w *WhileStatement) GetType() ast_pb.NodeType {
	return w.NodeType
}

func (w *WhileStatement) GetSrc() SrcNode {
	return w.Src
}

func (w *WhileStatement) GetCondition() Node[NodeType] {
	return w.Condition
}

func (w *WhileStatement) GetBody() *BodyNode {
	return w.Body
}

func (w *WhileStatement) GetKind() ast_pb.NodeType {
	return w.Kind
}

func (w *WhileStatement) GetImplemented() bool {
	return true
}

func (w *WhileStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (w *WhileStatement) GetNodes() []Node[NodeType] {
	return w.Body.Statements
}

func (w *WhileStatement) ToProto() NodeType {
	return ast_pb.While{}
}

func (w *WhileStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.WhileStatementContext,
) Node[NodeType] {
	w.Src = SrcNode{
		Id:          w.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(w.ASTBuilder)
	w.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, w, ctx.Expression())

	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(w.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, w, ctx.Statement().Block())
		w.Body = bodyNode

		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(w.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, w, uncheckedCtx)
				w.Body.Statements = append(w.Body.Statements, bodyNode)
			}
		}
	}

	return w
}
