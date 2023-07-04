package syntaxerrors

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
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

// Error returns the error message.
func (e SyntaxError) Error() error {
	return fmt.Errorf("syntax error: %s at line %d, column %d in context '%s'. Severity: %s", e.Message, e.Line, e.Column, e.Context, e.Severity.String())
}

// SyntaxErrorListener is a listener for syntax errors in Solidity contracts.
// It maintains a stack of contexts and a slice of SyntaxErrors.
type SyntaxErrorListener struct {
	*parser.BaseSolidityParserListener
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
		Severity: l.DetermineSeverity(msg, l.currentContext()),
		Context:  l.currentContext(),
	}

	// Add the error to the Errors slice
	l.Errors = append(l.Errors, err)
}

func (l *SyntaxErrorListener) currentContext() string {
	// If there are no contexts, return an empty string
	if len(l.contexts) == 0 {
		return ""
	}

	// Return the current context (the last one in the slice)
	return l.contexts[len(l.contexts)-1]
}
