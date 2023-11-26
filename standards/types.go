package standards

import (
	eip_pb "github.com/unpackdev/protos/dist/go/eip"
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
	Type string `json:"type"`

	// Indexed indicates whether the input is indexed.
	// This is particularly relevant for event parameters,
	// where indexed parameters can be used as a filter for event logs.
	Indexed bool `json:"indexed"`

	// Matched indicates whether the input has been matched via confidence check.
	Matched bool `json:"matched"`
}

// ToProto converts the Input to its protobuf representation.
func (i *Input) ToProto() *eip_pb.Input {
	return &eip_pb.Input{
		Type:    i.Type,
		Indexed: i.Indexed,
		Matched: i.Matched,
	}
}

// Output represents an output parameter for Ethereum functions and events.
type Output struct {
	// Type specifies the Ethereum data type of the output.
	Type string `json:"type"`

	// Matched indicates whether the output has been matched via confidence check.
	Matched bool `json:"matched"`
}

// ToProto converts the Output to its protobuf representation.
func (i *Output) ToProto() *eip_pb.Output {
	return &eip_pb.Output{
		Type:    i.Type,
		Matched: i.Matched,
	}
}

// Function represents an Ethereum smart contract function.
type Function struct {
	// Name specifies the name of the function.
	Name string `json:"name"`

	// Inputs is a slice of Input structs, representing the input parameters of the function.
	Inputs []Input `json:"inputs"`

	// Outputs is a slice of Output structs, representing the data types of the function's return values.
	Outputs []Output `json:"outputs"`

	// Matched indicates whether the input has been matched via confidence check.
	Matched bool `json:"matched"`
}

// ToProto converts the Function to its protobuf representation.
func (f *Function) ToProto() *eip_pb.Function {
	protoInputs := make([]*eip_pb.Input, 0)
	protoOutputs := make([]*eip_pb.Output, 0)

	for _, input := range f.Inputs {
		protoInputs = append(protoInputs, input.ToProto())
	}

	for _, output := range f.Outputs {
		protoOutputs = append(protoOutputs, output.ToProto())
	}

	return &eip_pb.Function{
		Name:    f.Name,
		Inputs:  protoInputs,
		Outputs: protoOutputs,
		Matched: f.Matched,
	}
}

// Event represents an Ethereum smart contract event.
type Event struct {
	// Name specifies the name of the event.
	Name string `json:"name"`

	// Inputs is a slice of Input structs, representing the input parameters of the event.
	Inputs []Input `json:"inputs"`

	// Outputs is a slice of Output structs, representing the data types of the event's return values.
	Outputs []Output `json:"outputs"`

	// Matched indicates whether the input has been matched via confidence check.
	Matched bool `json:"matched"`
}

// ToProto converts the Event to its protobuf representation.
func (e *Event) ToProto() *eip_pb.Event {
	protoInputs := make([]*eip_pb.Input, 0)
	protoOutputs := make([]*eip_pb.Output, 0)

	for _, input := range e.Inputs {
		protoInputs = append(protoInputs, input.ToProto())
	}

	for _, output := range e.Outputs {
		protoOutputs = append(protoOutputs, output.ToProto())
	}

	return &eip_pb.Event{
		Name:    e.Name,
		Inputs:  protoInputs,
		Outputs: protoOutputs,
		Matched: e.Matched,
	}
}

// ContractStandard represents a standard interface for Ethereum smart contracts,
// such as the ERC-20 or ERC-721 standards.
type ContractStandard struct {
	// Name specifies the name of the contract standard, e.g., "ERC-20 Token Standard".
	Name string `json:"name"`

	// Url specifies the URL of the contract standard, e.g., "https://eips.ethereum.org/EIPS/eip-20".
	Url string `json:"url"`

	// Type specifies the type of the contract standard, e.g., ERC20 or ERC721.
	Type Standard `json:"type"`

	// Stagnant indicates whether the contract standard is stagnant in terms of development.
	Stagnant bool `json:"stagnant"`

	// ABI specifies the ABI of the contract standard.
	ABI string `json:"abi"`

	// PackageName specifies the package name of the contract standard.
	PackageName string `json:"package_name"`

	// PackageOutputPath specifies the package output path of the contract standard.
	PackageOutputPath string `json:"package_output_path"`

	// Functions is a slice of Function structs, representing the functions defined in the contract standard.
	Functions []Function `json:"functions"`

	// Events is a slice of Event structs, representing the events defined in the contract standard.
	Events []Event `json:"events"`
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
		Url:       cs.Url,
		Type:      standard,
		Stagnant:  cs.Stagnant,
		Functions: protoFunctions,
		Events:    protoEvents,
	}
}

