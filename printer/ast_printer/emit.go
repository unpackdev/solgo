package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printEmit(node *ast.Emit, sb *strings.Builder, depth int) bool {
	success := true
	args := []string{}
	for _, arg := range node.GetArguments() {
		s, ok := Print(arg)
		success = ok && success
		args = append(args, s)
	}
	sb.WriteString("emit ")
	success = PrintRecursive(node.GetExpression(), sb, depth) && success
	sb.WriteString("(")
	writeSeperatedList(sb, ", ", args)
	sb.WriteString(")")
	return success
}
