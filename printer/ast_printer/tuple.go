package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printTupleExpression(node *ast.TupleExpression, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("(")
	components := []string{}
	for _, component := range node.GetComponents() {
		s, ok := Print(component)
		success = ok && success
		components = append(components, s)
	}
	sb.WriteString(strings.Join(components, ", "))
	sb.WriteString(")")
	return success
}
