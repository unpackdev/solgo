package ast

import (
	"fmt"
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type SimpleStatement[T NodeType] struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewSimpleStatement[T any](b *ASTBuilder) *SimpleStatement[T] {
	return &SimpleStatement[T]{
		ASTBuilder: b,
	}
}

func (s *SimpleStatement[T]) GetId() int64 {
	return s.Id
}

func (s *SimpleStatement[T]) GetType() ast_pb.NodeType {
	return s.NodeType
}

func (s *SimpleStatement[T]) GetSrc() SrcNode {
	return s.Src
}

func (s *SimpleStatement[T]) GetTypeDescription() TypeDescription {
	return TypeDescription{}
}

func (s *SimpleStatement[T]) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (s *SimpleStatement[T]) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.SimpleStatementContext,
) Node[NodeType] {

	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.VariableDeclarationStatementContext:
			varDeclar := NewVariableDeclarationStatement(s.ASTBuilder)
			varDeclar.Parse(unit, contractNode, fnNode, bodyNode, childCtx)
			return varDeclar
		case *parser.ExpressionStatementContext:
			expr := NewExpressionStatement(s.ASTBuilder)
			return expr.Parse(unit, contractNode, fnNode, bodyNode, childCtx)
		default:
			fmt.Println("Statement child type:", reflect.TypeOf(childCtx))
			panic("Unknown simple statement child type @ SimpleStatement.Parse")
		}
	}

	s.Id = s.GetNextID()
	s.Src = SrcNode{
		Id:          s.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	return s
}
