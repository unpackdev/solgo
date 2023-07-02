package solgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushContext(t *testing.T) {
	parser := &ContextualSolidityParser{
		contextStack: []string{},
	}

	parser.PushContext("ContractDefinition")

	assert.Equal(t, 1, len(parser.contextStack), "Expected context stack length to be 1")
	assert.Equal(t, "ContractDefinition", parser.contextStack[0], "Expected context to be 'ContractDefinition'")
}

func TestPopContext(t *testing.T) {
	parser := &ContextualSolidityParser{
		contextStack: []string{"ContractDefinition"},
	}

	parser.PopContext()

	assert.Equal(t, 0, len(parser.contextStack), "Expected context stack length to be 0")
}

func TestCurrentContext(t *testing.T) {
	parser := &ContextualSolidityParser{
		contextStack: []string{"ContractDefinition", "FunctionDeclaration"},
	}

	context := parser.CurrentContext()

	assert.Equal(t, "FunctionDeclaration", context, "Expected current context to be 'FunctionDeclaration'")
}

func TestCurrentContextEmpty(t *testing.T) {
	parser := &ContextualSolidityParser{
		contextStack: []string{},
	}

	context := parser.CurrentContext()

	assert.Equal(t, "", context, "Expected current context to be an empty string")
}
