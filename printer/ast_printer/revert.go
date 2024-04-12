package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printRevertStatement(node *ast.RevertStatement, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("revert")
	if node.GetExpression() != nil {
		sb.WriteString(" ")
		success = PrintRecursive(node.GetExpression(), sb, depth) && success
	}
	args := []string{}
	for _, arg := range node.GetArguments() {
		s, ok := Print(arg)
		success = success && ok
		args = append(args, s)
	}
	sb.WriteString("(")
	writeSeperatedList(sb, ", ", args)
	sb.WriteString(")")
	return success
}
