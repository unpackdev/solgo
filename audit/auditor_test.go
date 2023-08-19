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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAuditor(t *testing.T) {
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
	// Default path for solc will be used by automatically understanding what is the solc version from
	// pragmas in the main solidity contract file.
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
		config      *Config
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
			wantErr:     false,
			wantSolcErr: false,
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			auditor, err := NewAuditor(context.Background(), solc, slitherConfig, testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, auditor)
			assert.True(t, auditor.IsReady())
			assert.IsType(t, testCase.sources, auditor.GetSources())
			assert.IsType(t, slitherConfig, auditor.GetConfig())
			assert.IsType(t, &Slither{}, auditor.GetSlither())
			report, err := auditor.Analyze()
			if testCase.wantSolcErr {
				assert.Error(t, err)
				assert.Nil(t, report)
				return
			}

			assert.NoError(t, err)

			if testCase.wantErr {
				assert.NotEmpty(t, report.Error)
				assert.False(t, report.Success)
			} else {
				assert.NotNil(t, report.ToProto())
			}
		})
	}
}
