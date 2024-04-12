package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printIfStatement(node *ast.IfStatement, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("if (")
	if node.GetCondition() != nil {
		success = PrintRecursive(node.GetCondition(), sb, depth) && success
	}
	sb.WriteString(") {\n")
	if node.GetBody() != nil {
		success = PrintRecursive(node.GetBody(), sb, depth+1) && success
	}
	// if node.GetElse() != nil {
	// 	sb.WriteString(indentString("} else ", depth-1))
	// 	success = PrintRecursive(node.GetElse(), sb, depth) && success
	// }
	sb.WriteString(indentString("}", depth-1))
	return success
}
