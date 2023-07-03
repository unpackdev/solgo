package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContractNode_Children(t *testing.T) {
	stateVariableNode := &StateVariableNode{}
	structNode := &StructNode{}
	eventNode := &EventNode{}
	errorNode := &ErrorNode{}
	constructorNode := &ConstructorNode{}
	functionNode := &FunctionNode{}
	usingDirectiveNode := &UsingDirectiveNode{}

	tt := []struct {
		name     string
		contract *ContractNode
		expected []Node
	}{
		{
			name: "all nodes present",
			contract: &ContractNode{
				StateVariables: []*StateVariableNode{stateVariableNode},
				Structs:        []*StructNode{structNode},
				Events:         []*EventNode{eventNode},
				Errors:         []*ErrorNode{errorNode},
				Constructor:    constructorNode,
				Functions:      []*FunctionNode{functionNode},
				Using:          []*UsingDirectiveNode{usingDirectiveNode},
			},
			expected: []Node{
				functionNode,
				stateVariableNode,
				eventNode,
				errorNode,
				structNode,
				usingDirectiveNode,
				constructorNode,
			},
		},
		{
			name: "no nodes present",
			contract: &ContractNode{
				StateVariables: []*StateVariableNode{},
				Structs:        []*StructNode{},
				Events:         []*EventNode{},
				Errors:         []*ErrorNode{},
				Constructor:    nil,
				Functions:      []*FunctionNode{},
				Using:          []*UsingDirectiveNode{},
			},
			expected: []Node{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.contract.Children()
			assert.Equal(t, tc.expected, children)

			for i, child := range children {
				if i < len(tc.contract.Functions) {
					assert.IsType(t, &FunctionNode{}, child)
				} else if i < len(tc.contract.Functions)+len(tc.contract.StateVariables) {
					assert.IsType(t, &StateVariableNode{}, child)
				} else if i < len(tc.contract.Functions)+len(tc.contract.StateVariables)+len(tc.contract.Events) {
					assert.IsType(t, &EventNode{}, child)
				} else if i < len(tc.contract.Functions)+len(tc.contract.StateVariables)+len(tc.contract.Events)+len(tc.contract.Errors) {
					assert.IsType(t, &ErrorNode{}, child)
				} else if i < len(tc.contract.Functions)+len(tc.contract.StateVariables)+len(tc.contract.Events)+len(tc.contract.Errors)+len(tc.contract.Structs) {
					assert.IsType(t, &StructNode{}, child)
				} else if i < len(tc.contract.Functions)+len(tc.contract.StateVariables)+len(tc.contract.Events)+len(tc.contract.Errors)+len(tc.contract.Structs)+len(tc.contract.Using) {
					assert.IsType(t, &UsingDirectiveNode{}, child)
				} else {
					assert.IsType(t, &ConstructorNode{}, child)
				}
			}
			assert.Equal(t, len(tc.expected), len(children))
		})
	}
}
