package solgo

import (
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// AbiListener is a listener for the Solidity parser that constructs an ABI
// as it walks the parse tree.
type AbiListener struct {
	*parser.BaseSolidityParserListener            // Base listener that does nothing, to be extended by AbiListener
	parser                             *AbiParser // The parser that constructs the ABI
}

// NewAbiListener creates a new AbiListener with a new AbiParser.
func NewAbiListener() *AbiListener {
	return &AbiListener{parser: &AbiParser{abi: ABI{}}}
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
