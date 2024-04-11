package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printVariableDeclaration(node *ast.VariableDeclaration, sb *strings.Builder, depth int) bool {
	success := true
	isTuple := len(node.GetDeclarations()) > 1
	if isTuple {
		decls := []string{}
		for _, decl := range node.GetDeclarations() {
			s, ok := Print(decl)
			success = ok && success
			decls = append(decls, s)
		}
		writeSeperatedList(sb, ", ", decls)
	} else {
		PrintRecursive(node.GetDeclarations()[0], sb, depth)
	}
	sb.WriteString(" = ")
	PrintRecursive(node.GetInitialValue(), sb, depth)
	return success
}
