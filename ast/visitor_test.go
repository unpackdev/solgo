package ast

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

var (
	matchedCallDataErr = errors.New("matched_call_data")
)

func TestAsvVisitor(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	testCases := []struct {
		name                 string
		outputPath           string
		sources              *solgo.Sources
		unresolvedReferences int64
		visitFn              func(node Node[NodeType]) (bool, error)
		expectsErrors        bool
		disabled             bool
	}{
		{
			name:       "Function Node Visitor Test",
			outputPath: "ast/",
			sources: &solgo.Sources{
				SourceUnits: []*solgo.SourceUnit{
					{
						Name:    "Empty",
						Path:    tests.ReadContractFileForTest(t, "AbiComplex").Path,
						Content: tests.ReadContractFileForTest(t, "AbiComplex").Content,
					},
				},
				EntrySourceUnitName: "Empty",
				LocalSourcesPath:    buildFullPath("../sources/"),
			},
			unresolvedReferences: 0,
			visitFn: func(node Node[NodeType]) (bool, error) {
				//t.Logf("Visiting node id:%d - type:%s - actual_type: %T", node.GetId(), node.GetType(), node)

				var memoryParams []*Parameter

				// Seek only for a function type... It can be a different type such as
				// constructor, fallback, receive...
				if nodeCtx, ok := node.(*Function); ok {
					if nodeCtx.GetName() == "sub" {
						for _, parameter := range nodeCtx.GetParameters().Parameters {
							if parameter.StorageLocation == ast_pb.StorageLocation_MEMORY {
								memoryParams = append(memoryParams, parameter)
							}
						}

						if len(memoryParams) > 0 {
							for _, childCtx := range nodeCtx.GetNodes() {
								if fnCallCtx, fnOk := childCtx.(*FunctionCall); fnOk {
									for _, arg := range fnCallCtx.GetArguments() {
										if argCtx, ok := arg.(*PrimaryExpression); ok {
											for _, mParam := range memoryParams {
												if mParam.GetName() == argCtx.GetName() {

													// Now just for the sake of example, return false so, we stop iteration...
													// Error here is a trick to return test as successful
													return false, matchedCallDataErr
												}
											}
										}
									}
								}
							}
						}
					}
				}

				return true, nil
			},
			disabled: false,
		},
	}

	for _, testCase := range testCases {
		if testCase.disabled {
			continue
		}

		t.Run(testCase.name, func(t *testing.T) {
			parser, err := solgo.NewParserFromSources(context.TODO(), testCase.sources)
			assert.NoError(t, err)
			assert.NotNil(t, parser)

			astBuilder := NewAstBuilder(
				// We need to provide parser to the ast builder so that it can
				// access comments and other information from the parser.
				parser.GetParser(),

				// We need to provide sources to the ast builder so that it can
				// access the source code of the contracts.
				parser.GetSources(),
			)

			err = parser.RegisterListener(solgo.ListenerAst, astBuilder)
			assert.NoError(t, err)

			syntaxErrs := parser.Parse()
			if !testCase.expectsErrors {
				assert.Empty(t, syntaxErrs)
			}

			// This step is actually quite important as it resolves all the
			// references in the AST. Without this step, the AST will be
			// incomplete.
			errs := astBuilder.ResolveReferences()
			if testCase.expectsErrors {
				var errsExpected []error
				assert.Equal(t, errsExpected, errs)
			}
			assert.Equal(t, int(testCase.unresolvedReferences), astBuilder.GetResolver().GetUnprocessedCount())
			assert.Equal(t, len(astBuilder.GetResolver().GetUnprocessedNodes()), astBuilder.GetResolver().GetUnprocessedCount())

			visitStatus, visitErr := astBuilder.GetTree().ExecuteTypeVisit(
				ast_pb.NodeType_FUNCTION_DEFINITION,
				testCase.visitFn,
			)
			assert.False(t, visitStatus)
			assert.ErrorIs(t, visitErr, matchedCallDataErr)
		})
	}
}
