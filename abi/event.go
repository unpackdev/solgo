package abi

import (
	"fmt"

	"github.com/unpackdev/solgo/ir"
)

// processEvent processes an IR event and returns a Method representation of it.
// It extracts the name and parameters of the event and sets its type to "event" and state mutability to "view".
func (b *Builder) processEvent(unit *ir.Event) (*Method, error) {
	toReturn := &Method{
		Name:            unit.GetName(),
		Inputs:          make([]MethodIO, 0),
		Outputs:         make([]MethodIO, 0),
		Type:            "event",
		StateMutability: "view", // Events in Ethereum are view-only and don't modify state.
	}

	// Process parameters of the event.
	// Note: In Ethereum, event parameters are considered as inputs.
	for _, parameter := range unit.GetParameters() {
		if parameter.GetTypeDescription() == nil {
			return nil, fmt.Errorf("nil type description for event parameter %s", parameter.GetName())
		}

		methodIo := MethodIO{
			Name:    parameter.GetName(),
			Indexed: parameter.IsIndexed(),
		}
		toReturn.Inputs = append(
			toReturn.Inputs,
			b.buildMethodIO(methodIo, parameter.GetTypeDescription()),
		)
	}

	return toReturn, nil
}

func (b *Builder) GetEventAsAbi(unit *ir.Event) ([]*Method, error) {
	method, err := b.processEvent(unit)
	if err != nil {
		return nil, err
	}

	return []*Method{method}, nil
}
