package ast

import (
	"strings"
	"time"

	"github.com/txpull/solgo/parser"
)

// ASTBuilder is responsible for constructing the AST from the parser output.
type ASTBuilder struct {
	*parser.BaseSolidityParserListener // Embed the base listener
	// Add any additional state you need here.
	currentContract  *ContractNode
	currentInterface *InterfaceNode
	currentFunction  *FunctionNode
	astRoot          *RootNode
	errors           []error
	parseTime        time.Time
	pragmas          [][]string
}

func NewAstBuilder() *ASTBuilder {
	return &ASTBuilder{
		pragmas: make([][]string, 0),
	}
}

// EnterSourceUnit is called when the parser enters a source unit (i.e., a file).
func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	b.currentContract = nil

	b.astRoot = &RootNode{
		Contracts:  make([]*ContractNode, 0),
		Interfaces: make([]*InterfaceNode, 0),
	}

	b.errors = nil
	b.parseTime = time.Now()

	// QUESTION: Do we want to do anything else besides appending the pragmas?
	for _, pragma := range ctx.AllPragmaDirective() {
		for _, token := range pragma.AllPragmaToken() {
			pragmas := make([]string, 0)
			pragmas = append(pragmas, strings.TrimSpace(token.GetText()))
			b.pragmas = append(b.pragmas, pragmas)
		}
	}
}

// EnterContractDefinition is called when the parser enters a contract definition.
func (b *ASTBuilder) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	// Create a new ContractNode and set it as the current contract.
	b.currentContract = &ContractNode{
		Name:           ctx.Identifier().GetText(), // Get the contract name from the context
		StateVariables: make([]*StateVariableNode, 0),
		Kind:           "contract",
	}

	// Handle contract inheritance
	if ctx.InheritanceSpecifierList() != nil {
		for _, inheritance := range ctx.InheritanceSpecifierList().AllInheritanceSpecifier() {
			b.currentContract.Inherits = append(b.currentContract.Inherits, inheritance.GetText())
		}
	}

	// Handle contract kind (contract, interface, library)
	if ctx.Abstract() != nil {
		b.currentContract.Kind = "abstract"
	}

	// Add the contract to the root node.
	b.astRoot.Contracts = append(b.astRoot.Contracts, b.currentContract)
}

// ExitContractDefinition is called when the parser exits a contract definition.
func (b *ASTBuilder) ExitContractDefinition(ctx *parser.ContractDefinitionContext) {
	b.currentContract = nil
}

// EnterInterfaceDefinition is called when the parser enters an interface definition.
func (b *ASTBuilder) EnterInterfaceDefinition(ctx *parser.InterfaceDefinitionContext) {
	// Create a new InterfaceNode and set it as the current interface.
	currentInterface := &InterfaceNode{
		Name: ctx.Identifier().GetText(),
	}

	b.astRoot.Interfaces = append(b.astRoot.Interfaces, currentInterface)
}

func (b *ASTBuilder) ExitInterfaceDefinition(ctx *parser.InterfaceDefinitionContext) {
	b.currentInterface = nil
}

func (b *ASTBuilder) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	// Create a new StateVariableNode.
	variable := &StateVariableNode{
		Name: func() string {
			if ctx.Identifier() != nil {
				return ctx.Identifier().GetText()
			}
			return ""
		}(),
		Type: ctx.GetType_().GetText(),
	}

	// Determine the visibility of the state variable.
	if ctx.GetVisibilitySet() {
		if ctx.AllInternal() != nil {
			variable.Visibility = "internal"
		} else if ctx.AllPrivate() != nil {
			variable.Visibility = "private"
		} else if ctx.AllPublic() != nil {
			variable.Visibility = "public"
		}
	}

	// Determine if the state variable is constant.
	if ctx.GetConstantnessSet() {
		for _, constant := range ctx.AllConstant() {
			if constant.GetText() == "constant" {
				variable.IsConstant = true
			}
		}
	}

	// Determine if the state variable is immutable.
	if ctx.AllImmutable() != nil {
		for _, modifier := range ctx.AllImmutable() {
			if modifier.GetText() == "immutable" {
				variable.IsImmutable = true
			}
		}
	}

	// Get the initial value of the state variable, if any.
	if initialValue := ctx.GetInitialValue(); initialValue != nil {
		variable.InitialValue = initialValue.GetText()
	}

	// Add the variable to the current contract.
	b.currentContract.StateVariables = append(b.currentContract.StateVariables, variable)
}

