package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printUnaryPrefix(node *ast.UnaryPrefix, sb *strings.Builder, depth int) bool {
	sb.WriteString(getUnaryOperatorString(node.GetOperator()))
	success := PrintRecursive(node.GetExpression(), sb, depth)
	return success
}
