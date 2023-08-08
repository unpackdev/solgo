package abi

import "github.com/txpull/solgo/ir"

func (b *Builder) processConstructor(unit *ir.Constructor) *Method {
	toReturn := &Method{
		Name:            "",
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "constructor",
		StateMutability: b.normalizeStateMutability(unit.GetStateMutability()),
	}

	for _, parameter := range unit.GetParameters() {
		methodIo := MethodIO{
			Name: parameter.GetName(),
		}
		toReturn.Inputs = append(
			toReturn.Inputs,
			b.buildMethodIO(methodIo, parameter.GetTypeDescription()),
		)
	}

	for _, parameter := range unit.GetReturnStatements() {
		methodIo := MethodIO{
			Name: parameter.GetName(),
		}
		toReturn.Outputs = append(
			toReturn.Outputs,
			b.buildMethodIO(methodIo, parameter.GetTypeDescription()),
		)
	}
	return toReturn
}
