package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

func TestProcessReceive(t *testing.T) {
	// Create a mock Receive object
	mockReceive := &ir.Receive{
		Name: "mockReceive",
		Parameters: []*ir.Parameter{
			{
				Name: "param1",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
			{
				Name: "param2",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "uint256",
					TypeIdentifier: "t_uint256",
				},
			},
		},
	}

	// Create a Builder object
	builder := &Builder{}

	// Call the processReceive function
	result := builder.processReceive(mockReceive)

	// Assert that the returned Method object has the expected properties
	assert.Equal(t, "receive", result.Type)
	assert.Equal(t, "payable", result.StateMutability)
	assert.Equal(t, "", result.Name)
	assert.Equal(t, 2, len(result.Inputs))
	assert.Equal(t, "param1", result.Inputs[0].Name)
	assert.Equal(t, "param2", result.Inputs[1].Name)
	assert.Equal(t, "uint256", result.Inputs[0].Type)
	assert.Equal(t, "uint256", result.Inputs[1].Type)
	assert.Equal(t, "uint256", result.Inputs[0].InternalType)
}
