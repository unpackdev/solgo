package solc

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCompiler(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	compilerConfig, err := NewDefaultConfig()
	assert.NoError(t, err)
	assert.NotNil(t, compilerConfig)

	// Define multiple test cases
	testCases := []struct {
		name            string
		outputPath      string
		sources         *solgo.Sources
		wantErr         bool
		compilerVersion string
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
			wantErr:         false,
			compilerVersion: "0.8.0",
		},
		{
			name:       "Corrupted Contract Test",
			outputPath: "audits/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "CorruptedContract",
						Path:    tests.ReadContractFileForTest(t, "audits/CorruptedContract").Path,
						Content: tests.ReadContractFileForTest(t, "audits/CorruptedContract").Content,
					},
				},
				EntrySourceUnitName:  "CorruptedContract",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     buildFullPath("../sources/"),
			},
			wantErr:         true,
			compilerVersion: "0.8.0",
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
			compiler, err := NewCompiler(context.Background(), compilerConfig, testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, compiler)
			assert.NotNil(t, compiler.GetContext())
			assert.NotNil(t, compiler.GetSolc())
			assert.NotNil(t, compiler.GetSources())

			if testCase.compilerVersion != "" {
				compiler.SetCompilerVersion(testCase.compilerVersion)
				assert.Equal(t, testCase.compilerVersion, compiler.GetCompilerVersion())
			}

			compilerResults, err := compiler.Compile()
			if testCase.wantErr {
				if compilerResults != nil {
					assert.True(t, compilerResults.HasErrors())
					assert.False(t, compilerResults.HasWarnings())
					assert.GreaterOrEqual(t, len(compilerResults.GetWarnings()), 0)
					assert.GreaterOrEqual(t, len(compilerResults.GetErrors()), 1)
				}

				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, compilerResults)

			assert.NotEmpty(t, compilerResults.GetRequestedVersion())
			assert.NotEmpty(t, compilerResults.GetCompilerVersion())
			assert.NotEmpty(t, compilerResults.GetBytecode())
			assert.NotEmpty(t, compilerResults.GetABI())
			assert.NotEmpty(t, compilerResults.GetContractName())
			assert.GreaterOrEqual(t, len(compilerResults.GetWarnings()), 0)
			assert.GreaterOrEqual(t, len(compilerResults.GetErrors()), 0)
		})
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}
