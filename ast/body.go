package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

func (b *ASTBuilder) parseBodyElement(sourceUnit *ast_pb.SourceUnit, identifierNode *ast_pb.Node, bodyElement parser.IContractBodyElementContext) *ast_pb.Node {
	id := atomic.AddInt64(&b.nextID, 1) - 1
	toReturn := &ast_pb.Node{
		Id: id,
		Src: &ast_pb.Src{
			Line:        int64(bodyElement.GetStart().GetLine()),
			Start:       int64(bodyElement.GetStart().GetStart()),
			End:         int64(bodyElement.GetStop().GetStop()),
			Length:      int64(bodyElement.GetStop().GetStop() - bodyElement.GetStart().GetStart() + 1),
			ParentIndex: identifierNode.Id,
		},
	}

	if usingForDeclaration := bodyElement.UsingDirective(); usingForDeclaration != nil {
		toReturn = b.parseUsingForDeclaration(
			sourceUnit,
			toReturn,
			usingForDeclaration.(*parser.UsingDirectiveContext),
		)
	} else if stateVariableDeclaration := bodyElement.StateVariableDeclaration(); stateVariableDeclaration != nil {
		toReturn = b.parseStateVariableDeclaration(
			sourceUnit,
			toReturn,
			stateVariableDeclaration.(*parser.StateVariableDeclarationContext),
		)
	} else if eventDefinition := bodyElement.EventDefinition(); eventDefinition != nil {
		toReturn = b.parseEventDefinition(
			sourceUnit,
			toReturn,
			eventDefinition.(*parser.EventDefinitionContext),
		)
	} else if functionDefinition := bodyElement.FunctionDefinition(); functionDefinition != nil {
		toReturn = b.parseFunctionDefinition(
			sourceUnit,
			toReturn,
			functionDefinition.(*parser.FunctionDefinitionContext),
		)
	} else if constructorDefinition := bodyElement.ConstructorDefinition(); constructorDefinition != nil {
		toReturn = b.parseConstructorDefinition(
			sourceUnit,
			toReturn,
			constructorDefinition.(*parser.ConstructorDefinitionContext),
		)
	} else if enumDefinition := bodyElement.EnumDefinition(); enumDefinition != nil {
		toReturn = b.parseEnumDefinition(
			sourceUnit,
			toReturn,
			enumDefinition.(*parser.EnumDefinitionContext),
		)
	} else if structDefinition := bodyElement.StructDefinition(); structDefinition != nil {
		toReturn = b.parseStructDefinition(
			sourceUnit,
			toReturn,
			structDefinition.(*parser.StructDefinitionContext),
		)
	} else if modifierDefinition := bodyElement.ModifierDefinition(); modifierDefinition != nil {
		toReturn = b.parseModifierDefinition(
			sourceUnit,
			toReturn,
			modifierDefinition.(*parser.ModifierDefinitionContext),
		)
	} else if fallbackFunctionDefinition := bodyElement.FallbackFunctionDefinition(); fallbackFunctionDefinition != nil {
		toReturn = b.parseFallbackFunctionDefinition(
			sourceUnit,
			toReturn,
			fallbackFunctionDefinition.(*parser.FallbackFunctionDefinitionContext),
		)
	} else if receiveFunctionDefinition := bodyElement.ReceiveFunctionDefinition(); receiveFunctionDefinition != nil {
		toReturn = b.parseReceiveFunctionDefinition(
			sourceUnit,
			toReturn,
			receiveFunctionDefinition.(*parser.ReceiveFunctionDefinitionContext),
		)
	} else if errorDefinition := bodyElement.ErrorDefinition(); errorDefinition != nil {
		toReturn = b.parseErrorDefinition(
			sourceUnit,
			toReturn,
			errorDefinition.(*parser.ErrorDefinitionContext),
		)
	} else if userDefinedValue := bodyElement.UserDefinedValueTypeDefinition(); userDefinedValue != nil {
		zap.L().Warn(
			"User defined value type definition not implemented",
			zap.String("source_unit_name", sourceUnit.GetName()),
			zap.Int64("line", int64(bodyElement.GetStart().GetLine())),
		)
	}

	return toReturn
}
