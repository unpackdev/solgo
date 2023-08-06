package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestErrorMethods(t *testing.T) {
	// Create a new Error instance
	errorInstance := &Error{
		unit:     &ast.ErrorDefinition{},
		Id:       1,
		NodeType: ast_pb.NodeType(1),
		Name:     "errorName",
		Parameters: []*Parameter{
			{
				unit:            &ast.Parameter{},
				Id:              1,
				NodeType:        ast_pb.NodeType(1),
				Name:            "parameter",
				Type:            "type",
				TypeDescription: &ast.TypeDescription{},
			},
		},
		TypeDescription: &ast.TypeDescription{},
	}

	// Test GetAST method
	assert.IsType(t, &ast.ErrorDefinition{}, errorInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), errorInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), errorInstance.GetNodeType())

	// Test GetSrc method
	assert.IsType(t, ast.SrcNode{}, errorInstance.GetSrc())
}
