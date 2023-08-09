package eip

// TokenCount calculates and returns the total number of tokens (inputs and outputs)
// present in the functions and events of a given ContractStandard.
func TokenCount(cs ContractStandard) int {
	count := 0

	// Helper function to count tokens in a slice of types
	countTokens := func(types []string) int {
		tokenCount := 0
		for range types {
			// Add conditions here if there are specific token types to count
			tokenCount++
		}
		return tokenCount
	}

	// Count tokens in functions
	for _, function := range cs.Functions {
		for _, input := range function.Inputs {
			count++                                    // Count the input type itself
			count += countTokens([]string{input.Type}) // Count tokens in the input type
		}
		count += countTokens(function.Outputs) // Count tokens in the outputs
	}

	// Count tokens in events
	for _, event := range cs.Events {
		for _, input := range event.Inputs {
			count++                                    // Count the input type itself
			count += countTokens([]string{input.Type}) // Count tokens in the input type
		}
		count += countTokens(event.Outputs) // Count tokens in the outputs
	}

	return count
}
