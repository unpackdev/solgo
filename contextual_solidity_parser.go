package solgo

import "github.com/txpull/solgo/parser"

type ContextualSolidityParser struct {
	*parser.SolidityParser
	contextStack []string
}

func (p *ContextualSolidityParser) PushContext(context string) {
	p.contextStack = append(p.contextStack, context)
}

func (p *ContextualSolidityParser) PopContext() {
	p.contextStack = p.contextStack[:len(p.contextStack)-1]
}

func (p *ContextualSolidityParser) CurrentContext() string {
	if len(p.contextStack) == 0 {
		return ""
	}
	return p.contextStack[len(p.contextStack)-1]
}

// Override the methods corresponding to the grammar rules here
// For example, if you have a rule for function declarations:
func (p *ContextualSolidityParser) ContractDefinition() parser.IContractDefinitionContext {
	p.PushContext("ContractDefinition")
	defer p.PopContext()
	return p.SolidityParser.ContractDefinition() // Call the original method
}
