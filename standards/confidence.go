package standards

import (
	eip_pb "github.com/txpull/protos/dist/go/eip"
)

// ConfidenceLevel represents the confidence level of a discovery.
type ConfidenceLevel int

// String returns the string representation of the confidence level.
func (c ConfidenceLevel) String() string {
	switch c {
	case HighConfidence:
		return "high"
	case MediumConfidence:
		return "medium"
	case LowConfidence:
		return "low"
	case NoConfidence:
		return "no_confidence"
	default:
		return "unknown"
	}
}

// ConfidenceThreshold represents the threshold value for a confidence level.
type ConfidenceThreshold float64

// ToProto converts a ConfidenceLevel to its protobuf representation.
func (c ConfidenceLevel) ToProto() eip_pb.ConfidenceLevel {
	return eip_pb.ConfidenceLevel(c)
}

// ToProto converts a ConfidenceThreshold to its protobuf representation.
func (c ConfidenceThreshold) ToProto() eip_pb.ConfidenceThreshold {
	return eip_pb.ConfidenceThreshold(c)
}

const (
	// HighConfidenceThreshold represents a high confidence threshold value.
	HighConfidenceThreshold ConfidenceThreshold = 0.9

	// MediumConfidenceThreshold represents a medium confidence threshold value.
	MediumConfidenceThreshold ConfidenceThreshold = 0.5

	// LowConfidenceThreshold represents a low confidence threshold value.
	LowConfidenceThreshold ConfidenceThreshold = 0.1

	// NoConfidenceThreshold represents no confidence threshold value.
	NoConfidenceThreshold ConfidenceThreshold = 0.0

	// HighConfidence represents a high confidence level.
	HighConfidence ConfidenceLevel = 3

	// MediumConfidence represents a medium confidence level.
	MediumConfidence ConfidenceLevel = 2

	// LowConfidence represents a low confidence level.
	LowConfidence ConfidenceLevel = 1

	// NoConfidence represents no confidence level.
	NoConfidence ConfidenceLevel = 0
)

// CalculateDiscoveryConfidence calculates the confidence level and threshold based on the total confidence.
func CalculateDiscoveryConfidence(totalConfidence float64) (ConfidenceLevel, ConfidenceThreshold) {
	total := ConfidenceThreshold(totalConfidence)
	switch {
	case total >= HighConfidenceThreshold:
		return HighConfidence, HighConfidenceThreshold
	case total >= MediumConfidenceThreshold:
		return MediumConfidence, MediumConfidenceThreshold
	case total >= LowConfidenceThreshold:
		return LowConfidence, LowConfidenceThreshold
	default:
		return NoConfidence, NoConfidenceThreshold
	}
}

// ConfidenceCheck checks the confidence of a contract against a standard EIP.
func ConfidenceCheck(standard EIP, contract *ContractMatcher) (Discovery, bool) {
	toReturn := Discovery{
		Standard:         standard.GetType(),
		Confidence:       NoConfidence,
		ConfidencePoints: 0,
		Threshold:        NoConfidenceThreshold,
		MaximumTokens:    standard.TokenCount(),
		DiscoveredTokens: 0,
		Contract: &ContractMatcher{
			Name:      contract.Name,
			Functions: make([]Function, 0),
			Events:    make([]Event, 0),
		},
	}
	foundTokenCount := 0
	discoveredFunctions := map[string]bool{}
	discoveredEvents := map[string]bool{}

	for _, standardFunction := range standard.GetFunctions() {
		contractFn := Function{
			Name:    standardFunction.Name,
			Inputs:  make([]Input, 0),
			Outputs: make([]Output, 0),
		}

		for _, contractFunction := range contract.Functions {
			if _, found := discoveredFunctions[contractFunction.Name]; !found {
				if tokensFound, found := functionMatch(&contractFn, standardFunction, contractFunction); found {
					discoveredFunctions[contractFunction.Name] = true
					contractFn.Matched = true
					foundTokenCount += tokensFound
				}
			}
		}

		if !contractFn.Matched {
			contractFn.Matched = false

			if standardFunction.Inputs == nil {
				standardFunction.Inputs = make([]Input, 0)
			} else {
				contractFn.Inputs = standardFunction.Inputs
			}

			if standardFunction.Outputs == nil {
				standardFunction.Outputs = make([]Output, 0)
			} else {
				contractFn.Outputs = standardFunction.Outputs
			}
		}

		toReturn.Contract.Functions = append(toReturn.Contract.Functions, contractFn)
	}

	for _, event := range standard.GetEvents() {

		eventFn := Event{
			Name:    event.Name,
			Inputs:  make([]Input, 0),
			Outputs: make([]Output, 0),
		}

		for _, contractEvent := range contract.Events {
			if _, found := discoveredEvents[contractEvent.Name]; !found {
				if tokensFound, found := eventMatch(&eventFn, event, contractEvent); found {
					discoveredEvents[contractEvent.Name] = true
					eventFn.Matched = true
					foundTokenCount += tokensFound
				}
			}
		}

		if !eventFn.Matched {
			eventFn.Matched = false

			if event.Inputs == nil {
				event.Inputs = make([]Input, 0)
			} else {
				eventFn.Inputs = event.Inputs
			}

			if event.Outputs == nil {
				event.Outputs = make([]Output, 0)
			} else {
				eventFn.Outputs = event.Outputs
			}
		}

		toReturn.Contract.Events = append(toReturn.Contract.Events, eventFn)
	}

	toReturn.DiscoveredTokens = foundTokenCount

	// Calculate the total confidence based on the discovered tokens and maximum tokens
	confidencePoints := float64(foundTokenCount) / float64(standard.TokenCount())
	level, threshold := CalculateDiscoveryConfidence(confidencePoints)
	toReturn.Confidence = level
	toReturn.ConfidencePoints = confidencePoints
	toReturn.Threshold = threshold

	return toReturn, foundTokenCount > 0
}

