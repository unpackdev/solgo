package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

func TestProcessError(t *testing.T) {
	tests := []struct {
		name     string
		input    *ir.Error
		expected *Method
	}{
		{
			"SingleParameterError",
			&ir.Error{
				Name: "ErrorA",
				Parameters: []*ir.Parameter{
					{
						Name: "param1",
						TypeDescription: &ast.TypeDescription{
							TypeString: "uint256",
						},
					},
				},
			},
			&Method{
				Name:            "ErrorA",
				Type:            "error",
				StateMutability: "view",
				Inputs: []MethodIO{
					{
						Name:         "param1",
						Type:         "uint256",
						InternalType: "uint256",
						Indexed:      true,
					},
				},
				Outputs: []MethodIO{},
			},
		},
		{
			"MultipleParametersError",
			&ir.Error{
				Name: "ErrorB",
				Parameters: []*ir.Parameter{
					{
						Name: "param1",
						TypeDescription: &ast.TypeDescription{
							TypeString: "uint256",
						},
					},
					{
						Name: "param2",
						TypeDescription: &ast.TypeDescription{
							TypeString: "address",
						},
					},
				},
			},
			&Method{
				Name:            "ErrorB",
				Type:            "error",
				StateMutability: "view",
				Inputs: []MethodIO{
					{
						Name:         "param1",
						Type:         "uint256",
						InternalType: "uint256",
						Indexed:      true,
					},
					{
						Name:         "param2",
						Type:         "address",
						InternalType: "address",
						Indexed:      true,
					},
				},
				Outputs: []MethodIO{},
			},
		},
	}

	builder := &Builder{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.processError(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
