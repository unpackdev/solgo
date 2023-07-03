package ast

// Node is the interface that all AST nodes implement.
type Node interface {
	// Children returns the child nodes of this node.
	Children() []Node
}
