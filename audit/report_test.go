package audit

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/0x19/solc-switch"
	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestResponse(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

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

	compiler, err := solc.New(context.TODO(), solcConfig)
	assert.NoError(t, err)
	assert.NotNil(t, compiler)

	// Make sure to sync the releases...
	err = compiler.Sync()
	assert.NoError(t, err)

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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			auditor, err := NewAuditor(context.TODO(), compiler, slitherConfig, testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, auditor)
			assert.True(t, auditor.IsReady())
			assert.IsType(t, testCase.sources, auditor.GetSources())
			assert.IsType(t, slitherConfig, auditor.GetConfig())
			assert.IsType(t, &Slither{}, auditor.GetSlither())

			resp, err := auditor.Analyze()
			assert.NoError(t, err)
			if testCase.wantErr {
				assert.NotEmpty(t, resp.Error)
				assert.False(t, resp.Success)
			}

			// Test FilterDetectorsByImpact function
			for _, impact := range []ImpactLevel{ImpactHigh, ImpactMedium, ImpactLow, ImpactInfo} {
				detectors := resp.FilterDetectorsByImpact(impact)
				for _, d := range detectors {
					assert.Equal(t, impact, ImpactLevel(d.Impact))
				}
			}

			// Test HasError function
			if resp.HasError() {
				assert.NotNil(t, resp.Error)
				assert.NotEmpty(t, resp.Error)
			}

			// Test ElementsByType function
			// Just an example, you might have other types to test
			elements := resp.ElementsByType("function")
			for _, e := range elements {
				assert.Equal(t, "function", e.Type)
			}

			// Test UniqueImpactLevels function
			uniqueImpacts := resp.UniqueImpactLevels()
			assert.LessOrEqual(t, len(uniqueImpacts), 4) // As we have 4 predefined impact levels

			// Test DetectorsByCheck function
			// Assuming "reentrancy" as a check type for demonstration
			detectors := resp.DetectorsByCheck("reentrancy")
			for _, d := range detectors {
				assert.Equal(t, "reentrancy", d.Check)
			}

			// Test CountByImpactLevel function
			countMap := resp.CountByImpactLevel()
			for impact, count := range countMap {
				assert.True(t, count >= 0)
				assert.Contains(t, []ImpactLevel{ImpactHigh, ImpactMedium, ImpactLow, ImpactInfo}, impact)
				assert.NotEmpty(t, impact.String())
			}

			// Test HighConfidenceDetectors function
			highConfDetectors := resp.HighConfidenceDetectors()
			for _, d := range highConfDetectors {
				assert.Equal(t, "High", d.Confidence)
			}

			// Test HasIssues function
			hasIssues := resp.HasIssues()
			assert.Equal(t, len(resp.Results.Detectors) > 0, hasIssues)
		})
	}
}
