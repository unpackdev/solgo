package solgo

import "github.com/txpull/solgo/parser"

// ContextualSolidityParser is a wrapper around the SolidityParser that maintains a stack of contexts.
// This allows the parser to keep track of the current context (e.g., within a contract definition, function definition, etc.)
// as it parses a Solidity contract.
type ContextualSolidityParser struct {
	*parser.SolidityParser          // SolidityParser is the base parser from the Solidity parser.
	contextStack           []string // contextStack is a stack of contexts. Each context corresponds to a rule in the grammar.
}

// PushContext pushes a new context onto the context stack. This should be called when the parser enters a new rule.
func (p *ContextualSolidityParser) PushContext(context string) {
	p.contextStack = append(p.contextStack, context)
}

// PopContext pops the current context from the context stack. This should be called when the parser exits a rule.
func (p *ContextualSolidityParser) PopContext() {
	p.contextStack = p.contextStack[:len(p.contextStack)-1]
}

// CurrentContext returns the current context, i.e., the context at the top of the stack.
// If the stack is empty, it returns an empty string.
func (p *ContextualSolidityParser) CurrentContext() string {
	if len(p.contextStack) == 0 {
		return ""
	}
	return p.contextStack[len(p.contextStack)-1]
}

// ContractDefinition is called when the parser enters a contract definition.
// It pushes "ContractDefinition" onto the context stack, calls the original ContractDefinition method,
// and then pops the context from the stack before returning.
func (p *ContextualSolidityParser) ContractDefinition() parser.IContractDefinitionContext {
	p.PushContext("ContractDefinition")
	defer p.PopContext()
	return p.SolidityParser.ContractDefinition() // Call the original method
}
