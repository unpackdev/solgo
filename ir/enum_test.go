package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func TestEnumMethods(t *testing.T) {
	// Create a new Enum instance
	enumInstance := &Enum{
		unit:          &ast.EnumDefinition{},
		Id:            1,
		NodeType:      ast_pb.NodeType(1),
		Name:          "enumName",
		CanonicalName: "canonicalName",
		Members: []*Parameter{
			{
				unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "parameter",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
	}

	// Test GetAST method
	assert.IsType(t, &ast.EnumDefinition{}, enumInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), enumInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), enumInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "enumName", enumInstance.GetName())

	// Test GetCanonicalName method
	assert.Equal(t, "canonicalName", enumInstance.GetCanonicalName())

	// Test GetMembers method
	assert.IsType(t, []*Parameter{}, enumInstance.GetMembers())

	// Test GetSrc method
	assert.IsType(t, ast.SrcNode{}, enumInstance.GetSrc())
}
