package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printBody(node *ast.BodyNode, sb *strings.Builder, depth int) bool {
	success := true
	for _, stmt := range node.GetStatements() {
		sb.WriteString(indentString("", depth))
		success = PrintRecursive(stmt, sb, depth+1) && success
		writeStrings(sb, ";\n")
	}
	return success
}
