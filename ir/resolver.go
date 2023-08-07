package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func (r *Builder) byFunction(name string) *Function {
	for _, unit := range r.astBuilder.GetRoot().GetSourceUnits() {
		contract := unit.GetContract()
		for _, node := range contract.GetNodes() {
			if function, ok := node.(*ast.Function); ok {
				if function.GetName() == name {
					return r.processFunction(function, false)
				}
			}
		}
	}

	return nil
}

// byRecursiveSearch is a helper function that attempts to resolve a node by its name by recursively searching the node's children.
// It returns the resolved Node and its TypeDescription, or nil if the node cannot be found.
func (r *Builder) byRecursiveSearch(node ast.Node[ast.NodeType], name string) (ast.Node[ast.NodeType], *ast.TypeDescription) {
	if node == nil || node.GetNodes() == nil {
		return nil, nil
	}

	switch nodeCtx := node.(type) {
	case *ast.DoWhileStatement:
		for _, condition := range nodeCtx.GetCondition().GetNodes() {
			if primary, ok := condition.(*ast.PrimaryExpression); ok {
				if primary.GetName() == name {
					return primary, primary.GetTypeDescription()
				}
			}
		}

	case *ast.ForStatement:
		for _, condition := range nodeCtx.GetCondition().GetNodes() {
			if primary, ok := condition.(*ast.PrimaryExpression); ok {
				if primary.GetName() == name {
					return primary, primary.GetTypeDescription()
				}
			}
		}
	}

	for _, n := range node.GetNodes() {
		if n == nil {
			continue
		}

		// Needs to be here as there are no parent nodes available so it wont be captured by the
		// main function block.
		if n.GetType() == ast_pb.NodeType_IDENTIFIER {
			if nodeCtx, ok := n.(*ast.PrimaryExpression); ok {
				if nodeCtx.GetName() == name {
					return nodeCtx, nodeCtx.GetTypeDescription()
				}
			}
		}

		if node, nodeType := r.byRecursiveSearch(n, name); node != nil && nodeType != nil {
			return node, nodeType
		}
	}

	return nil, nil
}
