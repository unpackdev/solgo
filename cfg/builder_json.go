package cfg

import (
	"fmt"
	"github.com/goccy/go-json"
)

// ToJSON converts a specified contract or the entire graph to a JSON representation.
// This method is part of the Builder type which presumably builds or manages the CFG.
//
// If a contractName is provided, the method attempts to find and convert only the specified
// contract node within the graph to JSON. If the specified contract is not found, it returns
// an error.
//
// If no contractName is provided (i.e., an empty string), the method converts the entire graph
// to JSON, representing all nodes within the graph.
func (b *Builder) ToJSON(contractName string) ([]byte, error) {
	if contractName == "" {
		return json.Marshal(b.GetGraph().Nodes)
	}

	node, exists := b.GetGraph().Nodes[contractName]
	if !exists {
		return []byte{}, fmt.Errorf("contract %s not found in the graph", contractName)
	}
	return json.Marshal(node)
}
