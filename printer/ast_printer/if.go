package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printIfStatement(node *ast.IfStatement, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("if (")
	if node.GetCondition() != nil {
		success = PrintRecursive(node.GetCondition(), sb, depth) && success
	}
	sb.WriteString(") ")
	if node.GetBody() != nil {
		success = PrintRecursive(node.GetBody(), sb, depth) && success
	}
	return success
}
