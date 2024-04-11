package ast_printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func getStorageLocationString(storage ast_pb.StorageLocation) string {
	switch storage {
	case ast_pb.StorageLocation_DEFAULT:
		return ""
	case ast_pb.StorageLocation_MEMORY:
		return "memory"
	case ast_pb.StorageLocation_STORAGE:
		return "storage"
	case ast_pb.StorageLocation_CALLDATA:
		return "calldata"
	default:
		return ""
	}
}

func printParamter(node *ast.Parameter, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString(node.GetName())
	typeName, ok := Print(node.GetTypeName())
	success = ok && success
	ident := node.GetName()
	storage := getStorageLocationString(node.GetStorageLocation())
	writeSeperatedStrings(sb, " ", typeName, storage, ident)
	return success
}
