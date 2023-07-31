package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"go.uber.org/zap"
)

// Tree is a structure that represents an Abstract Syntax Tree (AST).
// It contains a reference to an ASTBuilder and the root node of the AST.
type Tree struct {
	*ASTBuilder

	// astRoot is the root node of the Abstract Syntax Tree.
	astRoot *RootNode
}

// NewTree creates a new Tree with the provided ASTBuilder.
func NewTree(b *ASTBuilder) *Tree {
	return &Tree{
		ASTBuilder: b,
	}
}

// SetRoot sets the root node of the Abstract Syntax Tree (AST).
func (t *Tree) SetRoot(root *RootNode) {
	t.astRoot = root
}

// AppendRootNodes appends the provided SourceUnit nodes to the root node of the AST.
func (t *Tree) AppendRootNodes(roots ...*SourceUnit[Node[ast_pb.SourceUnit]]) {
	t.astRoot.SourceUnits = append(t.astRoot.SourceUnits, roots...)
}

// GetRoot returns the root node of the Abstract Syntax Tree (AST).
func (t *Tree) GetRoot() *RootNode {
	return t.astRoot
}

// GetById attempts to find a node in the AST by its ID.
// It performs a recursive search through all nodes in the AST.
// Returns the found Node or nil if the node cannot be found.
func (t *Tree) GetById(id int64) Node[NodeType] {
	for _, node := range t.astRoot.GetNodes() {
		if node.GetId() == id {
			return node
		}

		if n := t.byRecursiveId(node, id); n != nil {
			return n
		}
	}
	return nil
}

// UpdateNodeReferenceById attempts to update the reference descriptor of a node in the AST by its ID.
// It performs a recursive search through all nodes in the AST.
// Returns true if the node was found and updated, false otherwise.
func (t *Tree) UpdateNodeReferenceById(nodeId int64, nodeRefId int64, typeRef *TypeDescription) bool {
	if nodeId == 0 || nodeRefId == 0 {
		zap.L().Warn(
			"Invalid arguments provided to UpdateNodeReferenceId",
			zap.Int64("nodeId", nodeId),
			zap.Int64("node_ref_id", nodeRefId),
			zap.Any("type_ref", typeRef),
		)
		return false
	}

	for _, child := range t.astRoot.GetNodes() {
		if n := t.byRecursiveReferenceUpdate(child, nodeId, nodeRefId, typeRef); n {
			return n
		}
	}

	return false
}

// byRecursiveReferenceUpdate is a helper function that attempts to update the reference descriptor of a node by its ID by recursively searching the node's children.
// Returns true if the node was found and updated, false otherwise.
func (t *Tree) byRecursiveReferenceUpdate(child Node[NodeType], nodeId int64, nodeRefId int64, typeRef *TypeDescription) bool {
	if child.GetId() == nodeId {
		child.SetReferenceDescriptor(nodeRefId, typeRef)
		return true
	}

	for _, c := range child.GetNodes() {
		if n := t.byRecursiveReferenceUpdate(c, nodeId, nodeRefId, typeRef); n {
			return n
		}
	}

	return false
}

// byRecursiveId is a helper function that attempts to find a node by its ID by recursively searching the node's children.
// Returns the found Node or nil if the node cannot be found.
func (t *Tree) byRecursiveId(node Node[NodeType], id int64) Node[NodeType] {
	if node.GetId() == id {
		return node
	}

	for _, child := range node.GetNodes() {
		if n := t.byRecursiveId(child, id); n != nil {
			return n
		}
	}

	return nil
}
