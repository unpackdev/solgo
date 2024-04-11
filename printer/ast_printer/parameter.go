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

func getVisibilityString(visibility ast_pb.Visibility) string {
	switch visibility {
	case ast_pb.Visibility_INTERNAL:
		return "internal"
	case ast_pb.Visibility_PUBLIC:
		return "public"
	case ast_pb.Visibility_EXTERNAL:
		return "external"
	case ast_pb.Visibility_PRIVATE:
		return "private"
	default:
		return ""
	}
}

func getStateMutabilityString(mut ast_pb.Mutability) string {
	switch mut {
	case ast_pb.Mutability_PURE:
		return "pure"
	case ast_pb.Mutability_VIEW:
		return "view"
	case ast_pb.Mutability_NONPAYABLE:
		return "nonpayable"
	case ast_pb.Mutability_PAYABLE:
		return "payable"
	default:
		return ""
	}
}

func printParameter(node *ast.Parameter, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString(node.GetName())
	typeName, ok := Print(node.GetTypeName())
	success = ok && success
	ident := node.GetName()
	storage := getStorageLocationString(node.GetStorageLocation())
	writeSeperatedStrings(sb, " ", typeName, storage, ident)
	return success
}
