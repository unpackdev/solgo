package eip

import (
	"fmt"
	"strings"

	eip_pb "github.com/txpull/protos/dist/go/eip"
)

// newFunction creates and returns a new Function struct with the provided name, inputs, and outputs.
func newFunction(name string, inputs []Input, outputs []string) Function {
	return Function{
		Name:    name,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

// newEvent creates and returns a new Event struct with the provided name, inputs, and outputs.
func newEvent(name string, inputs []Input, outputs []string) Event {
	return Event{
		Name:    name,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

// GetProtoStandardFromString converts a string representation of an Ethereum standard
// to its corresponding protobuf enum value. If the standard is not recognized,
// it returns an error.
//
// Parameters:
// s: The string representation of the Ethereum standard.
//
// Returns:
// - The corresponding protobuf enum value of the Ethereum standard.
// - An error if the standard is not recognized.
func GetProtoStandardFromString(s string) (eip_pb.Standard, error) {
	// Convert the string to uppercase to match the enum naming convention
	standardValue, ok := eip_pb.Standard_value[strings.ToUpper(s)]
	if !ok {
		return eip_pb.Standard_UNKNOWN, fmt.Errorf("unknown standard '%s'", s)
	}
	return eip_pb.Standard(standardValue), nil
}
