package abi

import "github.com/txpull/solgo/ir"

// processReceive processes the provided Receive unit from the IR and returns a Method representation.
// The returned Method will always have its Type set to "receive" and StateMutability set to "payable".
func (b *Builder) processReceive(unit *ir.Receive) *Method {
	toReturn := &Method{
		Name:            "", // Name is left empty for receive type
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "receive", // Type is always set to "receive" for this kind of method
		StateMutability: "payable", // Receive methods are always payable
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
