package ast

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestAstBuilderFromSourceAsString(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// All of the defined test cases can be discovered in sources_test.go file
	testCases := getSourceTestCases(t)

	for _, testCase := range testCases {
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
			assert.Empty(t, syntaxErrs)

			// This step is actually quite important as it resolves all the
			// references in the AST. Without this step, the AST will be
			// incomplete.
			errs := astBuilder.ResolveReferences()
			var errsExpected []error
			assert.Equal(t, errsExpected, errs)
			assert.Equal(t, int(testCase.unresolvedReferences), astBuilder.GetResolver().GetUnprocessedCount())
			assert.Equal(t, len(astBuilder.GetResolver().GetUnprocessedNodes()), astBuilder.GetResolver().GetUnprocessedCount())
			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				prettyJson, err := utils.ToJSONPretty(sourceUnit)
				assert.NoError(t, err)
				assert.NotEmpty(t, prettyJson)

				err = utils.WriteToFile(
					"../data/tests/"+testCase.outputPath+sourceUnit.GetName()+".solgo.ast.json",
					prettyJson,
				)
				assert.NoError(t, err)
			}

			prettyJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, prettyJson)
			err = utils.WriteToFile(
				"../data/tests/"+testCase.outputPath+testCase.sources.EntrySourceUnitName+".solgo.ast.json",
				prettyJson,
			)
			assert.NoError(t, err)
			//assert.Equal(t, testCase.expectedAst, string(prettyJson))

			astJson, err := astBuilder.ToJSON()
			assert.NoError(t, err)
			assert.NotEmpty(t, astJson)

			astPretty, _ := utils.ToJSONPretty(astBuilder.ToProto())
			err = utils.WriteToFile(
				"../data/tests/"+testCase.outputPath+testCase.sources.EntrySourceUnitName+".solgo.ast.proto.json",
				astPretty,
			)
			assert.NoError(t, err)
			assert.NotEmpty(t, astPretty)
			//assert.Equal(t, testCase.expectedProto, string(astPretty))

			// Basic check for root
			astRoot := astBuilder.GetRoot()

			assert.NotNil(t, astRoot.GetSrc())
			assert.NotNil(t, astRoot.GetTypeDescription())
			assert.NotNil(t, astRoot.GetComments())
			assert.GreaterOrEqual(t, astRoot.GetSourceUnitCount(), int32(0))

			// Zero is here for the first contract that's empty...
			assert.GreaterOrEqual(t, astRoot.EntrySourceUnit, int64(0))

			// We need to check that the entry source unit name is correct.
			for _, sourceUnit := range astRoot.GetSourceUnits() {
				if astRoot.GetEntrySourceUnit() == sourceUnit.GetId() {
					assert.Equal(t, sourceUnit.GetName(), testCase.sources.EntrySourceUnitName)
				}

				assert.True(t, astRoot.HasSourceUnits())
				assert.NotNil(t, astRoot.GetSourceUnitByName(sourceUnit.GetName()))
				assert.NotNil(t, astRoot.GetSourceUnitById(sourceUnit.GetId()))

				assert.GreaterOrEqual(t, len(sourceUnit.GetImports()), 0)
				assert.GreaterOrEqual(t, len(sourceUnit.GetPragmas()), 0)
				assert.NotNil(t, sourceUnit.GetContract())
				assert.NotEmpty(t, sourceUnit.GetLicense())

				// Recursive test against all nodes. A common place where we can add tests to check
				// if the AST is correct.
				recursiveTest(t, sourceUnit)
			}

		})
	}
}

