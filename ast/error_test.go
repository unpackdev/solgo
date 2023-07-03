package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNode_Children(t *testing.T) {
	errorValueNode := &ErrorValueNode{Name: "TestError", Type: "uint256", Code: 1}

	tt := []struct {
		name     string
		err      *ErrorNode
		expected []Node
	}{
		{
			name: "ErrorNode with values",
			err: &ErrorNode{
				Name:   "TestError",
				Values: []*ErrorValueNode{errorValueNode},
			},
			expected: []Node{errorValueNode},
		},
		{
			name: "ErrorNode with no values",
			err: &ErrorNode{
				Name:   "TestError",
				Values: []*ErrorValueNode{},
			},
			expected: []Node{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.err.Children()
			assert.Equal(t, tc.expected, children)

			for _, child := range children {
				assert.IsType(t, &ErrorValueNode{}, child)
			}
		})
	}
}

func TestErrorValueNode_Children(t *testing.T) {
	errorValueNode := &ErrorValueNode{Name: "TestError", Type: "uint256", Code: 1}
	children := errorValueNode.Children()
	assert.Nil(t, children)
}
