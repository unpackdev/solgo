package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printFor(node *ast.ForStatement, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("for (")
	if node.GetInitialiser() != nil {
		success = PrintRecursive(node.GetInitialiser(), sb, depth) && success
	}
	sb.WriteString("; ")
	if node.GetCondition() != nil {
		success = PrintRecursive(node.GetCondition(), sb, depth) && success
	}
	sb.WriteString("; ")
	if node.GetClosure() != nil {
		success = PrintRecursive(node.GetClosure(), sb, depth) && success
	}
	sb.WriteString(") {\n")
	if node.GetBody() != nil {
		success = PrintRecursive(node.GetBody(), sb, depth) && success
	}
	sb.WriteString("}")
	return success
}
