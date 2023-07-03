package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableNode_Children(t *testing.T) {
	tt := []struct {
		name     string
		node     *VariableNode
		expected []Node
	}{
		{
			name: "VariableNode with no children",
			node: &VariableNode{
				Name: "myVariable",
				Type: "uint256",
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.node.Children()

			assert.Equal(t, tc.expected, children, "Children should match the expected values")
		})
	}
}

func TestStateVariableNode_Children(t *testing.T) {
	tt := []struct {
		name     string
		node     *StateVariableNode
		expected []Node
	}{
		{
			name: "StateVariableNode with no children",
			node: &StateVariableNode{
				Name:         "myStateVariable",
				Type:         "address",
				Visibility:   "public",
				IsConstant:   false,
				IsImmutable:  false,
				InitialValue: "",
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.node.Children()

			assert.Equal(t, tc.expected, children, "Children should match the expected values")
		})
	}
}
