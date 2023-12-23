package cfg

import (
	"fmt"

	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

// Node represents a node in the graph, encapsulating information about a Solidity contract.
// It includes the contract's name, IR representation, import dependencies, inherited contracts,
// and a flag indicating if it's the entry contract.
type Node struct {
	Name          string              `json:"name"`
	Contract      *ir.Contract        `json:"-"`
	Imports       []*ir.Import        `json:"imports"`
	Inherits      []*ast.BaseContract `json:"inherits"`
	EntryContract bool                `json:"entry_contract"`
}

// IsEntryContract returns true if the node represents an entry contract in the graph.
func (n *Node) IsEntryContract() bool {
	return n.EntryContract
}

// GetContract returns the IR representation of the Solidity contract encapsulated by the node.
func (n *Node) GetContract() *ir.Contract {
	return n.Contract
}

// GetImports returns a slice of all the import dependencies of the contract.
func (n *Node) GetImports() []*ir.Import {
	return n.Imports
}

// GetInherits returns a slice of all the contracts inherited by the contract.
func (n *Node) GetInherits() []*ast.BaseContract {
	return n.Inherits
}

// GetName returns the name of the Solidity contract.
func (n *Node) GetName() string {
	return n.Name
}

// GetImportNames returns a slice of the absolute paths of all import dependencies of the contract.
func (n *Node) GetImportNames() []string {
	var names []string
	for _, imp := range n.Imports {
		names = append(names, imp.GetAbsolutePath())
	}
	return names
}

// GetInheritedContractNames returns a slice of names of all contracts inherited by the contract.
func (n *Node) GetInheritedContractNames() []string {
	var names []string
	for _, inherit := range n.Inherits {
		names = append(names, inherit.BaseName.GetName())
	}
	return names
}

// ToString provides a string representation of the Node, including its name, entry contract status,
// imports, and inherited contracts. This is useful for debugging and logging purposes.
func (n *Node) ToString() string {
	return fmt.Sprintf("Node{Name: %s, EntryContract: %t, Imports: %v, Inherits: %v}",
		n.Name, n.EntryContract, n.GetImportNames(), n.GetInheritedContractNames())
}
