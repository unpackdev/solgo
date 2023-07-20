package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseConstructorDefinition(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, ctx *parser.ConstructorDefinitionContext) *ast_pb.Node {
	node.NodeType = ast_pb.NodeType_FUNCTION_DEFINITION
	node.Kind = ast_pb.NodeType_CONSTRUCTOR
	node.StateMutability = ast_pb.Mutability_NONPAYABLE
	node.Visibility = ast_pb.Visibility_INTERNAL

	for _, payableCtx := range ctx.AllPayable() {
		if payableCtx.GetText() == "payable" {
			node.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	// If block is not empty we are going to assume that the function is implemented.
	// @TODO: Take assumption to the next level in the future.
	node.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	// Get function visibility state.
	for _, internalCtx := range ctx.AllInternal() {
		if internalCtx.GetText() == "internal" {
			node.Visibility = ast_pb.Visibility_INTERNAL
		}
	}

	for _, publicCtx := range ctx.AllPublic() {
		if publicCtx.GetText() == "public" {
			node.Visibility = ast_pb.Visibility_PUBLIC
		}
	}

	// Extract function parameters.
	if ctx.ParameterList() != nil {
		node.Parameters = b.traverseParameterList(node, ctx.ParameterList())
	}

	// Return arguments of cumstructor are pretty much useless as they are always empty.
	node.ReturnParameters = &ast_pb.ParametersList{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType:   ast_pb.NodeType_PARAMETER_LIST,
		Parameters: []*ast_pb.Parameter{},
	}

	return node
}
