package solgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSyntaxErrorListener(t *testing.T) {
	listener := NewSyntaxErrorListener()

	assert.NotNil(t, listener.DefaultErrorListener, "Expected DefaultErrorListener to be initialized")
	assert.NotNil(t, listener.Errors, "Expected Errors slice to be initialized")
	assert.Equal(t, 0, len(listener.Errors), "Expected Errors slice to be empty")
}

func TestSyntaxErrorListener_SyntaxError(t *testing.T) {
	listener := NewSyntaxErrorListener()

	listener.SyntaxError(nil, nil, 1, 1, "missing ';'", nil)

	assert.Equal(t, 1, len(listener.Errors), "Expected one error to be recorded")
	assert.Equal(t, 1, listener.Errors[0].Line, "Expected error line to be 1")
	assert.Equal(t, 1, listener.Errors[0].Column, "Expected error column to be 1")
	assert.Equal(t, "missing ';'", listener.Errors[0].Message, "Expected error message to be 'missing ';''")
	assert.Equal(t, SeverityHigh, listener.Errors[0].Severity, "Expected error severity to be SeverityHigh")
	assert.Equal(t, "", listener.Errors[0].Context, "Expected error context to be an empty string")
}

func TestDetermineSeverity(t *testing.T) {
	assert.Equal(t, SeverityHigh, determineSeverity("missing ';'", "FunctionDeclaration"), "Expected severity to be SeverityHigh for missing token in FunctionDeclaration context")
	assert.Equal(t, SeverityMedium, determineSeverity("mismatched input", "VariableDeclaration"), "Expected severity to be SeverityMedium for mismatched input in VariableDeclaration context")
	assert.Equal(t, SeverityMedium, determineSeverity("extraneous input", "Expression"), "Expected severity to be SeverityLow for extraneous input in Expression context")
}
