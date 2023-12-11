package abi

import (
	"regexp"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// normalizeStateMutability converts the provided Mutability value to its corresponding string representation.
func (b *Builder) normalizeStateMutability(m ast_pb.Mutability) string {
	switch m {
	case ast_pb.Mutability_PURE:
		return "pure"
	case ast_pb.Mutability_VIEW:
		return "view"
	case ast_pb.Mutability_NONPAYABLE:
		return "nonpayable"
	case ast_pb.Mutability_PAYABLE:
		return "payable"
	default:
		return "view"
	}
}

// isMappingType checks if the given type name represents a mapping type in Solidity.
func isMappingType(name string) bool {
	return strings.Contains(name, "mapping")
}

// isContractType checks if the given type name represents a contract type in Solidity.
func isContractType(name string) bool {
	return strings.Contains(name, "contract")
}

// isStructType checks if the given type name represents a enum type in Solidity.
func isEnumType(name string) bool {
	return strings.Contains(name, "enum")
}

// isStructType checks if the given type name represents a enum type in Solidity.
func isStructType(name string) bool {
	return strings.Contains(name, "struct")
}

// parseMappingType parses a Solidity ABI mapping type.
// It returns a boolean indicating success, a slice of key types, and a slice of value types.
func parseMappingType(typeName string) (bool, []string, []string) {
	re := regexp.MustCompile(`mapping\((\w+)\s*=>\s*(.+)\)`)
	matches := re.FindStringSubmatch(typeName)

	if len(matches) < 3 {
		return false, []string{}, []string{}
	}

	input := []string{matches[1]}
	output := []string{matches[2]}

	if isMappingType(output[0]) {
		_, nestedInput, nestedOutput := parseMappingType(output[0])
		input = append(input, nestedInput...)
		output = nestedOutput
	}

	return true, input, output
}

// normalizeTypeName normalizes a Solidity type name to its canonical form.
func normalizeTypeName(typeName string) string {
	isArray, _ := regexp.MatchString(`\[\d+\]`, typeName)
	isSlice := strings.HasPrefix(typeName, "[]")
	isSliceRight := strings.HasSuffix(typeName, "[]")

	switch {
	case isArray:
		numberPart := typeName[strings.Index(typeName, "[")+1 : strings.Index(typeName, "]")]
		typePart := typeName[strings.Index(typeName, "]")+1:]
		return "[" + numberPart + "]" + normalizeTypeName(typePart)

	case isSlice:
		typePart := typeName[2:]
		return "[]" + normalizeTypeName(typePart)
	case isSliceRight:
		typePart := typeName[:len(typeName)-2]
		return normalizeTypeName(typePart) + "[]"

	case strings.HasPrefix(typeName, "int_const"):
		return "int256"

	case strings.HasPrefix(typeName, "uint"):
		if typeName == "uint" {
			return "uint256"
		}
		return typeName
	case strings.HasPrefix(typeName, "int"):
		if typeName == "int" {
			return "int256"
		}
		return typeName
	case strings.HasPrefix(typeName, "enum"):
		return "uint8"
	case strings.HasPrefix(typeName, "bool"):
		return typeName
	case strings.HasPrefix(typeName, "bytes"):
		return typeName
	case typeName == "string":
		return "string"
	case typeName == "address":
		return "address"
	case typeName == "addresspayable":
		return "address"
	case typeName == "tuple":
		return "tuple"
	default:
		return typeName
	}
}

// normalizeTypeNameWithStatus normalizes a Solidity type name and returns a boolean indicating if the type is recognized.
func normalizeTypeNameWithStatus(typeName string) (string, bool) {
	isArray, _ := regexp.MatchString(`\[\d+\]`, typeName)
	isSlice := strings.HasPrefix(typeName, "[]")
	isSliceRight := strings.HasSuffix(typeName, "[]")

	switch {
	case isArray:
		numberPart := typeName[strings.Index(typeName, "[")+1 : strings.Index(typeName, "]")]
		typePart := typeName[strings.Index(typeName, "]")+1:]

		return "[" + numberPart + "]" + normalizeTypeName(typePart), true

	case isSlice:
		typePart := typeName[2:]
		return "[]" + normalizeTypeName(typePart), true

	case isSliceRight:
		typePart := typeName[:len(typeName)-2]
		return normalizeTypeName(typePart) + "[]", true

	case strings.HasPrefix(typeName, "uint"):
		if typeName == "uint" {
			return "uint256", true
		}
		return typeName, true
	case strings.HasPrefix(typeName, "int"):
		if typeName == "int" {
			return "int256", true
		}
		return typeName, true
	case strings.HasPrefix(typeName, "enum"):
		return "uint8", true
	case strings.HasPrefix(typeName, "bool"):
		return typeName, true
	case strings.HasPrefix(typeName, "bytes"):
		return typeName, true
	case typeName == "string":
		return "string", true
	case typeName == "address":
		return "address", true
	case typeName == "addresspayable":
		return "address", true
	case typeName == "tuple":
		return "tuple", true
	default:
		return typeName, false
	}
}
