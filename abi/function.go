// Package abi provides tools for building and parsing Ethereum ABI (Application Binary Interface) data.
package abi

import (
	"github.com/unpackdev/solgo/ir"
)

// processFunction processes an IR function and returns a Method representation of it.
// It extracts the name, input, and output parameters of the function, sets its type to "function",
// and normalizes its state mutability.
func (b *Builder) processFunction(unit *ir.Function) *Method {
	toReturn := &Method{
		Name:            unit.GetName(),
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "function",
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
