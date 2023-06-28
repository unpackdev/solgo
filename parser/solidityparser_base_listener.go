// Code generated from ../antlr/SolidityParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SolidityParser
import "github.com/antlr4-go/antlr/v4"

// BaseSolidityParserListener is a complete listener for a parse tree produced by SolidityParser.
type BaseSolidityParserListener struct{}

var _ SolidityParserListener = &BaseSolidityParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSolidityParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSolidityParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSolidityParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSolidityParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceUnit is called when production sourceUnit is entered.
func (s *BaseSolidityParserListener) EnterSourceUnit(ctx *SourceUnitContext) {}

// ExitSourceUnit is called when production sourceUnit is exited.
func (s *BaseSolidityParserListener) ExitSourceUnit(ctx *SourceUnitContext) {}

// EnterPragmaDirective is called when production pragmaDirective is entered.
func (s *BaseSolidityParserListener) EnterPragmaDirective(ctx *PragmaDirectiveContext) {}

// ExitPragmaDirective is called when production pragmaDirective is exited.
func (s *BaseSolidityParserListener) ExitPragmaDirective(ctx *PragmaDirectiveContext) {}

// EnterImportDirective is called when production importDirective is entered.
func (s *BaseSolidityParserListener) EnterImportDirective(ctx *ImportDirectiveContext) {}

// ExitImportDirective is called when production importDirective is exited.
func (s *BaseSolidityParserListener) ExitImportDirective(ctx *ImportDirectiveContext) {}

// EnterImportAliases is called when production importAliases is entered.
func (s *BaseSolidityParserListener) EnterImportAliases(ctx *ImportAliasesContext) {}

// ExitImportAliases is called when production importAliases is exited.
func (s *BaseSolidityParserListener) ExitImportAliases(ctx *ImportAliasesContext) {}

// EnterPath is called when production path is entered.
func (s *BaseSolidityParserListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseSolidityParserListener) ExitPath(ctx *PathContext) {}

// EnterSymbolAliases is called when production symbolAliases is entered.
func (s *BaseSolidityParserListener) EnterSymbolAliases(ctx *SymbolAliasesContext) {}

// ExitSymbolAliases is called when production symbolAliases is exited.
func (s *BaseSolidityParserListener) ExitSymbolAliases(ctx *SymbolAliasesContext) {}

// EnterContractDefinition is called when production contractDefinition is entered.
func (s *BaseSolidityParserListener) EnterContractDefinition(ctx *ContractDefinitionContext) {}

// ExitContractDefinition is called when production contractDefinition is exited.
func (s *BaseSolidityParserListener) ExitContractDefinition(ctx *ContractDefinitionContext) {}

// EnterInterfaceDefinition is called when production interfaceDefinition is entered.
func (s *BaseSolidityParserListener) EnterInterfaceDefinition(ctx *InterfaceDefinitionContext) {}

// ExitInterfaceDefinition is called when production interfaceDefinition is exited.
func (s *BaseSolidityParserListener) ExitInterfaceDefinition(ctx *InterfaceDefinitionContext) {}

// EnterLibraryDefinition is called when production libraryDefinition is entered.
func (s *BaseSolidityParserListener) EnterLibraryDefinition(ctx *LibraryDefinitionContext) {}

// ExitLibraryDefinition is called when production libraryDefinition is exited.
func (s *BaseSolidityParserListener) ExitLibraryDefinition(ctx *LibraryDefinitionContext) {}

// EnterInheritanceSpecifierList is called when production inheritanceSpecifierList is entered.
func (s *BaseSolidityParserListener) EnterInheritanceSpecifierList(ctx *InheritanceSpecifierListContext) {
}

// ExitInheritanceSpecifierList is called when production inheritanceSpecifierList is exited.
func (s *BaseSolidityParserListener) ExitInheritanceSpecifierList(ctx *InheritanceSpecifierListContext) {
}

// EnterInheritanceSpecifier is called when production inheritanceSpecifier is entered.
func (s *BaseSolidityParserListener) EnterInheritanceSpecifier(ctx *InheritanceSpecifierContext) {}

