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
