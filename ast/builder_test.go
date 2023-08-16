package ast

import (
	"context"
	"fmt"
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
		{
			name:       "SushiXSwap Nightmare Test",
			outputPath: "contracts/sushixswap/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Address",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/Address").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/Address").Content,
					},
					{
						Name:    "BentoAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/BentoAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/BentoAdapter").Content,
					},
					{
						Name:    "IBentoBoxMinimal",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IBentoBoxMinimal").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IBentoBoxMinimal").Content,
					},
					{
						Name:    "IERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IERC20").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IERC20").Content,
					},
					{
						Name:    "IImmutableState",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IImmutableState").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IImmutableState").Content,
					},
					{
						Name:    "ImmutableState",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ImmutableState").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ImmutableState").Content,
					},
					{
						Name:    "IPool",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IPool").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IPool").Content,
					},
					{
						Name:    "IStargateAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateAdapter").Content,
					},
					{
						Name:    "IStargateReceiver",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateReceiver").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateReceiver").Content,
					},
					{
						Name:    "IStargateRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateRouter").Content,
					},
					{
						Name:    "IStargateWidget",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateWidget").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateWidget").Content,
					},
					{
						Name:    "ISushiXSwap",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ISushiXSwap").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ISushiXSwap").Content,
					},
					{
						Name:    "ITridentRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentRouter").Content,
					},
					{
						Name:    "ITridentSwapAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentSwapAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentSwapAdapter").Content,
					},
					{
						Name:    "IUniswapV2Pair",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IUniswapV2Pair").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IUniswapV2Pair").Content,
					},
					{
						Name:    "IWETH",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IWETH").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IWETH").Content,
					},
					{
						Name:    "SafeERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeERC20").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeERC20").Content,
					},
					{
						Name:    "SafeMath",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeMath").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeMath").Content,
					},
					{
						Name:    "StargateAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/StargateAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/StargateAdapter").Content,
					},
					{
						Name:    "SushiLegacyAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiLegacyAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiLegacyAdapter").Content,
					},
					{
						Name:    "SushiXSwap",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiXSwap").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiXSwap").Content,
					},
					{
						Name:    "TokenAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/TokenAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/TokenAdapter").Content,
					},
					{
						Name:    "TridentSwapAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/TridentSwapAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/TridentSwapAdapter").Content,
					},
					{
						Name:    "UniswapV2Library",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/UniswapV2Library").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/UniswapV2Library").Content,
					},
				},
				EntrySourceUnitName: "SushiXSwap",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/sushixswap/SushiXSwap.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/sushixswap/SushiXSwap.solgo.ast.proto").Content,
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

				// Recursive test against all nodes. A common place where we can add tests to check
				// if the AST is correct.
				recursiveTest(t, sourceUnit)
			}

		})
	}
}

func recursiveTest(t *testing.T, node Node[NodeType]) {
	assert.NotNil(t, node.GetNodes(), fmt.Sprintf("Node %T has nil nodes", node))
	assert.GreaterOrEqual(t, node.GetId(), int64(0), fmt.Sprintf("Node %T has empty id", node))
	assert.NotNil(t, node.GetType(), fmt.Sprintf("Node %T has empty type", node))
	assert.NotNil(t, node.GetSrc(), fmt.Sprintf("Node %T has empty GetSrc()", node))
	assert.NotNil(t, node.GetTypeDescription(), fmt.Sprintf("Node %T has not defined GetTypeDescription()", node))

	if contract, ok := node.(*Contract); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)
		assert.NotNil(t, contract.IsAbstract())
		assert.NotNil(t, contract.GetKind())
		assert.NotNil(t, contract.IsFullyImplemented())

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}

	}

	if contract, ok := node.(*Library); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}
	}

	if contract, ok := node.(*Interface); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}
	}

	for _, childNode := range node.GetNodes() {
		recursiveTest(t, childNode)
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}

func TestAstReferenceSetDescriptor(t *testing.T) {
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
		{
			name:       "SushiXSwap Nightmare Test",
			outputPath: "contracts/sushixswap/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Address",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/Address").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/Address").Content,
					},
					{
						Name:    "BentoAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/BentoAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/BentoAdapter").Content,
					},
					{
						Name:    "IBentoBoxMinimal",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IBentoBoxMinimal").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IBentoBoxMinimal").Content,
					},
					{
						Name:    "IERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IERC20").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IERC20").Content,
					},
					{
						Name:    "IImmutableState",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IImmutableState").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IImmutableState").Content,
					},
					{
						Name:    "ImmutableState",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ImmutableState").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ImmutableState").Content,
					},
					{
						Name:    "IPool",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IPool").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IPool").Content,
					},
					{
						Name:    "IStargateAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateAdapter").Content,
					},
					{
						Name:    "IStargateReceiver",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateReceiver").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateReceiver").Content,
					},
					{
						Name:    "IStargateRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateRouter").Content,
					},
					{
						Name:    "IStargateWidget",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateWidget").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IStargateWidget").Content,
					},
					{
						Name:    "ISushiXSwap",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ISushiXSwap").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ISushiXSwap").Content,
					},
					{
						Name:    "ITridentRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentRouter").Content,
					},
					{
						Name:    "ITridentSwapAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentSwapAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/ITridentSwapAdapter").Content,
					},
					{
						Name:    "IUniswapV2Pair",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IUniswapV2Pair").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IUniswapV2Pair").Content,
					},
					{
						Name:    "IWETH",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/IWETH").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/IWETH").Content,
					},
					{
						Name:    "SafeERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeERC20").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeERC20").Content,
					},
					{
						Name:    "SafeMath",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeMath").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SafeMath").Content,
					},
					{
						Name:    "StargateAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/StargateAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/StargateAdapter").Content,
					},
					{
						Name:    "SushiLegacyAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiLegacyAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiLegacyAdapter").Content,
					},
					{
						Name:    "SushiXSwap",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiXSwap").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/SushiXSwap").Content,
					},
					{
						Name:    "TokenAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/TokenAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/TokenAdapter").Content,
					},
					{
						Name:    "TridentSwapAdapter",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/TridentSwapAdapter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/TridentSwapAdapter").Content,
					},
					{
						Name:    "UniswapV2Library",
						Path:    tests.ReadContractFileForTest(t, "contracts/sushixswap/UniswapV2Library").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/sushixswap/UniswapV2Library").Content,
					},
				},
				EntrySourceUnitName: "SushiXSwap",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/sushixswap/SushiXSwap.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/sushixswap/SushiXSwap.solgo.ast.proto").Content,
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

			// We need to check that the entry source unit name is correct.
			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				recursiveReferenceDescriptorSetTest(t, sourceUnit)
			}

		})
	}
}

func recursiveReferenceDescriptorSetTest(t *testing.T, node Node[NodeType]) {
	node.SetReferenceDescriptor(0, &TypeDescription{})

	for _, childNode := range node.GetNodes() {
		childNode.SetReferenceDescriptor(0, &TypeDescription{})
		recursiveReferenceDescriptorSetTest(t, childNode)
	}
}
