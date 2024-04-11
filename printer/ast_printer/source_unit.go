package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printSourceUnit(node ast.Node[ast.NodeType], sb *strings.Builder, depth int) bool {
	success := true
	for _, child := range node.GetNodes() {
		s, ok := Print(child.(ast.Node[ast.NodeType]))
		success = ok && success
		sb.WriteString(s)
		sb.WriteString("\n")
	}
	return success
}
