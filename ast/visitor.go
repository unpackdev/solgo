package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

type NodeVisitor struct {
	Visit     func(node Node[NodeType]) bool
	TypeVisit map[ast_pb.NodeType][]func(node Node[NodeType]) (bool, error)
}

// RegisterTypeVisit allows registering multiple visitation functions for a specific node type.
func (nv *NodeVisitor) RegisterTypeVisit(nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) (bool, error)) {
	if nv.TypeVisit == nil {
		nv.TypeVisit = make(map[ast_pb.NodeType][]func(node Node[NodeType]) (bool, error))
	}
	nv.TypeVisit[nodeType] = append(nv.TypeVisit[nodeType], visitFunc)
}

// Walk traverses the AST nodes starting from the initial set of nodes in the root.
func (t *Tree) Walk(visitor *NodeVisitor) error {
	return t.WalkNodes(t.astRoot.GetNodes(), visitor)
}

// Update WalkNode and WalkNodes to handle return values from walkRecursive
func (t *Tree) WalkNode(startNode Node[NodeType], visitor *NodeVisitor) error {
	_, err := t.walkRecursive(startNode, visitor)
	return err
}

func (t *Tree) WalkNodes(startNodes []Node[NodeType], visitor *NodeVisitor) error {
	for _, node := range startNodes {
		_, err := t.walkRecursive(node, visitor)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tree) ExecuteTypeVisit(nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) (bool, error)) (bool, error) {
	return t.executeTypeVisitRecursive(t.astRoot.GetNodes(), nodeType, visitFunc)
}

func (t *Tree) ExecuteCustomTypeVisit(nodes []Node[NodeType], nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) (bool, error)) (bool, error) {
	return t.executeTypeVisitRecursive(nodes, nodeType, visitFunc)
}

func (t *Tree) executeTypeVisitRecursive(nodes []Node[NodeType], nodeType ast_pb.NodeType, visitFunc func(node Node[NodeType]) (bool, error)) (bool, error) {
	for _, node := range nodes {
		if node == nil {
			continue
		}

		// Execute the visit function if the node type matches
		if node.GetType() == nodeType {
			if status, err := visitFunc(node); err != nil {
				return status, err
			} else if !status {
				return status, err
			}
		}

		// Recursively call this function for all child nodes
		status, err := t.executeTypeVisitRecursive(node.GetNodes(), nodeType, visitFunc)
		if err != nil {
			return status, err
		}
	}

	return false, nil
}

func (t *Tree) walkRecursive(node Node[NodeType], visitor *NodeVisitor) (bool, error) {
	if node == nil {
		return true, nil
	}

	// Execute all registered functions sequentially for this node type
	if visitFuncs, ok := visitor.TypeVisit[node.GetType()]; ok {
		for _, visitFunc := range visitFuncs {
			status, err := visitFunc(node)
			if err != nil {
				return status, err
			}
			if !status {
				return status, nil
			}
		}
	} else if visitor.Visit != nil {
		status := visitor.Visit(node)
		if !status {
			return status, nil
		}
	}

	// Recursively call this function for all child nodes
	for _, child := range node.GetNodes() {
		status, err := t.walkRecursive(child, visitor)
		if err != nil {
			return status, err
		}
		if !status {
			return status, nil
		}
	}

	return true, nil
}
