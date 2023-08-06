package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestStateVariableMethods(t *testing.T) {
	// Create a new StateVariable instance
	stateVariable := &StateVariable{
		Id:              1,
		ContractId:      2,
		Name:            "TestVariable",
		NodeType:        ast_pb.NodeType(1),
		Visibility:      ast_pb.Visibility(1),
		Constant:        true,
		StorageLocation: ast_pb.StorageLocation(1),
		StateMutability: ast_pb.Mutability(1),
		Type:            "TestType",
		TypeDescription: &ast.TypeDescription{TypeString: "TestTypeDescription"},
	}

	// Test GetId method
	assert.Equal(t, int64(1), stateVariable.GetId())

	// Test GetContractId method
	assert.Equal(t, int64(2), stateVariable.GetContractId())

	// Test GetName method
	assert.Equal(t, "TestVariable", stateVariable.GetName())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), stateVariable.GetNodeType())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), stateVariable.GetVisibility())

	// Test IsConstant method
	assert.Equal(t, true, stateVariable.IsConstant())

	// Test GetStorageLocation method
	assert.Equal(t, ast_pb.StorageLocation(1), stateVariable.GetStorageLocation())

	// Test GetStateMutability method
	assert.Equal(t, ast_pb.Mutability(1), stateVariable.GetStateMutability())

	// Test GetType method
	assert.Equal(t, "TestType", stateVariable.GetType())

	// Test GetTypeDescription method
	assert.Equal(t, &ast.TypeDescription{TypeString: "TestTypeDescription"}, stateVariable.GetTypeDescription())
}
