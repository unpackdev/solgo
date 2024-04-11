package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

// parameterList does not have the correct Node interface, we handle it separately
func printParameterList(node *ast.ParameterList, sb *strings.Builder, depth int) bool {
	success := true
	params := []string{}
	for _, param := range node.GetParameters() {
		s, ok := Print(param)
		success = ok && success
		params = append(params, s)
	}
	writeSeperatedList(sb, ", ", params)
	return success
}
