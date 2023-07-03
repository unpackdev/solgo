package ast

// VariableNode represents a variable definition in Solidity.
type VariableNode struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (v *VariableNode) Children() []Node {
	// Variables have no children.
	return nil
}

type StateVariableNode struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Visibility   string `json:"visibility"`
	IsConstant   bool   `json:"is_constant"`
	IsImmutable  bool   `json:"is_immutable"`
	InitialValue string `json:"initial_value"`
}

func (v *StateVariableNode) Children() []Node {
	// Variables have no children.
	return nil
}
