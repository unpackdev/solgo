package cfg

import (
	"errors"

	"github.com/unpackdev/solgo/ir"
)

// Variable represents a state variable within a smart contract.
// It encapsulates information about the state variable and its relationship
// to the contract it belongs to.
type Variable struct {
	Node          *Node             // Node is the contract where the state variable is defined.
	StateVariable *ir.StateVariable // StateVariable is the IR representation of the state variable.
	IsEntry       bool              // IsEntry indicates if this variable is from the entry contract.
}

// String returns the name of the state variable.
func (v *Variable) String() string {
	return v.StateVariable.GetName()
}

// GetNode returns the contract node associated with this state variable.
func (v *Variable) GetNode() *Node {
	return v.Node
}

// GetVariable returns the internal IR representation of the state variable.
func (v *Variable) GetVariable() *ir.StateVariable {
	return v.StateVariable
}

// GetContract returns the IR representation of the contract to which this variable belongs.
func (v *Variable) GetContract() *ir.Contract {
	return v.Node.Contract
}

// IsEntryContract checks if this variable is from the entry contract.
func (v *Variable) IsEntryContract() bool {
	return v.IsEntry
}

// GetOrderedStateVariables returns a slice of state variables in the order of their declaration.
// It fetches variables from the entry contract and follows the inheritance chain.
// Returns an error if the graph is not initialized or the entry contract is not found.
func (b *Builder) GetOrderedStateVariables() ([]*Variable, error) {
	if b.graph == nil {
		return nil, errors.New("graph is not initialized")
	}

	var entryContract *Node
	for _, node := range b.graph.Nodes {
		if node.EntryContract {
			entryContract = node
			break
		}
	}

	if entryContract == nil {
		return nil, errors.New("entry contract not found")
	}

	var variables []*Variable
	b.collectStateVariables(entryContract, &variables, entryContract.Name)

	return variables, nil
}

// collectStateVariables is a helper function to recursively collect state variables
// from a given node and its inherited contracts.
func (b *Builder) collectStateVariables(node *Node, variables *[]*Variable, entryContractName string) {
	// Collect state variables in reverse order of inheritance
	for i := len(node.Inherits) - 1; i >= 0; i-- {
		inheritedNode, exists := b.graph.Nodes[node.Inherits[i].BaseName.Name]
		if exists {
			b.collectStateVariables(inheritedNode, variables, entryContractName)
		}
	}

	// Add the state variables of the current node
	for _, stateVar := range node.Contract.GetStateVariables() {
		*variables = append(*variables, &Variable{
			Node:          node,
			StateVariable: stateVar,
			IsEntry:       node.Name == entryContractName,
		})
	}
}

// GetStorageStateVariables returns a slice of state variables relevant to storage.
// It follows the inheritance chain and collects variables from base contracts before the derived ones.
// Returns an error if the graph is not initialized or the entry contract is not found.
func (b *Builder) GetStorageStateVariables() ([]*Variable, error) {
	if b.graph == nil {
		return nil, errors.New("graph is not initialized")
	}

	entryContract := b.graph.GetNode(b.builder.GetRoot().GetEntryContract().GetName())
	if entryContract == nil {
		return nil, errors.New("entry contract not found")
	}

	var variables []*Variable
	b.collectStorageStateVariables(entryContract, &variables)

	return variables, nil
}

// collectStorageStateVariables is a helper function to recursively collect storage state variables
// from a given node and its base contracts.
func (b *Builder) collectStorageStateVariables(node *Node, variables *[]*Variable) {
	// First, add state variables of base contracts in the order of inheritance
	for _, baseNodeRef := range node.Inherits {
		baseNode := b.graph.GetNode(baseNodeRef.BaseName.GetName())
		if baseNode != nil {
			b.collectStorageStateVariables(baseNode, variables)
		}
	}

	// Then, add state variables of the current contract
	for _, stateVar := range node.Contract.GetStateVariables() {
		*variables = append(*variables, &Variable{
			Node:          node,
			StateVariable: stateVar,
			IsEntry:       node.Name == b.builder.GetRoot().GetEntryContract().GetName(),
		})
	}
}
