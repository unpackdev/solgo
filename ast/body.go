package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type BodyNode[T any] struct {
	*ASTBuilder

	Id          int64           `json:"id"`
	NodeType    ast_pb.NodeType `json:"node_type"`
	Kind        ast_pb.NodeType `json:"kind"`
	Src         SrcNode         `json:"src"`
	Implemented bool            `json:"implemented"`
	Statements  []T             `json:"statements"`
}

func NewBodyNode[T any](b *ASTBuilder) *BodyNode[T] {
	return &BodyNode[T]{
		ASTBuilder: b,
		Statements: make([]T, 0),
	}
}

func (b *BodyNode[T]) GetId() int64 {
	return b.Id
}

func (b *BodyNode[T]) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *BodyNode[T]) GetSrc() SrcNode {
	return b.Src
}

func (b *BodyNode[T]) Parse(
	unit *SourceUnit[Node],
	contractNode Node,
	bodyCtx parser.IContractBodyElementContext,
) Node {
	for _, bodyChildCtx := range bodyCtx.GetChildren() {
		switch childCtx := bodyChildCtx.(type) {
		case *parser.FunctionDefinitionContext:
			fn := NewFunctionNode[Node](b.ASTBuilder)
			return fn.Parse(unit, contractNode, bodyCtx, childCtx)
		}
	}

	// Could not find any function definitions so we'll just return the body node.
	b.Id = b.GetNextID()
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(bodyCtx.GetStart().GetLine()),
		Column:      int64(bodyCtx.GetStart().GetColumn()),
		Start:       int64(bodyCtx.GetStart().GetStart()),
		End:         int64(bodyCtx.GetStop().GetStop()),
		Length:      int64(bodyCtx.GetStop().GetStop() - bodyCtx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	return b
}
