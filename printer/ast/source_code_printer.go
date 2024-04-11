package printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
	"go.uber.org/zap"
)

// Print is a function that prints the AST nodes to source code
func Print(node ast.Node[ast.NodeType]) (string, bool) {
	sb := strings.Builder{}
	success := PrintRecursive(node, &sb, 0)
	return sb.String(), success
}

// PrintRecursive is a function that prints the AST nodes to source code recursively
func PrintRecursive(node ast.Node[ast.NodeType], sb *strings.Builder, depth int) bool {
	if node == nil {
		zap.S().Error("Node is nil")
		return false
	}
	switch node := node.(type) {
	case *ast.AndOperation:
		return printAndOperation(node, sb, depth)
	default:
		zap.S().Errorf("Unknown node type: %T\n", node)
		return false
	}
}

func writeSeperatedStrings(sb *strings.Builder, seperator string, s ...string) {
	count := 0
	for _, item := range s {
		// Skip empty strings
		if item == "" {
			continue
		}

		if count > 0 {
			sb.WriteString(seperator)
			sb.WriteString(item)
		} else {
			sb.WriteString(item)
		}
	}
}

func writeSeperatedList(sb *strings.Builder, seperator string, s []string) {
	count := 0
	for _, item := range s {
		// Skip empty strings
		if item == "" {
			continue
		}

		if count > 0 {
			sb.WriteString(seperator)
			sb.WriteString(item)
		} else {
			sb.WriteString(item)
		}
	}
}

func writeStrings(sb *strings.Builder, s ...string) {
	for _, item := range s {
		sb.WriteString(item)
	}
}