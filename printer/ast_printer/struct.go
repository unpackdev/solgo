package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printStructDefinition(node *ast.StructDefinition, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("struct ")
	sb.WriteString(node.GetName())
	sb.WriteString(" {\n")
	for _, member := range node.GetMembers() {
		sb.WriteString(indentString("", depth+1))
		success = PrintRecursive(member, sb, depth) && success
		sb.WriteString(";\n")
	}
	sb.WriteString(indentString("}\n", depth))
	return success
}