// ExitInheritanceSpecifier is called when production inheritanceSpecifier is exited.
func (s *BaseSolidityParserListener) ExitInheritanceSpecifier(ctx *InheritanceSpecifierContext) {}

// EnterContractBodyElement is called when production contractBodyElement is entered.
func (s *BaseSolidityParserListener) EnterContractBodyElement(ctx *ContractBodyElementContext) {}

// ExitContractBodyElement is called when production contractBodyElement is exited.
func (s *BaseSolidityParserListener) ExitContractBodyElement(ctx *ContractBodyElementContext) {}

// EnterNamedArgument is called when production namedArgument is entered.
func (s *BaseSolidityParserListener) EnterNamedArgument(ctx *NamedArgumentContext) {}

// ExitNamedArgument is called when production namedArgument is exited.
func (s *BaseSolidityParserListener) ExitNamedArgument(ctx *NamedArgumentContext) {}

// EnterCallArgumentList is called when production callArgumentList is entered.
func (s *BaseSolidityParserListener) EnterCallArgumentList(ctx *CallArgumentListContext) {}

// ExitCallArgumentList is called when production callArgumentList is exited.
func (s *BaseSolidityParserListener) ExitCallArgumentList(ctx *CallArgumentListContext) {}

// EnterIdentifierPath is called when production identifierPath is entered.
func (s *BaseSolidityParserListener) EnterIdentifierPath(ctx *IdentifierPathContext) {}

// ExitIdentifierPath is called when production identifierPath is exited.
func (s *BaseSolidityParserListener) ExitIdentifierPath(ctx *IdentifierPathContext) {}

// EnterModifierInvocation is called when production modifierInvocation is entered.
func (s *BaseSolidityParserListener) EnterModifierInvocation(ctx *ModifierInvocationContext) {}

