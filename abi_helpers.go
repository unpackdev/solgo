package solgo

import (
	"regexp"
	"strings"
)

// normalizeTypeName normalizes the type name in Solidity to its canonical form.
// For example, "uint" is normalized to "uint256", and "addresspayable" is normalized to "address".
// If the type name is not one of the special cases, it is returned as is.
func normalizeTypeName(typeName string) string {
	switch typeName {
	case "uint":
		return "uint256"
	case "int":
		return "int256"
	case "addresspayable":
		return "address"
	default:
		return typeName
	}
}

// isMappingType checks if the given type name represents a mapping type in Solidity.
// It returns true if the type name contains the string "mapping", and false otherwise.
func isMappingType(name string) bool {
	return strings.Contains(name, "mapping")
}

func isStructType(definedStructs map[string]MethodIO, typeName string) bool {
	_, exists := definedStructs[typeName]
	return exists
}

// parseMappingType parses a mapping type in Solidity ABI.
// It takes a string of the form "mapping(keyType => valueType)" and returns three values:
//   - A boolean indicating whether the parsing was successful. If the string is not a mapping type, this will be false.
//   - A slice of strings representing the types of the keys in the mapping. If the mapping is nested, this will contain multiple elements.
//   - A slice of strings representing the types of the values in the mapping. If the mapping is nested, the inner mappings will be flattened,
//     and this will contain the types of the innermost values.
func parseMappingType(abi string) (bool, []string, []string) {
	re := regexp.MustCompile(`mapping\((\w+)\s*=>\s*(.+)\)`)
	matches := re.FindStringSubmatch(abi)

	if len(matches) < 3 {
		return false, []string{}, []string{}
	}

	input := []string{matches[1]}
	output := []string{matches[2]}

	// If the output is another mapping, parse it recursively
	if isMappingType(output[0]) {
		_, nestedInput, nestedOutput := parseMappingType(output[0])
		input = append(input, nestedInput...)
		output = nestedOutput
	}

	return true, input, output
}
