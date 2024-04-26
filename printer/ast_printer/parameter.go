package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printParameter(node *ast.Parameter, sb *strings.Builder, depth int) bool {
	success := true
	typeName := ""
	ok := true
	if node.GetTypeName() != nil {
		typeName, ok = Print(node.GetTypeName())
	}
	success = ok && success
	ident := node.GetName()
	storage := getStorageLocationString(node.GetStorageLocation())
	writeSeperatedStrings(sb, " ", typeName, storage, ident)
	return success
}
