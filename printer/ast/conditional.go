package printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
	"go.uber.org/zap"
)

func printConditional(node ast.Conditional, sb *strings.Builder, depth int) bool {
	success := true
	if len(node.GetExpressions()) < 3 {
		zap.S().Error("Conditional node must have at least 3 expressions")
		return false
	}
	success = PrintRecursive(node.GetExpressions()[0], sb, depth) && success
	sb.WriteString(" ? ")
	success = PrintRecursive(node.GetExpressions()[1], sb, depth) && success
	sb.WriteString(" : ")
	success = PrintRecursive(node.GetExpressions()[2], sb, depth) && success
	return success
}
