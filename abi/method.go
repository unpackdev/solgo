package abi

import (
	abi_pb "github.com/txpull/protos/dist/go/abi"
)

// MethodIO represents an input or output parameter of a contract method or event.
type MethodIO struct {
	Indexed         bool       `json:"indexed,omitempty"`         // Indicates if the parameter is indexed. Only used by events.
	InternalType    string     `json:"internalType,omitempty"`    // Represents the internal Solidity type of the parameter.
	Name            string     `json:"name"`                      // Name of the parameter.
	Type            string     `json:"type"`                      // Type of the parameter.
	Components      []MethodIO `json:"components,omitempty"`      // Components of the parameter, used if it's a struct or tuple type.
	StateMutability string     `json:"stateMutability,omitempty"` // State mutability of the function (e.g., pure, view, nonpayable, payable).
	Inputs          []MethodIO `json:"inputs,omitempty"`          // Input parameters of the function.
	Outputs         []MethodIO `json:"outputs,omitempty"`         // Output parameters of the function.
}

// ToProto converts the MethodIO to its protobuf representation.
func (m *MethodIO) ToProto() *abi_pb.MethodIO {
	toReturn := &abi_pb.MethodIO{
		Indexed:         m.Indexed,
		InternalType:    m.InternalType,
		Name:            m.Name,
		Type:            m.Type,
		StateMutability: m.StateMutability,
		Inputs:          make([]*abi_pb.MethodIO, 0),
		Outputs:         make([]*abi_pb.MethodIO, 0),
		Components:      make([]*abi_pb.MethodIO, 0),
	}

	for _, input := range m.Inputs {
		toReturn.Inputs = append(toReturn.Inputs, input.ToProto())
	}

	for _, output := range m.Outputs {
		toReturn.Outputs = append(toReturn.Outputs, output.ToProto())
	}

	for _, component := range m.Components {
		toReturn.Components = append(toReturn.Components, component.ToProto())
	}

	return toReturn
}

// Method represents a contract function.
type Method struct {
	Components      []MethodIO `json:"components,omitempty"` // Components of the parameter, used if it's a struct or tuple type.
	Inputs          []MethodIO `json:"inputs"`               // Input parameters of the function.
	Outputs         []MethodIO `json:"outputs"`              // Output parameters of the function.
	Name            string     `json:"name"`                 // Name of the function.
	Type            string     `json:"type"`                 // Type of the method (always "function" for functions).
	StateMutability string     `json:"stateMutability"`      // State mutability of the function (e.g., pure, view, nonpayable, payable).
}

// ToProto converts the Method to its protobuf representation.
func (m *Method) ToProto() *abi_pb.Method {
	toReturn := &abi_pb.Method{
		Components:      make([]*abi_pb.MethodIO, 0),
		Inputs:          make([]*abi_pb.MethodIO, 0),
		Outputs:         make([]*abi_pb.MethodIO, 0),
		Name:            m.Name,
		Type:            m.Type,
		StateMutability: m.StateMutability,
	}

	for _, input := range m.Inputs {
		toReturn.Inputs = append(toReturn.Inputs, input.ToProto())
	}

	for _, output := range m.Outputs {
		toReturn.Outputs = append(toReturn.Outputs, output.ToProto())
	}

	for _, component := range m.Components {
		toReturn.Components = append(toReturn.Components, component.ToProto())
	}

	return toReturn
}
