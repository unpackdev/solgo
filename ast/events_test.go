package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventParameterNode_Children(t *testing.T) {
	eventParameterNode := &EventParameterNode{Name: "TestParameter", Type: "uint256", Indexed: true}
	children := eventParameterNode.Children()
	assert.Nil(t, children)
}

func TestEventNode_Children(t *testing.T) {
	eventParameterNode := &EventParameterNode{Name: "TestParameter", Type: "uint256", Indexed: true}

	tt := []struct {
		name     string
		event    *EventNode
		expected []Node
	}{
		{
			name: "EventNode with parameters",
			event: &EventNode{
				Name:       "TestEvent",
				Anonymous:  false,
				Parameters: []*EventParameterNode{eventParameterNode},
			},
			expected: []Node{eventParameterNode},
		},
		{
			name: "EventNode with no parameters",
			event: &EventNode{
				Name:       "TestEvent",
				Anonymous:  false,
				Parameters: []*EventParameterNode{},
			},
			expected: []Node{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.event.Children()
			assert.Equal(t, tc.expected, children)

			for _, child := range children {
				assert.IsType(t, &EventParameterNode{}, child)
			}
		})
	}
}
