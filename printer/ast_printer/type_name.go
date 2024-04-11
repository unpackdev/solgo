package ast_printer

import (
	"fmt"
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printTypeName(node *ast.TypeName, sb *strings.Builder, depth int) bool {
	success := true
	if node.ValueType != nil {
		keyType, ok := Print(node.KeyType)
		if !ok {
			success = false
		}
		valueType, ok := Print(node.ValueType)
		if !ok {
			success = false
		}
		typeName := fmt.Sprintf("mapping(%s => %s)", keyType, valueType)
		sb.WriteString(typeName)
	} else {
		sb.WriteString(node.GetName())
	}
	return success
}
