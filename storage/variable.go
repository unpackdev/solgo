package storage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unpackdev/solgo/ir"
)

// Variable represents a state variable within a smart contract.
// It embeds ir.StateVariable and provides additional methods
// to determine the characteristics of the variable.
type Variable struct {
	*ir.StateVariable `json:"state_variable"` // StateVariable is the underlying state variable from the IR.
	Contract          *ir.Contract            `json:"contract"`       // Contract is the contract to which this variable belongs.
	EntryContract     bool                    `json:"entry_contract"` // EntryContract indicates if this variable is part of the entry contract.
}

// IsAddressType checks if the variable is of an address type.
// Returns true if the variable type is prefixed with "address".
func (v *Variable) IsAddressType() bool {
	variableType := v.StateVariable.GetType()
	return strings.HasPrefix(variableType, "address")
}

// IsDynamicArray checks if the variable is a dynamic array.
// Currently returns false and needs implementation for dynamic array checking.
func (v *Variable) IsDynamicArray() bool {
	variableType := v.StateVariable.GetType()
	return strings.Contains(variableType, "[]")
}

// IsArrayType checks if the variable is of an array type.
// Returns true if the variable type includes square brackets, indicating an array.
func (v *Variable) IsArrayType() bool {
	variableType := v.StateVariable.GetType()
	return !strings.Contains(variableType, "[]") && strings.Contains(variableType, "[")
}

// IsMappingType checks if the variable is of a mapping type.
// Returns true if the variable type starts with "mapping".
func (v *Variable) IsMappingType() bool {
	variableType := v.StateVariable.GetType()
	return strings.HasPrefix(variableType, "mapping")
}

// IsStructType checks if the variable is of a struct type.
// Returns true if the variable type starts with "struct".
func (v *Variable) IsStructType() bool {
	variableType := v.StateVariable.GetType()
	return strings.HasPrefix(variableType, "struct")
}

// IsStructType checks if the variable is of a contract type.
// Returns true if the variable type starts with "contract".
func (v *Variable) IsContractType() bool {
	variableType := v.StateVariable.GetType()
	return strings.HasPrefix(variableType, "contract")
}

// GetArrayLength returns the length of an array variable.
// For a variable representing an array, it parses and returns the array length.
// Returns 0 for non-array variables or if the length is not explicitly defined.
func (v *Variable) GetArrayLength() (int64, error) {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type contains square brackets, indicating an array
	if strings.Contains(variableType, "[") && strings.Contains(variableType, "]") {
		// Find the position of the square brackets
		startIndex := strings.Index(variableType, "[")
		endIndex := strings.Index(variableType, "]")

		// Extract the substring between the square brackets
		lengthStr := variableType[startIndex+1 : endIndex]

		// If the length is not specified, it's a dynamic array (return 0)
		if lengthStr == "" {
			return 0, nil
		}

		// Convert the length string to an integer
		length, err := strconv.ParseInt(lengthStr, 10, 64)
		if err != nil {
			// Handle the error (e.g., log it, return 0, or return an error)
			// For simplicity, returning 0 here
			return 0, fmt.Errorf("failed to parse array length: %v", err)
		}

		return length, nil
	}

	return 0, nil
}

// GetStructMembers returns a slice of Variables representing the members of a struct.
// For a variable representing a struct, it parses and returns the struct members.
// Returns an empty slice for non-struct variables or if no members are defined.
func (v *Variable) GetStructMembers() []*Variable {
	variableType := v.StateVariable.GetType()
	if strings.HasPrefix(variableType, "struct") {
		// Implementation for parsing the members from the type
		// ...
		return []*Variable{}
	}
	return []*Variable{}
}
