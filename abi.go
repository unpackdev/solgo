package solgo

import (
	"strings"

	"github.com/txpull/solgo/parser"
)

type AbiTreeShapeListener struct {
	*parser.BaseSolidityParserListener
	abi []map[string]interface{}
}

func NewAbiTreeShapeListener() *AbiTreeShapeListener {
	return &AbiTreeShapeListener{
		abi: make([]map[string]interface{}, 0),
	}
}

func (l *AbiTreeShapeListener) EnterConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	constructor := make(map[string]interface{})
	constructor["type"] = "constructor"

	// Extract input parameters
	inputs := make([]map[string]string, 0)
	if ctx.GetArguments() != nil {
		for _, paramCtx := range ctx.GetArguments().AllParameterDeclaration() {
			param := make(map[string]string)
			if paramCtx.Identifier() != nil {
				param["name"] = paramCtx.Identifier().GetText()
			}
			param["type"] = paramCtx.TypeName().GetText()
			param["internalType"] = paramCtx.TypeName().GetText()
			inputs = append(inputs, param)
		}
	}
	constructor["inputs"] = inputs

	l.abi = append(l.abi, constructor)
}

func (l *AbiTreeShapeListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	function := make(map[string]interface{})

	function["name"] = ctx.Identifier().GetText()

	// Extract input parameters
	inputs := make([]map[string]string, 0)
	if ctx.GetArguments() != nil {
		for _, paramCtx := range ctx.GetArguments().AllParameterDeclaration() {
			param := make(map[string]string)
			if paramCtx.Identifier() != nil {
				param["name"] = paramCtx.Identifier().GetText()
			}
			param["type"] = paramCtx.TypeName().GetText()
			param["internalType"] = paramCtx.TypeName().GetText()
			inputs = append(inputs, param)
		}
	}
	function["inputs"] = inputs

	// Extract output parameters
	outputs := make([]map[string]string, 0)
	if ctx.GetReturnParameters() != nil {
		for _, paramCtx := range ctx.GetReturnParameters().GetParameters() {
			param := make(map[string]string)
			if paramCtx.Identifier() != nil {
				param["name"] = paramCtx.Identifier().GetText()
			}
			param["type"] = paramCtx.TypeName().GetText()
			outputs = append(outputs, param)
		}
	}
	function["outputs"] = outputs

	// Extract stateMutability
	if ctx.AllStateMutability() != nil {
		for _, stateMutabilityCtx := range ctx.AllStateMutability() {
			if stateMutabilityCtx != nil {
				if stateMutabilityCtx.Payable() != nil {
					function["payable"] = true
				}
				function["stateMutability"] = stateMutabilityCtx.GetText()
			}
		}
	}

	l.abi = append(l.abi, function)
}

func (l *AbiTreeShapeListener) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	// Maps can be returned in internal functions and external contract functions but not
	// in contract public functions.
	if strings.Contains(ctx.TypeName().GetText(), "mapping") {
		return
	}

	variable := make(map[string]interface{})
	variable["name"] = ctx.Identifier().GetText()
	variable["type"] = "function"
	variable["stateMutability"] = "view"

	inputs := make([]map[string]string, 0)
	variable["inputs"] = inputs

	outputs := make([]map[string]string, 0)

	// Some bug that needs to be fixed, for now this is the fix...
	switch ctx.TypeName().GetText() {
	case "addresspayable":
		param := make(map[string]string)
		param["name"] = ""
		param["type"] = "address"
		param["internalType"] = "address"
		outputs = append(outputs, param)
	default:
		param := make(map[string]string)
		param["name"] = ""
		param["type"] = ctx.TypeName().GetText()
		param["internalType"] = ctx.TypeName().GetText()
		outputs = append(outputs, param)
	}

	variable["outputs"] = outputs

	l.abi = append(l.abi, variable)
}

func (l *AbiTreeShapeListener) EnterEventDefinition(ctx *parser.EventDefinitionContext) {
	event := make(map[string]interface{})
	event["name"] = ctx.Identifier().GetText()
	event["type"] = "event"
	if ctx.Anonymous() != nil {
		event["anonymous"] = true
	} else {
		event["anonymous"] = false
	}

	inputs := make([]map[string]interface{}, 0)
	if ctx.GetParameters() != nil {
		for _, paramCtx := range ctx.GetParameters() {
			param := make(map[string]interface{})
			if paramCtx.Identifier() != nil {
				param["name"] = paramCtx.Identifier().GetText()
			}
			param["type"] = paramCtx.TypeName().GetText()
			param["internalType"] = paramCtx.TypeName().GetText()

			if paramCtx.Indexed() != nil && paramCtx.Indexed().GetText() == "indexed" {
				param["indexed"] = true
			} else {
				param["indexed"] = false
			}

			inputs = append(inputs, param)
		}

	}
	event["inputs"] = inputs

	l.abi = append(l.abi, event)
}

func (l *AbiTreeShapeListener) GetABI() []map[string]interface{} {
	return l.abi
}

func (s *SolGo) GetABI() ([]map[string]interface{}, error) {
	listener, err := s.GetListener(ListenerAbiTreeShape)
	if err != nil {
		return nil, err
	}

	return listener.(*AbiTreeShapeListener).GetABI(), nil
}
