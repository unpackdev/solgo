package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printEventDefinition(node *ast.EventDefinition, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("event ")
	sb.WriteString(node.GetName())
	sb.WriteString("(")
	printParameterList(node.GetParameters(), sb, depth)
	sb.WriteString(")")
	sb.WriteString(";\n")
	return success
}
