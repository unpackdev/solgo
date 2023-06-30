package solgo

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/txpull/solgo/parser"
)

// MethodIO represents an input or output parameter of a contract method or event.
type MethodIO struct {
	Indexed      bool                `json:"indexed,omitempty"`    // Used only by the events
	InternalType string              `json:"internalType"`         // The internal Solidity type of the parameter
	Name         string              `json:"name"`                 // The name of the parameter
	Type         string              `json:"type"`                 // The type of the parameter
	Components   []map[string]string `json:"components,omitempty"` // Components of the parameter, if it's a struct or tuple type
}

// IMethod is an interface that represents a contract method, event, or constructor.
type IMethod interface{}

// MethodConstructor represents a contract constructor.
type MethodConstructor struct {
	Inputs  []MethodIO `json:"inputs"`            // The input parameters of the constructor
	Type    string     `json:"type"`              // The type of the method (always "constructor" for constructors)
	Outputs []MethodIO `json:"outputs,omitempty"` // The output parameters of the constructor (always empty for constructors)
}

// MethodEvent represents a contract event.
type MethodEvent struct {
	Anonymous bool       `json:"anonymous"`      // Whether the event is anonymous
	Inputs    []MethodIO `json:"inputs"`         // The input parameters of the event
	Name      string     `json:"name,omitempty"` // The name of the event
	Type      string     `json:"type"`           // The type of the method (always "event" for events)
}

// Method represents a contract function.
type Method struct {
	Inputs          []MethodIO `json:"inputs"`          // The input parameters of the function
	Outputs         []MethodIO `json:"outputs"`         // The output parameters of the function
	Name            string     `json:"name"`            // The name of the function
	Type            string     `json:"type"`            // The type of the method (always "function" for functions)
	StateMutability string     `json:"stateMutability"` // The state mutability of the function (pure, view, nonpayable, payable)
}

// MethodVariable represents a contract state variable.
type MethodVariable struct {
	Inputs          []MethodIO `json:"inputs"`            // The input parameters of the variable (always empty for variables)
	Outputs         []MethodIO `json:"outputs,omitempty"` // The output parameters of the variable (always contains one element representing the variable itself)
	Name            string     `json:"name"`              // The name of the variable
	Type            string     `json:"type"`              // The type of the method (always "function" for variables)
	StateMutability string     `json:"stateMutability"`   // The state mutability of the variable (always "view" for variables)
}

// MethodFallbackOrReceive represents a contract fallback or receive function.
type MethodFallbackOrReceive struct {
	Type            string `json:"type"`                      // The type of the method (either "fallback" or "receive")
	StateMutability string `json:"stateMutability,omitempty"` // The state mutability of the function (nonpayable for fallback functions, payable for receive functions)
}

// ABI represents a contract ABI, which is a list of contract methods, events, and constructors.
type ABI []IMethod

// AbiParser is a parser that can parse a Solidity contract ABI
// and convert it into an ABI object that can be easily manipulated.
type AbiParser struct {
	abi ABI
}

