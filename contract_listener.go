package solgo

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
	contractName                       string                 // contractName is the name of the contract.
	implements                         []string               // implements is a list of interfaces that the contract implements.
	imports                            []string               // imports is a list of contracts that the contract imports.
	pragmas                            []string               // pragmas is a list of pragma directives in the contract.
	libraries                          []string               // libraries is a list of libraries used in the contract.
	license                            string                 // license is the SPDX license identifier, if present.
	comments                           []string               // comments is a list of comments in the contract.
}

// NewContractListener creates a new ContractListener. It takes a SolidityParser as an argument.
func NewContractListener(parser *parser.SolidityParser) *ContractListener {
	return &ContractListener{
		parser:     parser,
		implements: make([]string, 0),
		imports:    make([]string, 0),
		pragmas:    make([]string, 0),
		libraries:  make([]string, 0),
	}
}

// EnterEveryRule is called when the parser enters any rule in the grammar.
// It is used to search for license and any comments that code has.
// ANTLR parser by default have comments disabled to be parsed as tokens, so we need to
// search for them manually using the CommonTokenStream.
func (l *ContractListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// Search for comments and licenses only once
	if l.comments == nil || l.license == "" {
		l.searchForCommentsAndLicenses()
	}
}

// EnterPragmaDirective is called when the parser enters a pragma directive.
// It extracts all pragma tokens from the context and adds them to the pragmas slice.
func (l *ContractListener) EnterPragmaDirective(ctx *parser.PragmaDirectiveContext) {
	if ctx.AllPragmaToken() != nil {
		for _, pragma := range ctx.AllPragmaToken() {
			l.pragmas = append(l.pragmas, strings.TrimSpace(pragma.GetText()))
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
	l.imports = append(l.imports, importPath)
}

// EnterContractDefinition is called when the parser enters a contract definition.
// It extracts the contract name from the context and sets it to the contractName field.
func (l *ContractListener) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	l.contractName = ctx.Identifier().GetText() // get the contract name
}

// EnterInheritanceSpecifier is called when the parser enters an inheritance specifier.
// It extracts the name of the inherited contract/interface and adds it to the implements slice.
func (l *ContractListener) EnterInheritanceSpecifier(ctx *parser.InheritanceSpecifierContext) {
	l.implements = append(l.implements, ctx.GetName().GetText())
}

// EnterUsingDirective is called when the parser enters a using directive.
// It extracts the name of the library and adds it to the libraries slice.
func (l *ContractListener) EnterUsingDirective(ctx *parser.UsingDirectiveContext) {
	for _, identifierPath := range ctx.AllIdentifierPath() {
		l.implements = append(l.implements, identifierPath.GetText())
	}
}

// GetLicense returns the SPDX license identifier, if present.
func (l *ContractListener) GetLicense() string {
	return l.license
}

// GetName returns the name of the contract.
func (l *ContractListener) GetName() string {
	return l.contractName
}

// GetPragmas returns a slice of all pragma directives in the contract.
func (l *ContractListener) GetPragmas() []string {
	return l.pragmas
}

// GetImports returns a slice of all contracts that the contract imports.
func (l *ContractListener) GetImports() []string {
	return l.imports
}

// GetImplements returns a slice of all interfaces that the contract implements.
func (l *ContractListener) GetImplements() []string {
	return l.implements
}

// GetComments returns a slice of all comments in the contract.
func (l *ContractListener) GetComments() []string {
	return l.comments
}

// searchForCommentsAndLicenses searches for comments and SPDX license identifiers in the token stream.
// It adds found comments to the comments slice and sets the license field if an SPDX license identifier is found.
func (l *ContractListener) searchForCommentsAndLicenses() {
	stream := l.parser.GetTokenStream().(*antlr.CommonTokenStream)
	tokens := stream.GetAllTokens()

	for _, token := range tokens {
		if token.GetTokenType() == parser.SolidityLexerLINE_COMMENT {
			text := token.GetText()

			if strings.HasPrefix(text, "// SPDX-License-Identifier:") ||
				strings.HasPrefix(text, "/* SPDX-License-Identifier:") {
				license := strings.TrimSpace(text[27:]) // extract the license
				l.license = license
			} else {
				// It's a regular comment
				l.comments = append(l.comments, text)
			}
		}
	}
}
