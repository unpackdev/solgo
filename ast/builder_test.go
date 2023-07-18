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
			expected: `{"source_units":[{"id":0,"license":"","node_type":"SourceUnit","nodes":[],"src":{"line":1,"start":0,"end":0,"length":0,"index":0}}]}`,
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

			/* 			for contractName, contract := range testCase.contracts {

				// Register the contract information listener



				syntaxErrs := parser.Parse()
				assert.Empty(t, syntaxErrs)

				prettyJson, err := astBuilder.ToPrettyJSON()
				assert.NoError(t, err)
				assert.NotEmpty(t, prettyJson)

				//fmt.Println(string(prettyJson))

				err = astBuilder.WriteJSONToFile("../data/tests/ast/" + contractName + ".solgo.ast.json")
				assert.NoError(t, err)

				 				astJson, err := astBuilder.ToJSON()
				   				assert.NoError(t, err)
				   				assert.NotEmpty(t, astJson)
				   				ioutil.WriteFile(testCase.name+".json", []byte(astJson), 0777)
				   				assert.Equal(t, testCase.expected, astJson)
			} */
		})
	}
}
