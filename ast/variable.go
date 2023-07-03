package ast

// VariableNode represents a variable definition in Solidity.
type VariableNode struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Children returns an empty slice of nodes.
func (v *VariableNode) Children() []Node {
	return nil
}

// StateVariableNode represents a state variable definition in Solidity.
type StateVariableNode struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Visibility   string `json:"visibility"`
	IsConstant   bool   `json:"is_constant"`
	IsImmutable  bool   `json:"is_immutable"`
	InitialValue string `json:"initial_value"`
}

// Children returns an empty slice of nodes.
func (v *StateVariableNode) Children() []Node {
	return nil
}
