package ast

// RootNode represents the root of the AST and can contain multiple contracts.
type RootNode struct {
	Interfaces []*InterfaceNode `json:"interfaces"`
	Contracts  []*ContractNode  `json:"contracts"`
}

func (r *RootNode) Children() []Node {
	// Convert the slices of ContractNodes and InterfaceNodes to a slice of Nodes.
	nodes := make([]Node, len(r.Contracts)+len(r.Interfaces))
	for i, contract := range r.Contracts {
		nodes[i] = contract
	}
	for i, iface := range r.Interfaces {
		nodes[i+len(r.Contracts)] = iface
	}
	return nodes
}
