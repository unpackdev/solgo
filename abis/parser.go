package abis

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// AbiParser is a parser that can parse a Solidity contract ABI
// and convert it into an ABI object that can be easily manipulated.
type AbiParser struct {
	abi               ABI
	contractName      string
	definedStructs    map[string]MethodIO
	definedEnums      map[string]bool // map to keep track of defined enum types
	definedInterfaces map[string]bool
	definedLibraries  map[string]bool
	definedContracts  map[string]ContractDefinition
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
				Type: func() string {
					if _, ok := p.definedInterfaces[paramCtx.TypeName().GetText()]; ok {
						return "address"
					}
					ntype, found := normalizeTypeNameWithStatus(paramCtx.TypeName().GetText())
					if !found {
						return "address"
					}

					return ntype
				}(),
				InternalType: func() string {
					if _, ok := p.definedInterfaces[paramCtx.TypeName().GetText()]; ok {
						return fmt.Sprintf("contract %s", paramCtx.TypeName().GetText())
					}
					ntype, found := normalizeTypeNameWithStatus(paramCtx.TypeName().GetText())
					if !found {
						return fmt.Sprintf("contract %s", paramCtx.TypeName().GetText())
					}
					return ntype
				}(),
			})
		}
	}

	constructor := MethodConstructor{
		Type:    "constructor",
		Inputs:  inputs,
		Outputs: make([]MethodIO, 0),
		StateMutability: func() string {
			if ctx.AllPayable() != nil {
				for _, payable := range ctx.AllPayable() {
					if payable.GetText() == "payable" {
						return "payable"
					}
				}
			}
			return "nonpayable"
		}(),
	}

	p.abi = append(p.abi, constructor)
	dcontract := p.definedContracts[p.contractName]
	dcontract.Constructor = constructor
	p.definedContracts[p.contractName] = dcontract

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
			argumentName := func() string {
				if paramCtx.Identifier() != nil {
					return paramCtx.Identifier().GetText()
				}
				return ""
			}()

			if isStructType(p.definedStructs, paramCtx.TypeName().GetText()) {
				nestedComponent, err := p.getStructComponents(paramCtx.TypeName().GetText())
				if err != nil {
					zap.L().Error(
						"Unsuported argument type",
						zap.String("type", paramCtx.TypeName().GetText()),
						zap.String("name", argumentName),
					)

					return fmt.Errorf(
						"unsupported function: '%s' argument name: '%s' type: '%s'",
						ctx.Identifier().GetText(),
						argumentName,
						paramCtx.TypeName().GetText(),
					)
				}

				inputs = append(inputs, nestedComponent)
				continue
			}

			inputs = append(inputs, MethodIO{
				Name:         argumentName,
				Type:         normalizeTypeName(paramCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(paramCtx.TypeName().GetText()),
			})
		}
	}

	outputs := make([]MethodIO, 0)
	if ctx.GetReturnParameters() != nil {
		for _, paramCtx := range ctx.GetReturnParameters().GetParameters() {
			argumentName := func() string {
				if paramCtx.Identifier() != nil {
					return paramCtx.Identifier().GetText()
				}
				return ""
			}()

			if isStructType(p.definedStructs, paramCtx.TypeName().GetText()) {
				// Checking if the parameter is a struct...
				nestedComponent, err := p.getStructComponents(paramCtx.TypeName().GetText())
				if err != nil {
					zap.L().Error(
						"Unsuported argument type",
						zap.String("type", paramCtx.TypeName().GetText()),
						zap.String("name", argumentName),
					)

					return fmt.Errorf(
						"unsupported function: '%s' argument name: '%s' type: '%s'",
						ctx.Identifier().GetText(),
						argumentName,
						paramCtx.TypeName().GetText(),
					)
				}

				outputs = append(outputs, nestedComponent)
				continue
			}

			outputs = append(outputs, MethodIO{
				Name:         argumentName,
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
	typeName := ctx.TypeName().GetText()

	if isMappingType(typeName) {
		inputs := make([]MethodIO, 0)
		outputs := make([]MethodIO, 0)
		matched, inputList, outputList := parseMappingType(typeName)

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
	} else if isEnumType(p.definedEnums, typeName) {
		p.abi = append(p.abi, MethodVariable{
			Inputs: make([]MethodIO, 0),
			Outputs: []MethodIO{
				{
					Type: "uint8", // enums are represented as uint8 in the ABI
					InternalType: fmt.Sprintf(
						"enum %s.%s",
						p.contractName,
						typeName,
					),
				},
			},
			Name:            ctx.Identifier().GetText(),
			StateMutability: "view",
			Type:            "function",
		})
		return nil
	} else if isContractType(p.definedContracts, typeName) {
		p.abi = append(p.abi, MethodVariable{
			Inputs: make([]MethodIO, 0),
			Outputs: []MethodIO{
				{
					Type: "uint8", // enums are represented as uint8 in the ABI
					InternalType: fmt.Sprintf(
						"contract %s",
						typeName,
					),
				},
			},
			Name:            ctx.Identifier().GetText(),
			StateMutability: "view",
			Type:            "function",
		})
		return nil
	} else if isInterfaceType(p.definedInterfaces, typeName) {
		p.abi = append(p.abi, MethodVariable{
			Inputs: make([]MethodIO, 0),
			Outputs: []MethodIO{
				{
					Type: "uint8", // enums are represented as uint8 in the ABI
					InternalType: fmt.Sprintf(
						"contract %s",
						typeName,
					),
				},
			},
			Name:            ctx.Identifier().GetText(),
			StateMutability: "view",
			Type:            "function",
		})
		return nil
	}

	p.abi = append(p.abi, MethodVariable{
		Inputs: make([]MethodIO, 0),
		Outputs: []MethodIO{
			{
				Type:         normalizeTypeName(typeName),
				InternalType: normalizeTypeName(typeName),
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

// AppendStruct injects a struct definition into internal struct mapping for future use by functions.
// Structs are not part of the ABI in meaning that you get a view function immediately
// without declaring it like in regular types. Instead, structs are used as input and output types
// for functions and only there they are visible.
// Current function will only store initial function definitions and because we have forward declarations,
// we will need to process additionally all structs after all declarations are processed to ensure
// nested structs are processed correctly.
func (p *AbiParser) AppendStruct(ctx *parser.StructDefinitionContext) error {
	structName := ctx.Identifier().GetText()

	components := make([]MethodIO, 0)

	if ctx.AllStructMember() != nil {
		for _, memberCtx := range ctx.AllStructMember() {
			components = append(components, MethodIO{
				Name: func() string {
					if memberCtx.Identifier() != nil {
						return memberCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:         normalizeTypeName(memberCtx.TypeName().GetText()),
				InternalType: normalizeTypeName(memberCtx.TypeName().GetText()),
			})
		}
	}

	p.definedStructs[structName] = MethodIO{
		Components:   components,
		Name:         structName,
		Type:         "tuple",
		InternalType: fmt.Sprintf("struct %s.%s", p.contractName, structName),
	}

	return nil
}

// ResolveStruct iterates over the defined structs in the AbiParser and resolves their components.
// If a component is of a struct type, it retrieves the components of the nested struct and updates the component's type to "tuple".
// The component's InternalType is also updated to reflect the struct's name and the contract it belongs to.
// If a struct component cannot be resolved, it logs a debug message and returns an error.
// This function is useful for resolving nested structs and should be called after all structs have been defined.
func (p *AbiParser) ResolveStruct(ctx *parser.StructDefinitionContext) error {
	for structName, structIO := range p.definedStructs {
		for i, component := range structIO.Components {
			if isStructType(p.definedStructs, component.Type) {
				basicStructType := strings.TrimRight(component.Type, "[]")
				nestedComponent, err := p.getStructComponents(basicStructType)
				if err != nil {
					// Problematic is that if there are multiple passes to resolve structs,
					// we will get multiple errors for the same struct while at the same time at the last pass
					// we will get the correct result. This is because we are not sure if the struct is defined
					// before or after the struct that uses it.
					// Forward declarations... because of it, debug log is used instead of error/warn.
					zap.L().Debug(
						"Failed to discover struct nested component. Maybe it's not defined yet?",
						zap.String("contract", p.contractName),
						zap.String("struct", structName),
						zap.String("component_type", component.Type),
						zap.String("component_name", component.Name),
						zap.Error(err),
					)
					return err
				}

				structIO.Components[i].Components = nestedComponent.Components
				structIO.Components[i].Type = normalizeStructTypeName(p.definedStructs, component.Type)
				structIO.Components[i].InternalType = fmt.Sprintf("struct %s.%s", p.contractName, component.Type)
			}
		}
	}

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

// getStructComponents retrieves the components of a struct given its name.
// It returns a MethodIO object representing the components of the struct and an error.
// If the struct is not defined in the AbiParser's definedStructs map, it returns an empty MethodIO object and an error.
func (p *AbiParser) getStructComponents(structName string) (MethodIO, error) {
	components, exists := p.definedStructs[structName]
	if !exists {
		return MethodIO{}, fmt.Errorf("struct %s not defined", structName)
	}

	return components, nil
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

// ToStruct returns the ABI object.
func (p *AbiParser) ToStruct() ABI {
	return p.abi
}

// GetConstructor returns the constructor of the contract.
func (p *AbiParser) GetConstructor() (*MethodConstructor, error) {
	if p.contractName == "" {
		return nil, errors.New("contract name is not defined")
	}

	if _, ok := p.definedContracts[p.contractName]; !ok {
		return nil, errors.New("contract is not defined")
	}
	toReturn := p.definedContracts[p.contractName].Constructor
	return &toReturn, nil
}
