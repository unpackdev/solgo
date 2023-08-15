package predictor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestPredictor(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name       string
		modelPath  string
		sourceCode string
		wantErr    bool
	}{
		{
			name:       "Reentrancy Contract Test",
			modelPath:  "path_to_model",
			sourceCode: tests.ReadContractFileForTest(t, "audits/VulnerableBank").Content,
			wantErr:    false,
		},
		{
			name:       "FooBar Contract Test",
			modelPath:  "path_to_model",
			sourceCode: tests.ReadContractFileForTest(t, "audits/FooBar").Content,
			wantErr:    true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			predictor, err := NewPredictor(testCase.modelPath)
			assert.NoError(t, err)
			assert.NotNil(t, predictor)

		})
	}
}
