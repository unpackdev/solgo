package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printModifierDefinition(node *ast.ModifierDefinition, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("modifier ")
	sb.WriteString(node.GetName())
	sb.WriteString("(")
	printParameterList(node.GetParameters(), sb, depth)
	sb.WriteString(") ")
	visibility := getVisibilityString(node.GetVisibility())
	virtual := ""
	if node.IsVirtual() {
		virtual = "virtual"
	}
	writeSeperatedStrings(sb, " ", visibility, virtual)

	if node.GetBody() == nil {
		sb.WriteString(";\n")
	} else {
		sb.WriteString(" {\n")
		success = PrintRecursive(node.GetBody(), sb, depth+1) && success
		sb.WriteString(indentString("}\n", depth))
	}
	return success
}
