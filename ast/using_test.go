package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsingDirectiveNode_Children(t *testing.T) {
	tt := []struct {
		name     string
		node     *UsingDirectiveNode
		expected []Node
	}{
		{
			name: "UsingDirectiveNode with no children",
			node: &UsingDirectiveNode{
				Alias:      "MyAlias",
				Type:       "MyType",
				IsWildcard: false,
				IsGlobal:   true,
				IsUserDef:  true,
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
