package cfg

import (
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

// Graph represents a directed graph of Nodes, with each node representing a Solidity contract.
// The graph captures the relationships and dependencies among contracts within a project.
type Graph struct {
	Nodes map[string]*Node
}

// NewGraph creates and returns a new instance of a Graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// AddNode adds a new Node to the Graph. It takes the contract name, its IR representation,
// and a boolean indicating if it is an entry contract. If the node already exists, it does nothing.
func (g *Graph) AddNode(name string, contract *ir.Contract, isEntryContract bool) {
	if _, exists := g.Nodes[name]; !exists {
		g.Nodes[name] = &Node{
			Name:          name,
			Contract:      contract,
			EntryContract: isEntryContract,
			Imports:       []*ir.Import{},
			Inherits:      []*ast.BaseContract{},
		}
	}
}

// AddDependency creates a dependency edge from one node to another by adding an import.
// The dependency is defined from 'from' node to 'to' import.
func (g *Graph) AddDependency(from string, to *ir.Import) {
	fromNode, exists := g.Nodes[from]
	if !exists {
		fromNode = &Node{Name: from}
		g.Nodes[from] = fromNode
	}
	fromNode.Imports = append(fromNode.Imports, to)
}

// GetNodes returns all nodes present in the Graph.
func (g *Graph) GetNodes() map[string]*Node {
	return g.Nodes
}

// GetNode retrieves a node by name from the Graph. It returns nil if the node does not exist.
func (g *Graph) GetNode(name string) *Node {
	if node, exists := g.Nodes[name]; exists {
		return node
	}
	return nil
}

// NodeExists checks if a node with the given name exists in the Graph.
func (g *Graph) NodeExists(name string) bool {
	_, exists := g.Nodes[name]
	return exists
}

// AddInheritance adds an inheritance relationship from one node to a base contract.
// The relationship is added to the 'from' node's Inherits slice.
func (g *Graph) AddInheritance(from string, to *ast.BaseContract) {
	fromNode, exists := g.Nodes[from]
	if !exists {
		fromNode = &Node{Name: from}
		g.Nodes[from] = fromNode
	}

	fromNode.Inherits = append(fromNode.Inherits, to)
}

// CountNodes returns the total number of nodes in the Graph.
func (g *Graph) CountNodes() int {
	return len(g.Nodes)
}
