package ast

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnumNode_Children(t *testing.T) {
	tests := []struct {
		name         string
		memberValues []*EnumMemberNode
		expected     []Node
	}{
		{
			name:         "Empty Enum",
			memberValues: []*EnumMemberNode{},
			expected:     []Node{},
		},
		{
			name: "Enum with Members",
			memberValues: []*EnumMemberNode{
				{Name: "Value1"},
				{Name: "Value2"},
				{Name: "Value3"},
			},
			expected: []Node{
				&EnumMemberNode{Name: "Value1"},
				&EnumMemberNode{Name: "Value2"},
				&EnumMemberNode{Name: "Value3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			enum := &EnumNode{
				Name:         "Color",
				MemberValues: test.memberValues,
			}

			children := enum.Children()

			assert.Equal(t, test.expected, children)
		})
	}
}

func TestEnumNode_JSONMarshalling(t *testing.T) {
	tests := []struct {
		name     string
		enum     *EnumNode
		expected string
	}{
		{
			name: "Empty Enum",
			enum: &EnumNode{
				Name:         "Color",
				MemberValues: []*EnumMemberNode{},
			},
			expected: `{"name":"Color","memberValues":[]}`,
		},
		{
			name: "Enum with Members",
			enum: &EnumNode{
				Name: "Color",
				MemberValues: []*EnumMemberNode{
					{Name: "Red"},
					{Name: "Green"},
					{Name: "Blue"},
				},
			},
			expected: `{"name":"Color","memberValues":[{"name":"Red"},{"name":"Green"},{"name":"Blue"}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(test.enum)
			require.NoError(t, err)

			assert.JSONEq(t, test.expected, string(jsonBytes))

			unmarshaledEnum := &EnumNode{}
			err = json.Unmarshal(jsonBytes, unmarshaledEnum)
			require.NoError(t, err)

			assert.Equal(t, test.enum, unmarshaledEnum)
		})
	}
}