// EnterConstructorDefinition is called when the parser enters a constructor definition.
func (b *ASTBuilder) EnterConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	constructor := &ConstructorNode{
		Parameters: make([]*VariableNode, 0),
		Body:       make([]*StatementNode, 0),
	}

	if arguments := ctx.GetArguments(); arguments != nil {
		for _, parameterCtx := range arguments.AllParameterDeclaration() {
			constructor.Parameters = append(constructor.Parameters, &VariableNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type: parameterCtx.TypeName().GetText(),
			})
		}
	}

	if body := ctx.GetBody(); body != nil {
		for _, statementCtx := range body.AllStatement() {
			constructor.Body = append(constructor.Body, &StatementNode{
				Raw:     b.getTextSliceWithOriginalFormatting(statementCtx),
				TextRaw: statementCtx.GetText(),
			})
		}
	}

	b.currentContract.Constructor = constructor
}

func (b *ASTBuilder) EnterStructDefinition(ctx *parser.StructDefinitionContext) {
	// Create a new StructNode and set its name
	structNode := &StructNode{
		Name: ctx.GetName().GetText(),
	}

	// Handle struct members
	for _, memberCtx := range ctx.AllStructMember() {
		structNode.Members = append(structNode.Members, &StructMemberNode{
			Name: func() string {
				if memberCtx.Identifier() != nil {
					return memberCtx.Identifier().GetText()
				}
				return ""
			}(),
			Type: memberCtx.GetType_().GetText(),
		})
	}

	// Add the struct node to the current contract
	b.currentContract.Structs = append(b.currentContract.Structs, structNode)
}

func (b *ASTBuilder) EnterErrorDefinition(ctx *parser.ErrorDefinitionContext) {
	// Create a new ErrorNode and set its name.
	errorNode := &ErrorNode{
		Name:   ctx.GetName().GetText(),
		Values: make([]*ErrorValueNode, 0),
	}

	// Handle error values
	for _, errorParamCtx := range ctx.GetParameters() {
		errorValue := &ErrorValueNode{
			Name: func() string {
				if errorParamCtx.Identifier() != nil {
					return errorParamCtx.Identifier().GetText()
				}
				return ""
			}(),
			Type: errorParamCtx.GetType_().GetText(),
			Code: 0,
		}
		errorNode.Values = append(errorNode.Values, errorValue)
	}

	// Add the error node to the current contract.
	b.currentContract.Errors = append(b.currentContract.Errors, errorNode)
}

// EnterFunctionDefinition is called when the parser enters a function definition.
func (b *ASTBuilder) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	currentFunction := b.CreateFunction(ctx)

	if b.currentContract != nil {
		b.currentContract.Functions = append(b.currentContract.Functions, currentFunction)
	} else if b.currentInterface != nil {
		b.currentInterface.Functions = append(b.currentInterface.Functions, currentFunction)
	}
}

