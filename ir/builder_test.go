package ir

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/tests"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestIrBuilderFromSources(t *testing.T) {
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
		wantErr              bool
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
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/Empty.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/Empty.ir.proto").Content,
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
				EntrySourceUnitName:  "SimpleStorage",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/SimpleStorage.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/SimpleStorage.ir.proto").Content,
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
				EntrySourceUnitName:  "ERC20",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/ERC20.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/ERC20.ir.proto").Content,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/TokenSale.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/TokenSale.ir.proto").Content,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/Lottery.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/Lottery.ir.proto").Content,
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
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/TransparentUpgradeableProxy.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/TransparentUpgradeableProxy.ir.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:                 "Nil Sources",
			outputPath:           "ast/",
			unresolvedReferences: 0,
			wantErr:              true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := NewBuilderFromSources(context.TODO(), testCase.sources)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Nil(t, parser)
				return
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, parser)
			}
			assert.IsType(t, &Builder{}, parser)
			assert.IsType(t, &ast.ASTBuilder{}, parser.GetAstBuilder())
			assert.IsType(t, &solgo.Parser{}, parser.GetParser())
			assert.IsType(t, &solgo.Sources{}, parser.GetSources())

			// Important step which will parse the sources and build the AST including check for
			// reference errors and syntax errors.
			// If you wish to only parse the sources without checking for errors, use
			// parser.GetParser().Parse()
			assert.Empty(t, parser.Parse())

			// Now we can get into the business of building the intermediate representation
			assert.NoError(t, parser.Build())

			// Get the root node of the IR
			root := parser.GetRoot()
			assert.NotNil(t, root)

			pretty, err := parser.ToJSONPretty()
			assert.NoError(t, err)
			assert.NotNil(t, pretty)

			// Leaving it here for now to make unit tests pass...
			// This will be removed before final push to the main branch
			err = utils.WriteToFile(
				"../data/tests/ir/"+testCase.sources.EntrySourceUnitName+".ir.json",
				pretty,
			)
			assert.NoError(t, err)

			assert.NotNil(t, parser.ToProto())

			j, err := parser.ToJSON()
			assert.NotNil(t, j)
			assert.NoError(t, err)

			protoPretty, err := parser.ToProtoPretty()
			assert.NoError(t, err)
			assert.NotNil(t, protoPretty)

			// Leaving it here for now to make unit tests pass...
			// This will be removed before final push to the main branch
			err = utils.WriteToFile(
				"../data/tests/ir/"+testCase.sources.EntrySourceUnitName+".ir.proto.json",
				protoPretty,
			)
			assert.NoError(t, err)

			for _, eip := range root.GetStandards() {
				assert.NotNil(t, eip)
				assert.NotNil(t, eip.GetConfidence())
				assert.NotNil(t, eip.GetContractId())
				assert.NotNil(t, eip.GetContractName())
				assert.NotNil(t, eip.GetStandard())
				assert.NotNil(t, eip.ToProto())
			}

			assert.NotNil(t, root.HasStandard(standards.ERC1014))
			assert.NotNil(t, root.GetAST())

			for _, contract := range root.GetContracts() {
				assert.NotNil(t, contract)
				assert.NotNil(t, contract.GetAST())
				assert.NotNil(t, contract.GetSrc())
				assert.NotNil(t, contract.GetUnitSrc())
				assert.NotNil(t, contract.ToProto())

				if contract.GetFallback() != nil {
					assert.NotNil(t, contract.GetFallback().GetAST())
					assert.NotNil(t, contract.GetFallback().GetSrc())
					assert.NotNil(t, contract.GetFallback().ToProto())
				}

				if contract.GetReceive() != nil {
					assert.NotNil(t, contract.GetReceive().GetAST())
					assert.NotNil(t, contract.GetReceive().GetSrc())
					assert.NotNil(t, contract.GetReceive().ToProto())
				}

				if contract.GetConstructor() != nil {
					assert.NotNil(t, contract.GetConstructor().GetAST())
					assert.NotNil(t, contract.GetConstructor().GetSrc())
					assert.NotNil(t, contract.GetConstructor().ToProto())
				}

				for _, pragma := range contract.GetPragmas() {
					assert.NotNil(t, pragma)
					assert.NotNil(t, pragma.GetAST())
					assert.NotNil(t, pragma.GetSrc())
					assert.NotNil(t, pragma.ToProto())
				}

				for _, enum := range contract.GetEnums() {
					assert.NotNil(t, enum)
					assert.NotNil(t, enum.GetAST())
					assert.NotNil(t, enum.GetSrc())
					assert.NotNil(t, enum.ToProto())
				}

				for _, structs := range contract.GetStructs() {
					assert.NotNil(t, structs)
					assert.NotNil(t, structs.GetAST())
					assert.NotNil(t, structs.GetSrc())
					assert.NotNil(t, structs.ToProto())
				}

				for _, function := range contract.GetFunctions() {
					assert.NotNil(t, function)
					assert.NotNil(t, function.GetAST())
					assert.NotNil(t, function.GetSrc())
					assert.NotNil(t, function.GetBody().GetAST())
					assert.NotNil(t, function.GetBody().GetSrc())
					assert.NotNil(t, function.ToProto())
					assert.NotEmpty(t, function.GetSignature())

					for _, param := range function.GetParameters() {
						assert.NotNil(t, param)
						assert.NotNil(t, param.GetAST())
						assert.NotNil(t, param.GetSrc())
						assert.NotNil(t, param.ToProto())
					}

					for _, modifier := range function.GetModifiers() {
						assert.NotNil(t, modifier)
						assert.NotNil(t, modifier.GetAST())
						assert.NotNil(t, modifier.GetSrc())
						assert.NotNil(t, modifier.ToProto())
					}

					for _, override := range function.GetOverrides() {
						assert.NotNil(t, override)
						assert.NotNil(t, override.GetAST())
						assert.NotNil(t, override.GetSrc())
						assert.NotNil(t, override.ToProto())
					}

					for _, returns := range function.GetReturnStatements() {
						assert.NotNil(t, returns)
						assert.NotNil(t, returns.GetAST())
						assert.NotNil(t, returns.GetSrc())
						assert.NotNil(t, returns.ToProto())
					}

					for _, statement := range function.GetBody().GetStatements() {
						if node, ok := statement.(*FunctionCall); ok {
							assert.NotNil(t, node)
							assert.NotNil(t, node.GetAST())
							assert.NotNil(t, node.GetSrc())
							assert.NotNil(t, node.GetId())
							assert.Equal(t, len(node.GetNodes()), 0)
							assert.NotNil(t, node.ToProto())
							if node.GetExternalContract() != nil {
								assert.NotNil(t, node.GetExternalContract().GetSrc())
							}
							if node.GetReferenceStatement() != nil {
								assert.NotEmpty(t, node.GetReferenceStatement().GetSrc())
							}
						}
					}
				}

				for _, variable := range contract.GetStateVariables() {
					assert.NotNil(t, variable)
					assert.NotNil(t, variable.GetAST())
					assert.NotNil(t, variable.GetSrc())
					assert.NotNil(t, variable.ToProto())
				}

				for _, event := range contract.GetEvents() {
					assert.NotNil(t, event)
					assert.NotNil(t, event.GetAST())
					assert.NotNil(t, event.GetSrc())
					assert.NotNil(t, event.ToProto())
				}
			}
		})
	}
}

