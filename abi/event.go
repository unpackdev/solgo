package abi

import (
	"github.com/unpackdev/solgo/ir"
)

// processEvent processes an IR event and returns a Method representation of it.
// It extracts the name and parameters of the event and sets its type to "event" and state mutability to "view".
func (b *Builder) processEvent(unit *ir.Event) *Method {
	toReturn := &Method{
		Name:            unit.GetName(),
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "event",
		StateMutability: "view", // Events in Ethereum are view-only and don't modify state.
	}

	// Process parameters of the event.
	// Note: In Ethereum, event parameters are considered as outputs.
	for _, parameter := range unit.GetParameters() {
		toReturn.Outputs = append(toReturn.Outputs, MethodIO{
			Name:         parameter.GetName(),
			Type:         parameter.GetTypeDescription().TypeString,
			InternalType: parameter.GetTypeDescription().TypeString,
			Indexed:      true, // Parameters for events can be indexed.
		})
	}

	return toReturn
}
