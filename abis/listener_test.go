package abis

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAbiListener(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name     string
		contract string
		expected string
	}{
		{
			name:     "Dummy Contract",
			contract: tests.ReadContractFileForTest(t, "Dummy").Content,
			expected: tests.ReadJsonBytesForTest(t, "Dummy").Content,
		},
		{
			name:     "Mappings",
			contract: tests.ReadContractFileForTest(t, "Mappings").Content,
			expected: tests.ReadJsonBytesForTest(t, "Mappings").Content,
		},
		{
			name:     "Structs",
			contract: tests.ReadContractFileForTest(t, "Structs").Content,
			expected: tests.ReadJsonBytesForTest(t, "Structs").Content,
		},
		{
			name:     "Enums",
			contract: tests.ReadContractFileForTest(t, "Enums").Content,
			expected: tests.ReadJsonBytesForTest(t, "Enums").Content,
		},
		{
			name:     "ERC20 Token",
			contract: tests.ReadContractFileForTest(t, "ERC20_Token").Content,
			expected: tests.ReadJsonBytesForTest(t, "ERC20_Token").Content,
		},
		{
			name:     "Complex Abi",
			contract: tests.ReadContractFileForTest(t, "AbiComplex").Content,
			expected: tests.ReadJsonBytesForTest(t, "AbiComplex").Content,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.NewParser(context.TODO(), strings.NewReader(testCase.contract))
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			// Register the contract information listener
			abiListener := NewAbiListener()
			err = parser.RegisterListener(solgo.ListenerAbi, abiListener)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			assert.Empty(t, syntaxErrs)

			abiParser := abiListener.GetParser()

			abiJson, err := abiParser.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, abiJson)

			abi, err := abiParser.ToABI()
			assert.NoError(t, err)
			assert.NotEmpty(t, abi)

			// Assert the parsed contract matches the expected result
			assert.Equal(t, testCase.expected, abiJson)
		})
	}
}
