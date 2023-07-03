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

	// Define mock interface nodes
	interfaceNode1 := &InterfaceNode{Name: "Interface1"}
	interfaceNode2 := &InterfaceNode{Name: "Interface2"}

	tt := []struct {
		name     string
		root     *RootNode
		expected []Node
	}{
		{
			name: "RootNode with multiple contracts and interfaces",
			root: &RootNode{
				Contracts:  []*ContractNode{contractNode1, contractNode2},
				Interfaces: []*InterfaceNode{interfaceNode1, interfaceNode2},
			},
			expected: []Node{contractNode1, contractNode2, interfaceNode1, interfaceNode2},
		},
		{
			name: "RootNode with no contracts and interfaces",
			root: &RootNode{
				Contracts:  []*ContractNode{},
				Interfaces: []*InterfaceNode{},
			},
			expected: []Node{},
		},
		{
			name: "RootNode with contracts only",
			root: &RootNode{
				Contracts:  []*ContractNode{contractNode1, contractNode2, contractNode3},
				Interfaces: []*InterfaceNode{},
			},
			expected: []Node{contractNode1, contractNode2, contractNode3},
		},
		{
			name: "RootNode with interfaces only",
			root: &RootNode{
				Contracts:  []*ContractNode{},
				Interfaces: []*InterfaceNode{interfaceNode1, interfaceNode2},
			},
			expected: []Node{interfaceNode1, interfaceNode2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.root.Children()
			assert.Equal(t, tc.expected, children)

			// Assert the type of each child
			for _, child := range children {
				_, isContract := child.(*ContractNode)
				_, isInterface := child.(*InterfaceNode)
				assert.True(t, isContract || isInterface, "unexpected type %T", child)
			}
		})
	}
}
