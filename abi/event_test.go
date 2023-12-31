package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

func TestProcessEvent(t *testing.T) {
	builder := &Builder{}

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
			result := builder.processEvent(tc.input)
			assert.Equal(t, tc.expectedType, result.Type)
			assert.Equal(t, tc.expectedName, result.Name)
			assert.Equal(t, tc.expectedOutputs, len(result.Outputs))
			for _, output := range result.Outputs {
				assert.True(t, output.Indexed)
			}
		})
	}
}
