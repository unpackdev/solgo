package ast

import (
	"testing"

	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/tests"
)

// getSourceTestCases returns the test cases (bunch of contracts) for the ast builder.
func getSourceTestCases(t *testing.T) []struct {
	name                 string
	outputPath           string
	sources              *solgo.Sources
	expectedAst          string
	expectedProto        string
	unresolvedReferences int64
} {
	return []struct {
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
			name:       "PAPA Contract",
			outputPath: "contracts/papa/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "ERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/papa/Token").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/papa/Token").Content,
					},
				},
				EntrySourceUnitName: "Token",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/papa/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/papa/Token.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "BabyToken Contract - 0xadd33a83549e115e3171c645b15a16ec6d1b5352",
			outputPath: "contracts/babytoken/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "ERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/babytoken/Token").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/babytoken/Token").Content,
					},
				},
				EntrySourceUnitName: "BABYTOKEN",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/babytoken/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/babytoken/Token.solgo.ast.proto").Content,
			unresolvedReferences: 0,
		},
		{
			name:       "Hello Contract - 0xCAaa580D02751e02Eb79b6f5b24B2417B118868f",
			outputPath: "contracts/hello/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Token",
						Path:    tests.ReadContractFileForTest(t, "contracts/hello/Token").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/hello/Token").Content,
					},
				},
				EntrySourceUnitName: "NFTradeNFTToken",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/hello/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/hello/Token.solgo.ast.proto").Content,
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
}
