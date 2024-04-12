package ast_printer

import (
	"fmt"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
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
		if node.NodeType == ast_pb.NodeType_USER_DEFINED_PATH_NAME {
			userType := node.GetTree().GetById(node.GetReferencedDeclaration()).(*ast.EnumDefinition)
			sb.WriteString(userType.GetName())
		} else {
			sb.WriteString(node.GetName())
		}
	}
	return success
}
