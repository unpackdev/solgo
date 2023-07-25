package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type SimpleStatement[T NodeType] struct {
	*ASTBuilder

	Id          int64           `json:"id"`
	NodeType    ast_pb.NodeType `json:"node_type"`
	Src         SrcNode         `json:"src"`
	Assignments []int64         `json:"assignments"`
}

func NewSimpleStatement[T any](b *ASTBuilder) *SimpleStatement[T] {
	return &SimpleStatement[T]{
		ASTBuilder:  b,
		Assignments: make([]int64, 0),
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

func (s *SimpleStatement[T]) GetAssignments() []int64 {
	return s.Assignments
}

func (s *SimpleStatement[T]) ToProto() ast_pb.Statement {
	return ast_pb.Statement{}
}

func (s *SimpleStatement[T]) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], contractNode Node[NodeType], bodyNode *BodyNode[T], ctx *parser.SimpleStatementContext) {
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

	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.VariableDeclarationStatementContext:
			s.NodeType = ast_pb.NodeType_VARIABLE_DECLARATION_STATEMENT
			s.parseVariableDeclarationStatement(unit, contractNode, bodyNode, childCtx)
		default:
			panic("Unknown simple statement child type @ SimpleStatement.Parse")
		}
	}

	s.dumpNode(s)
}

func (s *SimpleStatement[T]) parseVariableDeclarationStatement(unit *SourceUnit[Node[ast_pb.SourceUnit]], contractNode Node[NodeType], bodyNode *BodyNode[T], ctx *parser.VariableDeclarationStatementContext) {

}
