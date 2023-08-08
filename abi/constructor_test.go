package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

func TestProcessConstructor(t *testing.T) {
	mockConstructor := &ir.Constructor{
		StateMutability: ast_pb.Mutability_NONPAYABLE,
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
	result := builder.processConstructor(mockConstructor)

	// Assert that the returned Method object has the expected properties
	assert.Equal(t, "constructor", result.Type)
	assert.Equal(t, builder.normalizeStateMutability(mockConstructor.GetStateMutability()), result.StateMutability)
	assert.Equal(t, "", result.Name)
	assert.Equal(t, 2, len(result.Inputs))
	assert.Equal(t, "inputParam1", result.Inputs[0].Name)
	assert.Equal(t, "inputParam2", result.Inputs[1].Name)
	assert.Equal(t, "uint256", result.Inputs[0].Type)
	assert.Equal(t, "uint256", result.Inputs[1].Type)
	assert.Equal(t, "uint256", result.Inputs[0].InternalType)
	assert.Equal(t, "uint256", result.Inputs[1].InternalType)
	assert.Equal(t, 0, len(result.Outputs))
}
