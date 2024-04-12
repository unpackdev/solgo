package ast_printer

import (
	"fmt"
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printStateVariableDeclaration(node *ast.StateVariableDeclaration, sb *strings.Builder, depth int) bool {
	success := true
	fmt.Printf("%+v\n\n", node.GetTypeName())
	typeName, ok := Print(node.GetTypeName())
	success = ok && success
	ident := node.GetName()
	storage := getStorageLocationString(node.GetStorageLocation())
	visibility := getVisibilityString(node.GetVisibility())
	override := ""
	if typeName == "addresspayable" {
		typeName = "address"
		storage = "payable"
	}
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
