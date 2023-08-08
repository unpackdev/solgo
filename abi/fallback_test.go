package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

func TestProcessFallback(t *testing.T) {
	// Create a mock Fallback object
	mockFallback := &ir.Fallback{
		StateMutability: ast_pb.Mutability_VIEW,
		Parameters: []*ir.Parameter{
			{
				Name: "inputParam1",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
			{
				Name: "inputParam2",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
		},
		ReturnStatements: []*ir.Parameter{
			{
				Name:            "outputParam1",
				TypeDescription: &ast.TypeDescription{
					// Add mock type description properties here
				},
			},
			{
				Name:            "outputParam2",
				TypeDescription: &ast.TypeDescription{
					// Add mock type description properties here
				},
			},
		},
	}

	// Create a Builder object
	builder := &Builder{}

	// Call the processFallback function
	result := builder.processFallback(mockFallback)

	// Assert that the returned Method object has the expected properties
	assert.Equal(t, "fallback", result.Type)
	assert.Equal(t, builder.normalizeStateMutability(mockFallback.GetStateMutability()), result.StateMutability)
	assert.Equal(t, "", result.Name)
	assert.Equal(t, 2, len(result.Inputs))
	assert.Equal(t, "inputParam1", result.Inputs[0].Name)
	assert.Equal(t, "inputParam2", result.Inputs[1].Name)
	assert.Equal(t, 2, len(result.Outputs))
	assert.Equal(t, "outputParam1", result.Outputs[0].Name)
	assert.Equal(t, "outputParam2", result.Outputs[1].Name)
	assert.Equal(t, "uint256", result.Inputs[0].Type)
	assert.Equal(t, "uint256", result.Inputs[1].Type)
	assert.Equal(t, "uint256", result.Inputs[0].InternalType)
	assert.Equal(t, "uint256", result.Inputs[1].InternalType)
}
