package audit

import (
	"context"
	"os"
	"testing"

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
	slitherConfig, err := NewDefaultConfig(os.TempDir())
	assert.NoError(t, err)

	// Define multiple test cases
	testCases := []struct {
		name       string
		outputPath string
		sources    *solgo.Sources
		wantErr    bool
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
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			auditor, err := NewAuditor(context.Background(), slitherConfig, testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, auditor)
			assert.True(t, auditor.IsReady())
			assert.IsType(t, testCase.sources, auditor.GetSources())
			assert.IsType(t, slitherConfig, auditor.GetConfig())
			assert.IsType(t, &Slither{}, auditor.GetSlither())

			response, err := auditor.Analyze()
			assert.NoError(t, err)
			if testCase.wantErr {
				assert.NotEmpty(t, response.Error)
				assert.False(t, response.Success)
			}
		})
	}
}
