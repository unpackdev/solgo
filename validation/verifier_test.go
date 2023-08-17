package validation

import (
	"context"
	"encoding/hex"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/solc"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestVerifier(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	compilerConfig, err := solc.NewDefaultConfig()
	assert.NoError(t, err)
	assert.NotNil(t, compilerConfig)

	// Define multiple test cases
	testCases := []struct {
		name               string
		outputPath         string
		sources            *solgo.Sources
		wantErr            bool
		wantBuildErr       bool
		wantNewVerifierErr bool
		wantDiff           bool
		diffCount          int
		wantNilResults     bool
		config             *solc.Config
		bytecode           []byte
	}{
		{
			name:               "No sources",
			outputPath:         "audits/",
			sources:            nil,
			wantErr:            true,
			wantDiff:           false,
			wantNewVerifierErr: true,
			config:             compilerConfig,
			bytecode:           []byte{0x01, 0x02, 0x03},
		},
		{
			name:       "No solc config",
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
			wantErr:            true,
			wantDiff:           false,
			wantNewVerifierErr: true,
			config:             nil,
			bytecode:           []byte{0x01, 0x02, 0x03},
		},
		{
			name:       "Sources should be corrupted and return error",
			outputPath: "audits/",
			sources: &solgo.Sources{
				EntrySourceUnitName:  "VulnerableBank",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			wantErr:      true,
			wantDiff:     false,
			wantBuildErr: true,
			config:       compilerConfig,
			bytecode:     []byte{0x01, 0x02, 0x03},
		},
		{
			name:       "Reentrancy Contract Test Bytecode Missmatch",
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
			wantErr:   true,
			wantDiff:  true,
			diffCount: 6,
			config:    compilerConfig,
			bytecode:  []byte{0x01, 0x02, 0x03},
		},
		{
			name:       "Reentrancy Contract Test Bytecode Match",
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
			bytecode: func() []byte {
				b, _ := hex.DecodeString("608060405234801561001057600080fd5b506105ca806100206000396000f3fe6080604052600436106100345760003560e01c806327e235e3146100395780633ccfd60b14610076578063d0e30db01461008d575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b91906102d8565b610097565b60405161006d9190610485565b60405180910390f35b34801561008257600080fd5b5061008b6100af565b005b610095610229565b005b60006020528060005260406000206000915090505481565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008111610135576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012c90610465565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405161015b90610410565b60006040518083038185875af1925050503d8060008114610198576040519150601f19603f3d011682016040523d82523d6000602084013e61019d565b606091505b50509050806101e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d890610445565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000341161026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390610425565b60405180910390fd5b346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102ba91906104bc565b92505081905550565b6000813590506102d28161057d565b92915050565b6000602082840312156102ea57600080fd5b60006102f8848285016102c3565b91505092915050565b600061030e6027836104ab565b91507f4465706f73697420616d6f756e742073686f756c64206265206772656174657260008301527f207468616e2030000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610374600f836104ab565b91507f5472616e73666572206661696c656400000000000000000000000000000000006000830152602082019050919050565b60006103b46014836104ab565b91507f496e73756666696369656e742062616c616e63650000000000000000000000006000830152602082019050919050565b60006103f46000836104a0565b9150600082019050919050565b61040a81610544565b82525050565b600061041b826103e7565b9150819050919050565b6000602082019050818103600083015261043e81610301565b9050919050565b6000602082019050818103600083015261045e81610367565b9050919050565b6000602082019050818103600083015261047e816103a7565b9050919050565b600060208201905061049a6000830184610401565b92915050565b600081905092915050565b600082825260208201905092915050565b60006104c782610544565b91506104d283610544565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156105075761050661054e565b5b828201905092915050565b600061051d82610524565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b61058681610512565b811461059157600080fd5b5056fea2646970667358221220e27ebb8a52fb9a56833f8bc098c1e940d6f786c2fdecc619def204c20ad5006664736f6c63430008000033")
				return b
			}(),
			config: compilerConfig,
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
			wantBuildErr:   false,
			wantErr:        true,
			wantNilResults: true,
			bytecode:       []byte{0x01, 0x02, 0x03},
			config:         compilerConfig,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			verifier, err := NewVerifier(context.Background(), testCase.config, testCase.sources)
			if testCase.wantBuildErr {
				assert.Error(t, err)
				assert.Nil(t, verifier)
				return
			}

			if testCase.wantNewVerifierErr {
				assert.Error(t, err)
				assert.Nil(t, verifier)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, verifier)
			assert.NotNil(t, verifier.GetContext())
			assert.NotNil(t, verifier.GetCompiler())
			assert.NotNil(t, verifier.GetSources())

			results, err := verifier.Verify(testCase.bytecode)
			if testCase.wantErr {
				assert.Error(t, err)
				if !testCase.wantNilResults {
					assert.NotNil(t, results)
					assert.False(t, results.IsVerified())
					assert.Equal(t, results.GetExpectedBytecode(), hex.EncodeToString(testCase.bytecode))
					assert.NotNil(t, results.GetCompilerResults())
					assert.NotEmpty(t, results.GetDiffPretty())
					if testCase.wantDiff {
						assert.Equal(t, len(results.GetDiffs()), testCase.diffCount)
					} else {
						assert.Equal(t, len(results.GetDiffs()), 0)
					}
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, results)
			assert.Equal(t, len(results.GetDiffs()), 0)
			assert.True(t, results.IsVerified())
			assert.NotNil(t, results.GetCompilerResults())
			assert.Empty(t, results.GetDiffPretty())
		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
