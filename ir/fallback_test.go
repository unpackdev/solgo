package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func TestFallbackMethods(t *testing.T) {
	// Create a new Fallback instance
	fallbackInstance := &Fallback{
		Unit:            &ast.Fallback{},
		Id:              1,
		NodeType:        ast_pb.NodeType(1),
		Name:            "fallback",
		Kind:            ast_pb.NodeType(1),
		Implemented:     true,
		Visibility:      ast_pb.Visibility(1),
		StateMutability: ast_pb.Mutability(1),
		Virtual:         true,
		Modifiers: []*Modifier{
			{
				Unit:          &ast.ModifierInvocation{},
				Id:            1,
				NodeType:      ast_pb.NodeType(1),
				Name:          "modifier",
				ArgumentTypes: []*ast.TypeDescription{},
			},
		},
		Parameters: []*Parameter{
			{
				Unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "parameter",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
		ReturnStatements: []*Parameter{
			{
				Unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "return",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
	}

	// Test GetAST method
	assert.IsType(t, &ast.Fallback{}, fallbackInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), fallbackInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), fallbackInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "fallback", fallbackInstance.GetName())

	// Test GetKind method
	assert.Equal(t, ast_pb.NodeType(1), fallbackInstance.GetKind())

	// Test IsImplemented method
	assert.Equal(t, true, fallbackInstance.IsImplemented())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), fallbackInstance.GetVisibility())

	// Test GetStateMutability method
	assert.Equal(t, ast_pb.Mutability(1), fallbackInstance.GetStateMutability())

	// Test IsVirtual method
	assert.Equal(t, true, fallbackInstance.IsVirtual())

	// Test GetModifiers method
	assert.IsType(t, []*Modifier{}, fallbackInstance.GetModifiers())

	// Test GetParameters method
	assert.IsType(t, []*Parameter{}, fallbackInstance.GetParameters())
}