// functionMatch matches a function from a contract to a standard function and returns the total token count and a boolean indicating if a match was found.
func functionMatch(newFn *Function, standardFunction, contractFunction Function) (int, bool) {
	totalTokenCount := 0

	if standardFunction.Name == contractFunction.Name {
		totalTokenCount++
		for _, sfnInput := range standardFunction.Inputs {
			newInput := Input{Type: sfnInput.Type, Indexed: sfnInput.Indexed}
			for _, fnInput := range contractFunction.Inputs {
				if standardInput, matched := inputMatch(standardFunction.Inputs, fnInput); matched {
					totalTokenCount += 2 // Counting the input match and type match...
					if standardInput.Indexed == fnInput.Indexed {
						totalTokenCount++
					}
					newInput.Matched = true
					break
				}
			}
			newFn.Inputs = append(newFn.Inputs, newInput)
		}

		for _, sfnOutput := range standardFunction.Outputs {
			newOutput := Output{Type: sfnOutput.Type}
			for range standardFunction.Outputs {
				for _, fnOutput := range contractFunction.Outputs {
					if _, matched := outputMatch(standardFunction.Outputs, fnOutput); matched {
						totalTokenCount += 2 // Counting the input match and type match...
					}
					newOutput.Matched = true
					break
				}
			}
			newFn.Outputs = append(newFn.Outputs, newOutput)
		}
	}

	return totalTokenCount, totalTokenCount > 0
}

// eventMatch matches an event from a contract to a standard event and returns the total token count and a boolean indicating if a match was found.
func eventMatch(newEvent *Event, standardEvent, event Event) (int, bool) {
	totalTokenCount := 0

	if standardEvent.Name == event.Name {
		totalTokenCount++
		newEvent.Name = event.Name
		for _, seInput := range standardEvent.Inputs {
			newInput := Input{Type: seInput.Type, Indexed: seInput.Indexed}
			for _, eventInput := range event.Inputs {
				if standardInput, matched := inputMatch(standardEvent.Inputs, eventInput); matched {
					totalTokenCount += 2 // Counting the input match and type match...
					if standardInput.Indexed == eventInput.Indexed {
						totalTokenCount++
					}
					newInput.Matched = true
					break
				}
			}
			newEvent.Inputs = append(newEvent.Inputs, newInput)
		}

		for _, seOutput := range standardEvent.Outputs {
			newOutput := Output{Type: seOutput.Type}
			for range event.Outputs {
				for _, fnOutput := range event.Outputs {
					if _, matched := outputMatch(standardEvent.Outputs, fnOutput); matched {
						totalTokenCount += 2 // Counting the input match and type match...
					}
					newOutput.Matched = true
					break
				}
			}
			newEvent.Outputs = append(newEvent.Outputs, newOutput)
		}
	}

	return totalTokenCount, totalTokenCount > 0
}

// inputMatch matches an input to a list of inputs and returns the matched input and a boolean indicating if a match was found.
func inputMatch(inputs []Input, nodeInput Input) (*Input, bool) {
	for _, input := range inputs {
		if input.Type == nodeInput.Type {
			return &input, true
		}
	}

	return nil, false
}

// outputMatch matches an output to a list of outputs and returns the matched output and a boolean indicating if a match was found.
func outputMatch(outputs []Output, nodeOutput Output) (*Output, bool) {
	for _, output := range outputs {
		if output.Type == nodeOutput.Type {
			return &output, true
		}
	}

	return nil, false
}
