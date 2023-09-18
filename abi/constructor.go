package abi

import "github.com/unpackdev/solgo/ir"

// processConstructor processes an IR constructor and returns a Method representation of it.
// It extracts the input and output parameters of the constructor and normalizes its state mutability.
func (b *Builder) processConstructor(unit *ir.Constructor) *Method {
	// Initialize a new Method structure for the constructor.
	toReturn := &Method{
		Name:            "", // Constructors in Ethereum don't have a name.
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

	// Process return statements of the constructor.
	// Note: In Ethereum, constructors don't return values. This might be specific to the IR representation.
	/* for _, parameter := range unit.GetReturnStatements() {
		methodIo := MethodIO{
			Name: parameter.GetName(),
		}
		toReturn.Outputs = append(
			toReturn.Outputs,
			b.buildMethodIO(methodIo, parameter.GetTypeDescription()),
		)
	} */

	return toReturn
}
