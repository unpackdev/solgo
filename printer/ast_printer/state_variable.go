package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printStateVariableDeclaration(node *ast.StateVariableDeclaration, sb *strings.Builder, depth int) bool {
	success := true
	typeName, ok := Print(node.GetTypeName())
	success = ok && success
	ident := node.GetName()
	storage := getStorageLocationString(node.GetStorageLocation())
	visibility := getVisibilityString(node.GetVisibility())
	override := ""
	if node.Override {
		override = "override"
	}
	writeSeperatedStrings(sb, " ", visibility, storage, typeName, override, ident)
	if node.GetInitialValue() != nil {
		sb.WriteString(" = ")
		success = PrintRecursive(node.GetInitialValue(), sb, depth) && success
	}
	sb.WriteString(";\n")
	return success
}
