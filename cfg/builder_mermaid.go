package cfg

import (
	"fmt"
	"strings"
)

// ToMermaid generates a string representation of the control flow graph (CFG) in Mermaid syntax.
// Mermaid is a tool that generates diagrams and flowcharts from text in a similar manner as markdown.
//
// The function iterates through each node in the graph and constructs the Mermaid syntax accordingly.
// It represents each contract in the graph as a node in the Mermaid graph. Entry contracts are
// distinguished with a special notation. The function also visualizes dependencies and inheritance
// relationships between contracts using arrows in the Mermaid syntax.
//
// If the graph is nil or contains no nodes, the function returns a default string indicating that
// no contracts are found.
func (b *Builder) ToMermaid() string {
	if b.graph == nil || len(b.graph.Nodes) == 0 {
		return "graph LR\n    No_Contracts[No contracts found]"
	}

	var mermaidGraph strings.Builder
	mermaidGraph.WriteString("graph LR\n")
	for _, node := range b.graph.Nodes {
		nodeName := node.Name

		if node.EntryContract {
			mermaidGraph.WriteString(fmt.Sprintf("    %s[(%s - Entry)]\n", nodeName, nodeName))
		} else {
			mermaidGraph.WriteString(fmt.Sprintf("    %s[%s]\n", nodeName, nodeName))
		}

		for _, imp := range node.Imports {
			importedName := imp.GetAbsolutePath()
			mermaidGraph.WriteString(fmt.Sprintf("    %s --> %s\n", nodeName, importedName))
		}

		for _, inherit := range node.Inherits {
			inheritedName := inherit.BaseName.Name
			mermaidGraph.WriteString(fmt.Sprintf("    %s -->|inherits| %s\n", nodeName, inheritedName))
		}
	}

	return mermaidGraph.String()
}
