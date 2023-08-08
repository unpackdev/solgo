package abi

import "github.com/txpull/solgo/ir"

func (b *Builder) processReceive(unit *ir.Receive) *Method {
	toReturn := &Method{
		Name:            "",
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "receive",
		StateMutability: "payable", // receive is always payable
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

	return toReturn
}
