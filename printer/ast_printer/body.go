package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printBody(node *ast.BodyNode, sb *strings.Builder, depth int) bool {
	success := true
	stmts := []string{}
	for _, stmt := range node.GetStatements() {
		s, ok := Print(stmt)
		success = ok && success
		stmts = append(stmts, indentString(s, depth+1))
	}
	writeSeperatedList(sb, "\n;", stmts)
	return success
}
