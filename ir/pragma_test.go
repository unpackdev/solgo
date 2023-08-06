package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
)

func TestPragmaMethods(t *testing.T) {
	// Create a new Pragma instance
	pragmaInstance := &Pragma{
		Id:       1,
		NodeType: ast_pb.NodeType(1),
		Literals: []string{"TestLiteral1", "TestLiteral2"},
		Text:     "pragma solidity ^0.8.0;",
	}

	// Test GetId method
	assert.Equal(t, int64(1), pragmaInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), pragmaInstance.GetNodeType())

	// Test GetLiterals method
	assert.Equal(t, []string{"TestLiteral1", "TestLiteral2"}, pragmaInstance.GetLiterals())

	// Test GetText method
	assert.Equal(t, "pragma solidity ^0.8.0;", pragmaInstance.GetText())

	// Test GetVersion method
	assert.Equal(t, "^0.8.0", pragmaInstance.GetVersion())
}