// InjectConstructor injects a constructor definition into the ABI.
// It takes a ConstructorDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectConstructor(ctx *parser.ConstructorDefinitionContext) error {
	inputs := make([]MethodIO, 0)
	if ctx.GetArguments() != nil {
		for _, paramCtx := range ctx.GetArguments().AllParameterDeclaration() {
			inputs = append(inputs, MethodIO{
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}
	}

	p.abi = append(p.abi, MethodConstructor{
		Type:    "constructor",
		Inputs:  inputs,
		Outputs: make([]MethodIO, 0),
	})
	return nil
}

// InjectEvent injects an event definition into the ABI.
// It takes an EventDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectEvent(ctx *parser.EventDefinitionContext) error {
	inputs := make([]MethodIO, 0)
	if ctx.GetParameters() != nil {
		for _, paramCtx := range ctx.GetParameters() {
			inputs = append(inputs, MethodIO{
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				Indexed:      paramCtx.Indexed() != nil && paramCtx.Indexed().GetText() == "indexed",
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}

	}

	p.abi = append(p.abi, MethodEvent{
		Anonymous: ctx.Anonymous() != nil && ctx.Anonymous().GetText() == "anonymous",
		Inputs:    inputs,
		Name:      ctx.Identifier().GetText(),
		Type:      "event",
	})
	return nil
}

// InjectFunction injects a function definition into the ABI.
// It takes a FunctionDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectFunction(ctx *parser.FunctionDefinitionContext) error {
	inputs := make([]MethodIO, 0)
	if ctx.GetArguments() != nil {
		for _, paramCtx := range ctx.GetArguments().AllParameterDeclaration() {
			inputs = append(inputs, MethodIO{
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}

	}

	outputs := make([]MethodIO, 0)
	if ctx.GetReturnParameters() != nil {
		for _, paramCtx := range ctx.GetReturnParameters().GetParameters() {
			outputs = append(outputs, MethodIO{
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}
	}

	// Default state mutability is nonpayable
	stateMutability := "nonpayable"

	if ctx.AllStateMutability() != nil && len(ctx.AllStateMutability()) > 0 {
		for _, stateMutabilityCtx := range ctx.AllStateMutability() {
			if stateMutabilityCtx != nil {
				stateMutability = stateMutabilityCtx.GetText()
			}
		}
	}

	p.abi = append(p.abi, Method{
		Inputs:          inputs,
		Outputs:         outputs,
		Name:            ctx.Identifier().GetText(),
		Type:            "function",
		StateMutability: stateMutability,
	})
	return nil
}

// InjectStateVariable injects a state variable definition into the ABI.
// It takes a StateVariableDeclarationContext (from the parser) as input.
func (p *AbiParser) InjectStateVariable(ctx *parser.StateVariableDeclarationContext) error {
	if isMappingType(ctx.TypeName().GetText()) {
		inputs := make([]MethodIO, 0)
		outputs := make([]MethodIO, 0)
		matched, inputList, outputList := parseMappingType(ctx.TypeName().GetText())

		if matched {
			for _, input := range inputList {
				inputs = append(inputs, MethodIO{
					Type:         normalizeTypeName(input),
					InternalType: normalizeTypeName(input),
				})
			}

			for _, output := range outputList {
				outputs = append(outputs, MethodIO{
					Type:         normalizeTypeName(output),
					InternalType: normalizeTypeName(output),
				})
			}
		}

		p.abi = append(p.abi, MethodVariable{
			Inputs:          inputs,
			Outputs:         outputs,
			Name:            ctx.Identifier().GetText(),
			StateMutability: "view",
			Type:            "function",
		})
		return nil
	}

	p.abi = append(p.abi, MethodVariable{
		Outputs: []MethodIO{
			{
				Type:         normalizeTypeName(ctx.TypeName().GetText()),
				InternalType: normalizeTypeName(ctx.TypeName().GetText()),
			},
		},
		Name:            ctx.Identifier().GetText(),
		StateMutability: "view",
		Type:            "function",
	})
	return nil
}

// InjectError injects an error definition into the ABI.
// It takes an ErrorDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectError(ctx *parser.ErrorDefinitionContext) error {
	inputs := make([]MethodIO, 0)
	if ctx.GetParameters() != nil {
		for _, paramCtx := range ctx.GetParameters() {
			inputs = append(inputs, MethodIO{
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}

	}

	p.abi = append(p.abi, MethodVariable{
		Inputs:          inputs,
		Name:            ctx.Identifier().GetText(),
		StateMutability: "view",
		Type:            "error",
	})
	return nil
}

// InjectModifier injects a modifier definition into the ABI.
// It takes a ModifierDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectModifier(ctx *parser.ModifierDefinitionContext) error {
	return errors.New("modifier injection is not yet supported")
}

// InjectEnum injects an enum definition into the ABI.
// It takes an EnumDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectEnum(ctx *parser.EnumDefinitionContext) error {
	return errors.New("enum injection is not yet supported")
}

// InjectStruct injects a struct definition into the ABI.
// It takes a StructDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectStruct(ctx *parser.StructDefinitionContext) error {
	return errors.New("struct injection is not yet supported")
}

// InjectContract injects a contract definition into the ABI.
// It takes a ContractDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectContract(ctx *parser.ContractDefinitionContext) error {
	return errors.New("contract injection is not yet supported")
}

// InjectFallback injects a fallback function definition into the ABI.
// It takes a FallbackFunctionDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectFallback(ctx *parser.FallbackFunctionDefinitionContext) error {
	stateMutability := "nonpayable"

	if ctx.AllStateMutability() != nil && len(ctx.AllStateMutability()) > 0 {
		for _, stateMutabilityCtx := range ctx.AllStateMutability() {
			if stateMutabilityCtx != nil {
				stateMutability = stateMutabilityCtx.GetText()
			}
		}
	}

	p.abi = append(p.abi, MethodFallbackOrReceive{
		Type:            "fallback",
		StateMutability: stateMutability,
	})
	return nil
}

// InjectReceive injects a receive function definition into the ABI.
// It takes a ReceiveFunctionDefinitionContext (from the parser) as input.
func (p *AbiParser) InjectReceive(ctx *parser.ReceiveFunctionDefinitionContext) error {
	p.abi = append(p.abi, MethodFallbackOrReceive{
		Type:            "receive",
		StateMutability: "payable",
	})
	return nil
}

// ToJSON converts the ABI object into a JSON string.
func (p *AbiParser) ToJSON() (string, error) {
	abiJSON, err := json.Marshal(p.abi)
	if err != nil {
		return "", err
	}

	return string(abiJSON), nil
}

// ToABI converts the ABI object into an ethereum/go-ethereum ABI object.
func (p *AbiParser) ToABI() (*abi.ABI, error) {
	jsonData, err := p.ToJSON()
	if err != nil {
		return nil, err
	}

	toReturn, err := abi.JSON(strings.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	return &toReturn, nil
}
