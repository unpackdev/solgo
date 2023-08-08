package abi

// MethodIO represents an input or output parameter of a contract method or event.
type MethodIO struct {
	Indexed         bool       `json:"indexed,omitempty"`         // Used only by the events
	InternalType    string     `json:"internalType,omitempty"`    // The internal Solidity type of the parameter
	Name            string     `json:"name"`                      // The name of the parameter
	Type            string     `json:"type"`                      // The type of the parameter
	Components      []MethodIO `json:"components,omitempty"`      // Components of the parameter, if it's a struct or tuple type
	StateMutability string     `json:"stateMutability,omitempty"` // The state mutability of the function (pure, view, nonpayable, payable)
	Inputs          []MethodIO `json:"inputs,omitempty"`          // The input parameters of the function
	Outputs         []MethodIO `json:"outputs,omitempty"`         // The output parameters of the function
}

// Method represents a contract function.
type Method struct {
	Inputs          []MethodIO `json:"inputs"`          // The input parameters of the function
	Outputs         []MethodIO `json:"outputs"`         // The output parameters of the function
	Name            string     `json:"name"`            // The name of the function
	Type            string     `json:"type"`            // The type of the method (always "function" for functions)
	StateMutability string     `json:"stateMutability"` // The state mutability of the function (pure, view, nonpayable, payable)
}
