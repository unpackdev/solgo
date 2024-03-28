package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

func TestProcessStateVariable(t *testing.T) {
	builder := &Builder{
		resolver: &TypeResolver{
			parser:         nil,
			processedTypes: make(map[string]bool),
		},
	}

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
			expectedOutput: "uint256",
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
			assert.Equal(t, tc.expectedType, result.Type)
			assert.Equal(t, tc.input.Name, result.Name)
			assert.Equal(t, tc.expectedOutput, result.Outputs[0].Type)
		})
	}
}
