package cfg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/tests"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCfgBuilder(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// Define multiple test cases
	testCases := []struct {
		name       string
		outputPath string
		sources    *solgo.Sources
		wantErr    bool
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
				EntrySourceUnitName:  "Empty",
				MaskLocalSourcesPath: false,
				LocalSourcesPath:     utils.GetLocalSourcesPath(),
			},
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
				EntrySourceUnitName:  "SimpleStorage",
				MaskLocalSourcesPath: true,
				LocalSourcesPath:     utils.GetLocalSourcesPath(),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser, err := ir.NewBuilderFromSources(context.TODO(), testCase.sources)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Nil(t, parser)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, parser)
			assert.IsType(t, &ir.Builder{}, parser)
			assert.IsType(t, &ast.ASTBuilder{}, parser.GetAstBuilder())
			assert.IsType(t, &solgo.Parser{}, parser.GetParser())
			assert.IsType(t, &solgo.Sources{}, parser.GetSources())

			// Important step which will parse the sources and build the AST including check for
			// reference errors and syntax errors.
			// If you wish to only parse the sources without checking for errors, use
			// parser.GetParser().Parse()
			assert.Empty(t, parser.Parse())

			// Now we can get into the business of building the intermediate representation
			assert.NoError(t, parser.Build())

			// Now we can get into the business of building the control flow graph
			builder, err := NewBuilder(context.Background(), parser)
			assert.NoError(t, err)
			assert.NotNil(t, builder)
			assert.IsType(t, &Builder{}, builder)

		})
	}
}
