package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo"
)

func TestContinueStatement(t *testing.T) {
	// Initialize the ASTBuilder and ContinueStatement
	b := NewAstBuilder(nil, &solgo.Sources{})
	cs := NewContinueStatement(b)

	// Define test cases
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Test NewContinueStatement",
			test: func(t *testing.T) {
				assert.Equal(t, b.GetNextID(), cs.GetId()+1)
				assert.Equal(t, ast_pb.NodeType_CONTINUE, cs.GetType())
			},
		},
		{
			name: "Test GetId",
			test: func(t *testing.T) {
				assert.Equal(t, cs.Id, cs.GetId())
			},
		},
		{
			name: "Test GetType",
			test: func(t *testing.T) {
				assert.Equal(t, cs.NodeType, cs.GetType())
			},
		},
		{
			name: "Test GetSrc",
			test: func(t *testing.T) {
				assert.Equal(t, cs.Src, cs.GetSrc())
			},
		},
		{
			name: "Test GetTypeDescription",
			test: func(t *testing.T) {
				assert.Nil(t, cs.GetTypeDescription())
			},
		},
		{
			name: "Test GetNodes",
			test: func(t *testing.T) {
				assert.Nil(t, cs.GetNodes())
			},
		},
		{
			name: "Test ToProto",
			test: func(t *testing.T) {
				//assert.Equal(t, ast_pb.Continue{}, cs.ToProto())
			},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}
