package syntaxerrors

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/parser"
	"github.com/txpull/solgo/tests"
)

func TestSyntaxErrorListener(t *testing.T) {
	testCases := []struct {
		name     string
		contract string
		expected []SyntaxError
	}{
		{
			name:     "Randomly Corrupted Contract",
			contract: tests.ReadContractFileForTest(t, "BuggyContract").Content,
			expected: []SyntaxError{
				{
					Line:     9,
					Column:   4,
					Message:  "missing ';' at '}'",
					Severity: SeverityError,
					Context:  "",
				},
				{
					Line:     17,
					Column:   12,
					Message:  "mismatched input '(' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: SeverityError,
					Context:  "",
				},
				{
					Line:     17,
					Column:   27,
					Message:  "mismatched input ')' expecting {';', '='}",
					Severity: SeverityError,
					Context:  "",
				},
				{
					Line:     18,
					Column:   14,
					Message:  "extraneous input '=' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: SeverityError,
					Context:  "",
				},
				{
					Line:     24,
					Column:   4,
					Message:  "missing ';' at '}'",
					Severity: SeverityError,
					Context:  "",
				},
				{
					Line:     25,
					Column:   0,
					Message:  "extraneous input '}' expecting {<EOF>, 'abstract', 'address', 'bool', 'bytes', 'contract', 'enum', 'error', Fixed, FixedBytes, 'from', Function, 'global', 'import', 'interface', 'library', 'mapping', 'pragma', 'revert', SignedIntegerType, 'string', 'struct', 'type', Ufixed, UnsignedIntegerType, 'using', Identifier}",
					Severity: SeverityError,
					Context:  "",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create an ANTLR input stream from the contract string
			input := antlr.NewInputStream(tc.contract)

			// Create a Solidity lexer
			lexer := parser.NewSolidityLexer(input)

			// Create an ANTLR token stream from the lexer
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			// Create a new SyntaxErrorListener
			listener := NewSyntaxErrorListener()

			// Create a ContextualSolidityParser with the token stream and listener
			parser := NewContextualSolidityParser(tokens, listener)

			// Parse the contract
			parser.SourceUnit()

			// Check that the errors match the expected errors
			assert.Equal(t, tc.expected, listener.Errors)
		})
	}
}