func TestIrBuilderFromJSON(t *testing.T) {
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
		wantErr              bool
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
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/Empty.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/Empty.ir.proto").Content,
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
				EntrySourceUnitName:  "SimpleStorage",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/SimpleStorage.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/SimpleStorage.ir.proto").Content,
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
				EntrySourceUnitName:  "ERC20",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/ERC20.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/ERC20.ir.proto").Content,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/TokenSale.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/TokenSale.ir.proto").Content,
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
				LocalSourcesPath:    "../sources/",
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/Lottery.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/Lottery.ir.proto").Content,
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
			expectedAst:          tests.ReadJsonBytesForTest(t, "ir/TransparentUpgradeableProxy.ir").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ir/TransparentUpgradeableProxy.ir.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:                 "Nil Sources",
			outputPath:           "ast/",
			unresolvedReferences: 0,
			wantErr:              true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := NewBuilderFromSources(context.TODO(), testCase.sources)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Nil(t, parser)
				return
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, parser)
			}
			assert.IsType(t, &Builder{}, parser)
			assert.IsType(t, &ast.ASTBuilder{}, parser.GetAstBuilder())
			assert.IsType(t, &solgo.Parser{}, parser.GetParser())
			assert.IsType(t, &solgo.Sources{}, parser.GetSources())

			// Important step which will parse the sources and build the AST including check for
			// reference errors and syntax errors.
			// If you wish to only parse the sources without checking for errors, use
			// parser.GetParser().Parse()
			assert.Empty(t, parser.Parse())

			// Now we can get into the business of building the intermediate representation
			assert.NoError(t, parser.Build())

			// ^ Just so we can get latest changes in the IR and AST prior importing AST from JSON
			astJson, err := parser.GetAstBuilder().ToJSON()
			assert.NoError(t, err)
			assert.NotNil(t, astJson)

			parser, err = NewBuilderFromJSON(context.TODO(), astJson)
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Now we can get into the business of building the intermediate representation
			assert.NoError(t, parser.Build())

			// Get the root node of the IR
			root := parser.GetRoot()
			assert.NotNil(t, root)

			pretty, err := parser.ToJSONPretty()
			assert.NoError(t, err)
			assert.NotNil(t, pretty)

			assert.NotNil(t, parser.ToProto())

			j, err := parser.ToJSON()
			assert.NotNil(t, j)
			assert.NoError(t, err)

			for _, eip := range root.GetStandards() {
				assert.NotNil(t, eip)
				assert.NotNil(t, eip.GetConfidence())
				assert.NotNil(t, eip.GetContractId())
				assert.NotNil(t, eip.GetContractName())
				assert.NotNil(t, eip.GetStandard())
				assert.NotNil(t, eip.ToProto())
			}

			assert.NotNil(t, root.HasStandard(standards.ERC1014))
			assert.NotNil(t, root.GetAST())

			for _, contract := range root.GetContracts() {
				assert.NotNil(t, contract)
				assert.NotNil(t, contract.GetAST())
				assert.NotNil(t, contract.GetSrc())
				assert.NotNil(t, contract.GetUnitSrc())
				assert.NotNil(t, contract.ToProto())

				if contract.GetFallback() != nil {
					assert.NotNil(t, contract.GetFallback().GetAST())
					assert.NotNil(t, contract.GetFallback().GetSrc())
					assert.NotNil(t, contract.GetFallback().ToProto())
				}

				if contract.GetReceive() != nil {
					assert.NotNil(t, contract.GetReceive().GetAST())
					assert.NotNil(t, contract.GetReceive().GetSrc())
					assert.NotNil(t, contract.GetReceive().ToProto())
				}

				if contract.GetConstructor() != nil {
					assert.NotNil(t, contract.GetConstructor().GetAST())
					assert.NotNil(t, contract.GetConstructor().GetSrc())
					assert.NotNil(t, contract.GetConstructor().ToProto())
				}

				for _, pragma := range contract.GetPragmas() {
					assert.NotNil(t, pragma)
					assert.NotNil(t, pragma.GetAST())
					assert.NotNil(t, pragma.GetSrc())
					assert.NotNil(t, pragma.ToProto())
				}

				for _, enum := range contract.GetEnums() {
					assert.NotNil(t, enum)
					assert.NotNil(t, enum.GetAST())
					assert.NotNil(t, enum.GetSrc())
					assert.NotNil(t, enum.ToProto())
				}

				for _, structs := range contract.GetStructs() {
					assert.NotNil(t, structs)
					assert.NotNil(t, structs.GetAST())
					assert.NotNil(t, structs.GetSrc())
					assert.NotNil(t, structs.ToProto())
				}

				for _, function := range contract.GetFunctions() {
					assert.NotNil(t, function)
					assert.NotNil(t, function.GetAST())
					assert.NotNil(t, function.GetSrc())
					assert.NotNil(t, function.GetBody().GetAST())
					assert.NotNil(t, function.GetBody().GetSrc())
					assert.NotNil(t, function.ToProto())

					if function.GetKind() == ast_pb.NodeType_KIND_FUNCTION {
						assert.NotEmpty(t, function.GetSignature())
					}

					for _, param := range function.GetParameters() {
						assert.NotNil(t, param)
						assert.NotNil(t, param.GetAST())
						assert.NotNil(t, param.GetSrc())
						assert.NotNil(t, param.ToProto())
					}

					for _, modifier := range function.GetModifiers() {
						assert.NotNil(t, modifier)
						assert.NotNil(t, modifier.GetAST())
						assert.NotNil(t, modifier.GetSrc())
						assert.NotNil(t, modifier.ToProto())
					}

					for _, override := range function.GetOverrides() {
						assert.NotNil(t, override)
						assert.NotNil(t, override.GetAST())
						assert.NotNil(t, override.GetSrc())
						assert.NotNil(t, override.ToProto())
					}

					for _, returns := range function.GetReturnStatements() {
						assert.NotNil(t, returns)
						assert.NotNil(t, returns.GetAST())
						assert.NotNil(t, returns.GetSrc())
						assert.NotNil(t, returns.ToProto())
					}

					for _, statement := range function.GetBody().GetStatements() {
						if node, ok := statement.(*FunctionCall); ok {
							assert.NotNil(t, node)
							assert.NotNil(t, node.GetAST())
							assert.NotNil(t, node.GetSrc())
							assert.NotNil(t, node.GetId())
							assert.Equal(t, len(node.GetNodes()), 0)
							assert.NotNil(t, node.ToProto())
							if node.GetExternalContract() != nil {
								assert.NotNil(t, node.GetExternalContract().GetSrc())
							}
							if node.GetReferenceStatement() != nil {
								assert.NotEmpty(t, node.GetReferenceStatement().GetSrc())
							}
						}
					}
				}

				for _, variable := range contract.GetStateVariables() {
					assert.NotNil(t, variable)
					assert.NotNil(t, variable.GetAST())
					assert.NotNil(t, variable.GetSrc())
					assert.NotNil(t, variable.ToProto())
				}

				for _, event := range contract.GetEvents() {
					assert.NotNil(t, event)
					assert.NotNil(t, event.GetAST())
					assert.NotNil(t, event.GetSrc())
					assert.NotNil(t, event.ToProto())
				}
			}
		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