func (b *ASTBuilder) EnterFallbackFunctionDefinition(ctx *parser.FallbackFunctionDefinitionContext) {
	// Create a new FunctionNode for the fallback function.
	fallbackFunction := &FunctionNode{
		Name:             "fallback",
		Parameters:       make([]*VariableNode, 0),
		ReturnParameters: make([]*VariableNode, 0),
		Body:             make([]*StatementNode, 0),
		Visibility:       make([]string, 0),
		Mutability:       make([]string, 0),
		Modifiers:        make([]string, 0),
		Overrides:        false,
		IsVirtual:        false,
		IsFallback:       true,
	}

	// Handle virtual modifier
	if ctx.GetVirtualSet() {
		fallbackFunction.IsVirtual = true
	}

	if ctx.GetVisibilitySet() {
		for _, externalCtx := range ctx.AllExternal() {
			visibility := externalCtx.GetText()
			fallbackFunction.Visibility = append(fallbackFunction.Visibility, visibility)
		}
	}

	if ctx.GetMutabilitySet() {
		for _, mutabilityCtx := range ctx.AllStateMutability() {
			mutability := mutabilityCtx.GetText()
			fallbackFunction.Mutability = append(fallbackFunction.Mutability, mutability)
		}
	}

	if len(fallbackFunction.Mutability) == 0 {
		fallbackFunction.Mutability = append(fallbackFunction.Mutability, "nonpayable")
	}

	// Handle override specifier
	if ctx.GetOverrideSpecifierSet() {
		for _, overrideCtx := range ctx.AllOverrideSpecifier() {
			fallbackFunction.Overrides = !overrideCtx.IsEmpty()
		}
	}

	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := modifierCtx.GetText()
		fallbackFunction.Modifiers = append(fallbackFunction.Modifiers, modifier)
	}

	// Handle function parameters
	if parameters := ctx.AllParameterList(); parameters != nil {
		for _, parameterCtx := range parameters {
			for _, param := range parameterCtx.AllParameterDeclaration() {
				fallbackFunction.Parameters = append(fallbackFunction.Parameters, &VariableNode{
					Name: func() string {
						if param.Identifier() != nil {
							return param.Identifier().GetText()
						}
						return ""
					}(),
					Type: param.TypeName().GetText(),
				})
			}
		}
	}

	// Handle function body
	if body := ctx.GetBody(); body != nil {
		for _, statementCtx := range body.AllStatement() {
			statement := &StatementNode{
				Raw:     b.getTextSliceWithOriginalFormatting(statementCtx),
				TextRaw: statementCtx.GetText(),
			}
			fallbackFunction.Body = append(fallbackFunction.Body, statement)
		}
	}

	// Add the fallback function to the current contract.
	b.currentContract.Functions = append(b.currentContract.Functions, fallbackFunction)
}

func (b *ASTBuilder) EnterReceiveFunctionDefinition(ctx *parser.ReceiveFunctionDefinitionContext) {
	// Create a new FunctionNode and set it as the current function.
	receiveFn := &FunctionNode{
		Name:             "receive",
		Parameters:       make([]*VariableNode, 0),
		ReturnParameters: make([]*VariableNode, 0),
		Body:             make([]*StatementNode, 0),
		Visibility:       make([]string, 0),
		Mutability:       make([]string, 0),
		Modifiers:        make([]string, 0),
		Overrides:        false,
		IsVirtual:        false,
		IsReceive:        true,
	}

	// Handle virtual modifier
	if ctx.GetVirtualSet() {
		receiveFn.IsVirtual = true
	}

	// Handle override specifier
	if ctx.GetOverrideSpecifierSet() {
		for _, overrideCtx := range ctx.AllOverrideSpecifier() {
			receiveFn.Overrides = !overrideCtx.IsEmpty()
		}
	}

	// Handle function body
	if body := ctx.GetBody(); body != nil {
		for _, statementCtx := range body.AllStatement() {
			statement := &StatementNode{
				Raw:     b.getTextSliceWithOriginalFormatting(statementCtx),
				TextRaw: statementCtx.GetText(),
			}
			receiveFn.Body = append(receiveFn.Body, statement)
		}
	}

	// Add the receive function to the current contract.
	b.currentContract.Functions = append(b.currentContract.Functions, receiveFn)
}

func (b *ASTBuilder) EnterEventDefinition(ctx *parser.EventDefinitionContext) {
	event := &EventNode{
		Name:      ctx.GetName().GetText(),
		Anonymous: ctx.Anonymous() != nil,
	}

	if ctx.AllEventParameter() != nil {
		event.Parameters = make([]*EventParameterNode, 0)

		for _, parameterCtx := range ctx.AllEventParameter() {
			event.Parameters = append(event.Parameters, &EventParameterNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:    parameterCtx.GetType_().GetText(),
				Indexed: parameterCtx.Indexed() != nil,
			})
		}
	}

	b.currentContract.Events = append(b.currentContract.Events, event)
}

// GetTree returns the root of the AST.
func (b *ASTBuilder) GetTree() Node {
	return b.astRoot
}

// GetPragmas returns the pragmas found in the source unit.
func (b *ASTBuilder) GetPragmas() [][]string {
	return b.pragmas
}

// GetErrors returns the errors found during parsing.
func (b *ASTBuilder) GetErrors() []error {
	return b.errors
}

// GetParseTime returns the time it took to parse the source unit.
func (b *ASTBuilder) GetParseTime() time.Duration {
	return time.Since(b.parseTime)
}
