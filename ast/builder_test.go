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
		name                 string
		sources              solgo.Sources
		expected             string
		unresolvedReferences int64
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
				EntrySourceUnitName: "Empty",
			},
			expected:             tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast").Content,
			unresolvedReferences: 0,
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
				EntrySourceUnitName: "SimpleStorage",
			},
			expected:             tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast").Content,
			unresolvedReferences: 4,
		},
		{
			name: "OpenZeppelin ERC20 Test",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "SafeMath",
						Path:    "SafeMath.sol",
						Content: tests.ReadContractFileForTest(t, "ast/SafeMath").Content,
					},
					{
						Name:    "IERC20",
						Path:    "IERC20.sol",
						Content: tests.ReadContractFileForTest(t, "ast/IERC20").Content,
					},
					{
						Name:    "IERC20Metadata",
						Path:    "IERC20Metadata.sol",
						Content: tests.ReadContractFileForTest(t, "ast/IERC20Metadata").Content,
					},
					{
						Name:    "Context",
						Path:    "Context.sol",
						Content: tests.ReadContractFileForTest(t, "ast/Context").Content,
					},
					{
						Name:    "ERC20",
						Path:    "ERC20.sol",
						Content: tests.ReadContractFileForTest(t, "ast/ERC20").Content,
					},
				},
				EntrySourceUnitName: "ERC20",
			},
			expected:             tests.ReadJsonBytesForTest(t, "ast/ERC20.solgo.ast").Content,
			unresolvedReferences: 15,
		},

		{
			name: "Token Sale ERC20 Test",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "SafeMath",
						Path:    "SafeMath.sol",
						Content: tests.ReadContractFileForTest(t, "ast/SafeMath").Content,
					},
					{
						Name:    "IERC20",
						Path:    "IERC20.sol",
						Content: tests.ReadContractFileForTest(t, "ast/IERC20").Content,
					},
					{
						Name:    "TokenSale",
						Path:    "TokenSale.sol",
						Content: tests.ReadContractFileForTest(t, "ast/TokenSale").Content,
					},
				},
				EntrySourceUnitName: "TokenSale",
			},
			expected:             tests.ReadJsonBytesForTest(t, "ast/TokenSale.solgo.ast").Content,
			unresolvedReferences: 15,
		},
		{
			name: "Lottery Test",
			sources: solgo.Sources{
				SourceUnits: []solgo.SourceUnit{
					{
						Name:    "Lottery",
						Path:    "Lottery.sol",
						Content: tests.ReadContractFileForTest(t, "ast/Lottery").Content,
					},
				},
				EntrySourceUnitName: "Lottery",
			},
			expected: tests.ReadJsonBytesForTest(t, "ast/Lottery.solgo.ast").Content,
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

			// This step is actually quite important as it resolves all the
			// references in the AST. Without this step, the AST will be
			// incomplete.
			err = astBuilder.ResolveReferences()
			assert.NoError(t, err)

			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				prettyJson, err := astBuilder.ToPrettyJSON(sourceUnit)
				assert.NoError(t, err)
				assert.NotEmpty(t, prettyJson)

				err = astBuilder.WriteToFile(
					"../data/tests/ast/"+sourceUnit.GetName()+".solgo.ast.json",
					prettyJson,
				)
				assert.NoError(t, err)
			}

			prettyJson, err := astBuilder.ToPrettyJSON(astBuilder.GetRoot())
			assert.NoError(t, err)
			assert.NotEmpty(t, prettyJson)
			err = astBuilder.WriteToFile(
				"../data/tests/ast/"+testCase.sources.EntrySourceUnitName+".solgo.ast.json",
				prettyJson,
			)
			assert.NoError(t, err)

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)

			// Zero is here for the first contract that's empty...
			assert.GreaterOrEqual(t, astBuilder.GetRoot().EntrySourceUnit, int64(0))

			// We need to check that the entry source unit name is correct.
			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				if astBuilder.GetRoot().EntrySourceUnit == sourceUnit.GetId() {
					assert.Equal(t, sourceUnit.GetName(), testCase.sources.EntrySourceUnitName)
				}
			}

			//assert.Equal(t, testCase.expected, string(astJson))
			//fmt.Println(string(prettyJson))
		})
	}
}
