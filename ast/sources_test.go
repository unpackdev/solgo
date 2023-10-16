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
	expectsErrors        bool
	disabled             bool
} {
	return []struct {
		name                 string
		outputPath           string
		sources              *solgo.Sources
		expectedAst          string
		expectedProto        string
		unresolvedReferences int64
		expectsErrors        bool
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
				EntrySourceUnitName: "Empty",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Empty.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
		},
		{
			name:       "L1Vault - 0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41",
			outputPath: "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "L1Vault",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1Vault").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1Vault").Content,
					},
					{
						Name:    "Strings",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Strings").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Strings").Content,
					},
					{
						Name:    "Math",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Math").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Math").Content,
					},
					{
						Name:    "AccessControlUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AccessControlUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AccessControlUpgradeable").Content,
					},
					{
						Name:    "IAccessControlUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IAccessControlUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IAccessControlUpgradeable").Content,
					},
					{
						Name:    "draft-IERC1822Upgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/draft-IERC1822Upgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/draft-IERC1822Upgradeable").Content,
					},
					{
						Name:    "ERC1967UpgradeUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC1967UpgradeUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC1967UpgradeUpgradeable").Content,
					},
					{
						Name:    "IBeaconUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IBeaconUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IBeaconUpgradeable").Content,
					},
					{
						Name:    "Initializable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Initializable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Initializable").Content,
					},
					{
						Name:    "UUPSUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/UUPSUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/UUPSUpgradeable").Content,
					},
					{
						Name:    "PausableUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/PausableUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/PausableUpgradeable").Content,
					},
					{
						Name:    "AddressUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AddressUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AddressUpgradeable").Content,
					},
					{
						Name:    "ContextUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ContextUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ContextUpgradeable").Content,
					},
					{
						Name:    "StorageSlotUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/StorageSlotUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/StorageSlotUpgradeable").Content,
					},
					{
						Name:    "StringsUpgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/StringsUpgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/StringsUpgradeable").Content,
					},
					{
						Name:    "ERC165Upgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC165Upgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC165Upgradeable").Content,
					},
					{
						Name:    "IERC165Upgradeable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IERC165Upgradeable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IERC165Upgradeable").Content,
					},
					{
						Name:    "Multicallable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Multicallable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Multicallable").Content,
					},
					{
						Name:    "ERC20",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC20").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/ERC20").Content,
					},
					{
						Name:    "SafeTransferLib",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/SafeTransferLib").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/SafeTransferLib").Content,
					},
					{
						Name:    "AffineGovernable",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AffineGovernable").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/AffineGovernable").Content,
					},
					{
						Name:    "BaseStrategy",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BaseStrategy").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BaseStrategy").Content,
					},
					{
						Name:    "BaseVault",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BaseVault").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BaseVault").Content,
					},
					{
						Name:    "BridgeEscrow",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BridgeEscrow").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/BridgeEscrow").Content,
					},
					{
						Name:    "WormholeRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/WormholeRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/WormholeRouter").Content,
					},
					{
						Name:    "L1BridgeEscrow",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1BridgeEscrow").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1BridgeEscrow").Content,
					},
					{
						Name:    "L1WormholeRouter",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1WormholeRouter").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1WormholeRouter").Content,
					},
					{
						Name:    "IRootChainManager",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IRootChainManager").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IRootChainManager").Content,
					},
					{
						Name:    "IWormhole",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IWormhole").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/IWormhole").Content,
					},
					{
						Name:    "Constants",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Constants").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Constants").Content,
					},
					{
						Name:    "Unchecked",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Unchecked").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/Unchecked").Content,
					},
				},
				EntrySourceUnitName: "L1Vault",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1Vault.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0x3b07A1A5de80f9b22DE0EC6C44C6E59DDc1C5f41/L1Vault.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "MintPassExtension - 0x7637a7E82e6af52ABeb27667489E110193D60b42",
			outputPath: "contracts/1:0x7637a7E82e6af52ABeb27667489E110193D60b42/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "MintPassExtension",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x7637a7E82e6af52ABeb27667489E110193D60b42/MintPassExtension").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x7637a7E82e6af52ABeb27667489E110193D60b42/MintPassExtension").Content,
					},
				},
				EntrySourceUnitName: "MintPassExtension",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0x7637a7E82e6af52ABeb27667489E110193D60b42/MintPassExtension.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0x7637a7E82e6af52ABeb27667489E110193D60b42/MintPassExtension.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
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
			disabled:             false,
		},
		{
			name:       "MultiSigWallet - 0x01b72e781732D92717DCEbE35a4243A10b9BaB44",
			outputPath: "contracts/1:0x01b72e781732D92717DCEbE35a4243A10b9BaB44/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "MultiSigWallet",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x01b72e781732D92717DCEbE35a4243A10b9BaB44/MultiSigWallet").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x01b72e781732D92717DCEbE35a4243A10b9BaB44/MultiSigWallet").Content,
					},
				},
				EntrySourceUnitName: "MultiSigWallet",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0x01b72e781732D92717DCEbE35a4243A10b9BaB44/MultiSigWallet.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0x01b72e781732D92717DCEbE35a4243A10b9BaB44/MultiSigWallet.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             true, // v0.4.26+commit.4563c3fc
		},
		{
			name:       "LockProxy - 0xf6378141BC900020a438F3914e4C3ceA29907b27",
			outputPath: "contracts/1:0xf6378141BC900020a438F3914e4C3ceA29907b27/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "LockProxy",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0xf6378141BC900020a438F3914e4C3ceA29907b27/LockProxy").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0xf6378141BC900020a438F3914e4C3ceA29907b27/LockProxy").Content,
					},
				},
				EntrySourceUnitName: "LockProxy",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0xf6378141BC900020a438F3914e4C3ceA29907b27/LockProxy.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0xf6378141BC900020a438F3914e4C3ceA29907b27/LockProxy.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             true, // v0.5.17+commit.d19bba13
		},
		{
			name:       "Payment - 0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed",
			outputPath: "contracts/1:0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Payment",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed/Payment").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed/Payment").Content,
					},
				},
				EntrySourceUnitName: "Payment",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed/Payment.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0x4892e397641530E7CCF1d07e94a5eAc68A2760Ed/Payment.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "FRAX3CRVLevSwap - 0xd747740FfAC8A6397bA80676299c4e3105999a9A",
			outputPath: "contracts/1:0xd747740FfAC8A6397bA80676299c4e3105999a9A/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "FRAX3CRVLevSwap",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0xd747740FfAC8A6397bA80676299c4e3105999a9A/FRAX3CRVLevSwap").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0xd747740FfAC8A6397bA80676299c4e3105999a9A/FRAX3CRVLevSwap").Content,
					},
				},
				EntrySourceUnitName: "FRAX3CRVLevSwap",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0xd747740FfAC8A6397bA80676299c4e3105999a9A/FRAX3CRVLevSwap.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0xd747740FfAC8A6397bA80676299c4e3105999a9A/FRAX3CRVLevSwap.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "Qatar_Ecuador - 0x275659c6e77f9c5f6d3fc93adb388017d00500a7",
			outputPath: "contracts/1:0x275659c6e77f9c5f6d3fc93adb388017d00500a7/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Qatar_Ecuador",
						Path:    tests.ReadContractFileForTest(t, "contracts/1:0x275659c6e77f9c5f6d3fc93adb388017d00500a7/Qatar_Ecuador").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/1:0x275659c6e77f9c5f6d3fc93adb388017d00500a7/Qatar_Ecuador").Content,
					},
				},
				EntrySourceUnitName: "Qatar_Ecuador",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/1:0x275659c6e77f9c5f6d3fc93adb388017d00500a7/Qatar_Ecuador.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/1:0x275659c6e77f9c5f6d3fc93adb388017d00500a7/Qatar_Ecuador.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
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
			disabled:             false,
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
			disabled:             false,
		},
		{
			name:       "BlockchainLottery Contract - 0xb334015E66d7203d8891Ce5E78eAeEFB6B3aE392",
			outputPath: "contracts/blottery/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Lottery",
						Path:    tests.ReadContractFileForTest(t, "contracts/blottery/Lottery").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/blottery/Lottery").Content,
					},
				},
				EntrySourceUnitName: "BlockchainLottery",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/blottery/Lottery.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/blottery/Lottery.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
		},
		{
			name:       "SeaGod Contract - 0xb334015E66d7203d8891Ce5E78eAeEFB6B3aE392",
			outputPath: "contracts/seagod/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Seagod",
						Path:    tests.ReadContractFileForTest(t, "contracts/seagod/Seagod").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/seagod/Seagod").Content,
					},
				},
				EntrySourceUnitName: "SeaGod",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/seagod/Seagod.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/seagod/Seagod.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        true,
			disabled:             false,
		},
		{
			name:       "RickRolledToken Contract - 0xe81473ff76c60dafb526ca037d0a1bc282d42a4d",
			outputPath: "contracts/rick/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Token",
						Path:    tests.ReadContractFileForTest(t, "contracts/rick/Token").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/rick/Token").Content,
					},
				},
				EntrySourceUnitName: "RickRolledToken",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/rick/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/rick/Token.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "PTM Contract - 0x9A4e2AB29f9edE0c362f82F873F9d727810480F2",
			outputPath: "contracts/ptm/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Ptm",
						Path:    tests.ReadContractFileForTest(t, "contracts/ptm/Ptm").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/ptm/Ptm").Content,
					},
				},
				EntrySourceUnitName: "PTM",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/ptm/Ptm.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/ptm/Ptm.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "KnoxLpLocker Contract - 0x09D10fbcEbd414DE4683856fF72a3587761A1587",
			outputPath: "contracts/knox/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "KnoxLpLocker",
						Path:    tests.ReadContractFileForTest(t, "contracts/knox/Knox").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/knox/Knox").Content,
					},
				},
				EntrySourceUnitName: "KnoxLpLocker",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/knox/Knox.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/knox/Knox.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "AdminUpgradeabilityProxy Contract - 0xA567D9B111b570cc5b68eDef188056FFfD1e2813",
			outputPath: "contracts/adminproxy/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "AdminUpgradeabilityProxy",
						Path:    tests.ReadContractFileForTest(t, "contracts/adminproxy/Admin").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/adminproxy/Admin").Content,
					},
				},
				EntrySourceUnitName: "AdminUpgradeabilityProxy",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/adminproxy/Admin.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/adminproxy/Admin.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        true,
			disabled:             true, // 0.5.10 contract, has some issues with the parser
		},
		{
			name:       "RouterV2 Contract - 0x8A99Ad90E77e376AD4Cec21231AF855a87771fD0",
			outputPath: "contracts/router/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "RouterV2",
						Path:    tests.ReadContractFileForTest(t, "contracts/router/RouterV2").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/router/RouterV2").Content,
					},
				},
				EntrySourceUnitName: "RouterV2",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/router/RouterV2.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/router/RouterV2.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "ItemsMarketplace Contract - 0x7C3a812bBfC759bf85097211253e63f9e5F49439",
			outputPath: "contracts/0x7C3a812bBfC759bf85097211253e63f9e5F49439/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Contract",
						Path:    tests.ReadContractFileForTest(t, "contracts/0x7C3a812bBfC759bf85097211253e63f9e5F49439/Contract").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/0x7C3a812bBfC759bf85097211253e63f9e5F49439/Contract").Content,
					},
				},
				EntrySourceUnitName: "ItemsMarketplace",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/0x7C3a812bBfC759bf85097211253e63f9e5F49439/Contract.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/0x7C3a812bBfC759bf85097211253e63f9e5F49439/Contract.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "CRAB  Contract - 0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2",
			outputPath: "contracts/0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "CRAB",
						Path:    tests.ReadContractFileForTest(t, "contracts/0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2/Contract").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2/Contract").Content,
					},
				},
				EntrySourceUnitName: "CRAB",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2/CRAB.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/0x9ba77c0489c0a2D16F0C8314189acDA4d3af8Aa2/CRAB.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "Pledge  Contract - 0x16Ca8d09D693201d54a2882c05B8421102fc00B2",
			outputPath: "contracts/0x16Ca8d09D693201d54a2882c05B8421102fc00B2/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Pledge",
						Path:    tests.ReadContractFileForTest(t, "contracts/0x16Ca8d09D693201d54a2882c05B8421102fc00B2/Contract").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/0x16Ca8d09D693201d54a2882c05B8421102fc00B2/Contract").Content,
					},
				},
				EntrySourceUnitName: "Pledge",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/0x16Ca8d09D693201d54a2882c05B8421102fc00B2/Pledge.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/0x16Ca8d09D693201d54a2882c05B8421102fc00B2/Pledge.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
			disabled:             false,
		},
		{
			name:       "VirtualX - 0xe301C9525Ade8c368329a055212Fd56b202c1E3C",
			outputPath: "contracts/56:0xe301C9525Ade8c368329a055212Fd56b202c1E3C/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "VirtualX",
						Path:    tests.ReadContractFileForTest(t, "contracts/56:0xe301C9525Ade8c368329a055212Fd56b202c1E3C/VirtualX").Path,
						Content: tests.ReadContractFileForTest(t, "contracts/56:0xe301C9525Ade8c368329a055212Fd56b202c1E3C/VirtualX").Content,
					},
				},
				EntrySourceUnitName: "VirtualX",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/56:0xe301C9525Ade8c368329a055212Fd56b202c1E3C/VirtualX.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/56:0xe301C9525Ade8c368329a055212Fd56b202c1E3C/VirtualX.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			expectsErrors:        false,
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
				EntrySourceUnitName: "SimpleStorage",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/SimpleStorage.solgo.ast.proto").Content,
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
				EntrySourceUnitName: "ERC20",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/ERC20.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/ERC20.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
				LocalSources:        true,
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Token.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Token.solgo.ast.proto").Content,
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
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/TokenSale.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/TokenSale.solgo.ast.proto").Content,
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
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			expectedAst:          tests.ReadJsonBytesForTest(t, "ast/Lottery.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "ast/Lottery.solgo.ast.proto").Content,
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
			expectedAst:          tests.ReadJsonBytesForTest(t, "contracts/cheelee/TransparentUpgradeableProxy.solgo.ast").Content,
			expectedProto:        tests.ReadJsonBytesForTest(t, "contracts/cheelee/TransparentUpgradeableProxy.solgo.ast.proto").Content,
			unresolvedReferences: 0,
			disabled:             false,
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
			disabled:             false,
		},
	}
}
