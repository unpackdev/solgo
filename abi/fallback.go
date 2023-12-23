// Package abi provides tools for building and parsing Ethereum ABI (Application Binary Interface) data.
package abi

import (
	"github.com/unpackdev/solgo/ir"
)

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

	return toReturn
}
