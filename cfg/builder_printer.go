package cfg

import (
	"fmt"
	"strings"
)

// Print displays information about a specific contract or the entry contract in the graph.
// If the contractName is provided, it prints details of the specified contract.
// If the contractName is empty, it finds and prints details of the entry contract in the graph.
//
// The function prints the contract's name, its status as an entry contract, and recursively
// prints information about its imports and inherited contracts, if any.
func (b *Builder) Print(contractName string) {
	if contractName == "" {
		for _, node := range b.graph.Nodes {
			if node.EntryContract {
				b.printNode(node.Name, 0)
				return
			}
		}
	} else {
		b.printNode(contractName, 0)
	}
}

// printNode is a helper function used by Print to recursively print details of a contract and its dependencies.
// It prints the contract's name, its status as an entry contract, and details of its imports and inherited contracts.
// The depth parameter is used for indentation purposes to represent the level of recursion and relationship in the graph.
func (b *Builder) printNode(name string, depth int) {
	var output strings.Builder

	node, exists := b.graph.Nodes[name]
	if !exists {
		output.WriteString(fmt.Sprintf("%sContract %s not found in the graph.\n", strings.Repeat("  ", depth), name))
		fmt.Print(output.String())
		return
	}

	indent := strings.Repeat("  ", depth)
	output.WriteString(fmt.Sprintf("%sNode %s:\n", indent, node.Name))
	output.WriteString(fmt.Sprintf("%s  Entry Contract: %v\n", indent, node.EntryContract))

	if len(node.Imports) == 0 && len(node.Inherits) == 0 {
		output.WriteString(fmt.Sprintf("%s  No imports or base contracts\n", indent))
	} else {
		for _, imp := range node.Imports {
			output.WriteString(fmt.Sprintf("%s  Imports: %s (ID: %d, File: %s)\n", indent, imp.GetAbsolutePath(), imp.GetId(), imp.GetFile()))
		}
		for _, inherit := range node.Inherits {
			inheritedContractName := inherit.BaseName.Name
			output.WriteString(fmt.Sprintf("%s  Inherits: %s\n", indent, inheritedContractName))
			b.printNode(inheritedContractName, depth+1)
		}
	}

	fmt.Print(output.String())
}
