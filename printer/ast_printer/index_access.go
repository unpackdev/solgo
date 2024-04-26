package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printIndexAccess(node *ast.IndexAccess, sb *strings.Builder, depth int) bool {
	success := true
	success = PrintRecursive(node.GetBaseExpression(), sb, depth) && success
	sb.WriteString("[")
	success = PrintRecursive(node.GetIndexExpression(), sb, depth) && success
	sb.WriteString("]")
	return success
}
