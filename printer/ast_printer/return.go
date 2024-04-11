package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printReturn(node *ast.ReturnStatement, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("return")
	if node.GetExpression() != nil {
		sb.WriteString(" ")
		success = PrintRecursive(node.GetExpression(), sb, depth) && success
	}
	return success
}
