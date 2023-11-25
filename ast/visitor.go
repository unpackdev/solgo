package ast

import ast_pb "github.com/unpackdev/protos/dist/go/ast"

type NodeVisitor struct {
	Visit     func(node Node[NodeType]) bool
	TypeVisit map[ast_pb.NodeType][]func(node Node[NodeType]) bool
}

// RegisterTypeVisit allows registering multiple visitation functions for a specific node type.
func (nv *NodeVisitor) RegisterTypeVisit(nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) bool) {
	if nv.TypeVisit == nil {
		nv.TypeVisit = make(map[ast_pb.NodeType][]func(node Node[NodeType]) bool)
	}
	nv.TypeVisit[nodeType] = append(nv.TypeVisit[nodeType], visitFunc)
}

// Walk traverses the AST nodes starting from the initial set of nodes in the root.
func (t *Tree) Walk(visitor *NodeVisitor) {
	t.WalkNodes(t.astRoot.GetNodes(), visitor)
}

// WalkNode traverses the AST nodes starting from the specified node.
func (t *Tree) WalkNode(startNode Node[NodeType], visitor *NodeVisitor) {
	// Start the recursive walk from the specified node
	t.walkRecursive(startNode, visitor)
}

// WalkNode traverses the AST nodes starting from the specified node.
func (t *Tree) WalkNodes(startNode []Node[NodeType], visitor *NodeVisitor) {
	// Start the recursive walk from the specified nodes
	for _, node := range t.astRoot.GetNodes() {
		t.walkRecursive(node, visitor)
	}
}

func (t *Tree) ExecuteTypeVisit(nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) bool) {
	t.executeTypeVisitRecursive(t.astRoot.GetNodes(), nodeType, visitFunc)
}

func (t *Tree) executeTypeVisitRecursive(nodes []Node[NodeType], nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) bool) {
	for _, node := range nodes {
		if node == nil {
			continue
		}

		// Execute the visit function if the node type matches
		if node.GetType() == nodeType {
			if !visitFunc(node) {
				return // Stop if the visit function returns false
			}
		}

		// Recursively call this function for all child nodes
		t.executeTypeVisitRecursive(node.GetNodes(), nodeType, visitFunc)
	}
}

func (t *Tree) walkRecursive(node Node[NodeType], visitor *NodeVisitor) {
	if node == nil {
		return
	}

	continueTraversal := true

	// Execute all registered functions sequentially for this node type
	if visitFuncs, ok := visitor.TypeVisit[node.GetType()]; ok {
		for _, visitFunc := range visitFuncs {
			if !visitFunc(node) {
				continueTraversal = false
				break
			}
		}
	} else if visitor.Visit != nil {
		continueTraversal = visitor.Visit(node)
	}

	if continueTraversal {
		for _, child := range node.GetNodes() {
			t.walkRecursive(child, visitor)
		}
	}
}
