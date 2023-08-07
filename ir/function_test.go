package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestFunctionMethods(t *testing.T) {
	// Create a new Function instance
	functionInstance := &Function{
		unit:                    &ast.Function{},
		Id:                      1,
		NodeType:                ast_pb.NodeType(1),
		Name:                    "functionName",
		Kind:                    ast_pb.NodeType(1),
		Implemented:             true,
		Visibility:              ast_pb.Visibility(1),
		StateMutability:         ast_pb.Mutability(1),
		Virtual:                 true,
		ReferencedDeclarationId: 1,
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
	assert.IsType(t, &ast.Function{}, functionInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), functionInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), functionInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "functionName", functionInstance.GetName())

	// Test GetKind method
	assert.Equal(t, ast_pb.NodeType(1), functionInstance.GetKind())

	// Test IsImplemented method
	assert.Equal(t, true, functionInstance.IsImplemented())

	// Test GetVisibility method
	assert.Equal(t, ast_pb.Visibility(1), functionInstance.GetVisibility())

	// Test GetStateMutability method
	assert.Equal(t, ast_pb.Mutability(1), functionInstance.GetStateMutability())

	// Test IsVirtual method
	assert.Equal(t, true, functionInstance.IsVirtual())

	// Test GetModifiers method
	assert.IsType(t, []*Modifier{}, functionInstance.GetModifiers())

	// Test GetParameters method
	assert.IsType(t, []*Parameter{}, functionInstance.GetParameters())

	// Test GetReturnStatements method
	assert.IsType(t, []*Parameter{}, functionInstance.GetReturnStatements())
}
