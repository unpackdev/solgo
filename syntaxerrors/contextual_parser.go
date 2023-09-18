package syntaxerrors

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/unpackdev/solgo/parser"
)

// ContextualParser is a wrapper around the SolidityParser that maintains a stack of contexts.
// It provides methods for parsing function and contract definitions, and adds the current context to the SyntaxErrorListener.
type ContextualParser struct {
	*parser.SolidityParser
	SyntaxErrorListener *SyntaxErrorListener
}

// NewContextualParser creates a new ContextualSolidityParser.
// It takes a token stream and a SyntaxErrorListener, and returns a pointer to a ContextualSolidityParser.
func NewContextualParser(tokens antlr.TokenStream, listener *SyntaxErrorListener) *ContextualParser {
	parser := parser.NewSolidityParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	return &ContextualParser{
		SolidityParser:      parser,
		SyntaxErrorListener: listener,
	}
}

// SourceUnit parses a function definition and adds the "SourceUnit" context to the SyntaxErrorListener.
func (p *ContextualParser) SourceUnit() parser.ISourceUnitContext {
	p.SyntaxErrorListener.PushContext("SourceUnit")
	defer p.SyntaxErrorListener.PopContext()
	return p.SolidityParser.SourceUnit() // Call the original method
}
