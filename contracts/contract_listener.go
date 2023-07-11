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
	l.contractInfo.Name = ctx.Identifier().GetText() // get the contract name
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
