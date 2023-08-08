package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

func TestProcessStateVariable(t *testing.T) {
	// Create a Builder object
	builder := &Builder{}

	testCases := []struct {
		name           string
		input          *ir.StateVariable
		expectedType   string
		expectedOutput string
	}{
		{
			name: "mapping case",
			input: &ir.StateVariable{
				Name: "mockMappingVar",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "mapping(uint256 => uint256)",
					TypeIdentifier: "t_mapping",
				},
				StateMutability: ast_pb.Mutability_VIEW,
			},
			expectedType:   "function",
			expectedOutput: "mapping", // Adjust as per actual expected output
		},
		{
			name: "contract case",
			input: &ir.StateVariable{
				Name: "mockContractVar",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "contract Test.Contract",
					TypeIdentifier: "t_contract",
				},
				StateMutability: ast_pb.Mutability_VIEW,
			},
			expectedType:   "function",
			expectedOutput: "address",
		},
		{
			name: "enum case",
			input: &ir.StateVariable{
				Name: "mockEnumVar",
				TypeDescription: &ast.TypeDescription{
					TypeString:     "enum Test.Enum",
					TypeIdentifier: "t_enum",
				},
				StateMutability: ast_pb.Mutability_VIEW,
			},
			expectedType:   "function",
			expectedOutput: "uint8",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.processStateVariable(tc.input)

			// Common assertions for all cases
			assert.Equal(t, tc.expectedType, result.Type)
			assert.Equal(t, tc.input.Name, result.Name)

			// Specific assertions based on the test case
			//assert.Equal(t, tc.expectedOutput, result.Outputs[0].Type)
		})
	}
}
