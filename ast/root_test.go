package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootNode_Children(t *testing.T) {
	// Define mock contract nodes
	contractNode1 := &ContractNode{Name: "Contract1"}
	contractNode2 := &ContractNode{Name: "Contract2"}
	contractNode3 := &ContractNode{Name: "Contract3"}

	tt := []struct {
		name     string
		root     *RootNode
		expected []Node
	}{
		{
			name: "RootNode with multiple contracts and interfaces",
			root: &RootNode{
				Contracts: []*ContractNode{contractNode1, contractNode2},
			},
			expected: []Node{contractNode1, contractNode2},
		},
		{
			name: "RootNode with no contracts and interfaces",
			root: &RootNode{
				Contracts: []*ContractNode{},
			},
			expected: []Node{},
		},
		{
			name: "RootNode with contracts only",
			root: &RootNode{
				Contracts: []*ContractNode{contractNode1, contractNode2, contractNode3},
			},
			expected: []Node{contractNode1, contractNode2, contractNode3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.root.Children()
			assert.Equal(t, tc.expected, children)

			// Assert the type of each child
			for _, child := range children {
				_, isContract := child.(*ContractNode)
				assert.True(t, isContract, "unexpected type %T", child)
			}
		})
	}
}
