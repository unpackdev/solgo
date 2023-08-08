package abi

import "github.com/txpull/solgo/ir"

func (b *Builder) processStateVariable(stateVar *ir.StateVariable) *Method {
	toReturn := &Method{
		Name:            stateVar.GetName(),
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "function",
		StateMutability: b.normalizeStateMutability(stateVar.GetStateMutability()),
	}

	typeName := b.resolver.ResolveType(stateVar.GetTypeDescription())

	switch typeName {
	case "mapping":
		inputList, outputList := b.resolver.ResolveMappingType(stateVar.GetTypeDescription())
		toReturn.Inputs = append(toReturn.Inputs, inputList...)
		toReturn.Outputs = append(toReturn.Outputs, outputList...)
	case "contract":
		toReturn.Outputs = append(toReturn.Outputs, MethodIO{
			Type:         "address",
			InternalType: stateVar.GetTypeDescription().GetString(),
		})
	case "enum":
		toReturn.Outputs = append(toReturn.Outputs, MethodIO{
			Type:         "uint8", // enums are represented as uint8 in the ABI
			InternalType: stateVar.GetTypeDescription().GetString(),
		})
	default:
		toReturn.Outputs = append(toReturn.Outputs, MethodIO{
			Type:         typeName,
			InternalType: stateVar.GetTypeDescription().GetString(),
		})
	}

	return toReturn
}