func recursiveTest(t *testing.T, node Node[NodeType]) {
	assert.NotNil(t, node.GetNodes(), fmt.Sprintf("Node %T has nil nodes", node))
	assert.GreaterOrEqual(t, node.GetId(), int64(0), fmt.Sprintf("Node %T has empty id", node))
	assert.NotNil(t, node.GetType(), fmt.Sprintf("Node %T has empty type", node))
	assert.NotNil(t, node.GetSrc(), fmt.Sprintf("Node %T has empty GetSrc()", node))
	assert.NotNil(t, node.GetTypeDescription(), fmt.Sprintf("Node %T has not defined GetTypeDescription()", node))

	if contract, ok := node.(*Contract); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)
		assert.NotNil(t, contract.IsAbstract())
		assert.NotNil(t, contract.GetKind())
		assert.NotNil(t, contract.IsFullyImplemented())

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}

	}

	if contract, ok := node.(*Library); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}
	}

	if contract, ok := node.(*Interface); ok {
		assert.GreaterOrEqual(t, len(contract.GetBaseContracts()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStateVariables()), 0)
		assert.GreaterOrEqual(t, len(contract.GetStructs()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEnums()), 0)
		assert.GreaterOrEqual(t, len(contract.GetErrors()), 0)
		assert.GreaterOrEqual(t, len(contract.GetEvents()), 0)
		assert.GreaterOrEqual(t, len(contract.GetFunctions()), 0)
		assert.GreaterOrEqual(t, len(contract.GetContractDependencies()), 0)
		assert.GreaterOrEqual(t, len(contract.GetLinearizedBaseContracts()), 0)

		if contract.GetConstructor() != nil {
			assert.NotNil(t, contract.GetConstructor().GetSrc())
		}

		if contract.GetReceive() != nil {
			assert.NotNil(t, contract.GetReceive().GetSrc())
		}

		if contract.GetFallback() != nil {
			assert.NotNil(t, contract.GetFallback().GetSrc())
		}

		for _, base := range contract.GetBaseContracts() {
			assert.GreaterOrEqual(t, base.GetId(), int64(0))
			assert.NotNil(t, base.GetType())
			assert.NotNil(t, base.GetSrc())
		}
	}

	if while, ok := node.(*WhileStatement); ok {
		assert.NotNil(t, while.GetCondition())
		assert.NotNil(t, while.GetBody())
	}

	if variable, ok := node.(*StateVariableDeclaration); ok {
		assert.GreaterOrEqual(t, variable.GetReferencedDeclaration(), int64(0))
	}

	if variable, ok := node.(*VariableDeclaration); ok {
		assert.GreaterOrEqual(t, len(variable.GetAssignments()), 0)
	}

	if using, ok := node.(*UsingDirective); ok {
		if using.GetPathNode() != nil {
			assert.NotEmpty(t, using.GetPathNode().Name)
		}
		assert.GreaterOrEqual(t, using.GetReferencedDeclaration(), int64(0))
		assert.NotNil(t, using.GetLibraryName())
	}

	if typeName, ok := node.(*TypeName); ok {
		assert.GreaterOrEqual(t, typeName.GetReferencedDeclaration(), int64(0))
		assert.NotNil(t, typeName.GetStateMutability())
		assert.NotNil(t, typeName.GetTypeDescription().GetIdentifier())
		assert.NotNil(t, typeName.GetTypeDescription().GetString())
	}

	if try, ok := node.(*TryStatement); ok {
		assert.NotNil(t, try.GetImplemented())
	}

	if strct, ok := node.(*StructDefinition); ok {
		assert.NotEmpty(t, strct.GetSourceUnitName())
	}

	if revert, ok := node.(*RevertStatement); ok {
		assert.GreaterOrEqual(t, len(revert.GetArguments()), 0)
		if revert.GetExpression() != nil {
			assert.NotNil(t, revert.GetExpression().GetSrc())
		}
	}

	if enum, ok := node.(*EnumDefinition); ok {
		assert.NotEmpty(t, enum.GetSourceUnitName())
	}

	if errNode, ok := node.(*ErrorDefinition); ok {
		assert.NotEmpty(t, errNode.GetSourceUnitName())
	}

	for _, childNode := range node.GetNodes() {
		recursiveTest(t, childNode)
	}
}

func buildFullPath(relativePath string) string {
	absPath, _ := filepath.Abs(relativePath)
	return absPath
}

func TestAstReferenceSetDescriptor(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	assert.NoError(t, err)

	// Replace the global logger.
	zap.ReplaceGlobals(logger)

	// All of the defined test cases can be discovered in sources_test.go file
	testCases := getSourceTestCases(t)

	for _, testCase := range testCases {
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
			assert.Empty(t, syntaxErrs)

			assert.NotNil(t, astBuilder.GetParser())
			assert.NotNil(t, astBuilder.GetTree())

			// This step is actually quite important as it resolves all the
			// references in the AST. Without this step, the AST will be
			// incomplete.
			errs := astBuilder.ResolveReferences()
			var errsExpected []error
			assert.Equal(t, errsExpected, errs)
			assert.Equal(t, int(testCase.unresolvedReferences), astBuilder.GetResolver().GetUnprocessedCount())

			astBuilder.GetRoot().SetReferenceDescriptor(0, nil)
			astBuilder.GetRoot().SetReferenceDescriptor(0, &TypeDescription{})

			// We need to check that the entry source unit name is correct.
			for _, sourceUnit := range astBuilder.GetRoot().GetSourceUnits() {
				recursiveReferenceDescriptorSetTest(t, sourceUnit)
			}

		})
	}
}

// recursiveReferenceDescriptorSetTest is a recursive test that checks if all the reference descriptors
// functions exist and won't return any panics of some sort... It's just basic test.
func recursiveReferenceDescriptorSetTest(t *testing.T, node Node[NodeType]) {
	node.SetReferenceDescriptor(0, nil)
	node.SetReferenceDescriptor(0, &TypeDescription{})

	for _, childNode := range node.GetNodes() {
		childNode.SetReferenceDescriptor(0, nil)
		childNode.SetReferenceDescriptor(0, &TypeDescription{})
		recursiveReferenceDescriptorSetTest(t, childNode)
	}
}
