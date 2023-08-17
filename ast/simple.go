package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// SimpleStatement represents a simple statement in the AST.
type SimpleStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

// NewSimpleStatement creates a new instance of SimpleStatement using the provided ASTBuilder.
// This instance is more like a placeholder for the actual statements that are returned from Parse()
func NewSimpleStatement(b *ASTBuilder) *SimpleStatement {
	return &SimpleStatement{
		ASTBuilder: b,
	}
}

// Parse parses the SimpleStatement node from the provided context.
func (s *SimpleStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	parentNode Node[NodeType],
	ctx *parser.SimpleStatementContext,
) Node[NodeType] {
	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.VariableDeclarationStatementContext:
			varDeclar := NewVariableDeclarationStatement(s.ASTBuilder)
			varDeclar.Parse(unit, contractNode, fnNode, bodyNode, childCtx)
			return varDeclar
		case *parser.ExpressionStatementContext:
			return parseExpressionStatement(
				s.ASTBuilder,
				unit, contractNode, fnNode, bodyNode, parentNode, childCtx,
			)
		default:
			zap.L().Warn(
				"Unknown simple statement child type @ SimpleStatement.Parse",
				zap.String("child", fmt.Sprintf("%T", childCtx)),
			)
		}
	}

	return nil
}
