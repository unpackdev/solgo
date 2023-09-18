package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func TestConstructorMethods(t *testing.T) {
	// Create a new Constructor instance
	constructorInstance := &Constructor{
		unit: &ast.Constructor{},

		Id:              1,
		NodeType:        ast_pb.NodeType(1),
		Kind:            ast_pb.NodeType(1),
		Name:            "constructor",
		Implemented:     true,
		Visibility:      ast_pb.Visibility(1),
		StateMutability: ast_pb.Mutability(1),
		Virtual:         true,
		Modifiers: []*Modifier{
			{
				unit:          &ast.ModifierInvocation{},
				Id:            1,
				NodeType:      ast_pb.NodeType(1),
				Name:          "modifier",
				ArgumentTypes: []*ast.TypeDescription{},
			},
		},
		Parameters: []*Parameter{
			{
				unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "parameter",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
		ReturnStatements: []*Parameter{
			{
				unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "return",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
	}

	// Test GetAST method
	assert.IsType(t, &ast.Constructor{}, constructorInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), constructorInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), constructorInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "constructor", constructorInstance.GetName())

	// Test GetKind method
	assert.Equal(t, ast_pb.NodeType(1), constructorInstance.GetKind())

	// Test IsImplemented method
	assert.Equal(t, true, constructorInstance.IsImplemented())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), constructorInstance.GetVisibility())

	// Test GetStateMutability method
	assert.Equal(t, ast_pb.Mutability(1), constructorInstance.GetStateMutability())

	// Test IsVirtual method
	assert.Equal(t, true, constructorInstance.IsVirtual())

	// Test GetModifiers method
	assert.IsType(t, []*Modifier{}, constructorInstance.GetModifiers())

	// Test GetParameters method
	assert.IsType(t, []*Parameter{}, constructorInstance.GetParameters())

	// Test GetReturnStatements method
	assert.IsType(t, []*Parameter{}, constructorInstance.GetReturnStatements())
}
