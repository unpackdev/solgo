package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatementNode_Children(t *testing.T) {
	statementNode1 := &StatementNode{
		Raw:     []string{"return", "true", ";"},
		TextRaw: "return true;",
	}

	statementNode2 := &StatementNode{
		Raw:     []string{"emit", "Event()", ";"},
		TextRaw: "emit Event();",
	}

	tt := []struct {
		name      string
		statement *StatementNode
		expected  []Node
	}{
		{
			name:      "StatementNode with 'return' statement",
			statement: statementNode1,
			expected:  nil,
		},
		{
			name:      "StatementNode with 'emit' statement",
			statement: statementNode2,
			expected:  nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			children := tc.statement.Children()
			assert.Equal(t, tc.expected, children)
		})
	}
}
