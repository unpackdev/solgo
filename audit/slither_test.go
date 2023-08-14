package audit

import (
	"context"
	"path/filepath"
	"testing"

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
	slitherConfig, err := NewDefaultConfig(
		buildFullPath("../data/tests/auditor-tests/"),
	)
	assert.NoError(t, err)

	// Define multiple test cases
	testCases := []struct {
		name                 string
		outputPath           string
		sources              *solgo.Sources
		expectedAbi          string
		expectedProto        string
		unresolvedReferences int64
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
			//expectedAbi:          tests.ReadJsonBytesForTest(t, "abi/Empty.abi").Content,
			//expectedProto:        tests.ReadJsonBytesForTest(t, "abi/Empty.abi.proto").Content,
			unresolvedReferences: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			slither, err := NewSlither(ctx, slitherConfig)
			assert.NoError(t, err)
			assert.NotNil(t, slither)

			version, err := slither.Version()
			assert.NoError(t, err)
			assert.NotEmpty(t, version)

			response, raw, err := slither.Analyze(testCase.sources)
			assert.NoError(t, err)
			assert.NotEmpty(t, raw)
			assert.NotNil(t, response)

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
