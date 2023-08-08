package abi

import (
	"github.com/txpull/solgo/ir"
)

func (b *Builder) processError(unit *ir.Error) *Method {
	toReturn := &Method{
		Name:            unit.GetName(),
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "error",
		StateMutability: "view",
	}

	for _, parameter := range unit.GetParameters() {
		toReturn.Inputs = append(toReturn.Inputs, MethodIO{
			Name:         parameter.GetName(),
			Type:         parameter.GetTypeDescription().TypeString,
			InternalType: parameter.GetTypeDescription().TypeString,
			Indexed:      true,
		})
	}

	return toReturn
}
