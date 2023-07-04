package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/syntaxerrors"
	"github.com/txpull/solgo/tests"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")
	assert.Equal(t, input, solgo.GetInput(), "input reader is not set correctly")
	assert.NotNil(t, solgo.GetLexer(), "lexer is not initialized")
	assert.NotNil(t, solgo.GetParser(), "parser is not initialized")
}

func TestParse(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	errs := solgo.Parse()
	assert.Empty(t, errs, "syntax errors encountered during parsing")
}

func TestParseWithError(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	errs := solgo.Parse()
	assert.NotEmpty(t, errs, "expected syntax errors, but none were encountered")
}

func TestGetTree(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	tree := solgo.GetTree()
	assert.NotNil(t, tree, "parse tree is not generated")
}

func TestGetTokenStream(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	tokenStream := solgo.GetTokenStream()
	assert.NotNil(t, tokenStream, "token stream is not generated")
}

func TestGetInputStream(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	inputStream := solgo.GetInputStream()
	assert.NotNil(t, inputStream, "input stream is not generated")
}

func TestGetLexer(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	lexer := solgo.GetLexer()
	assert.NotNil(t, lexer, "lexer is not generated")
}

func TestGetInput(t *testing.T) {
	ctx := context.Background()
	input := strings.NewReader("contract Test {}")

	solgo, err := New(ctx, input)
	assert.NoError(t, err, "error creating SolGo instance")

	assert.Equal(t, input, solgo.GetInput(), "input reader is not returned correctly")
}

func TestNew_SyntaxErrors(t *testing.T) {
	testCases := []struct {
		name     string
		contract string
		expected []syntaxerrors.SyntaxError
	}{
		{
			name:     "Randomly Corrupted Contract",
			contract: tests.ReadContractFileForTestFromRootPath(t, "BuggyContract").Content,
			expected: []syntaxerrors.SyntaxError{
				{
					Line:     9,
					Column:   4,
					Message:  "missing ';' at '}'",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     17,
					Column:   12,
					Message:  "mismatched input '(' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     17,
					Column:   27,
					Message:  "mismatched input ')' expecting {';', '='}",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     18,
					Column:   14,
					Message:  "extraneous input '=' expecting {'constant', 'error', 'from', 'global', 'immutable', 'internal', 'override', 'private', 'public', 'revert', Identifier}",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     24,
					Column:   4,
					Message:  "missing ';' at '}'",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
				{
					Line:     25,
					Column:   0,
					Message:  "extraneous input '}' expecting {<EOF>, 'abstract', 'address', 'bool', 'bytes', 'contract', 'enum', 'error', Fixed, FixedBytes, 'from', Function, 'global', 'import', 'interface', 'library', 'mapping', 'pragma', 'revert', SignedIntegerType, 'string', 'struct', 'type', Ufixed, UnsignedIntegerType, 'using', Identifier}",
					Severity: syntaxerrors.SeverityError,
					Context:  "SourceUnit",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new SolGo instance
			solGo, err := New(context.Background(), strings.NewReader(tc.contract))
			assert.NoError(t, err)
			assert.NotNil(t, solGo)

			syntaxErrors := solGo.Parse()

			// Check that the syntax errors match the expected syntax errors
			assert.Equal(t, tc.expected, syntaxErrors)
		})
	}
}
