package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printContract(node *ast.Contract, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("contract ")
	sb.WriteString(node.GetName())
	baseContracts := []string{}
	for _, base := range node.GetBaseContracts() {
		baseContracts = append(baseContracts, base.BaseName.GetName())
	}
	if len(baseContracts) > 0 {
		sb.WriteString(" is ")
		writeSeperatedStrings(sb, ", ", baseContracts...)
	}

	sb.WriteString(" {\n")
	for _, child := range node.GetNodes() {
		success = PrintRecursive(child, sb, depth+1) && success
	}
	sb.WriteString("}\n")
	return success
}
