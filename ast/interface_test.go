package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceNode_Children(t *testing.T) {
	// Define mock function nodes
	functionNode1 := &FunctionNode{Name: "Function1"}
	functionNode2 := &FunctionNode{Name: "Function2"}
	functionNode3 := &FunctionNode{Name: "Function3"}

	tt := []struct {
		name     string
		iface    *InterfaceNode
		expected []Node
	}{
		{
			name: "InterfaceNode with multiple functions",
			iface: &InterfaceNode{
				Name:      "TestInterface1",
				Functions: []*FunctionNode{functionNode1, functionNode2, functionNode3},
			},
			expected: []Node{functionNode1, functionNode2, functionNode3},
		},
		{
			name: "InterfaceNode with no functions",
			iface: &InterfaceNode{
				Name:      "TestInterface2",
				Functions: []*FunctionNode{},
			},
			expected: []Node{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.iface.Children()
			assert.Equal(t, tc.expected, children)

			// Assert the type of each child
			for _, child := range children {
				_, ok := child.(*FunctionNode)
				assert.True(t, ok, "unexpected type %T", child)
			}
		})
	}
}
