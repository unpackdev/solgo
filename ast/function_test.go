package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionNode_Children(t *testing.T) {
	// Define mock nodes
	paramNode1 := &VariableNode{Name: "param1", Type: "uint256"}
	paramNode2 := &VariableNode{Name: "param2", Type: "address"}

	retParamNode1 := &VariableNode{Name: "retParam1", Type: "bool"}
	retParamNode2 := &VariableNode{Name: "retParam2", Type: "string"}

	statementNode1 := &StatementNode{Raw: []string{"return", "true", ";"}, TextRaw: "return true;"}
	statementNode2 := &StatementNode{Raw: []string{"emit", "Event()", ";"}, TextRaw: "emit Event();"}

	tt := []struct {
		name     string
		function *FunctionNode
		expected []Node
	}{
		{
			name: "FunctionNode with multiple parameters and body",
			function: &FunctionNode{
				Name:             "TestFunction1",
				Parameters:       []*VariableNode{paramNode1, paramNode2},
				ReturnParameters: []*VariableNode{retParamNode1, retParamNode2},
				Body:             []*StatementNode{statementNode1, statementNode2},
			},
			expected: []Node{paramNode1, paramNode2, retParamNode1, retParamNode2, statementNode1, statementNode2},
		},
		{
			name: "FunctionNode with only parameters",
			function: &FunctionNode{
				Name:             "TestFunction2",
				Parameters:       []*VariableNode{paramNode1, paramNode2},
				ReturnParameters: []*VariableNode{},
				Body:             []*StatementNode{},
			},
			expected: []Node{paramNode1, paramNode2},
		},
		{
			name: "FunctionNode with only return parameters",
			function: &FunctionNode{
				Name:             "TestFunction3",
				Parameters:       []*VariableNode{},
				ReturnParameters: []*VariableNode{retParamNode1, retParamNode2},
				Body:             []*StatementNode{},
			},
			expected: []Node{retParamNode1, retParamNode2},
		},
		{
			name: "FunctionNode with only body",
			function: &FunctionNode{
				Name:             "TestFunction4",
				Parameters:       []*VariableNode{},
				ReturnParameters: []*VariableNode{},
				Body:             []*StatementNode{statementNode1, statementNode2},
			},
			expected: []Node{statementNode1, statementNode2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.function.Children()
			assert.Equal(t, tc.expected, children)

			// Assert the type of each child
			for _, child := range children {
				switch child := child.(type) {
				case *VariableNode:
					if contains(tc.function.Parameters, child) {
						assert.Contains(t, tc.function.Parameters, child)
					} else if contains(tc.function.ReturnParameters, child) {
						assert.Contains(t, tc.function.ReturnParameters, child)
					} else {
						t.Errorf("unexpected variable node: %v", child)
					}
				case *StatementNode:
					assert.Contains(t, tc.function.Body, child)
				default:
					t.Errorf("unexpected type %T", child)
				}
			}
		})
	}
}

// contains is a helper function that checks if a slice of *VariableNode contains a *VariableNode
func contains(nodes []*VariableNode, node *VariableNode) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}