// ExitModifierInvocation is called when production modifierInvocation is exited.
func (s *BaseSolidityParserListener) ExitModifierInvocation(ctx *ModifierInvocationContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseSolidityParserListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseSolidityParserListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseSolidityParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseSolidityParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseSolidityParserListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseSolidityParserListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterConstructorDefinition is called when production constructorDefinition is entered.
func (s *BaseSolidityParserListener) EnterConstructorDefinition(ctx *ConstructorDefinitionContext) {}

// ExitConstructorDefinition is called when production constructorDefinition is exited.
func (s *BaseSolidityParserListener) ExitConstructorDefinition(ctx *ConstructorDefinitionContext) {}

// EnterStateMutability is called when production stateMutability is entered.
func (s *BaseSolidityParserListener) EnterStateMutability(ctx *StateMutabilityContext) {}

// ExitStateMutability is called when production stateMutability is exited.
func (s *BaseSolidityParserListener) ExitStateMutability(ctx *StateMutabilityContext) {}

// EnterOverrideSpecifier is called when production overrideSpecifier is entered.
func (s *BaseSolidityParserListener) EnterOverrideSpecifier(ctx *OverrideSpecifierContext) {}

// ExitOverrideSpecifier is called when production overrideSpecifier is exited.
func (s *BaseSolidityParserListener) ExitOverrideSpecifier(ctx *OverrideSpecifierContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseSolidityParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseSolidityParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterModifierDefinition is called when production modifierDefinition is entered.
func (s *BaseSolidityParserListener) EnterModifierDefinition(ctx *ModifierDefinitionContext) {}

// ExitModifierDefinition is called when production modifierDefinition is exited.
func (s *BaseSolidityParserListener) ExitModifierDefinition(ctx *ModifierDefinitionContext) {}

// EnterFallbackFunctionDefinition is called when production fallbackFunctionDefinition is entered.
func (s *BaseSolidityParserListener) EnterFallbackFunctionDefinition(ctx *FallbackFunctionDefinitionContext) {
}

// ExitFallbackFunctionDefinition is called when production fallbackFunctionDefinition is exited.
func (s *BaseSolidityParserListener) ExitFallbackFunctionDefinition(ctx *FallbackFunctionDefinitionContext) {
}

// EnterReceiveFunctionDefinition is called when production receiveFunctionDefinition is entered.
func (s *BaseSolidityParserListener) EnterReceiveFunctionDefinition(ctx *ReceiveFunctionDefinitionContext) {
}

// ExitReceiveFunctionDefinition is called when production receiveFunctionDefinition is exited.
func (s *BaseSolidityParserListener) ExitReceiveFunctionDefinition(ctx *ReceiveFunctionDefinitionContext) {
}

// EnterStructDefinition is called when production structDefinition is entered.
func (s *BaseSolidityParserListener) EnterStructDefinition(ctx *StructDefinitionContext) {}

// ExitStructDefinition is called when production structDefinition is exited.
func (s *BaseSolidityParserListener) ExitStructDefinition(ctx *StructDefinitionContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseSolidityParserListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseSolidityParserListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BaseSolidityParserListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BaseSolidityParserListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterUserDefinedValueTypeDefinition is called when production userDefinedValueTypeDefinition is entered.
func (s *BaseSolidityParserListener) EnterUserDefinedValueTypeDefinition(ctx *UserDefinedValueTypeDefinitionContext) {
}

// ExitUserDefinedValueTypeDefinition is called when production userDefinedValueTypeDefinition is exited.
func (s *BaseSolidityParserListener) ExitUserDefinedValueTypeDefinition(ctx *UserDefinedValueTypeDefinitionContext) {
}

// EnterStateVariableDeclaration is called when production stateVariableDeclaration is entered.
func (s *BaseSolidityParserListener) EnterStateVariableDeclaration(ctx *StateVariableDeclarationContext) {
}

// ExitStateVariableDeclaration is called when production stateVariableDeclaration is exited.
func (s *BaseSolidityParserListener) ExitStateVariableDeclaration(ctx *StateVariableDeclarationContext) {
}

// EnterConstantVariableDeclaration is called when production constantVariableDeclaration is entered.
func (s *BaseSolidityParserListener) EnterConstantVariableDeclaration(ctx *ConstantVariableDeclarationContext) {
}

// ExitConstantVariableDeclaration is called when production constantVariableDeclaration is exited.
func (s *BaseSolidityParserListener) ExitConstantVariableDeclaration(ctx *ConstantVariableDeclarationContext) {
}

// EnterEventParameter is called when production eventParameter is entered.
func (s *BaseSolidityParserListener) EnterEventParameter(ctx *EventParameterContext) {}

// ExitEventParameter is called when production eventParameter is exited.
func (s *BaseSolidityParserListener) ExitEventParameter(ctx *EventParameterContext) {}

// EnterEventDefinition is called when production eventDefinition is entered.
func (s *BaseSolidityParserListener) EnterEventDefinition(ctx *EventDefinitionContext) {}

// ExitEventDefinition is called when production eventDefinition is exited.
func (s *BaseSolidityParserListener) ExitEventDefinition(ctx *EventDefinitionContext) {}

// EnterErrorParameter is called when production errorParameter is entered.
func (s *BaseSolidityParserListener) EnterErrorParameter(ctx *ErrorParameterContext) {}

// ExitErrorParameter is called when production errorParameter is exited.
func (s *BaseSolidityParserListener) ExitErrorParameter(ctx *ErrorParameterContext) {}

// EnterErrorDefinition is called when production errorDefinition is entered.
func (s *BaseSolidityParserListener) EnterErrorDefinition(ctx *ErrorDefinitionContext) {}

// ExitErrorDefinition is called when production errorDefinition is exited.
func (s *BaseSolidityParserListener) ExitErrorDefinition(ctx *ErrorDefinitionContext) {}

// EnterUserDefinableOperator is called when production userDefinableOperator is entered.
func (s *BaseSolidityParserListener) EnterUserDefinableOperator(ctx *UserDefinableOperatorContext) {}

// ExitUserDefinableOperator is called when production userDefinableOperator is exited.
func (s *BaseSolidityParserListener) ExitUserDefinableOperator(ctx *UserDefinableOperatorContext) {}

// EnterUsingDirective is called when production usingDirective is entered.
func (s *BaseSolidityParserListener) EnterUsingDirective(ctx *UsingDirectiveContext) {}

// ExitUsingDirective is called when production usingDirective is exited.
func (s *BaseSolidityParserListener) ExitUsingDirective(ctx *UsingDirectiveContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseSolidityParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseSolidityParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterElementaryTypeName is called when production elementaryTypeName is entered.
func (s *BaseSolidityParserListener) EnterElementaryTypeName(ctx *ElementaryTypeNameContext) {}

// ExitElementaryTypeName is called when production elementaryTypeName is exited.
func (s *BaseSolidityParserListener) ExitElementaryTypeName(ctx *ElementaryTypeNameContext) {}

// EnterFunctionTypeName is called when production functionTypeName is entered.
func (s *BaseSolidityParserListener) EnterFunctionTypeName(ctx *FunctionTypeNameContext) {}

// ExitFunctionTypeName is called when production functionTypeName is exited.
func (s *BaseSolidityParserListener) ExitFunctionTypeName(ctx *FunctionTypeNameContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseSolidityParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseSolidityParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterDataLocation is called when production dataLocation is entered.
func (s *BaseSolidityParserListener) EnterDataLocation(ctx *DataLocationContext) {}

// ExitDataLocation is called when production dataLocation is exited.
func (s *BaseSolidityParserListener) ExitDataLocation(ctx *DataLocationContext) {}

// EnterUnaryPrefixOperation is called when production UnaryPrefixOperation is entered.
func (s *BaseSolidityParserListener) EnterUnaryPrefixOperation(ctx *UnaryPrefixOperationContext) {}

// ExitUnaryPrefixOperation is called when production UnaryPrefixOperation is exited.
func (s *BaseSolidityParserListener) ExitUnaryPrefixOperation(ctx *UnaryPrefixOperationContext) {}

// EnterPrimaryExpression is called when production PrimaryExpression is entered.
func (s *BaseSolidityParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production PrimaryExpression is exited.
func (s *BaseSolidityParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterOrderComparison is called when production OrderComparison is entered.
func (s *BaseSolidityParserListener) EnterOrderComparison(ctx *OrderComparisonContext) {}

// ExitOrderComparison is called when production OrderComparison is exited.
func (s *BaseSolidityParserListener) ExitOrderComparison(ctx *OrderComparisonContext) {}

// EnterConditional is called when production Conditional is entered.
func (s *BaseSolidityParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production Conditional is exited.
func (s *BaseSolidityParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterPayableConversion is called when production PayableConversion is entered.
func (s *BaseSolidityParserListener) EnterPayableConversion(ctx *PayableConversionContext) {}

// ExitPayableConversion is called when production PayableConversion is exited.
func (s *BaseSolidityParserListener) ExitPayableConversion(ctx *PayableConversionContext) {}

// EnterAssignment is called when production Assignment is entered.
func (s *BaseSolidityParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production Assignment is exited.
func (s *BaseSolidityParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterUnarySuffixOperation is called when production UnarySuffixOperation is entered.
func (s *BaseSolidityParserListener) EnterUnarySuffixOperation(ctx *UnarySuffixOperationContext) {}

// ExitUnarySuffixOperation is called when production UnarySuffixOperation is exited.
func (s *BaseSolidityParserListener) ExitUnarySuffixOperation(ctx *UnarySuffixOperationContext) {}

// EnterShiftOperation is called when production ShiftOperation is entered.
func (s *BaseSolidityParserListener) EnterShiftOperation(ctx *ShiftOperationContext) {}

// ExitShiftOperation is called when production ShiftOperation is exited.
func (s *BaseSolidityParserListener) ExitShiftOperation(ctx *ShiftOperationContext) {}

// EnterBitAndOperation is called when production BitAndOperation is entered.
func (s *BaseSolidityParserListener) EnterBitAndOperation(ctx *BitAndOperationContext) {}

// ExitBitAndOperation is called when production BitAndOperation is exited.
func (s *BaseSolidityParserListener) ExitBitAndOperation(ctx *BitAndOperationContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseSolidityParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseSolidityParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterIndexRangeAccess is called when production IndexRangeAccess is entered.
func (s *BaseSolidityParserListener) EnterIndexRangeAccess(ctx *IndexRangeAccessContext) {}

// ExitIndexRangeAccess is called when production IndexRangeAccess is exited.
func (s *BaseSolidityParserListener) ExitIndexRangeAccess(ctx *IndexRangeAccessContext) {}

// EnterIndexAccess is called when production IndexAccess is entered.
func (s *BaseSolidityParserListener) EnterIndexAccess(ctx *IndexAccessContext) {}

// ExitIndexAccess is called when production IndexAccess is exited.
func (s *BaseSolidityParserListener) ExitIndexAccess(ctx *IndexAccessContext) {}

// EnterAddSubOperation is called when production AddSubOperation is entered.
func (s *BaseSolidityParserListener) EnterAddSubOperation(ctx *AddSubOperationContext) {}

// ExitAddSubOperation is called when production AddSubOperation is exited.
func (s *BaseSolidityParserListener) ExitAddSubOperation(ctx *AddSubOperationContext) {}

// EnterBitOrOperation is called when production BitOrOperation is entered.
func (s *BaseSolidityParserListener) EnterBitOrOperation(ctx *BitOrOperationContext) {}

// ExitBitOrOperation is called when production BitOrOperation is exited.
func (s *BaseSolidityParserListener) ExitBitOrOperation(ctx *BitOrOperationContext) {}

// EnterExpOperation is called when production ExpOperation is entered.
func (s *BaseSolidityParserListener) EnterExpOperation(ctx *ExpOperationContext) {}

// ExitExpOperation is called when production ExpOperation is exited.
func (s *BaseSolidityParserListener) ExitExpOperation(ctx *ExpOperationContext) {}

// EnterAndOperation is called when production AndOperation is entered.
func (s *BaseSolidityParserListener) EnterAndOperation(ctx *AndOperationContext) {}

// ExitAndOperation is called when production AndOperation is exited.
func (s *BaseSolidityParserListener) ExitAndOperation(ctx *AndOperationContext) {}

// EnterInlineArray is called when production InlineArray is entered.
func (s *BaseSolidityParserListener) EnterInlineArray(ctx *InlineArrayContext) {}

// ExitInlineArray is called when production InlineArray is exited.
func (s *BaseSolidityParserListener) ExitInlineArray(ctx *InlineArrayContext) {}

// EnterOrOperation is called when production OrOperation is entered.
func (s *BaseSolidityParserListener) EnterOrOperation(ctx *OrOperationContext) {}

// ExitOrOperation is called when production OrOperation is exited.
func (s *BaseSolidityParserListener) ExitOrOperation(ctx *OrOperationContext) {}

// EnterMemberAccess is called when production MemberAccess is entered.
func (s *BaseSolidityParserListener) EnterMemberAccess(ctx *MemberAccessContext) {}

// ExitMemberAccess is called when production MemberAccess is exited.
func (s *BaseSolidityParserListener) ExitMemberAccess(ctx *MemberAccessContext) {}

// EnterMulDivModOperation is called when production MulDivModOperation is entered.
func (s *BaseSolidityParserListener) EnterMulDivModOperation(ctx *MulDivModOperationContext) {}

// ExitMulDivModOperation is called when production MulDivModOperation is exited.
func (s *BaseSolidityParserListener) ExitMulDivModOperation(ctx *MulDivModOperationContext) {}

// EnterFunctionCallOptions is called when production FunctionCallOptions is entered.
func (s *BaseSolidityParserListener) EnterFunctionCallOptions(ctx *FunctionCallOptionsContext) {}

// ExitFunctionCallOptions is called when production FunctionCallOptions is exited.
func (s *BaseSolidityParserListener) ExitFunctionCallOptions(ctx *FunctionCallOptionsContext) {}

// EnterNewExpr is called when production NewExpr is entered.
func (s *BaseSolidityParserListener) EnterNewExpr(ctx *NewExprContext) {}

// ExitNewExpr is called when production NewExpr is exited.
func (s *BaseSolidityParserListener) ExitNewExpr(ctx *NewExprContext) {}

// EnterBitXorOperation is called when production BitXorOperation is entered.
func (s *BaseSolidityParserListener) EnterBitXorOperation(ctx *BitXorOperationContext) {}

// ExitBitXorOperation is called when production BitXorOperation is exited.
func (s *BaseSolidityParserListener) ExitBitXorOperation(ctx *BitXorOperationContext) {}

// EnterTuple is called when production Tuple is entered.
func (s *BaseSolidityParserListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production Tuple is exited.
func (s *BaseSolidityParserListener) ExitTuple(ctx *TupleContext) {}

// EnterEqualityComparison is called when production EqualityComparison is entered.
func (s *BaseSolidityParserListener) EnterEqualityComparison(ctx *EqualityComparisonContext) {}

// ExitEqualityComparison is called when production EqualityComparison is exited.
func (s *BaseSolidityParserListener) ExitEqualityComparison(ctx *EqualityComparisonContext) {}

// EnterMetaType is called when production MetaType is entered.
func (s *BaseSolidityParserListener) EnterMetaType(ctx *MetaTypeContext) {}

// ExitMetaType is called when production MetaType is exited.
func (s *BaseSolidityParserListener) ExitMetaType(ctx *MetaTypeContext) {}

// EnterAssignOp is called when production assignOp is entered.
func (s *BaseSolidityParserListener) EnterAssignOp(ctx *AssignOpContext) {}

// ExitAssignOp is called when production assignOp is exited.
func (s *BaseSolidityParserListener) ExitAssignOp(ctx *AssignOpContext) {}

// EnterTupleExpression is called when production tupleExpression is entered.
func (s *BaseSolidityParserListener) EnterTupleExpression(ctx *TupleExpressionContext) {}

// ExitTupleExpression is called when production tupleExpression is exited.
func (s *BaseSolidityParserListener) ExitTupleExpression(ctx *TupleExpressionContext) {}

// EnterInlineArrayExpression is called when production inlineArrayExpression is entered.
func (s *BaseSolidityParserListener) EnterInlineArrayExpression(ctx *InlineArrayExpressionContext) {}

// ExitInlineArrayExpression is called when production inlineArrayExpression is exited.
func (s *BaseSolidityParserListener) ExitInlineArrayExpression(ctx *InlineArrayExpressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSolidityParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSolidityParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSolidityParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSolidityParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralWithSubDenomination is called when production literalWithSubDenomination is entered.
func (s *BaseSolidityParserListener) EnterLiteralWithSubDenomination(ctx *LiteralWithSubDenominationContext) {
}

// ExitLiteralWithSubDenomination is called when production literalWithSubDenomination is exited.
func (s *BaseSolidityParserListener) ExitLiteralWithSubDenomination(ctx *LiteralWithSubDenominationContext) {
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSolidityParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSolidityParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSolidityParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSolidityParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterHexStringLiteral is called when production hexStringLiteral is entered.
func (s *BaseSolidityParserListener) EnterHexStringLiteral(ctx *HexStringLiteralContext) {}

// ExitHexStringLiteral is called when production hexStringLiteral is exited.
func (s *BaseSolidityParserListener) ExitHexStringLiteral(ctx *HexStringLiteralContext) {}

// EnterUnicodeStringLiteral is called when production unicodeStringLiteral is entered.
func (s *BaseSolidityParserListener) EnterUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) {}

// ExitUnicodeStringLiteral is called when production unicodeStringLiteral is exited.
func (s *BaseSolidityParserListener) ExitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseSolidityParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseSolidityParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSolidityParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSolidityParserListener) ExitBlock(ctx *BlockContext) {}

// EnterUncheckedBlock is called when production uncheckedBlock is entered.
func (s *BaseSolidityParserListener) EnterUncheckedBlock(ctx *UncheckedBlockContext) {}

// ExitUncheckedBlock is called when production uncheckedBlock is exited.
func (s *BaseSolidityParserListener) ExitUncheckedBlock(ctx *UncheckedBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSolidityParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSolidityParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseSolidityParserListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseSolidityParserListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseSolidityParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseSolidityParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseSolidityParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseSolidityParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseSolidityParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseSolidityParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseSolidityParserListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseSolidityParserListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseSolidityParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseSolidityParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseSolidityParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseSolidityParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseSolidityParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseSolidityParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseSolidityParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseSolidityParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseSolidityParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseSolidityParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterEmitStatement is called when production emitStatement is entered.
func (s *BaseSolidityParserListener) EnterEmitStatement(ctx *EmitStatementContext) {}

// ExitEmitStatement is called when production emitStatement is exited.
func (s *BaseSolidityParserListener) ExitEmitStatement(ctx *EmitStatementContext) {}

// EnterRevertStatement is called when production revertStatement is entered.
func (s *BaseSolidityParserListener) EnterRevertStatement(ctx *RevertStatementContext) {}

// ExitRevertStatement is called when production revertStatement is exited.
func (s *BaseSolidityParserListener) ExitRevertStatement(ctx *RevertStatementContext) {}

// EnterAssemblyStatement is called when production assemblyStatement is entered.
func (s *BaseSolidityParserListener) EnterAssemblyStatement(ctx *AssemblyStatementContext) {}

// ExitAssemblyStatement is called when production assemblyStatement is exited.
func (s *BaseSolidityParserListener) ExitAssemblyStatement(ctx *AssemblyStatementContext) {}

// EnterAssemblyFlags is called when production assemblyFlags is entered.
func (s *BaseSolidityParserListener) EnterAssemblyFlags(ctx *AssemblyFlagsContext) {}

// ExitAssemblyFlags is called when production assemblyFlags is exited.
func (s *BaseSolidityParserListener) ExitAssemblyFlags(ctx *AssemblyFlagsContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseSolidityParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseSolidityParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// EnterVariableDeclarationTuple is called when production variableDeclarationTuple is entered.
func (s *BaseSolidityParserListener) EnterVariableDeclarationTuple(ctx *VariableDeclarationTupleContext) {
}

// ExitVariableDeclarationTuple is called when production variableDeclarationTuple is exited.
func (s *BaseSolidityParserListener) ExitVariableDeclarationTuple(ctx *VariableDeclarationTupleContext) {
}

// EnterVariableDeclarationStatement is called when production variableDeclarationStatement is entered.
func (s *BaseSolidityParserListener) EnterVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// ExitVariableDeclarationStatement is called when production variableDeclarationStatement is exited.
func (s *BaseSolidityParserListener) ExitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseSolidityParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseSolidityParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterMappingType is called when production mappingType is entered.
func (s *BaseSolidityParserListener) EnterMappingType(ctx *MappingTypeContext) {}

// ExitMappingType is called when production mappingType is exited.
func (s *BaseSolidityParserListener) ExitMappingType(ctx *MappingTypeContext) {}

// EnterMappingKeyType is called when production mappingKeyType is entered.
func (s *BaseSolidityParserListener) EnterMappingKeyType(ctx *MappingKeyTypeContext) {}

// ExitMappingKeyType is called when production mappingKeyType is exited.
func (s *BaseSolidityParserListener) ExitMappingKeyType(ctx *MappingKeyTypeContext) {}

// EnterYulStatement is called when production yulStatement is entered.
func (s *BaseSolidityParserListener) EnterYulStatement(ctx *YulStatementContext) {}

// ExitYulStatement is called when production yulStatement is exited.
func (s *BaseSolidityParserListener) ExitYulStatement(ctx *YulStatementContext) {}

// EnterYulBlock is called when production yulBlock is entered.
func (s *BaseSolidityParserListener) EnterYulBlock(ctx *YulBlockContext) {}

// ExitYulBlock is called when production yulBlock is exited.
func (s *BaseSolidityParserListener) ExitYulBlock(ctx *YulBlockContext) {}

// EnterYulVariableDeclaration is called when production yulVariableDeclaration is entered.
func (s *BaseSolidityParserListener) EnterYulVariableDeclaration(ctx *YulVariableDeclarationContext) {
}

// ExitYulVariableDeclaration is called when production yulVariableDeclaration is exited.
func (s *BaseSolidityParserListener) ExitYulVariableDeclaration(ctx *YulVariableDeclarationContext) {}

// EnterYulAssignment is called when production yulAssignment is entered.
func (s *BaseSolidityParserListener) EnterYulAssignment(ctx *YulAssignmentContext) {}

// ExitYulAssignment is called when production yulAssignment is exited.
func (s *BaseSolidityParserListener) ExitYulAssignment(ctx *YulAssignmentContext) {}

// EnterYulIfStatement is called when production yulIfStatement is entered.
func (s *BaseSolidityParserListener) EnterYulIfStatement(ctx *YulIfStatementContext) {}

// ExitYulIfStatement is called when production yulIfStatement is exited.
func (s *BaseSolidityParserListener) ExitYulIfStatement(ctx *YulIfStatementContext) {}

// EnterYulForStatement is called when production yulForStatement is entered.
func (s *BaseSolidityParserListener) EnterYulForStatement(ctx *YulForStatementContext) {}

// ExitYulForStatement is called when production yulForStatement is exited.
func (s *BaseSolidityParserListener) ExitYulForStatement(ctx *YulForStatementContext) {}

// EnterYulSwitchCase is called when production yulSwitchCase is entered.
func (s *BaseSolidityParserListener) EnterYulSwitchCase(ctx *YulSwitchCaseContext) {}

// ExitYulSwitchCase is called when production yulSwitchCase is exited.
func (s *BaseSolidityParserListener) ExitYulSwitchCase(ctx *YulSwitchCaseContext) {}

// EnterYulSwitchStatement is called when production yulSwitchStatement is entered.
func (s *BaseSolidityParserListener) EnterYulSwitchStatement(ctx *YulSwitchStatementContext) {}

// ExitYulSwitchStatement is called when production yulSwitchStatement is exited.
func (s *BaseSolidityParserListener) ExitYulSwitchStatement(ctx *YulSwitchStatementContext) {}

// EnterYulFunctionDefinition is called when production yulFunctionDefinition is entered.
func (s *BaseSolidityParserListener) EnterYulFunctionDefinition(ctx *YulFunctionDefinitionContext) {}

// ExitYulFunctionDefinition is called when production yulFunctionDefinition is exited.
func (s *BaseSolidityParserListener) ExitYulFunctionDefinition(ctx *YulFunctionDefinitionContext) {}

// EnterYulPath is called when production yulPath is entered.
func (s *BaseSolidityParserListener) EnterYulPath(ctx *YulPathContext) {}

// ExitYulPath is called when production yulPath is exited.
func (s *BaseSolidityParserListener) ExitYulPath(ctx *YulPathContext) {}

// EnterYulFunctionCall is called when production yulFunctionCall is entered.
func (s *BaseSolidityParserListener) EnterYulFunctionCall(ctx *YulFunctionCallContext) {}

// ExitYulFunctionCall is called when production yulFunctionCall is exited.
func (s *BaseSolidityParserListener) ExitYulFunctionCall(ctx *YulFunctionCallContext) {}

// EnterYulBoolean is called when production yulBoolean is entered.
func (s *BaseSolidityParserListener) EnterYulBoolean(ctx *YulBooleanContext) {}

// ExitYulBoolean is called when production yulBoolean is exited.
func (s *BaseSolidityParserListener) ExitYulBoolean(ctx *YulBooleanContext) {}

// EnterYulLiteral is called when production yulLiteral is entered.
func (s *BaseSolidityParserListener) EnterYulLiteral(ctx *YulLiteralContext) {}

// ExitYulLiteral is called when production yulLiteral is exited.
func (s *BaseSolidityParserListener) ExitYulLiteral(ctx *YulLiteralContext) {}

// EnterYulExpression is called when production yulExpression is entered.
func (s *BaseSolidityParserListener) EnterYulExpression(ctx *YulExpressionContext) {}

// ExitYulExpression is called when production yulExpression is exited.
func (s *BaseSolidityParserListener) ExitYulExpression(ctx *YulExpressionContext) {}
