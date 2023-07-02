package solgo

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