// ContractMatcher represents an Ethereum smart contract that attempts to confirm to a standard interface,
// such as the ERC-20 or ERC-721 standards. Used while performing a contract standard detection.
type ContractMatcher struct {
	// Name of the contract.
	Name string `json:"name"`

	// Functions is a slice of Function structs, representing the functions defined in the contract standard.
	Functions []Function `json:"functions"`

	// Events is a slice of Event structs, representing the events defined in the contract standard.
	Events []Event `json:"events"`
}

// ToProto converts the Event to its protobuf representation.
func (c *ContractMatcher) ToProto() *eip_pb.Contract {
	protoFns := make([]*eip_pb.Function, 0)
	protoEvents := make([]*eip_pb.Event, 0)

	for _, fn := range c.Functions {
		protoFns = append(protoFns, fn.ToProto())
	}

	for _, event := range c.Events {
		protoEvents = append(protoEvents, event.ToProto())
	}

	return &eip_pb.Contract{
		Name:      c.Name,
		Functions: protoFns,
		Events:    protoEvents,
	}
}

type FunctionMatcher struct {
	// Name of the contract.
	Name string `json:"name"`

	// Functions is a slice of Function structs, representing the functions defined in the contract standard.
	Functions []Function `json:"functions"`
}

// Discovery represents a contract standard discovery response.
type Discovery struct {
	// Confidence specifies the confidence level of the discovery.
	Confidence ConfidenceLevel `json:"confidence"`

	// ConfidencePoints specifies the confidence points of the discovery.
	ConfidencePoints float64 `json:"confidence_points"`

	// Threshold specifies the threshold level of the discovery.
	Threshold ConfidenceThreshold `json:"threshold"`

	// MaximumTokens specifies the maximum number of tokens in the standard.
	// This is basically a standard TokenCount() function response value.
	MaximumTokens int `json:"maximum_tokens"`

	// DiscoverdTokens specifies the number of tokens discovered in the standard.
	// The more tokens discovered, the higher the confidence level.
	DiscoveredTokens int `json:"discovered_tokens"`

	// ContractStandard that is being scanned.
	Standard Standard `json:"standard"`

	// Contract that is being scanned including mathed functions and events.
	Contract *ContractMatcher `json:"contract"`
}

// ToProto converts the Discovery to its protobuf representation.
func (d *Discovery) ToProto() *eip_pb.Discovery {
	return &eip_pb.Discovery{
		Standard:         d.Standard.ToProto(),
		Confidence:       d.Confidence.ToProto(),
		ConfidencePoints: int32(d.ConfidencePoints * 100),
		Threshold:        d.Threshold.ToProto(),
		MaximumTokens:    int32(d.MaximumTokens),
		DiscoveredTokens: int32(d.DiscoveredTokens),
		Contract:         d.Contract.ToProto(),
	}
}

type FunctionDiscovery struct {
	// Confidence specifies the confidence level of the discovery.
	Confidence ConfidenceLevel `json:"confidence"`

	// ConfidencePoints specifies the confidence points of the discovery.
	ConfidencePoints float64 `json:"confidence_points"`

	// Threshold specifies the threshold level of the discovery.
	Threshold ConfidenceThreshold `json:"threshold"`

	// MaximumTokens specifies the maximum number of tokens in the standard.
	// This is basically a standard TokenCount() function response value.
	MaximumTokens int `json:"maximum_tokens"`

	// DiscoverdTokens specifies the number of tokens discovered in the standard.
	// The more tokens discovered, the higher the confidence level.
	DiscoveredTokens int `json:"discovered_tokens"`

	// ContractStandard that is being scanned.
	Standard Standard `json:"standard"`

	// Contract that is being scanned including mathed functions and events.
	Function *Function `json:"function"`
}
