package ast

// ContractNode represents a contract definition in Solidity.
type ContractNode struct {
	Name           string                `json:"name"`
	StateVariables []*StateVariableNode  `json:"variables"`
	Structs        []*StructNode         `json:"structs"`
	Events         []*EventNode          `json:"events"`
	Errors         []*ErrorNode          `json:"errors"`
	Constructor    *ConstructorNode      `json:"constructor"`
	Functions      []*FunctionNode       `json:"functions"`
	Kind           string                `json:"kind"`
	Inherits       []string              `json:"inherits"`
	Using          []*UsingDirectiveNode `json:"using"`
}

func (c *ContractNode) Children() []Node {
	nodes := make([]Node, len(c.Functions)+len(c.StateVariables)+len(c.Events)+len(c.Errors)+len(c.Structs)+len(c.Using))
	if c.Constructor != nil {
		nodes = append(nodes, c.Constructor)
	}
	for i, function := range c.Functions {
		nodes[i] = function
	}
	for i, variable := range c.StateVariables {
		nodes[i+len(c.Functions)] = variable
	}
	for i, event := range c.Events {
		nodes[i+len(c.Functions)+len(c.StateVariables)] = event
	}
	for i, err := range c.Errors {
		nodes[i+len(c.Functions)+len(c.StateVariables)+len(c.Events)] = err
	}
	for i, strct := range c.Structs {
		nodes[i+len(c.Functions)+len(c.StateVariables)+len(c.Events)+len(c.Errors)] = strct
	}
	for i, using := range c.Using {
		nodes[i+len(c.Functions)+len(c.StateVariables)+len(c.Events)+len(c.Errors)+len(c.Structs)] = using
	}
	return nodes
}
