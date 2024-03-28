package standards

// TokenCount calculates and returns the total number of tokens (inputs and outputs)
// present in the functions and events of a given ContractStandard.
func TokenCount(cs ContractStandard) int {
	count := 0

	for _, function := range cs.Functions {
		count++

		for _, input := range function.Inputs {
			count++
			if len(input.Type) > 0 {
				count++
			}
			count++ // Indexed is always counted as it's a boolean
		}

		for _, output := range function.Outputs {
			count++
			if len(output.Type) > 0 {
				count++
			}
		}
	}

	for _, event := range cs.Events {
		count++

		for _, input := range event.Inputs {
			count++
			if len(input.Type) > 0 {
				count++
			}
			count++ // Indexed is always counted as it's a boolean
		}

		for _, output := range event.Outputs {
			count++
			if len(output.Type) > 0 {
				count++
			}
		}
	}

	return count
}

// FunctionTokenCount calculates the total number of tokens present in a given Ethereum smart contract function.
// It considers the function name as the initial token, then iterates over all inputs and outputs of the function.
// For each input and output, it increments the count by one for the parameter itself, an additional one if the type
// is specified (non-empty), and another for the 'Indexed' attribute (for inputs only), acknowledging it as a boolean.
// This count provides an estimate of the complexity or size of the function in terms of its components.
func FunctionTokenCount(fn Function) int {
	count := 1 // Assuming function name...

	for _, input := range fn.Inputs {
		count++
		if len(input.Type) > 0 {
			count++
		}
		count++ // Indexed is always counted as it's a boolean
	}

	for _, output := range fn.Outputs {
		count++
		if len(output.Type) > 0 {
			count++
		}
	}

	return count
}
