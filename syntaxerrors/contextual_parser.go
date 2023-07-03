package syntaxerrors

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
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

// FunctionDefinition parses a function definition and adds the "FunctionDeclaration" context to the SyntaxErrorListener.
func (p *ContextualParser) FunctionDefinition(ctx *parser.FunctionDefinitionContext) parser.IFunctionDefinitionContext {
	p.SyntaxErrorListener.PushContext("FunctionDeclaration")
	defer p.SyntaxErrorListener.PopContext()
	return p.SolidityParser.FunctionDefinition() // Call the original method
}

// ContractDefinition parses a contract definition and adds the "ContractDefinition" context to the SyntaxErrorListener.
func (p *ContextualParser) ContractDefinition() parser.IContractDefinitionContext {
	p.SyntaxErrorListener.PushContext("ContractDefinition")
	defer p.SyntaxErrorListener.PopContext()
	return p.SolidityParser.ContractDefinition() // Call the original method
}
