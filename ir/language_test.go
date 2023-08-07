package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguageMethods(t *testing.T) {
	// Create new Language instances
	solidityInstance := LanguageSolidity
	vyperInstance := LanguageVyper

	// Test String method for solidityInstance
	assert.Equal(t, "solidity", solidityInstance.String())

	// Test String method for vyperInstance
	assert.Equal(t, "vyper", vyperInstance.String())
}
