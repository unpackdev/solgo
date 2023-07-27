package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type SimpleStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewSimpleStatement(b *ASTBuilder) *SimpleStatement {
	return &SimpleStatement{
		ASTBuilder: b,
	}
}

func (s *SimpleStatement) GetId() int64 {
	return s.Id
}

func (s *SimpleStatement) GetType() ast_pb.NodeType {
	return s.NodeType
}

func (s *SimpleStatement) GetSrc() SrcNode {
	return s.Src
}

func (s *SimpleStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (s *SimpleStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (s *SimpleStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

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
			panic(
				fmt.Sprintf(
					"Unknown simple statement child type @ SimpleStatement.Parse: %T",
					childCtx,
				),
			)
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
