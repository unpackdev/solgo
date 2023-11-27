package storage

import (
	"strings"

	"github.com/unpackdev/solgo/ir"
)

type Variable struct {
	*ir.StateVariable `json:"state_variable"`
	Contract          *ir.Contract `json:"contract"`
	EntryContract     bool         `json:"entry_contract"`
}

func (v *Variable) IsAddressType() bool {
	variableType := v.StateVariable.GetType()
	return strings.HasPrefix(variableType, "address")
}

func (v *Variable) IsDynamicArray() bool {
	return false
}

func (v *Variable) IsArrayType() bool {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type contains square brackets, indicating an array
	return strings.Contains(variableType, "[]") || strings.Contains(variableType, "[")
}

func (v *Variable) IsMappingType() bool {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type starts with "mapping"
	return strings.HasPrefix(variableType, "mapping")
}

func (v *Variable) IsStructType() bool {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type starts with "mapping"
	return strings.HasPrefix(variableType, "struct")
}

func (v *Variable) GetArrayLength() int64 {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type contains square brackets, indicating an array
	if strings.Contains(variableType, "[]") || strings.Contains(variableType, "[") {
		// Parse the array length from the type
		// ...
		return 1
	}

	return 0
}

func (v *Variable) GetStructMembers() []*Variable {
	variableType := v.StateVariable.GetType()

	// Check if the variable's type starts with "struct"
	if strings.HasPrefix(variableType, "struct") {
		// Parse the members from the type
		// ...
		return []*Variable{}
	}

	return []*Variable{}
}
