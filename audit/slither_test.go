package audit

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/0x19/solc-switch"
	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/tests"
	"github.com/txpull/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSlither(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Global configuration for the slither as we'd need to define it only once for
	// this particular test suite.
	// Default configuration accepts temporary path so it can be tweaked as you wish.
	// There are no defaults and this parameter is necessary to be set!
	slitherConfig, err := NewDefaultConfig(os.TempDir())
	assert.NoError(t, err)

	solcConfig, err := solc.NewDefaultConfig()
	assert.NoError(t, err)
	assert.NotNil(t, solcConfig)

	// Preparation of solc repository. In the tests, this is required as we need to due to CI/CD permissions
	// have ability to set the releases path to the local repository.
	// @TODO: in the future investigate permissions between different go modules.
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	assert.NotEmpty(t, cwd)

	releasesPath := filepath.Join(cwd, "..", "data", "solc", "releases")
	err = solcConfig.SetReleasesPath(releasesPath)
	assert.NoError(t, err)

	solc, err := solc.New(context.TODO(), solcConfig)
	assert.NoError(t, err)
	assert.NotNil(t, solc)

	// Define multiple test cases
	testCases := []struct {
		name        string
		outputPath  string
		sources     *solgo.Sources
		wantErr     bool
		wantSolcErr bool
	}{
		{
			name:       "Reentrancy Contract Test",
			outputPath: "audits/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "VulnerableBank",
						Path:    tests.ReadContractFileForTest(t, "audits/VulnerableBank").Path,
						Content: tests.ReadContractFileForTest(t, "audits/VulnerableBank").Content,
					},
				},
				EntrySourceUnitName:  "VulnerableBank",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			wantErr: false,
		},
		{
			name:       "FooBar Contract Test",
			outputPath: "audits/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "FooBar",
						Path:    tests.ReadContractFileForTest(t, "audits/FooBar").Path,
						Content: tests.ReadContractFileForTest(t, "audits/FooBar").Content,
					},
				},
				EntrySourceUnitName:  "FooBar",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			wantErr:     true,
			wantSolcErr: true,
		},
		{
			name:       "Empty Contract Test",
			outputPath: "audits/",
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
			wantSolcErr: true,
		},
		{
			name:       "Simple Storage Contract Test",
			outputPath: "audits/",
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
		},
		{
			name:       "OpenZeppelin ERC20 Test",
			outputPath: "audits/",
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
		},
		{
			name:       "Token Sale ERC20 Test",
			outputPath: "audits/",
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
		},
		{
			name:       "Lottery Test",
			outputPath: "audits/",
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
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			slither, err := NewSlither(ctx, solc, slitherConfig)
			assert.NoError(t, err)
			assert.NotNil(t, slither)

			assert.True(t, slither.IsInstalled())

			version, err := slither.Version()
			assert.NoError(t, err)
			assert.NotEmpty(t, version)

			response, raw, err := slither.Analyze(testCase.sources)
			if testCase.wantSolcErr {
				assert.Error(t, err)
				assert.Empty(t, raw)
				assert.Nil(t, response)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, raw)
			assert.NotNil(t, response)

			if testCase.wantErr {
				assert.NotEmpty(t, response.Error)
				assert.False(t, response.Success)
			}

			err = utils.WriteToFile(
				"../data/tests/audits/"+testCase.sources.EntrySourceUnitName+".slither.raw.json",
				raw,
			)
			assert.NoError(t, err)
		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
