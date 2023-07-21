package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseFallbackFunctionDefinition(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, ctx *parser.FallbackFunctionDefinitionContext) *ast_pb.Node {
	node.Name = "fallback"
	node.NodeType = ast_pb.NodeType_FUNCTION_DEFINITION
	node.Kind = ast_pb.NodeType_FALLBACK
	// If block is not empty we are going to assume that the function is implemented.
	// @TODO: Take assumption to the next level in the future.
	node.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			node.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	for _, virtual := range ctx.AllVirtual() {
		if virtual.GetText() == "external" {
			node.Virtual = true
		}
	}

	// Get function state mutability.
	for _, stateMutability := range ctx.AllStateMutability() {
		if stateMutability.GetText() == "" {
			node.StateMutability = ast_pb.Mutability_IMMUTABLE
		} else if stateMutability.GetText() == "payable" {
			node.StateMutability = ast_pb.Mutability_PAYABLE
		} else if stateMutability.GetText() == "pure" {
			node.StateMutability = ast_pb.Mutability_PURE
		} else if stateMutability.GetText() == "view" {
			node.StateMutability = ast_pb.Mutability_VIEW
		} else {
			node.StateMutability = ast_pb.Mutability_MUTABLE
		}
	}

	if node.StateMutability == ast_pb.Mutability_M_DEFAULT {
		node.StateMutability = ast_pb.Mutability_NONPAYABLE
	}

	node.Parameters = b.traverseParameterList(sourceUnit, node, ctx.ParameterList(0))
	if node.Parameters == nil {
		node.Parameters = &ast_pb.ParametersList{
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
	}

	node.ReturnParameters = b.traverseParameterList(sourceUnit, node, ctx.GetReturnParameters())
	if node.ReturnParameters == nil {
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
	}

	// And now we are going to the big league. We are going to traverse the function body.
	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(ctx.Block().GetStart().GetLine()),
				Column:      int64(ctx.Block().GetStart().GetColumn()),
				Start:       int64(ctx.Block().GetStart().GetStart()),
				End:         int64(ctx.Block().GetStop().GetStop()),
				Length:      int64(ctx.Block().GetStop().GetStop() - ctx.Block().GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range ctx.Block().AllStatement() {
			if statement.IsEmpty() {
				continue
			}

			// Parent index statement in this case is used only to be able provide
			// index to the parent node. It is not used for anything else.
			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
				sourceUnit, node, bodyNode, parentIndexStmt, statement,
			))
		}

		node.Body = bodyNode
	}

	if ctx.Block() != nil && len(ctx.Block().AllUncheckedBlock()) > 0 {
		for _, uncheckedBlockCtx := range ctx.Block().AllUncheckedBlock() {
			bodyNode := &ast_pb.Body{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(ctx.Block().GetStart().GetLine()),
					Column:      int64(ctx.Block().GetStart().GetColumn()),
					Start:       int64(ctx.Block().GetStart().GetStart()),
					End:         int64(ctx.Block().GetStop().GetStop()),
					Length:      int64(ctx.Block().GetStop().GetStop() - ctx.Block().GetStart().GetStart() + 1),
					ParentIndex: node.Id,
				},
				NodeType: ast_pb.NodeType_UNCHECKED_BLOCK,
			}

			if uncheckedBlockCtx.Block() != nil && !uncheckedBlockCtx.Block().IsEmpty() {
				for _, statement := range uncheckedBlockCtx.Block().AllStatement() {
					if statement.IsEmpty() {
						continue
					}

					// Parent index statement in this case is used only to be able provide
					// index to the parent node. It is not used for anything else.
					parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

					bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
						sourceUnit, node, bodyNode, parentIndexStmt, statement,
					))
				}
			}

			node.Body = bodyNode
		}
	}

	return node
}
