package ast

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAstBuilderFromSourceAsString(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name     string
		sources  solgo.Sources
		expected string
	}{
		{
			name: "Empty Contract Test",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "Empty",
						Path:    tests.ReadContractFileForTest(t, "Empty").Path,
						Content: tests.ReadContractFileForTest(t, "Empty").Content,
					},
				},
				BaseSourceUnit: "Empty",
			},
			expected: string(tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast")),
		},
		{
			name: "Simple Storage Contract Test",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "MathLib",
						Path:    "MathLib.sol",
						Content: tests.ReadContractFileForTest(t, "ast/MathLib").Content,
					},
					{
						Name:    "SimpleStorage",
						Path:    "SimpleStorage.sol",
						Content: tests.ReadContractFileForTest(t, "ast/SimpleStorage").Content,
					},
				},
				BaseSourceUnit: "SimpleStorage",
			},
			expected: string(tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast")),
		},
		{
			name: "Simple Token",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "IERC20",
						Path:    "IERC20.sol",
						Content: tests.ReadContractFileForTest(t, "ast/IERC20").Content,
					},
				},
				BaseSourceUnit: "TokenSale",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.NewParserFromSources(context.TODO(), testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			astBuilder := NewAstBuilder(
				// We need to provide parser to the ast builder so that it can
				// access comments and other information from the parser.
				parser.GetParser(),

				// We need to provide sources to the ast builder so that it can
				// access the source code of the contracts.
				parser.GetSources(),
			)

			err = parser.RegisterListener(solgo.ListenerAst, astBuilder)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			prettyJson, err := astBuilder.ToPrettyJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, prettyJson)

			err = astBuilder.WritePrettyJSONToFile("../data/tests/ast/" + testCase.sources.BaseSourceUnit + ".solgo.ast.json")
			assert.NoError(t, err)

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)
			//assert.Equal(t, testCase.expected, string(astJson))
		})
	}
}
