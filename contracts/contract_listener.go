package contracts

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
)

// ContractListener is a listener for the Solidity parser that extracts information
// about contracts, including the contract name, implemented interfaces, imported contracts,
// pragmas, and comments. It also extracts the SPDX license identifier if present.
// This listener is designed to be used in conjunction with the Solidity parser to
// provide a convenient interface for working with Solidity contracts.
type ContractListener struct {
	*parser.BaseSolidityParserListener                        // BaseSolidityParserListener is the base listener from the Solidity parser.
	parser                             *parser.SolidityParser // parser is the Solidity parser instance.
	contractInfo                       ContractInfo           // contractInfo is the contract information extracted from the listener.
	hasProxyModifier                   bool
	hasAddressStateVariable            bool
	hasDelegateCall                    bool
}

// NewContractListener creates a new ContractListener. It takes a SolidityParser as an argument.
func NewContractListener(parser *parser.SolidityParser) *ContractListener {
	return &ContractListener{
		parser:       parser,
		contractInfo: ContractInfo{},
	}
}

// EnterEveryRule is called when the parser enters any rule in the grammar.
// It is used to search for license and any comments that code has.
// ANTLR parser by default have comments disabled to be parsed as tokens, so we need to
// search for them manually using the CommonTokenStream.
func (l *ContractListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// Search for comments and licenses only once
	if l.contractInfo.Comments == nil || l.contractInfo.License == "" {
		l.searchForCommentsAndLicenses()
	}
}

// EnterPragmaDirective is called when the parser enters a pragma directive.
// It extracts all pragma tokens from the context and adds them to the pragmas slice.
func (l *ContractListener) EnterPragmaDirective(ctx *parser.PragmaDirectiveContext) {
	if ctx.AllPragmaToken() != nil {
		for _, pragma := range ctx.AllPragmaToken() {
			l.contractInfo.Pragmas = append(l.contractInfo.Pragmas, strings.TrimSpace(pragma.GetText()))
		}
	}
}

// EnterImportDirective is called when the parser enters an import directive.
// It extracts the import path from the context and adds it to the imports slice.
func (l *ContractListener) EnterImportDirective(ctx *parser.ImportDirectiveContext) {
	importPath := ctx.GetText()
	importPath = strings.ReplaceAll(importPath, "import", "")
	importPath = strings.ReplaceAll(importPath, "\"", "")
	importPath = strings.ReplaceAll(importPath, ";", "")
	l.contractInfo.Imports = append(l.contractInfo.Imports, importPath)
}

// EnterContractDefinition is called when the parser enters a contract definition.
// It extracts the contract name from the context and sets it to the contractName field.
func (l *ContractListener) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	l.contractInfo.Name = ctx.Identifier().GetText()
	l.contractInfo.IsContract = true
	l.contractInfo.IsAbstract = ctx.Abstract() != nil
}

func (l *ContractListener) EnterInterfaceDefinition(ctx *parser.InterfaceDefinitionContext) {
	l.contractInfo.Name = ctx.Identifier().GetText() // get the contract name
	l.contractInfo.IsInterface = true
}

func (l *ContractListener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	l.contractInfo.Name = ctx.Identifier().GetText() // get the contract name
	l.contractInfo.IsLibrary = true
}

// EnterInheritanceSpecifier is called when the parser enters an inheritance specifier.
// It extracts the name of the inherited contract/interface and adds it to the implements slice.
func (l *ContractListener) EnterInheritanceSpecifier(ctx *parser.InheritanceSpecifierContext) {
	l.contractInfo.Implements = append(l.contractInfo.Implements, ctx.GetName().GetText())
}

// EnterUsingDirective is called when the parser enters a using directive.
// It extracts the name of the library and adds it to the libraries slice.
func (l *ContractListener) EnterUsingDirective(ctx *parser.UsingDirectiveContext) {
	for _, identifierPath := range ctx.AllIdentifierPath() {
		l.contractInfo.Implements = append(l.contractInfo.Implements, identifierPath.GetText())
	}
}

// EnterFunctionDefinition is called when the parser enters a function definition.
// It checks if there are any modifiers in the function definition and if there's a proxy modifier.
func (l *ContractListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	if ctx.AllModifierInvocation() != nil {
		for _, modifier := range ctx.AllModifierInvocation() {
			if strings.Contains(modifier.GetText(), "proxy") {
				l.hasProxyModifier = true
			}
		}
	}
}

// EnterFallbackFunctionDefinition is called when the parser enters a fallback function definition.
// It checks if there's a delegatecall in the fallback function. We use this later on to determine
// if the contract is a proxy.
func (l *ContractListener) EnterFallbackFunctionDefinition(ctx *parser.FallbackFunctionDefinitionContext) {
	if strings.Contains(ctx.GetText(), "delegatecall") {
		l.hasDelegateCall = true
	}
}

