package syntaxerrors

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
)

// ContextualSolidityParser is a wrapper around the SolidityParser that maintains a stack of contexts.
type ContextualSolidityParser struct {
	*parser.SolidityParser
	SyntaxErrorListener *SyntaxErrorListener
}

// NewContextualSolidityParser creates a new ContextualSolidityParser.
func NewContextualSolidityParser(tokens antlr.TokenStream, listener *SyntaxErrorListener) *ContextualSolidityParser {
	parser := parser.NewSolidityParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	return &ContextualSolidityParser{
		SolidityParser:      parser,
		SyntaxErrorListener: listener,
	}
}

func (p *ContextualSolidityParser) FunctionDefinition(ctx *parser.FunctionDefinitionContext) parser.IFunctionDefinitionContext {
	fmt.Println("OIOIIIII")
	p.SyntaxErrorListener.PushContext("FunctionDeclaration")
	defer p.SyntaxErrorListener.PopContext()
	return p.SolidityParser.FunctionDefinition() // Call the original method
}

func (p *ContextualSolidityParser) ContractDefinition() parser.IContractDefinitionContext {
	p.SyntaxErrorListener.PushContext("ContractDefinition")
	defer p.SyntaxErrorListener.PopContext()
	return p.SolidityParser.ContractDefinition() // Call the original method
}

func (p *ContextualSolidityParser) ExitContractDefinition(ctx *parser.ContractDefinitionContext) {
	p.SyntaxErrorListener.PopContext()
}
