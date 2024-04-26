package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printContinueStatement(node *ast.ContinueStatement, sb *strings.Builder, depth int) bool {
	sb.WriteString("continue")
	return true
}
