package abi

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
)

// nolint:unused
func dumpNode(whatever interface{}) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
	os.Exit(1)
}

// nolint:unused
func dumpNodeNoExit(whatever interface{}) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
}

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
// It returns true if the type name contains the string "mapping", and false otherwise.
func isMappingType(name string) bool {
	return strings.Contains(name, "mapping")
}

// isContractType checks if the given type name represents a contract type in Solidity.
// It returns true if the type name contains the string "contract", and false otherwise.
func isContractType(name string) bool {
	return strings.Contains(name, "contract")
}

// parseMappingType parses a mapping type in Solidity ABI.
// It takes a string of the form "mapping(keyType => valueType)" and returns three values:
//   - A boolean indicating whether the parsing was successful. If the string is not a mapping type, this will be false.
//   - A slice of strings representing the types of the keys in the mapping. If the mapping is nested, this will contain multiple elements.
//   - A slice of strings representing the types of the values in the mapping. If the mapping is nested, the inner mappings will be flattened,
//     and this will contain the types of the innermost values.
func parseMappingType(typeName string) (bool, []string, []string) {
	re := regexp.MustCompile(`mapping\((\w+)\s*=>\s*(.+)\)`)
	matches := re.FindStringSubmatch(typeName)

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

// normalizeTypeName normalizes the type name in Solidity to its canonical form.
// For example, "uint" is normalized to "uint256", and "addresspayable" is normalized to "address".
// If the type name is not one of the special cases, it is returned as is.
func normalizeTypeName(typeName string) string {
	isArray, _ := regexp.MatchString(`\[\d+\]`, typeName)
	isSlice := strings.HasPrefix(typeName, "[]")

	switch {
	case isArray:
		numberPart := typeName[strings.Index(typeName, "[")+1 : strings.Index(typeName, "]")]
		typePart := typeName[strings.Index(typeName, "]")+1:]
		return "[" + numberPart + "]" + normalizeTypeName(typePart)

	case isSlice:
		typePart := typeName[2:]
		return "[]" + normalizeTypeName(typePart)

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

func normalizeTypeNameWithStatus(typeName string) (string, bool) {
	isArray, _ := regexp.MatchString(`\[\d+\]`, typeName)
	isSlice := strings.HasPrefix(typeName, "[]")

	switch {
	case isArray:
		numberPart := typeName[strings.Index(typeName, "[")+1 : strings.Index(typeName, "]")]
		typePart := typeName[strings.Index(typeName, "]")+1:]

		return "[" + numberPart + "]" + normalizeTypeName(typePart), true

	case isSlice:
		typePart := typeName[2:]
		return "[]" + normalizeTypeName(typePart), true

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
