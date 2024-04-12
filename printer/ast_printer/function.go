package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printFunction(node *ast.Function, sb *strings.Builder, depth int) bool {
	success := true
	visibility := getVisibilityString(node.GetVisibility())
	funcName := node.GetName()
	mutability := getStateMutabilityString(node.GetStateMutability())
	virtual := ""
	if node.IsVirtual() {
		virtual = "virtual "
	}

	writeStrings(sb, "function ", funcName, "(")
	printParameterList(node.GetParameters(), sb, depth)
	writeSeperatedStrings(sb, " ", ")", visibility, virtual, mutability)

	paramBuilder := strings.Builder{}
	printParameterList(node.GetReturnParameters(), &paramBuilder, depth)

	if paramBuilder.String() != "" {
		writeStrings(sb, " returns (", paramBuilder.String(), ")")
	}
	sb.WriteString(" ")

	if node.GetBody() != nil {
		success = PrintRecursive(node.GetBody(), sb, depth) && success
	}

	return success
}
