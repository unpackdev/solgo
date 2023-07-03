package ast

// InterfaceNode represents an interface definition in Solidity.
type InterfaceNode struct {
	Name      string          `json:"name"`
	Functions []*FunctionNode `json:"functions"`
}

func (i *InterfaceNode) Children() []Node {
	nodes := make([]Node, len(i.Functions))
	for i, function := range i.Functions {
		nodes[i] = function
	}
	return nodes
}
