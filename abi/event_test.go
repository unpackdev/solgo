package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

func TestProcessEvent(t *testing.T) {
	builder := &Builder{
		resolver: &TypeResolver{
			parser:         nil,
			processedTypes: make(map[string]bool),
		},
	}

	testCases := []struct {
		name            string
		input           *ir.Event
		expectedType    string
		expectedName    string
		expectedOutputs int
	}{
		{
			name: "basic event case",
			input: &ir.Event{
				Name: "basicEvent",
				Parameters: []*ir.Parameter{
					{
						Name: "param1",
						TypeDescription: &ast.TypeDescription{
							TypeString:     "uint256",
							TypeIdentifier: "t_uint256",
						},
						Indexed: true,
					},
				},
			},
			expectedType:    "event",
			expectedName:    "basicEvent",
			expectedOutputs: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := builder.processEvent(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedType, result.Type)
			assert.Equal(t, tc.expectedName, result.Name)
			assert.Equal(t, tc.expectedOutputs, len(result.Inputs))
			for _, output := range result.Outputs {
				assert.True(t, output.Indexed)
			}
		})
	}
}
