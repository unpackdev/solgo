package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

func TestProcessFunction(t *testing.T) {
	// Create a Builder object
	builder := &Builder{}

	testCases := []struct {
		name            string
		input           *ir.Function
		expectedType    string
		expectedName    string
		expectedInputs  int
		expectedOutputs int
	}{
		{
			name: "basic function case",
			input: &ir.Function{
				Name: "basicFunction",
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
							TypeString:     "address",
							TypeIdentifier: "t_address",
						},
					},
				},
				ReturnStatements: []*ir.Parameter{
					{
						Name: "return1",
						TypeDescription: &ast.TypeDescription{
							TypeString:     "bool",
							TypeIdentifier: "t_bool",
						},
					},
				},
				StateMutability: ast_pb.Mutability_NONPAYABLE,
			},
			expectedType:    "function",
			expectedName:    "basicFunction",
			expectedInputs:  2,
			expectedOutputs: 1,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.processFunction(tc.input)

			// Assertions
			assert.Equal(t, tc.expectedType, result.Type)
			assert.Equal(t, tc.expectedName, result.Name)
			assert.Equal(t, tc.expectedInputs, len(result.Inputs))
			assert.Equal(t, tc.expectedOutputs, len(result.Outputs))
		})
	}
}
