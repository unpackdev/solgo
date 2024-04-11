package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printConstructor(node *ast.Constructor, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("constructor(")
	success = printParameterList(node.GetParameters(), sb, depth) && success
	sb.WriteString(") \n")
	success = PrintRecursive(node.GetBody(), sb, depth) && success
	return success
}