// EnterReceiveFunctionDefinition is called when the parser enters a receive function definition.
// It checks if there's a delegatecall in the receive function. We use this later on to determine
// if the contract is a proxy.
func (l *ContractListener) EnterReceiveFunctionDefinition(ctx *parser.ReceiveFunctionDefinitionContext) {
	if strings.Contains(ctx.GetText(), "delegatecall") {
		l.hasDelegateCall = true
	}
}

// EnterStateVariableDeclaration is called when the parser enters a state variable declaration.
// It checks if there's a state variable that could be storing the implementation address.
// We use this later on to determine if the contract is a proxy.
func (l *ContractListener) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	// Check if there's a state variable that could be storing the implementation address
	if strings.Contains(ctx.GetText(), "address") {
		l.hasAddressStateVariable = true
	}
}

// ExitContractDefinition is called when the parser exits a contract definition.
// It checks if the contract is a proxy and sets the IsProxy field to true if it is.
// It also sets the ProxyConfidence field based on current dummy algorithm.
func (l *ContractListener) ExitContractDefinition(ctx *parser.ContractDefinitionContext) {
	if l.hasProxyModifier && l.hasAddressStateVariable && l.hasDelegateCall {
		l.contractInfo.IsProxy = true
		l.contractInfo.ProxyConfidence = 100
	} else if l.hasProxyModifier && l.hasAddressStateVariable {
		l.contractInfo.IsProxy = true
		l.contractInfo.ProxyConfidence = 50
	} else if l.hasDelegateCall {
		l.contractInfo.IsProxy = true
		l.contractInfo.ProxyConfidence = 50
	} else if l.hasProxyModifier {
		l.contractInfo.IsProxy = true
		l.contractInfo.ProxyConfidence = 25
	}

	// Following exception is if there are imports from openzeppelin that we are sure about
	// that are proxies, but the listener doesn't detect them as proxies.
	for _, imp := range l.contractInfo.Imports {
		if strings.Contains(imp, "openzeppelin/contracts-upgradeable") {
			l.contractInfo.IsProxy = true
			l.contractInfo.ProxyConfidence = 100
		}
	}
}

// GetLicense returns the SPDX license identifier, if present.
func (l *ContractListener) GetLicense() string {
	return l.contractInfo.License
}

// GetName returns the name of the contract.
func (l *ContractListener) GetName() string {
	return l.contractInfo.Name
}

// GetPragmas returns a slice of all pragma directives in the contract.
func (l *ContractListener) GetPragmas() []string {
	return l.contractInfo.Pragmas
}

// GetImports returns a slice of all contracts that the contract imports.
func (l *ContractListener) GetImports() []string {
	return l.contractInfo.Imports
}

// GetImplements returns a slice of all interfaces that the contract implements.
func (l *ContractListener) GetImplements() []string {
	return l.contractInfo.Implements
}

// GetComments returns a slice of all comments in the contract.
func (l *ContractListener) GetComments() []string {
	return l.contractInfo.Comments
}

// GetIsProxy returns true if the contract is a proxy, false otherwise.
func (l *ContractListener) GetIsProxy() bool {
	return l.contractInfo.IsProxy
}

// GetProxyConfidence returns the confidence of the proxy detection algorithm.
func (l *ContractListener) GetProxyConfidence() int16 {
	return l.contractInfo.ProxyConfidence
}

func (l *ContractListener) IsContract() bool {
	return l.contractInfo.IsContract
}

func (l *ContractListener) IsInterface() bool {
	return l.contractInfo.IsInterface
}

func (l *ContractListener) IsLibrary() bool {
	return l.contractInfo.IsLibrary
}

func (l *ContractListener) IsAbstract() bool {
	return l.contractInfo.IsAbstract
}

// GetInfoForTests returns a map of all information extracted from the contract.
// This is used for testing purposes only
func (l *ContractListener) ToStruct() ContractInfo {
	return l.contractInfo
}

// searchForCommentsAndLicenses searches for comments and SPDX license identifiers in the token stream.
// It adds found comments to the comments slice and sets the license field if an SPDX license identifier is found.
func (l *ContractListener) searchForCommentsAndLicenses() {
	stream := l.parser.GetTokenStream().(*antlr.CommonTokenStream)
	tokens := stream.GetAllTokens()

	for _, token := range tokens {
		if token.GetTokenType() == parser.SolidityLexerLINE_COMMENT {
			text := token.GetText()

			if strings.HasPrefix(text, "// SPDX-License-Identifier:") {
				l.contractInfo.License = strings.TrimSpace(text[27:])
			} else {
				// It's a regular comment
				l.contractInfo.Comments = append(l.contractInfo.Comments, text)
			}
		}
		if token.GetTokenType() == parser.SolidityLexerCOMMENT {
			text := token.GetText()

			if strings.HasPrefix(text, "/* SPDX-License-Identifier:") {
				l.contractInfo.License = strings.TrimSpace(text[27 : len(text)-2])
			} else {
				// It's a regular comment
				l.contractInfo.Comments = append(l.contractInfo.Comments, text)
			}
		}
	}
}
