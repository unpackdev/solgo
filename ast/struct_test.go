package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructMemberNode_Children(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		node     *StructMemberNode
		expected []Node
	}{
		{
			name:     "StructMemberNode with no children",
			node:     &StructMemberNode{Name: "myMember", Type: "uint256"},
			expected: nil,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.node.Children()
			assert.Equal(t, tc.expected, children)
		})
	}
}

func TestStructNode_AddMember(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name            string
		membersToAdd    []*StructMemberNode
		expectedMembers []*StructMemberNode
	}{
		{
			name:            "Add single member",
			membersToAdd:    []*StructMemberNode{{Name: "myMember", Type: "uint256"}},
			expectedMembers: []*StructMemberNode{{Name: "myMember", Type: "uint256"}},
		},
		{
			name: "Add multiple members",
			membersToAdd: []*StructMemberNode{
				{Name: "member1", Type: "uint256"},
				{Name: "member2", Type: "address"},
			},
			expectedMembers: []*StructMemberNode{
				{Name: "member1", Type: "uint256"},
				{Name: "member2", Type: "address"},
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new StructNode
			structNode := &StructNode{Name: "MyStruct"}

			// Add members to the struct
			for _, member := range tc.membersToAdd {
				structNode.AddMember(member)
			}

			// Assert that the members have been added to the struct
			assert.Equal(t, tc.expectedMembers, structNode.Members)
		})
	}
}

func TestStructNode_Children(t *testing.T) {
	// Create a StructNode
	structNode := &StructNode{Name: "MyStruct"}

	// Create StructMemberNodes
	memberNode1 := &StructMemberNode{Name: "myMember1", Type: "uint256"}
	memberNode2 := &StructMemberNode{Name: "myMember2", Type: "address"}

	// Add the members to the struct
	structNode.AddMember(memberNode1)
	structNode.AddMember(memberNode2)

	// Call the Children() method
	children := structNode.Children()

	// Assert that the result contains the member nodes
	assert.Equal(t, []Node{memberNode1, memberNode2}, children)
}
