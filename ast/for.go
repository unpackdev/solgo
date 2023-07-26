package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ForStatement struct {
	*ASTBuilder

	Id          int64           `json:"id"`
	NodeType    ast_pb.NodeType `json:"node_type"`
	Src         SrcNode         `json:"src"`
	Initialiser Node[NodeType]  `json:"initialiser"`
	Condition   Node[NodeType]  `json:"condition"`
	Closure     Node[NodeType]  `json:"closure"`
	Body        *BodyNode       `json:"body"`
}

func NewForStatement(b *ASTBuilder) *ForStatement {
	return &ForStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_FOR_STATEMENT,
	}
}

func (f *ForStatement) GetId() int64 {
	return f.Id
}

func (f *ForStatement) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *ForStatement) GetSrc() SrcNode {
	return f.Src
}

func (f *ForStatement) GetInitialiser() Node[NodeType] {
	return f.Initialiser
}

func (f *ForStatement) GetCondition() Node[NodeType] {
	return f.Condition
}

func (f *ForStatement) GetClosure() Node[NodeType] {
	return f.Closure
}

func (f *ForStatement) GetBody() *BodyNode {
	return f.Body
}

func (f *ForStatement) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

func (f *ForStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (f *ForStatement) ToProto() NodeType {
	return ast_pb.For{}
}

// https://docs.soliditylang.org/en/v0.8.19/grammar.html#a4.SolidityParser.forStatement
func (f *ForStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ForStatementContext,
) Node[NodeType] {
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	if ctx.SimpleStatement() != nil {
		statement := NewSimpleStatement(f.ASTBuilder)
		f.Initialiser = statement.Parse(
			unit, contractNode, fnNode, bodyNode, ctx.SimpleStatement().(*parser.SimpleStatementContext),
		)
	}

	if ctx.ExpressionStatement() != nil {
		expr := NewExpressionStatement(f.ASTBuilder)
		f.Condition = expr.Parse(
			unit, contractNode, fnNode, bodyNode, ctx.ExpressionStatement().(*parser.ExpressionStatementContext),
		)
	}

	if ctx.Expression() != nil {
		expression := NewExpression(f.ASTBuilder)
		f.Closure = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())
	}

	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Statement().Block())
		f.Body = bodyNode

		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(f.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, f, uncheckedCtx)
				f.Body.Statements = append(f.Body.Statements, bodyNode)
			}
		}
	}

	return f
}
