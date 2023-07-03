package ast

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAstBuilder(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name     string
		contract string
		expected string
	}{
		{
			name:     "Empty Contract",
			contract: tests.ReadContractFileForTest(t, "Empty").Content,
			expected: ``,
		},
		{
			name:     "Complex AST Contract",
			contract: tests.ReadContractFileForTest(t, "AstComplex").Content,
			expected: ``,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.New(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			astBuilder := NewAstBuilder()
			err = parser.RegisterListener(solgo.ListenerAst, astBuilder)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			//jsonResponse, _ := json.MarshalIndent(astBuilder.GetTree(), "", "  ")
			//fmt.Println(string(jsonResponse))
		})
	}
}
