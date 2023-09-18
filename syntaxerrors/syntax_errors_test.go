package syntaxerrors

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/parser"
	"github.com/unpackdev/solgo/tests"
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
					Context:  "SourceUnit",
				},
				{
					Line:     17,
					Column:   12,
					Message:  "mismatched input '(' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     17,
					Column:   27,
					Message:  "mismatched input ')' expecting {';', '='}",
					Severity: SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     18,
					Column:   14,
					Message:  "extraneous input '=' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     24,
					Column:   4,
					Message:  "missing ';' at '}'",
					Severity: SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     25,
					Column:   0,
					Message:  "extraneous input '}' expecting {<EOF>, 'abstract', 'address', 'bool', 'bytes', 'contract', 'enum', 'error', Fixed, FixedBytes, 'from', Function, 'global', 'import', 'interface', 'library', 'mapping', 'pragma', 'revert', SignedIntegerType, 'string', 'struct', 'type', Ufixed, UnsignedIntegerType, 'using', Identifier}",
					Severity: SeverityError,
					Context:  "SourceUnit",
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

			// Create a ContextualParser with the token stream and listener
			parser := NewContextualParser(tokens, listener)

			// Parse the contract
			parser.SourceUnit()

			// Check that the errors match the expected errors
			assert.Equal(t, tc.expected, listener.Errors)

			for _, err := range listener.Errors {
				assert.Equal(t, err.Context, "SourceUnit")
				assert.Equal(t, err.Severity.String(), SeverityError.String())
				assert.Equal(t, err.Line > 0, true)
				assert.Equal(t, err.Message != "", true)
				assert.NotEmpty(t, err.Error())
			}
		})
	}
}
