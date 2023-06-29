package solgo

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// SeverityLevel represents the severity of a syntax error.
type SeverityLevel int

const (
	// SeverityHigh represents a high severity error.
	SeverityHigh SeverityLevel = iota
	// SeverityMedium represents a medium severity error.
	SeverityMedium
	// SeverityLow represents a low severity error.
	SeverityLow
)

// SyntaxError represents a syntax error in a Solidity contract.
type SyntaxError struct {
	// Line is the line number where the error occurred.
	Line int
	// Column is the column number where the error occurred.
	Column int
	// Message is the error message.
	Message string
	// Severity is the severity level of the error.
	Severity SeverityLevel
	// Context is the context in which the error occurred.
	Context string
}

// SyntaxErrorListener is a listener for syntax errors in Solidity contracts.
// It extends the DefaultErrorListener from the ANTLR4 parser.
type SyntaxErrorListener struct {
	// DefaultErrorListener is the base error listener from the ANTLR4 parser.
	*antlr.DefaultErrorListener
	// Errors is a slice of SyntaxErrors.
	Errors []SyntaxError
}

// NewSyntaxErrorListener creates a new SyntaxErrorListener.
func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               []SyntaxError{},
	}
}

// SyntaxError is called when a syntax error is encountered.
// It creates a SyntaxError with the line number, column number, error message, severity level, and context,
// and adds it to the Errors slice.
func (s *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	context := ""
	if parser, ok := recognizer.(*ContextualSolidityParser); ok {
		context = parser.CurrentContext()
	}
	severity := determineSeverity(msg, context)
	s.Errors = append(s.Errors, SyntaxError{
		Line:     line,
		Column:   column,
		Message:  msg,
		Severity: severity,
		Context:  context,
	})
}

// determineSeverity determines the severity level of a syntax error based on the error message and context.
// It returns SeverityHigh for missing tokens and no viable alternative errors, and for errors in function declarations.
// It returns SeverityMedium for mismatched input and extraneous input errors, and for errors in variable declarations.
// It returns SeverityLow for all other errors.
func determineSeverity(msg string, context string) SeverityLevel {
	// High severity errors
	if strings.Contains(msg, "missing") || strings.Contains(msg, "no viable alternative") || context == "FunctionDeclaration" {
		return SeverityHigh
	}
	// Medium severity errors
	if strings.Contains(msg, "mismatched") || strings.Contains(msg, "extraneous input") || context == "VariableDeclaration" {
		return SeverityMedium
	}
	// Low severity errors
	return SeverityLow
}
