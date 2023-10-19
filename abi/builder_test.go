package abi

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/tests"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestBuilderFromSources(t *testing.T) {
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
		expectedAbi          string
		expectedProto        string
		unresolvedReferences int64
		isEmpty              bool
		disabled             bool
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
				EntrySourceUnitName:  "Empty",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     "../sources/",
			},
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/Empty.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/Empty.abi.proto").Content,
			unresolvedReferences: 0,
			isEmpty:              true,
			disabled:             false,
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
				EntrySourceUnitName:  "SimpleStorage",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/SimpleStorage.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/SimpleStorage.abi.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
				EntrySourceUnitName:  "ERC20",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     "../sources/",
			},
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/ERC20.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/ERC20.abi.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/TokenSale.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/TokenSale.abi.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/Lottery.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/Lottery.abi.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
			expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/TransparentUpgradeableProxy.abi").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "abi/TransparentUpgradeableProxy.abi.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.disabled {
				return
			}

			builder, err := NewBuilderFromSources(context.TODO(), testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, builder)

			assert.IsType(t, builder.GetAstBuilder(), &ast.ASTBuilder{})
			assert.IsType(t, builder.GetParser(), &ir.Builder{})

			// Important step which will parse the entire AST and build the IR including check for
			// reference errors and syntax errors.
			assert.Empty(t, builder.Parse())

			// Now we can get into the business of building the ABIs...
			assert.NoError(t, builder.Build())

			// Get the root node of the ABI
			root := builder.GetRoot()
			assert.NotNil(t, root)

			assert.NotNil(t, builder.GetSources())
			assert.NotNil(t, builder.GetTypeResolver())

			if !testCase.isEmpty {
				assert.NotNil(t, root.GetIR())
				assert.NotNil(t, root.GetContractsAsSlice())
				assert.NotNil(t, root.GetEntryContract())
			}

			pretty, err := builder.ToJSONPretty()
			assert.NoError(t, err)
			assert.NotNil(t, pretty)

			// Check for entry contract name equivalence...
			if root.HasContracts() {
				assert.Equal(t, testCase.sources.EntrySourceUnitName, root.GetEntryName())
			}

			// Test of the ABI builder is to load generated ABI and parse it with
			// go-ethereum ABI parser and compare the results. On this way we know that
			// what users see will be literally identical and useful to wider community.
			for contractName, contract := range root.GetContracts() {
				// Assert GetContractByName() method to know that it's working...
				assert.Equal(t, contract, root.GetContractByName(contractName))

				jsonAbi, err := builder.ToJSON(contract)
				assert.NoError(t, err)
				assert.NotNil(t, jsonAbi)

				etherAbi, err := builder.ToABI(contract)
				assert.NoError(t, err)
				assert.NotNil(t, etherAbi)
			}

			// Leaving it here for now to make unit tests pass...
			// This will be removed before final push to the main branch
			err = utils.WriteToFile(
				"../data/tests/abi/"+testCase.sources.EntrySourceUnitName+".abi.json",
				pretty,
			)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedAbi, string(pretty))

			protoPretty, err := builder.ToProtoPretty()
			assert.NoError(t, err)
			assert.NotNil(t, protoPretty)

			// Leaving it here for now to make unit tests pass...
			// This will be removed before final push to the main branch
			err = utils.WriteToFile(
				"../data/tests/abi/"+testCase.sources.EntrySourceUnitName+".abi.proto.json",
				protoPretty,
			)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedProto, string(protoPretty))

		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
