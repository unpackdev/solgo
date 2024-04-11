package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printMemberAccessExpression(node *ast.MemberAccessExpression, sb *strings.Builder, depth int) bool {
	success := true
	success = PrintRecursive(node.GetExpression(), sb, depth) && success
	sb.WriteString(".")
	sb.WriteString(node.GetMemberName())
	return success
}
