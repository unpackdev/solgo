package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
	"go.uber.org/zap"
)

func printInlineArray(node *ast.InlineArray, sb *strings.Builder, depth int) bool {
	success := true
	if len(node.GetExpressions()) < 3 {
		zap.S().Error("Conditional node must have at least 3 expressions")
		return false
	}
	sb.WriteString("[")
	items := []string{}
	for _, item := range node.GetExpressions() {
		s, ok := Print(item)
		if !ok {
			success = false
		}
		items = append(items, s)
	}
	writeSeperatedList(sb, ", ", items)
	sb.WriteString("]")
	return success
}
