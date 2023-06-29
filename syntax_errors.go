package solgo

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type SeverityLevel int

const (
	SeverityHigh SeverityLevel = iota
	SeverityMedium
	SeverityLow
)

type SyntaxError struct {
	Line     int
	Column   int
	Message  string
	Severity SeverityLevel
	Context  string
}

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []SyntaxError
}

func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               []SyntaxError{},
	}
}

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
