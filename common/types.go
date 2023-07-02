package common

// ContractInfo contains information about a contract
type ContractInfo struct {
	Comments   []string `json:"comments"`   // Comments associated with the contract
	License    string   `json:"license"`    // License information of the contract
	Pragmas    []string `json:"pragmas"`    // Pragmas specified in the contract
	Imports    []string `json:"imports"`    // Imported dependencies of the contract
	Name       string   `json:"name"`       // Name of the contract
	Implements []string `json:"implements"` // Interfaces implemented by the contract
}

// MethodIO represents an input or output parameter of a contract method or event.
type MethodIO struct {
	Indexed      bool       `json:"indexed,omitempty"`    // Used only by the events
	InternalType string     `json:"internalType"`         // The internal Solidity type of the parameter
	Name         string     `json:"name"`                 // The name of the parameter
	Type         string     `json:"type"`                 // The type of the parameter
	Components   []MethodIO `json:"components,omitempty"` // Components of the parameter, if it's a struct or tuple type
}

// IMethod is an interface that represents a contract method, event, or constructor.
type IMethod interface{}

// MethodConstructor represents a contract constructor.
type MethodConstructor struct {
	Inputs  []MethodIO `json:"inputs"`            // The input parameters of the constructor
	Type    string     `json:"type"`              // The type of the method (always "constructor" for constructors)
	Outputs []MethodIO `json:"outputs,omitempty"` // The output parameters of the constructor (always empty for constructors)
}

// MethodEvent represents a contract event.
type MethodEvent struct {
	Anonymous bool       `json:"anonymous"`      // Whether the event is anonymous
	Inputs    []MethodIO `json:"inputs"`         // The input parameters of the event
	Name      string     `json:"name,omitempty"` // The name of the event
	Type      string     `json:"type"`           // The type of the method (always "event" for events)
}

// Method represents a contract function.
type Method struct {
	Inputs          []MethodIO `json:"inputs"`          // The input parameters of the function
	Outputs         []MethodIO `json:"outputs"`         // The output parameters of the function
	Name            string     `json:"name"`            // The name of the function
	Type            string     `json:"type"`            // The type of the method (always "function" for functions)
	StateMutability string     `json:"stateMutability"` // The state mutability of the function (pure, view, nonpayable, payable)
}

// MethodVariable represents a contract state variable.
type MethodVariable struct {
	Inputs          []MethodIO `json:"inputs"`          // The input parameters of the variable (always empty for variables)
	Outputs         []MethodIO `json:"outputs"`         // The output parameters of the variable (always contains one element representing the variable itself)
	Name            string     `json:"name"`            // The name of the variable
	Type            string     `json:"type"`            // The type of the method (always "function" for variables)
	StateMutability string     `json:"stateMutability"` // The state mutability of the variable (always "view" for variables)
}

// MethodFallbackOrReceive represents a contract fallback or receive function.
type MethodFallbackOrReceive struct {
	Type            string `json:"type"`                      // The type of the method (either "fallback" or "receive")
	StateMutability string `json:"stateMutability,omitempty"` // The state mutability of the function (nonpayable for fallback functions, payable for receive functions)
}

// ABI represents a contract ABI, which is a list of contract methods, events, and constructors.
type ABI []IMethod
