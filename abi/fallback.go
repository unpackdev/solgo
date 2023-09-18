// Package abi provides tools for building and parsing Ethereum ABI (Application Binary Interface) data.
package abi

import "github.com/unpackdev/solgo/ir"

// processFallback processes an IR fallback function and returns a Method representation of it.
// It extracts the input and output parameters of the fallback function and normalizes its state mutability.
func (b *Builder) processFallback(unit *ir.Fallback) *Method {
	toReturn := &Method{
		Name:            "", // In Ethereum, fallback functions don't have a name.
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "fallback",
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

	// Process return statements of the fallback function.
	// Note: In Ethereum, fallback functions can have return values.
	/* 	for _, parameter := range unit.GetReturnStatements() {
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
