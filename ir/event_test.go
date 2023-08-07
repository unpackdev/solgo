package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestEventMethods(t *testing.T) {
	// Create a new Event instance
	eventInstance := &Event{
		unit:      &ast.EventDefinition{},
		Id:        1,
		NodeType:  ast_pb.NodeType(1),
		Name:      "eventName",
		Anonymous: true,
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
	}

	// Test GetAST method
	assert.IsType(t, &ast.EventDefinition{}, eventInstance.GetAST())

	// Test GetId method
	assert.Equal(t, int64(1), eventInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), eventInstance.GetNodeType())

	// Test GetName method
	assert.Equal(t, "eventName", eventInstance.GetName())

	// Test IsAnonymous method
	assert.Equal(t, true, eventInstance.IsAnonymous())

	// Test GetParameters method
	assert.IsType(t, []*Parameter{}, eventInstance.GetParameters())
}
