package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printPayableConversion(node *ast.PayableConversion, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("payable(")
	args := []string{}
	for _, arg := range node.GetArguments() {
		s, ok := Print(arg)
		success = ok && success
		args = append(args, s)
	}
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(")")
	return success
}
