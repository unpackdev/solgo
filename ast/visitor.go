package ast

import ast_pb "github.com/unpackdev/protos/dist/go/ast"

type NodeVisitor struct {
	// Visit applies a function to each node.
	Visit func(node Node[NodeType]) bool

	// TypeVisit holds specific visitation functions for different node types.
	TypeVisit map[ast_pb.NodeType]func(node Node[NodeType]) bool
}

// Walk traverses the AST nodes starting from the initial set of nodes in the root.
func (t *Tree) Walk(visitor NodeVisitor) {
	t.WalkNodes(t.astRoot.GetNodes(), visitor)
}

// WalkNode traverses the AST nodes starting from the specified node.
func (t *Tree) WalkNode(startNode Node[NodeType], visitor NodeVisitor) {
	// Start the recursive walk from the specified node
	t.walkRecursive(startNode, visitor)
}

// WalkNode traverses the AST nodes starting from the specified node.
func (t *Tree) WalkNodes(startNode []Node[NodeType], visitor NodeVisitor) {
	// Start the recursive walk from the specified nodes
	for _, node := range t.astRoot.GetNodes() {
		t.walkRecursive(node, visitor)
	}
}

func (t *Tree) walkRecursive(node Node[NodeType], visitor NodeVisitor) {
	if node == nil {
		return
	}

	continueTraversal := true

	// Check if there's a specific visit function for this node type
	if typeVisitFunc, ok := visitor.TypeVisit[node.GetType()]; ok {
		continueTraversal = typeVisitFunc(node)
	} else if visitor.Visit != nil {
		continueTraversal = visitor.Visit(node)
	}

	if continueTraversal {
		for _, child := range node.GetNodes() {
			t.walkRecursive(child, visitor)
		}
	}
}
