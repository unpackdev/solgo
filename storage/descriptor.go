package storage

import (
	"math/big"
	"strings"

	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/ir"
)

type Variable struct {
	Contract      *ir.Contract
	EntryContract bool
	*ir.StateVariable
}

type Descriptor struct {
	Contract                     *contracts.Contract
	Block                        *big.Int
	StateVariables               map[string][]*Variable
	TargetVariables              map[string][]*Variable
	ConstantStorageSlotVariables map[string][]*Variable
}

func (s *Descriptor) GetStateVariables() map[string][]*Variable {
	return s.StateVariables
}

func (s *Descriptor) GetTargetVariables() map[string][]*Variable {
	return s.TargetVariables
}

func (s *Descriptor) GetConstantStorageSlotVariables() map[string][]*Variable {
	return s.ConstantStorageSlotVariables
}

func (s *Descriptor) GetContract() *contracts.Contract {
	return s.Contract
}

func (s *Descriptor) GetBlock() *big.Int {
	return s.Block
}

func (v *Variable) IsElementaryType() bool {
	variableType := v.StateVariable.GetType()

	// List of elementary types in Solidity
	elementaryTypes := []string{
		"bool", "address", "bytes32", // Fixed-size byte arrays are elementary
		"int", "uint", // These include all sizes, e.g., int256, uint256, etc.
		// Include all sizes explicitly if necessary
		"int8", "int16", "int32", "int64", "int128", "int256",
		"uint8", "uint16", "uint32", "uint64", "uint128", "uint256",
		//... other elementary types like fixed and ufixed of various sizes
	}

	// Check if the variable's type is in the list of elementary types
	for _, t := range elementaryTypes {
		if variableType == t {
			return true
		}
	}

	return false
}

func (v *Variable) IsPackedType() bool {
	// Assuming StateVariable has a method GetType that returns the type of the variable
	variableType := v.StateVariable.GetType()

	// List of types that are typically packed
	packedTypes := []string{
		"bool",    // 1 byte
		"int8",    // 1 byte
		"uint8",   // 1 byte
		"int16",   // 2 bytes
		"uint16",  // 2 bytes
		"int24",   // 3 bytes
		"uint24",  // 3 bytes
		"int32",   // 4 bytes
		"uint32",  // 4 bytes
		"int40",   // 5 bytes
		"uint40",  // 5 bytes
		"int48",   // 6 bytes
		"uint48",  // 6 bytes
		"int56",   // 7 bytes
		"uint56",  // 7 bytes
		"int64",   // 8 bytes
		"uint64",  // 8 bytes
		"int72",   // 9 bytes
		"uint72",  // 9 bytes
		"int80",   // 10 bytes
		"uint80",  // 10 bytes
		"int88",   // 11 bytes
		"uint88",  // 11 bytes
		"int96",   // 12 bytes
		"uint96",  // 12 bytes
		"int104",  // 13 bytes
		"uint104", // 13 bytes
		"int112",  // 14 bytes
		"uint112", // 14 bytes
		"int120",  // 15 bytes
		"uint120", // 15 bytes
		"int128",  // 16 bytes
		"uint128", // 16 bytes
		// ... Continue for types up to 31 bytes
	}
	// Check if the variable's type is in the list of packed types
	for _, t := range packedTypes {
		if variableType == t {
			return true
		}
	}

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
