package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// SimpleStatement represents a simple statement in the AST.
type SimpleStatement struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

// NewSimpleStatement creates a new instance of SimpleStatement using the provided ASTBuilder.
func NewSimpleStatement(b *ASTBuilder) *SimpleStatement {
	return &SimpleStatement{
		ASTBuilder: b,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the SimpleStatement node.
func (s *SimpleStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the SimpleStatement node.
func (s *SimpleStatement) GetId() int64 {
	return s.Id
}

// GetType returns the NodeType of the SimpleStatement node.
func (s *SimpleStatement) GetType() ast_pb.NodeType {
	return s.NodeType
}

// GetSrc returns the source information of the SimpleStatement node.
func (s *SimpleStatement) GetSrc() SrcNode {
	return s.Src
}

// GetTypeDescription returns the type description of the SimpleStatement node.
func (s *SimpleStatement) GetTypeDescription() *TypeDescription {
	return nil
}

// GetNodes returns an empty list of child nodes for the SimpleStatement node.
func (s *SimpleStatement) GetNodes() []Node[NodeType] {
	return nil
}

// ToProto always returns nil for the SimpleStatement node.
func (s *SimpleStatement) ToProto() NodeType {
	return nil
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
