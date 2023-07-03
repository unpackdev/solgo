package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructorNode_Children(t *testing.T) {
	// Define mock variable and statement nodes
	variableNode1 := &VariableNode{}
	variableNode2 := &VariableNode{}
	statementNode1 := &StatementNode{}
	statementNode2 := &StatementNode{}

	tt := []struct {
		name        string
		constructor *ConstructorNode
		expected    []Node
	}{
		{
			name: "both non-empty",
			constructor: &ConstructorNode{
				Parameters: []*VariableNode{variableNode1, variableNode2},
				Body:       []*StatementNode{statementNode1, statementNode2},
			},
			expected: []Node{variableNode1, variableNode2, statementNode1, statementNode2},
		},
		{
			name: "both empty",
			constructor: &ConstructorNode{
				Parameters: []*VariableNode{},
				Body:       []*StatementNode{},
			},
			expected: []Node{},
		},
		{
			name: "parameters non-empty and body empty",
			constructor: &ConstructorNode{
				Parameters: []*VariableNode{variableNode1, variableNode2},
				Body:       []*StatementNode{},
			},
			expected: []Node{variableNode1, variableNode2},
		},
		{
			name: "parameters empty and body non-empty",
			constructor: &ConstructorNode{
				Parameters: []*VariableNode{},
				Body:       []*StatementNode{statementNode1, statementNode2},
			},
			expected: []Node{statementNode1, statementNode2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.constructor.Children()
			assert.Equal(t, tc.expected, children)

			for i, child := range children {
				if i < len(tc.constructor.Parameters) {
					assert.IsType(t, &VariableNode{}, child)
				} else {
					assert.IsType(t, &StatementNode{}, child)
				}
			}
		})
	}
}
