package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

func TestProcessFallback(t *testing.T) {
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
				Name: "outputParam1",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
			{
				Name: "outputParam2",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
		},
	}

	builder := &Builder{}
	result := builder.processFallback(mockFallback)

	// Assert that the returned Method object has the expected properties
	assert.Equal(t, "fallback", result.Type)
	assert.Equal(t, builder.normalizeStateMutability(mockFallback.GetStateMutability()), result.StateMutability)
	assert.Equal(t, "", result.Name)
	assert.Equal(t, 0, len(result.Inputs))
	assert.Equal(t, 0, len(result.Outputs))
}
