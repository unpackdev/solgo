package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func TestStructMethods(t *testing.T) {
	// Create a new Struct instance
	structInstance := &Struct{
		Id:                      1,
		NodeType:                ast_pb.NodeType(1),
		Kind:                    ast_pb.NodeType(1),
		Name:                    "TestStruct",
		CanonicalName:           "TestCanonicalName",
		ReferencedDeclarationId: 2,
		Visibility:              ast_pb.Visibility(1),
		StorageLocation:         ast_pb.StorageLocation(1),
		Members:                 []*Parameter{},
		Type:                    "TestType",
		TypeDescription:         &ast.TypeDescription{TypeString: "TestTypeDescription"},
	}

	// Test GetId method
	assert.Equal(t, int64(1), structInstance.GetId())

	// Test GetName method
	assert.Equal(t, "TestStruct", structInstance.GetName())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), structInstance.GetNodeType())

	// Test GetKind method
	assert.Equal(t, ast_pb.NodeType(1), structInstance.GetKind())

	// Test GetCanonicalName method
	assert.Equal(t, "TestCanonicalName", structInstance.GetCanonicalName())

	// Test GetReferencedDeclarationId method
	assert.Equal(t, int64(2), structInstance.GetReferencedDeclarationId())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), structInstance.GetVisibility())

	// Test GetStorageLocation method
	assert.Equal(t, ast_pb.StorageLocation(1), structInstance.GetStorageLocation())

	// Test GetMembers method
	assert.Equal(t, []*Parameter{}, structInstance.GetMembers())

	// Test GetType method
	assert.Equal(t, "TestType", structInstance.GetType())

	// Test GetTypeDescription method
	assert.Equal(t, &ast.TypeDescription{TypeString: "TestTypeDescription"}, structInstance.GetTypeDescription())
}
