// Code generated from ../antlr/SolidityParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SolidityParser
import "github.com/antlr4-go/antlr/v4"

// SolidityParserListener is a complete listener for a parse tree produced by SolidityParser.
type SolidityParserListener interface {
	antlr.ParseTreeListener

	// EnterSourceUnit is called when entering the sourceUnit production.
	EnterSourceUnit(c *SourceUnitContext)

	// EnterPragmaDirective is called when entering the pragmaDirective production.
	EnterPragmaDirective(c *PragmaDirectiveContext)

	// EnterImportDirective is called when entering the importDirective production.
	EnterImportDirective(c *ImportDirectiveContext)

	// EnterImportAliases is called when entering the importAliases production.
	EnterImportAliases(c *ImportAliasesContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterSymbolAliases is called when entering the symbolAliases production.
	EnterSymbolAliases(c *SymbolAliasesContext)

	// EnterContractDefinition is called when entering the contractDefinition production.
	EnterContractDefinition(c *ContractDefinitionContext)

	// EnterInterfaceDefinition is called when entering the interfaceDefinition production.
	EnterInterfaceDefinition(c *InterfaceDefinitionContext)

	// EnterLibraryDefinition is called when entering the libraryDefinition production.
	EnterLibraryDefinition(c *LibraryDefinitionContext)

	// EnterInheritanceSpecifierList is called when entering the inheritanceSpecifierList production.
	EnterInheritanceSpecifierList(c *InheritanceSpecifierListContext)

	// EnterInheritanceSpecifier is called when entering the inheritanceSpecifier production.
	EnterInheritanceSpecifier(c *InheritanceSpecifierContext)

	// EnterContractBodyElement is called when entering the contractBodyElement production.
	EnterContractBodyElement(c *ContractBodyElementContext)

	// EnterNamedArgument is called when entering the namedArgument production.
	EnterNamedArgument(c *NamedArgumentContext)

	// EnterCallArgumentList is called when entering the callArgumentList production.
	EnterCallArgumentList(c *CallArgumentListContext)

	// EnterIdentifierPath is called when entering the identifierPath production.
	EnterIdentifierPath(c *IdentifierPathContext)

	// EnterModifierInvocation is called when entering the modifierInvocation production.
	EnterModifierInvocation(c *ModifierInvocationContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterConstructorDefinition is called when entering the constructorDefinition production.
	EnterConstructorDefinition(c *ConstructorDefinitionContext)

	// EnterStateMutability is called when entering the stateMutability production.
	EnterStateMutability(c *StateMutabilityContext)

	// EnterOverrideSpecifier is called when entering the overrideSpecifier production.
	EnterOverrideSpecifier(c *OverrideSpecifierContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterModifierDefinition is called when entering the modifierDefinition production.
	EnterModifierDefinition(c *ModifierDefinitionContext)

	// EnterFallbackFunctionDefinition is called when entering the fallbackFunctionDefinition production.
	EnterFallbackFunctionDefinition(c *FallbackFunctionDefinitionContext)

	// EnterReceiveFunctionDefinition is called when entering the receiveFunctionDefinition production.
	EnterReceiveFunctionDefinition(c *ReceiveFunctionDefinitionContext)

	// EnterStructDefinition is called when entering the structDefinition production.
	EnterStructDefinition(c *StructDefinitionContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterUserDefinedValueTypeDefinition is called when entering the userDefinedValueTypeDefinition production.
	EnterUserDefinedValueTypeDefinition(c *UserDefinedValueTypeDefinitionContext)

	// EnterStateVariableDeclaration is called when entering the stateVariableDeclaration production.
	EnterStateVariableDeclaration(c *StateVariableDeclarationContext)

	// EnterConstantVariableDeclaration is called when entering the constantVariableDeclaration production.
	EnterConstantVariableDeclaration(c *ConstantVariableDeclarationContext)

	// EnterEventParameter is called when entering the eventParameter production.
	EnterEventParameter(c *EventParameterContext)

	// EnterEventDefinition is called when entering the eventDefinition production.
	EnterEventDefinition(c *EventDefinitionContext)

	// EnterErrorParameter is called when entering the errorParameter production.
	EnterErrorParameter(c *ErrorParameterContext)

	// EnterErrorDefinition is called when entering the errorDefinition production.
	EnterErrorDefinition(c *ErrorDefinitionContext)

	// EnterUserDefinableOperator is called when entering the userDefinableOperator production.
	EnterUserDefinableOperator(c *UserDefinableOperatorContext)

	// EnterUsingDirective is called when entering the usingDirective production.
	EnterUsingDirective(c *UsingDirectiveContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterElementaryTypeName is called when entering the elementaryTypeName production.
	EnterElementaryTypeName(c *ElementaryTypeNameContext)

	// EnterFunctionTypeName is called when entering the functionTypeName production.
	EnterFunctionTypeName(c *FunctionTypeNameContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterDataLocation is called when entering the dataLocation production.
	EnterDataLocation(c *DataLocationContext)

	// EnterUnaryPrefixOperation is called when entering the UnaryPrefixOperation production.
	EnterUnaryPrefixOperation(c *UnaryPrefixOperationContext)

	// EnterPrimaryExpression is called when entering the PrimaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterOrderComparison is called when entering the OrderComparison production.
	EnterOrderComparison(c *OrderComparisonContext)

	// EnterConditional is called when entering the Conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterPayableConversion is called when entering the PayableConversion production.
	EnterPayableConversion(c *PayableConversionContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterUnarySuffixOperation is called when entering the UnarySuffixOperation production.
	EnterUnarySuffixOperation(c *UnarySuffixOperationContext)

	// EnterShiftOperation is called when entering the ShiftOperation production.
	EnterShiftOperation(c *ShiftOperationContext)

	// EnterBitAndOperation is called when entering the BitAndOperation production.
	EnterBitAndOperation(c *BitAndOperationContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterIndexRangeAccess is called when entering the IndexRangeAccess production.
	EnterIndexRangeAccess(c *IndexRangeAccessContext)

	// EnterIndexAccess is called when entering the IndexAccess production.
	EnterIndexAccess(c *IndexAccessContext)

	// EnterAddSubOperation is called when entering the AddSubOperation production.
	EnterAddSubOperation(c *AddSubOperationContext)

	// EnterBitOrOperation is called when entering the BitOrOperation production.
	EnterBitOrOperation(c *BitOrOperationContext)

	// EnterExpOperation is called when entering the ExpOperation production.
	EnterExpOperation(c *ExpOperationContext)

	// EnterAndOperation is called when entering the AndOperation production.
	EnterAndOperation(c *AndOperationContext)

	// EnterInlineArray is called when entering the InlineArray production.
	EnterInlineArray(c *InlineArrayContext)

	// EnterOrOperation is called when entering the OrOperation production.
	EnterOrOperation(c *OrOperationContext)

	// EnterMemberAccess is called when entering the MemberAccess production.
	EnterMemberAccess(c *MemberAccessContext)

	// EnterMulDivModOperation is called when entering the MulDivModOperation production.
	EnterMulDivModOperation(c *MulDivModOperationContext)

	// EnterFunctionCallOptions is called when entering the FunctionCallOptions production.
	EnterFunctionCallOptions(c *FunctionCallOptionsContext)

	// EnterNewExpr is called when entering the NewExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterBitXorOperation is called when entering the BitXorOperation production.
	EnterBitXorOperation(c *BitXorOperationContext)

	// EnterTuple is called when entering the Tuple production.
	EnterTuple(c *TupleContext)

	// EnterEqualityComparison is called when entering the EqualityComparison production.
	EnterEqualityComparison(c *EqualityComparisonContext)

	// EnterMetaType is called when entering the MetaType production.
	EnterMetaType(c *MetaTypeContext)

	// EnterAssignOp is called when entering the assignOp production.
	EnterAssignOp(c *AssignOpContext)

	// EnterTupleExpression is called when entering the tupleExpression production.
	EnterTupleExpression(c *TupleExpressionContext)

	// EnterInlineArrayExpression is called when entering the inlineArrayExpression production.
	EnterInlineArrayExpression(c *InlineArrayExpressionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralWithSubDenomination is called when entering the literalWithSubDenomination production.
	EnterLiteralWithSubDenomination(c *LiteralWithSubDenominationContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterHexStringLiteral is called when entering the hexStringLiteral production.
	EnterHexStringLiteral(c *HexStringLiteralContext)

	// EnterUnicodeStringLiteral is called when entering the unicodeStringLiteral production.
	EnterUnicodeStringLiteral(c *UnicodeStringLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterUncheckedBlock is called when entering the uncheckedBlock production.
	EnterUncheckedBlock(c *UncheckedBlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterEmitStatement is called when entering the emitStatement production.
	EnterEmitStatement(c *EmitStatementContext)

	// EnterRevertStatement is called when entering the revertStatement production.
	EnterRevertStatement(c *RevertStatementContext)

	// EnterAssemblyStatement is called when entering the assemblyStatement production.
	EnterAssemblyStatement(c *AssemblyStatementContext)

	// EnterAssemblyFlags is called when entering the assemblyFlags production.
	EnterAssemblyFlags(c *AssemblyFlagsContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclarationTuple is called when entering the variableDeclarationTuple production.
	EnterVariableDeclarationTuple(c *VariableDeclarationTupleContext)

	// EnterVariableDeclarationStatement is called when entering the variableDeclarationStatement production.
	EnterVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterMappingType is called when entering the mappingType production.
	EnterMappingType(c *MappingTypeContext)

	// EnterMappingKeyType is called when entering the mappingKeyType production.
	EnterMappingKeyType(c *MappingKeyTypeContext)

	// EnterYulStatement is called when entering the yulStatement production.
	EnterYulStatement(c *YulStatementContext)

	// EnterYulBlock is called when entering the yulBlock production.
	EnterYulBlock(c *YulBlockContext)

	// EnterYulVariableDeclaration is called when entering the yulVariableDeclaration production.
	EnterYulVariableDeclaration(c *YulVariableDeclarationContext)

	// EnterYulAssignment is called when entering the yulAssignment production.
	EnterYulAssignment(c *YulAssignmentContext)

	// EnterYulIfStatement is called when entering the yulIfStatement production.
	EnterYulIfStatement(c *YulIfStatementContext)

	// EnterYulForStatement is called when entering the yulForStatement production.
	EnterYulForStatement(c *YulForStatementContext)

	// EnterYulSwitchCase is called when entering the yulSwitchCase production.
	EnterYulSwitchCase(c *YulSwitchCaseContext)

	// EnterYulSwitchStatement is called when entering the yulSwitchStatement production.
	EnterYulSwitchStatement(c *YulSwitchStatementContext)

	// EnterYulFunctionDefinition is called when entering the yulFunctionDefinition production.
	EnterYulFunctionDefinition(c *YulFunctionDefinitionContext)

	// EnterYulPath is called when entering the yulPath production.
	EnterYulPath(c *YulPathContext)

	// EnterYulFunctionCall is called when entering the yulFunctionCall production.
	EnterYulFunctionCall(c *YulFunctionCallContext)

	// EnterYulBoolean is called when entering the yulBoolean production.
	EnterYulBoolean(c *YulBooleanContext)

	// EnterYulLiteral is called when entering the yulLiteral production.
	EnterYulLiteral(c *YulLiteralContext)

	// EnterYulExpression is called when entering the yulExpression production.
	EnterYulExpression(c *YulExpressionContext)

	// ExitSourceUnit is called when exiting the sourceUnit production.
	ExitSourceUnit(c *SourceUnitContext)

	// ExitPragmaDirective is called when exiting the pragmaDirective production.
	ExitPragmaDirective(c *PragmaDirectiveContext)

	// ExitImportDirective is called when exiting the importDirective production.
	ExitImportDirective(c *ImportDirectiveContext)

	// ExitImportAliases is called when exiting the importAliases production.
	ExitImportAliases(c *ImportAliasesContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitSymbolAliases is called when exiting the symbolAliases production.
	ExitSymbolAliases(c *SymbolAliasesContext)

	// ExitContractDefinition is called when exiting the contractDefinition production.
	ExitContractDefinition(c *ContractDefinitionContext)

	// ExitInterfaceDefinition is called when exiting the interfaceDefinition production.
	ExitInterfaceDefinition(c *InterfaceDefinitionContext)

	// ExitLibraryDefinition is called when exiting the libraryDefinition production.
	ExitLibraryDefinition(c *LibraryDefinitionContext)

	// ExitInheritanceSpecifierList is called when exiting the inheritanceSpecifierList production.
	ExitInheritanceSpecifierList(c *InheritanceSpecifierListContext)

	// ExitInheritanceSpecifier is called when exiting the inheritanceSpecifier production.
	ExitInheritanceSpecifier(c *InheritanceSpecifierContext)

	// ExitContractBodyElement is called when exiting the contractBodyElement production.
	ExitContractBodyElement(c *ContractBodyElementContext)

	// ExitNamedArgument is called when exiting the namedArgument production.
	ExitNamedArgument(c *NamedArgumentContext)

	// ExitCallArgumentList is called when exiting the callArgumentList production.
	ExitCallArgumentList(c *CallArgumentListContext)

	// ExitIdentifierPath is called when exiting the identifierPath production.
	ExitIdentifierPath(c *IdentifierPathContext)

	// ExitModifierInvocation is called when exiting the modifierInvocation production.
	ExitModifierInvocation(c *ModifierInvocationContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitConstructorDefinition is called when exiting the constructorDefinition production.
	ExitConstructorDefinition(c *ConstructorDefinitionContext)

	// ExitStateMutability is called when exiting the stateMutability production.
	ExitStateMutability(c *StateMutabilityContext)

	// ExitOverrideSpecifier is called when exiting the overrideSpecifier production.
	ExitOverrideSpecifier(c *OverrideSpecifierContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitModifierDefinition is called when exiting the modifierDefinition production.
	ExitModifierDefinition(c *ModifierDefinitionContext)

	// ExitFallbackFunctionDefinition is called when exiting the fallbackFunctionDefinition production.
	ExitFallbackFunctionDefinition(c *FallbackFunctionDefinitionContext)

	// ExitReceiveFunctionDefinition is called when exiting the receiveFunctionDefinition production.
	ExitReceiveFunctionDefinition(c *ReceiveFunctionDefinitionContext)

	// ExitStructDefinition is called when exiting the structDefinition production.
	ExitStructDefinition(c *StructDefinitionContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitUserDefinedValueTypeDefinition is called when exiting the userDefinedValueTypeDefinition production.
	ExitUserDefinedValueTypeDefinition(c *UserDefinedValueTypeDefinitionContext)

	// ExitStateVariableDeclaration is called when exiting the stateVariableDeclaration production.
	ExitStateVariableDeclaration(c *StateVariableDeclarationContext)

	// ExitConstantVariableDeclaration is called when exiting the constantVariableDeclaration production.
	ExitConstantVariableDeclaration(c *ConstantVariableDeclarationContext)

	// ExitEventParameter is called when exiting the eventParameter production.
	ExitEventParameter(c *EventParameterContext)

	// ExitEventDefinition is called when exiting the eventDefinition production.
	ExitEventDefinition(c *EventDefinitionContext)

	// ExitErrorParameter is called when exiting the errorParameter production.
	ExitErrorParameter(c *ErrorParameterContext)

	// ExitErrorDefinition is called when exiting the errorDefinition production.
	ExitErrorDefinition(c *ErrorDefinitionContext)

	// ExitUserDefinableOperator is called when exiting the userDefinableOperator production.
	ExitUserDefinableOperator(c *UserDefinableOperatorContext)

	// ExitUsingDirective is called when exiting the usingDirective production.
	ExitUsingDirective(c *UsingDirectiveContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitElementaryTypeName is called when exiting the elementaryTypeName production.
	ExitElementaryTypeName(c *ElementaryTypeNameContext)

	// ExitFunctionTypeName is called when exiting the functionTypeName production.
	ExitFunctionTypeName(c *FunctionTypeNameContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitDataLocation is called when exiting the dataLocation production.
	ExitDataLocation(c *DataLocationContext)

	// ExitUnaryPrefixOperation is called when exiting the UnaryPrefixOperation production.
	ExitUnaryPrefixOperation(c *UnaryPrefixOperationContext)

	// ExitPrimaryExpression is called when exiting the PrimaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitOrderComparison is called when exiting the OrderComparison production.
	ExitOrderComparison(c *OrderComparisonContext)

	// ExitConditional is called when exiting the Conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitPayableConversion is called when exiting the PayableConversion production.
	ExitPayableConversion(c *PayableConversionContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitUnarySuffixOperation is called when exiting the UnarySuffixOperation production.
	ExitUnarySuffixOperation(c *UnarySuffixOperationContext)

	// ExitShiftOperation is called when exiting the ShiftOperation production.
	ExitShiftOperation(c *ShiftOperationContext)

	// ExitBitAndOperation is called when exiting the BitAndOperation production.
	ExitBitAndOperation(c *BitAndOperationContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitIndexRangeAccess is called when exiting the IndexRangeAccess production.
	ExitIndexRangeAccess(c *IndexRangeAccessContext)

	// ExitIndexAccess is called when exiting the IndexAccess production.
	ExitIndexAccess(c *IndexAccessContext)

	// ExitAddSubOperation is called when exiting the AddSubOperation production.
	ExitAddSubOperation(c *AddSubOperationContext)

	// ExitBitOrOperation is called when exiting the BitOrOperation production.
	ExitBitOrOperation(c *BitOrOperationContext)

	// ExitExpOperation is called when exiting the ExpOperation production.
	ExitExpOperation(c *ExpOperationContext)

	// ExitAndOperation is called when exiting the AndOperation production.
	ExitAndOperation(c *AndOperationContext)

	// ExitInlineArray is called when exiting the InlineArray production.
	ExitInlineArray(c *InlineArrayContext)

	// ExitOrOperation is called when exiting the OrOperation production.
	ExitOrOperation(c *OrOperationContext)

	// ExitMemberAccess is called when exiting the MemberAccess production.
	ExitMemberAccess(c *MemberAccessContext)

	// ExitMulDivModOperation is called when exiting the MulDivModOperation production.
	ExitMulDivModOperation(c *MulDivModOperationContext)

	// ExitFunctionCallOptions is called when exiting the FunctionCallOptions production.
	ExitFunctionCallOptions(c *FunctionCallOptionsContext)

	// ExitNewExpr is called when exiting the NewExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitBitXorOperation is called when exiting the BitXorOperation production.
	ExitBitXorOperation(c *BitXorOperationContext)

	// ExitTuple is called when exiting the Tuple production.
	ExitTuple(c *TupleContext)

	// ExitEqualityComparison is called when exiting the EqualityComparison production.
	ExitEqualityComparison(c *EqualityComparisonContext)

	// ExitMetaType is called when exiting the MetaType production.
	ExitMetaType(c *MetaTypeContext)

	// ExitAssignOp is called when exiting the assignOp production.
	ExitAssignOp(c *AssignOpContext)

	// ExitTupleExpression is called when exiting the tupleExpression production.
	ExitTupleExpression(c *TupleExpressionContext)

	// ExitInlineArrayExpression is called when exiting the inlineArrayExpression production.
	ExitInlineArrayExpression(c *InlineArrayExpressionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralWithSubDenomination is called when exiting the literalWithSubDenomination production.
	ExitLiteralWithSubDenomination(c *LiteralWithSubDenominationContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitHexStringLiteral is called when exiting the hexStringLiteral production.
	ExitHexStringLiteral(c *HexStringLiteralContext)

	// ExitUnicodeStringLiteral is called when exiting the unicodeStringLiteral production.
	ExitUnicodeStringLiteral(c *UnicodeStringLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitUncheckedBlock is called when exiting the uncheckedBlock production.
	ExitUncheckedBlock(c *UncheckedBlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitEmitStatement is called when exiting the emitStatement production.
	ExitEmitStatement(c *EmitStatementContext)

	// ExitRevertStatement is called when exiting the revertStatement production.
	ExitRevertStatement(c *RevertStatementContext)

	// ExitAssemblyStatement is called when exiting the assemblyStatement production.
	ExitAssemblyStatement(c *AssemblyStatementContext)

	// ExitAssemblyFlags is called when exiting the assemblyFlags production.
	ExitAssemblyFlags(c *AssemblyFlagsContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclarationTuple is called when exiting the variableDeclarationTuple production.
	ExitVariableDeclarationTuple(c *VariableDeclarationTupleContext)

	// ExitVariableDeclarationStatement is called when exiting the variableDeclarationStatement production.
	ExitVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitMappingType is called when exiting the mappingType production.
	ExitMappingType(c *MappingTypeContext)

	// ExitMappingKeyType is called when exiting the mappingKeyType production.
	ExitMappingKeyType(c *MappingKeyTypeContext)

	// ExitYulStatement is called when exiting the yulStatement production.
	ExitYulStatement(c *YulStatementContext)

	// ExitYulBlock is called when exiting the yulBlock production.
	ExitYulBlock(c *YulBlockContext)

	// ExitYulVariableDeclaration is called when exiting the yulVariableDeclaration production.
	ExitYulVariableDeclaration(c *YulVariableDeclarationContext)

	// ExitYulAssignment is called when exiting the yulAssignment production.
	ExitYulAssignment(c *YulAssignmentContext)

	// ExitYulIfStatement is called when exiting the yulIfStatement production.
	ExitYulIfStatement(c *YulIfStatementContext)

	// ExitYulForStatement is called when exiting the yulForStatement production.
	ExitYulForStatement(c *YulForStatementContext)

	// ExitYulSwitchCase is called when exiting the yulSwitchCase production.
	ExitYulSwitchCase(c *YulSwitchCaseContext)

	// ExitYulSwitchStatement is called when exiting the yulSwitchStatement production.
	ExitYulSwitchStatement(c *YulSwitchStatementContext)

	// ExitYulFunctionDefinition is called when exiting the yulFunctionDefinition production.
	ExitYulFunctionDefinition(c *YulFunctionDefinitionContext)

	// ExitYulPath is called when exiting the yulPath production.
	ExitYulPath(c *YulPathContext)

	// ExitYulFunctionCall is called when exiting the yulFunctionCall production.
	ExitYulFunctionCall(c *YulFunctionCallContext)

	// ExitYulBoolean is called when exiting the yulBoolean production.
	ExitYulBoolean(c *YulBooleanContext)

	// ExitYulLiteral is called when exiting the yulLiteral production.
	ExitYulLiteral(c *YulLiteralContext)

	// ExitYulExpression is called when exiting the yulExpression production.
	ExitYulExpression(c *YulExpressionContext)
}
