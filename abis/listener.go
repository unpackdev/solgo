package abis

import (
	"fmt"

	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// AbiListener is a listener for the Solidity parser that constructs an ABI
// as it walks the parse tree. It embeds the BaseSolidityParserListener and uses an AbiParser
// to construct the ABI.
type AbiListener struct {
	*parser.BaseSolidityParserListener            // Base listener that does nothing, to be extended by AbiListener
	parser                             *AbiParser // The parser that constructs the ABI
}

// NewAbiListener creates a new AbiListener with a new AbiParser.
// It returns a pointer to the newly created AbiListener.
func NewAbiListener() *AbiListener {
	return &AbiListener{parser: &AbiParser{
		abi:               ABI{},
		definedStructs:    make(map[string]MethodIO),
		definedEnums:      make(map[string]bool), // map to keep track of defined enum types
		definedInterfaces: make(map[string]bool),
		definedLibraries:  make(map[string]bool),
		definedContracts:  make(map[string]ContractDefinition),
	}}
}

// EnterContractDefinition is called when the parser enters a contract definition.
// It extracts the contract name from the context and sets it to the contractName field.
// Later on we use contract name to define internalType of the tuple/struct type.
func (l *AbiListener) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	l.parser.contractName = ctx.Identifier().GetText() // get the contract name
	fmt.Println("contract name: ", l.parser.contractName)

	if _, ok := l.parser.definedContracts[l.parser.contractName]; !ok {
		l.parser.definedContracts[l.parser.contractName] = ContractDefinition{
			IsAbstract: func() bool {
				return ctx.Abstract() != nil
			}(),
		}
	}
}

func (l *AbiListener) EnterInterfaceDefinition(ctx *parser.InterfaceDefinitionContext) {
	if _, ok := l.parser.definedInterfaces[ctx.Identifier().GetText()]; !ok {
		l.parser.definedInterfaces[ctx.Identifier().GetText()] = true
	}
}

func (l *AbiListener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	if _, ok := l.parser.definedLibraries[ctx.Identifier().GetText()]; !ok {
		l.parser.definedLibraries[ctx.Identifier().GetText()] = true
	}
}

// EnterStateVariableDeclaration is called when the parser enters a state variable declaration.
// It injects the state variable into the ABI.
func (l *AbiListener) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	if err := l.parser.InjectStateVariable(ctx); err != nil {
		zap.L().Error(
			"failed to inject state variable",
			zap.Error(err),
		)
	}
}

// EnterConstructorDefinition is called when the parser enters a constructor definition.
// It injects the constructor into the ABI.
func (l *AbiListener) EnterConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	if err := l.parser.InjectConstructor(ctx); err != nil {
		zap.L().Error(
			"failed to inject constructor",
			zap.Error(err),
		)
	}
}

// EnterFunctionDefinition is called when the parser enters a function definition.
// It injects the function into the ABI.
func (l *AbiListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	if err := l.parser.InjectFunction(ctx); err != nil {
		zap.L().Error(
			"failed to inject event",
			zap.Error(err),
		)
	}
}

// EnterEventDefinition is called when the parser enters an event definition.
// It injects the event into the ABI.
func (l *AbiListener) EnterEventDefinition(ctx *parser.EventDefinitionContext) {
	if err := l.parser.InjectEvent(ctx); err != nil {
		zap.L().Error(
			"failed to inject event",
			zap.Error(err),
		)
	}
}

// EnterErrorDefinition is called when the parser enters an error definition.
// It injects the error into the ABI.
func (l *AbiListener) EnterErrorDefinition(ctx *parser.ErrorDefinitionContext) {
	if err := l.parser.InjectError(ctx); err != nil {
		zap.L().Error(
			"failed to inject error",
			zap.Error(err),
		)
	}
}

// EnterStructDefinition is called when the parser enters a struct definition.
// It appends the struct to the ABI for future resolution.
func (l *AbiListener) EnterStructDefinition(ctx *parser.StructDefinitionContext) {
	if err := l.parser.AppendStruct(ctx); err != nil {
		zap.L().Error(
			"failed to append struct",
			zap.Error(err),
		)
	}
}

// ExitStructDefinition is called when the parser exits a struct definition.
// It resolves the components of the struct in the ABI.
func (l *AbiListener) ExitStructDefinition(ctx *parser.StructDefinitionContext) {
	if err := l.parser.ResolveStruct(ctx); err != nil {
		zap.L().Error(
			"failed to resolve struct components",
			zap.Error(err),
		)
	}
}

// EnterEnumDefinition is a method of the AbiParser struct that is called when the parser encounters an enum definition in the Solidity code.
// It takes a context of type EnumDefinitionContext, which contains the information about the enum definition in the code.
// The method uses this context to extract the enum's name and its values, and stores them in the AbiParser's map of defined enums for future reference.
// This allows the parser to recognize the enum type when it is used later in the code.
func (p *AbiListener) EnterEnumDefinition(ctx *parser.EnumDefinitionContext) {
	if ctx.AllIdentifier() != nil {
		for _, id := range ctx.AllIdentifier() {
			p.parser.definedEnums[id.GetText()] = true
		}
	}
}

// EnterFallbackFunctionDefinition is called when the parser enters a fallback function definition.
// It injects the fallback function into the ABI.
func (l *AbiListener) EnterFallbackFunctionDefinition(ctx *parser.FallbackFunctionDefinitionContext) {
	if err := l.parser.InjectFallback(ctx); err != nil {
		zap.L().Error(
			"failed to inject fallback",
			zap.Error(err),
		)
	}
}

// EnterReceiveFunctionDefinition is called when the parser enters a receive function definition.
// It injects the receive function into the ABI.
func (l *AbiListener) EnterReceiveFunctionDefinition(ctx *parser.ReceiveFunctionDefinitionContext) {
	if err := l.parser.InjectReceive(ctx); err != nil {
		zap.L().Error(
			"failed to inject receive",
			zap.Error(err),
		)
	}
}

// GetParser returns the AbiParser associated with the AbiListener.
func (l *AbiListener) GetParser() *AbiParser {
	return l.parser
}
