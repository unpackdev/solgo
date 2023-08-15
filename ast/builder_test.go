package ast

import (
	"context"
	"path/filepath"
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
		outputPath           string
		sources              *solgo.Sources
		expectedAst          string
		expectedProto        string
		unresolvedReferences int64
	}{
		{
			name:       "Empty Contract Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Empty",
						Path:    tests.ReadContractFileForTest(t, "Empty").Path,
						Content: tests.ReadContractFileForTest(t, "Empty").Content,
					},
				},
				EntrySourceUnitName: "Empty",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Simple Storage Contract Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
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
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "OpenZeppelin ERC20 Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
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
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/ERC20.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/ERC20.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Token With Reference Resolution",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Token",
						Path:    "Token.sol",
						Content: tests.ReadContractFileForTest(t, "ast/Token").Content,
					},
				},
				EntrySourceUnitName: "Token",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Token.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Token Sale ERC20 Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "TokenSale",
						Path:    "TokenSale.sol",
						Content: tests.ReadContractFileForTest(t, "ast/TokenSale").Content,
					},
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
				},
				EntrySourceUnitName: "TokenSale",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/TokenSale.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/TokenSale.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Lottery Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Lottery",
						Path:    "Lottery.sol",
						Content: tests.ReadContractFileForTest(t, "ast/Lottery").Content,
					},
				},
				EntrySourceUnitName: "Lottery",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Lottery.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Lottery.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Cheelee Test", // Took this one as I could discover ipfs metadata :joy:
			outputPath: "contracts/cheelee/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Import",
						Path:    "Import.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/Import").Content,
					},
					{
						Name:    "BeaconProxy",
						Path:    "BeaconProxy.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/BeaconProxy").Content,
					},
					{
						Name:    "UpgradeableBeacon",
						Path:    "UpgradeableBeacon.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/UpgradeableBeacon").Content,
					},
					{
						Name:    "ERC1967Proxy",
						Path:    "ERC1967Proxy.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/ERC1967Proxy").Content,
					},
					{
						Name:    "TransparentUpgradeableProxy",
						Path:    "TransparentUpgradeableProxy.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/TransparentUpgradeableProxy").Content,
					},
					{
						Name:    "ProxyAdmin",
						Path:    "ProxyAdmin.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/ProxyAdmin").Content,
					},
					{
						Name:    "IBeacon",
						Path:    "IBeacon.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/IBeacon").Content,
					},
					{
						Name:    "Proxy",
						Path:    "Proxy.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/Proxy").Content,
					},
					{
						Name:    "ERC1967Upgrade",
						Path:    "ERC1967Upgrade.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/ERC1967Upgrade").Content,
					},
					{
						Name:    "Address",
						Path:    "Address.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/Address").Content,
					},
					{
						Name:    "StorageSlot",
						Path:    "StorageSlot.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/StorageSlot").Content,
					},
					{
						Name:    "Ownable",
						Path:    "Ownable.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/Ownable").Content,
					},
					{
						Name:    "Context",
						Path:    "Context.sol",
						Content: tests.ReadContractFileForTest(t, "contracts/cheelee/Context").Content,
					},
				},
				EntrySourceUnitName: "TransparentUpgradeableProxy",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/cheelee/TransparentUpgradeableProxy.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/cheelee/TransparentUpgradeableProxy.solgo.ast.proto").Content,
			unresolvedReferences: 0,
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
			errs := astBuilder.ResolveReferences()
			var errsExpected []error
			assert.Equal(t, errsExpected, errs)
			assert.Equal(t, int(testCase.unresolvedReferences), astBuilder.GetResolver().GetUnprocessedCount())

			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				prettyJson, err := astBuilder.ToPrettyJSON(sourceUnit)
				assert.NoError(t, err)
				assert.NotEmpty(t, prettyJson)

				err = astBuilder.WriteToFile(
					"../data/tests/"+testCase.outputPath+sourceUnit.GetName()+".solgo.ast.json",
					prettyJson,
				)
				assert.NoError(t, err)
			}

			prettyJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, prettyJson)
			err = astBuilder.WriteToFile(
				"../data/tests/"+testCase.outputPath+testCase.sources.EntrySourceUnitName+".solgo.ast.json",
				prettyJson,
			)
			assert.NoError(t, err)
			//assert.Equal(t, testCase.expectedAst, string(prettyJson))

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)

			astPretty, _ := astBuilder.ToPrettyJSON(astBuilder.ToProto())
			err = astBuilder.WriteToFile(
				"../data/tests/"+testCase.outputPath+testCase.sources.EntrySourceUnitName+".solgo.ast.proto.json",
				astPretty,
			)
			assert.NoError(t, err)
			assert.NotEmpty(t, astPretty)
			//assert.Equal(t, testCase.expectedProto, string(astPretty))

			// Zero is here for the first contract that's empty...
			assert.GreaterOrEqual(t, astBuilder.GetRoot().EntrySourceUnit, int64(0))

			// We need to check that the entry source unit name is correct.
			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				if astBuilder.GetRoot().EntrySourceUnit == sourceUnit.GetId() {
					assert.Equal(t, sourceUnit.GetName(), testCase.sources.EntrySourceUnitName)
				}
			}

		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
