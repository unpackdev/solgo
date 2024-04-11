package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printPrimaryExpression(node *ast.PrimaryExpression, sb *strings.Builder, depth int) bool {
	s := ""
	if node.GetValue() == "" {
		s = node.GetName()
	} else {
		s = node.GetValue()
	}
	sb.WriteString(s)
	return true
}
