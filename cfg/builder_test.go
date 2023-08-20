package cfg

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/goccy/go-graphviz"
	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
	"github.com/txpull/solgo/tests"
	"github.com/txpull/solgo/utils"
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
			builder := NewBuilder(context.Background(), parser)
			assert.NotNil(t, builder)
			assert.IsType(t, &Builder{}, builder)
			assert.NotNil(t, builder.GetGraphviz())

			defer builder.Close()

			graph, err := builder.Build()
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Nil(t, graph)
				return
			}

			// Close the graph to free up resources.
			// It is intentionally not closed in case that the graph is needed for further
			// processing.
			defer graph.Close()

			assert.NoError(t, err)
			assert.NotNil(t, graph)

			outputPath := filepath.Join("..", "data", "tests", "cfg", testCase.sources.EntrySourceUnitName)

			// Save the graph to a file as a png file.
			err = builder.SaveAs(graph, graphviz.PNG, outputPath+".png")
			assert.NoError(t, err)

			// Save the graph to a file as a dot file.
			dot, err := builder.GenerateDOT(graph)
			assert.NoError(t, err)
			assert.NotEmpty(t, dot)

			err = utils.WriteToFile(outputPath+".dot", []byte(dot))
			assert.NoError(t, err)

		})

	}

}

func TestNewBuilder(t *testing.T) {
	parser := &ir.Builder{}
	ctx := context.Background()

	builder := NewBuilder(ctx, parser)
	assert.NotNil(t, builder)
	assert.Equal(t, ctx, builder.ctx)
	assert.Equal(t, parser, builder.builder)
	assert.NotNil(t, builder.viz)
}

func TestClose(t *testing.T) {
	parser := &ir.Builder{}
	builder := NewBuilder(context.Background(), parser)

	err := builder.Close()
	assert.NoError(t, err)
}

func TestGetGraphviz(t *testing.T) {
	parser := &ir.Builder{}
	builder := NewBuilder(context.Background(), parser)
	viz := builder.GetGraphviz()
	assert.NotNil(t, viz)
	assert.IsType(t, &graphviz.Graphviz{}, viz)
}

func TestBuild(t *testing.T) {
	parser := &ir.Builder{}
	builder := NewBuilder(context.Background(), parser)
	graph, err := builder.Build()
	assert.Nil(t, graph)
	assert.Error(t, err)
	assert.Equal(t, "root node is not set", err.Error())
	builder.viz = nil
	graph, err = builder.Build()
	assert.Nil(t, graph)
	assert.Error(t, err)
	assert.Equal(t, "graphviz instance is not set", err.Error())
}

func TestGenerateDOT(t *testing.T) {
	parser := &ir.Builder{}
	builder := NewBuilder(context.Background(), parser)
	dot, err := builder.GenerateDOT(nil)
	assert.Empty(t, dot)
	assert.Error(t, err)
	assert.Equal(t, "graph is not set", err.Error())

	builder.viz = nil
	dot, err = builder.GenerateDOT(nil)
	assert.Empty(t, dot)
	assert.Error(t, err)
	assert.Equal(t, "graphviz instance is not set", err.Error())
}

func TestSaveAs(t *testing.T) {
	parser := &ir.Builder{}
	builder := NewBuilder(context.Background(), parser)
	err := builder.SaveAs(nil, graphviz.PNG, "test.png")
	assert.Error(t, err)
	assert.Equal(t, "graph is not set", err.Error())

	builder.viz = nil
	err = builder.SaveAs(nil, graphviz.PNG, "test.png")
	assert.Error(t, err)
	assert.Equal(t, "graphviz instance is not set", err.Error())
}
