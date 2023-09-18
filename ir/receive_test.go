package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

func TestReceiveMethods(t *testing.T) {
	// Create a new Receive instance
	receiveInstance := &Receive{
		Id:              1,
		NodeType:        ast_pb.NodeType(1),
		Name:            "TestReceive",
		Kind:            ast_pb.NodeType(1),
		Implemented:     true,
		Visibility:      ast_pb.Visibility(1),
		StateMutability: ast_pb.Mutability(1),
		Virtual:         true,
		Modifiers:       []*Modifier{{Id: 1, Name: "TestModifier"}},
		Parameters:      []*Parameter{{Id: 1, Name: "TestParameter"}},
	}

	// Test GetId method
	assert.Equal(t, int64(1), receiveInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), receiveInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "TestReceive", receiveInstance.GetName())

	// Test GetKind method
	assert.Equal(t, ast_pb.NodeType(1), receiveInstance.GetKind())

	// Test IsImplemented method
	assert.Equal(t, true, receiveInstance.IsImplemented())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), receiveInstance.GetVisibility())

	// Test GetStateMutability method
	assert.Equal(t, ast_pb.Mutability(1), receiveInstance.GetStateMutability())

	// Test IsVirtual method
	assert.Equal(t, true, receiveInstance.IsVirtual())

	// Test GetModifiers method
	assert.Equal(t, []*Modifier{{Id: 1, Name: "TestModifier"}}, receiveInstance.GetModifiers())

	// Test GetParameters method
	assert.Equal(t, []*Parameter{{Id: 1, Name: "TestParameter"}}, receiveInstance.GetParameters())
}
