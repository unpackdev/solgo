package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printFunctionCall(node *ast.FunctionCall, sb *strings.Builder, depth int) bool {
	success := true
	if node.GetExpression() != nil {
		success = PrintRecursive(node.GetExpression(), sb, depth) && success
	}
	args := []string{}
	if node.GetArguments() != nil {
		for _, arg := range node.GetArguments() {
			s, ok := Print(arg)
			success = ok && success
			args = append(args, s)
		}
	}
	sb.WriteString("(")
	writeSeperatedList(sb, ", ", args)
	sb.WriteString(")")
	return success
}
