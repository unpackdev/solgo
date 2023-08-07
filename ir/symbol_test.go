package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbols(t *testing.T) {
	symbol := &Symbol{
		Id:           1,
		Name:         "TestVariable",
		AbsolutePath: "TestAbsolutePath.sol",
	}

	// Test GetId method
	assert.Equal(t, int64(1), symbol.GetId())

	// Test GetName method
	assert.Equal(t, "TestVariable", symbol.GetName())

	// Test GetAbsolutePath method
	assert.Equal(t, "TestAbsolutePath.sol", symbol.GetAbsolutePath())
}
