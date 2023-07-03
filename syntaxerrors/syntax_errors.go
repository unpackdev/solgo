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
// It includes the line and column where the error occurred, the error message, the severity of the error, and the context in which the error occurred.
type SyntaxError struct {
	Line     int
	Column   int
	Message  string
	Severity SeverityLevel
	Context  string
}

// SyntaxErrorListener is a listener for syntax errors in Solidity contracts.
// It maintains a stack of contexts and a slice of SyntaxErrors.
type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors   []SyntaxError
	contexts []string
}

// NewSyntaxErrorListener creates a new SyntaxErrorListener.
// It returns a pointer to a SyntaxErrorListener with an empty slice of SyntaxErrors and an empty stack of contexts.
func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               []SyntaxError{},
	}
}

// PushContext adds a context to the stack.
func (l *SyntaxErrorListener) PushContext(ctx string) {
	l.contexts = append(l.contexts, ctx)
}

// PopContext removes the most recent context from the stack.
func (l *SyntaxErrorListener) PopContext() {
	if len(l.contexts) > 0 {
		// Remove the last context
		l.contexts = l.contexts[:len(l.contexts)-1]
	}
}

// SyntaxError handles a syntax error.
// It creates a new SyntaxError with the given parameters and adds it to the Errors slice.
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
