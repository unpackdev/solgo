package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) traverseFunctionDefinition(node *ast_pb.Node, fd *parser.FunctionDefinitionContext) *ast_pb.Node {
	// Extract the function name.
	node.Name = fd.Identifier().GetText()

	// Set the function type and its kind.
	node.NodeType = ast_pb.NodeType_FUNCTION_DEFINITION
	node.Kind = ast_pb.NodeType_FUNCTION_DEFINITION

	// If block is not empty we are going to assume that the function is implemented.
	// @TODO: Take assumption to the next level in the future.
	node.Implemented = !fd.Block().IsEmpty()

	// Get function visibility state.
	for _, visibility := range fd.AllVisibility() {
		node.Visibility = visibility.GetText()
	}

	// Get function state mutability.
	for _, stateMutability := range fd.AllStateMutability() {
		node.StateMutability = stateMutability.GetText()
	}

	// Get function modifiers.
	for _, modifier := range fd.AllModifierInvocation() {
		_ = modifier
		//node.Modifiers = append(node.Modifiers, modifier.GetText())
	}

	// Check if function is virtual.
	for _, virtual := range fd.AllVirtual() {
		node.Virtual = virtual.GetText() == "virtual"
	}

	// Check if function is override.
	// @TODO: Implement override specification.
	for _, override := range fd.AllOverrideSpecifier() {
		_ = override
	}

	// Extract function parameters.
	if len(fd.AllParameterList()) > 0 {
		node.Parameters = b.traverseParameterList(node, fd.AllParameterList()[0])
	}

	// Extract function return parameters.
	node.ReturnParameters = b.traverseParameterList(node, fd.GetReturnParameters())

	// And now we are going to the big league. We are going to traverse the function body.
	if !fd.Block().IsEmpty() {
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(fd.Block().GetStart().GetLine()),
				Column:      int64(fd.Block().GetStart().GetColumn()),
				Start:       int64(fd.Block().GetStart().GetStart()),
				End:         int64(fd.Block().GetStop().GetStop()),
				Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range fd.Block().AllStatement() {
			if statement.IsEmpty() {
				continue
			}
			bodyNode.Statements = append(bodyNode.Statements, b.traverseStatement(node, bodyNode, statement))
		}

		node.Body = bodyNode
	}

	return node
}
