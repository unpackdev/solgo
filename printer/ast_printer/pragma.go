package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printPragma(node *ast.Pragma, sb *strings.Builder, depth int) bool {
	sb.WriteString(node.GetText())
	sb.WriteString("\n")
	return true
}
