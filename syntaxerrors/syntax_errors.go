package syntaxerrors

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// SeverityLevel represents the severity of a syntax error.
type SeverityLevel int

const (
	// SeverityInfo represents a syntax error of informational level.
	SeverityInfo SeverityLevel = iota

	// SeverityWarning represents a syntax error of warning level.
	SeverityWarning

	// SeverityError represents a syntax error of error level.
	SeverityError
)

// SyntaxError represents a syntax error in a Solidity contract.
type SyntaxError struct {
	Line     int
	Column   int
	Message  string
	Severity SeverityLevel
	Context  string
}

// SyntaxErrorListener is a listener for syntax errors in Solidity contracts.
type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors   []SyntaxError
	contexts []string
}

// NewSyntaxErrorListener creates a new SyntaxErrorListener.
func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               []SyntaxError{},
	}
}

func (l *SyntaxErrorListener) PushContext(ctx string) {
	l.contexts = append(l.contexts, ctx)
}

func (l *SyntaxErrorListener) PopContext() {
	if len(l.contexts) > 0 {
		// Remove the last context
		l.contexts = l.contexts[:len(l.contexts)-1]
	}
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Replace the error message if needed
	msg = ReplaceErrorMessage(msg, "Semicolon", "';'")

	// Create a new SyntaxError
	err := SyntaxError{
		Line:     line,
		Column:   column,
		Message:  msg,
		Severity: l.determineSeverity(msg, l.currentContext()),
		Context:  l.currentContext(),
	}

	// Add the error to the Errors slice
	l.Errors = append(l.Errors, err)
}

func (l *SyntaxErrorListener) determineSeverity(msg, context string) SeverityLevel {
	// Replace the error message if needed
	msg = ReplaceErrorMessage(msg, "missing Semicolon at '}'", "missing ';' at '}'")

	// High severity errors
	highSeverityErrors := []string{
		"missing",
		"mismatched",
		"no viable alternative",
		"extraneous input",
	}
	for _, err := range highSeverityErrors {
		if strings.Contains(msg, err) {
			return SeverityError
		}
	}

	// Medium severity errors
	mediumSeverityErrors := []string{
		"cannot find symbol",
		"method not found",
	}
	for _, err := range mediumSeverityErrors {
		if strings.Contains(msg, err) {
			return SeverityWarning
		}
	}

	// Context-specific rules
	if context == "FunctionDeclaration" && strings.Contains(msg, "expected") {
		return SeverityError
	}

	// Default to low severity
	return SeverityInfo
}

func (l *SyntaxErrorListener) currentContext() string {
	// If there are no contexts, return an empty string
	if len(l.contexts) == 0 {
		return ""
	}

	// Return the current context (the last one in the slice)
	return l.contexts[len(l.contexts)-1]
}

// ReplaceErrorMessage replaces a specific error message with a new message.
func ReplaceErrorMessage(originalMsg, oldText, newText string) string {
	return strings.ReplaceAll(originalMsg, oldText, newText)
}
