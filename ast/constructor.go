package ast

// ConstructorNode represents a constructor definition in Solidity.
type ConstructorNode struct {
	Parameters []*VariableNode  `json:"parameters"`
	Body       []*StatementNode `json:"body"`
}

func (c *ConstructorNode) Children() []Node {
	nodes := make([]Node, len(c.Parameters)+len(c.Body))
	for i, parameter := range c.Parameters {
		nodes[i] = parameter
	}
	for i, statement := range c.Body {
		nodes[i+len(c.Parameters)] = statement
	}
	return nodes
}
