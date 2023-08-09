package eip

import (
	eip_pb "github.com/txpull/protos/dist/go/eip"
)

// Constants representing common Ethereum data types.
const (
	// TypeString represents the Ethereum "string" data type.
	TypeString = "string"

	// TypeAddress represents the Ethereum "address" data type.
	TypeAddress = "address"

	// TypeUint256 represents the Ethereum "uint256" data type.
	TypeUint256 = "uint256"

	// TypeBool represents the Ethereum "bool" data type.
	TypeBool = "bool"

	// TypeBytes represents the Ethereum "bytes" data type.
	TypeBytes = "bytes"

	// TypeBytes32 represents the Ethereum "bytes32" data type.
	TypeBytes32 = "bytes32"

	// TypeAddressArray represents an array of Ethereum "address" data types.
	TypeAddressArray = "address[]"

	// TypeUint256Array represents an array of Ethereum "uint256" data types.
	TypeUint256Array = "uint256[]"
)

// Input represents an input parameter for Ethereum functions and events.
type Input struct {
	// Type specifies the Ethereum data type of the input.
	Type string

	// Indexed indicates whether the input is indexed.
	// This is particularly relevant for event parameters,
	// where indexed parameters can be used as a filter for event logs.
	Indexed bool
}

// ToProto converts the Input to its protobuf representation.
func (i *Input) ToProto() *eip_pb.Input {
	return &eip_pb.Input{
		Type:    i.Type,
		Indexed: i.Indexed,
	}
}

// Function represents an Ethereum smart contract function.
type Function struct {
	// Name specifies the name of the function.
	Name string

	// Inputs is a slice of Input structs, representing the input parameters of the function.
	Inputs []Input

	// Outputs is a slice of strings, representing the data types of the function's return values.
	Outputs []string
}

// ToProto converts the Function to its protobuf representation.
func (f *Function) ToProto() *eip_pb.Function {
	protoInputs := make([]*eip_pb.Input, len(f.Inputs))
	for idx, input := range f.Inputs {
		protoInputs[idx] = input.ToProto()
	}

	return &eip_pb.Function{
		Name:    f.Name,
		Inputs:  protoInputs,
		Outputs: f.Outputs,
	}
}

// Event represents an Ethereum smart contract event.
type Event struct {
	// Name specifies the name of the event.
	Name string

	// Inputs is a slice of Input structs, representing the input parameters of the event.
	Inputs []Input

	// Outputs is a slice of strings, representing the data types of the event's return values.
	Outputs []string
}

// ToProto converts the Event to its protobuf representation.
func (e *Event) ToProto() *eip_pb.Event {
	protoInputs := make([]*eip_pb.Input, len(e.Inputs))
	for idx, input := range e.Inputs {
		protoInputs[idx] = input.ToProto()
	}

	return &eip_pb.Event{
		Name:    e.Name,
		Inputs:  protoInputs,
		Outputs: e.Outputs,
	}
}

// ContractStandard represents a standard interface for Ethereum smart contracts,
// such as the ERC-20 or ERC-721 standards.
type ContractStandard struct {
	// Name specifies the name of the contract standard, e.g., "ERC-20 Token Standard".
	Name string

	// Type specifies the type of the contract standard, e.g., ERC20 or ERC721.
	Type Standard

	// Functions is a slice of Function structs, representing the functions defined in the contract standard.
	Functions []Function

	// Events is a slice of Event structs, representing the events defined in the contract standard.
	Events []Event
}

// ToProto converts the ContractStandard to its protobuf representation.
func (cs *ContractStandard) ToProto() *eip_pb.ContractStandard {
	protoFunctions := make([]*eip_pb.Function, len(cs.Functions))
	for idx, function := range cs.Functions {
		protoFunctions[idx] = function.ToProto()
	}

	protoEvents := make([]*eip_pb.Event, len(cs.Events))
	for idx, event := range cs.Events {
		protoEvents[idx] = event.ToProto()
	}

	standard, err := GetProtoStandardFromString(cs.Type.String())
	if err != nil {
		panic(err)
	}

	return &eip_pb.ContractStandard{
		Name:      cs.Name,
		Type:      standard,
		Functions: protoFunctions,
		Events:    protoEvents,
	}
}
