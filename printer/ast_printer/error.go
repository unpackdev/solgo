package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printErrorDefinition(node *ast.ErrorDefinition, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("error ")
	sb.WriteString(node.GetName())
	sb.WriteString("(")
	success = printParameterList(node.GetParameters(), sb, depth) && success
	sb.WriteString(")")
	sb.WriteString(";\n")
	return success
}
