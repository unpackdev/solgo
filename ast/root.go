package ast

// RootNode represents the root of the AST and can contain multiple contracts.
type RootNode struct {
	Contracts []*ContractNode `json:"contracts"`
}

func (r *RootNode) Children() []Node {
	nodes := make([]Node, len(r.Contracts))
	for i, contract := range r.Contracts {
		nodes[i] = contract
	}
	return nodes
}
