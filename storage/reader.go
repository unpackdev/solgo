package storage

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ir"
	"go.uber.org/zap"
)

type Reader struct {
	ctx           context.Context
	storage       *Storage
	descriptor    *Descriptor
	storageLayout map[string]*SlotInfo
}

func NewReader(ctx context.Context, s *Storage, d *Descriptor) (*Reader, error) {
	return &Reader{
		ctx:           ctx,
		storage:       s,
		descriptor:    d,
		storageLayout: make(map[string]*SlotInfo),
	}, nil
}

func (r *Reader) GetDescriptor() *Descriptor {
	return r.descriptor
}

func (r *Reader) GetStorageVariables() error {
	contract := r.descriptor.Contract

	if contract == nil {
		return fmt.Errorf("failed to get storage variables as contract is not set")
	}

	descriptor := contract.GetDescriptor()
	if descriptor == nil {
		return fmt.Errorf("failed to get storage variables as contract descriptor is not set (parsing did not occur or failed)")
	}

	detector := descriptor.Detector
	if detector == nil {
		return fmt.Errorf("failed to get storage variables as contract detector is not set (parsing did not occur or failed)")
	}

	ir := detector.GetIR()
	if ir == nil {
		return fmt.Errorf("failed to get storage variables as contract IR is not set (parsing did not occur or failed)")
	}

	for _, subContract := range ir.GetRoot().GetContracts() {
		if subContract != nil {
			for _, stateVariable := range subContract.GetStateVariables() {
				if stateVariable == nil {
					continue
				}

				zap.L().Debug(
					"Found state variable",
					zap.String("contract_address", contract.GetAddress().Hex()),
					zap.String("contract_name", subContract.GetName()),
					zap.String("name", stateVariable.GetName()),
					zap.String("type", stateVariable.GetType()),
					zap.Any("visibility", stateVariable.GetVisibility()),
					zap.Any("storage", stateVariable.GetStateMutability()),
					zap.Any("constant", stateVariable.IsConstant()),
				)

				r.descriptor.StateVariables[subContract.Name] = append(
					r.descriptor.StateVariables[subContract.Name],
					&Variable{
						Contract:      subContract,
						EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
						StateVariable: stateVariable,
					},
				)

				if !stateVariable.IsConstant() && stateVariable.GetStateMutability() == ast_pb.Mutability_IMMUTABLE {
					r.descriptor.TargetVariables[subContract.Name] = append(
						r.descriptor.TargetVariables[subContract.Name],
						&Variable{
							Contract:      subContract,
							EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
							StateVariable: stateVariable,
						},
					)
				} else {
					r.descriptor.ConstantStorageSlotVariables[subContract.Name] = append(
						r.descriptor.ConstantStorageSlotVariables[subContract.Name],
						&Variable{
							Contract:      subContract,
							EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
							StateVariable: stateVariable,
						},
					)
				}
			}
		}
	}

	return nil
}

func (r *Reader) GetStorageLayout() (map[string]*SlotInfo, error) {
	currentSlot := int64(0)

	for contractName, variables := range r.descriptor.GetTargetVariables() {
		for _, variable := range variables {
			slot := calculateSlot(variable, currentSlot)
			size := calculateSize(variable)
			offset := calculateOffset(variable)

			slotInfo := &SlotInfo{
				Name:   variable.GetName(),
				Type:   variable.GetType(),
				Slot:   slot,
				Size:   size,
				Offset: offset,
			}

			r.storageLayout[contractName] = slotInfo
			currentSlot += size // Update currentSlot for the next variable
		}
	}

	return r.storageLayout, nil
}

func calculateSize(variable *Variable) int64 {
	// Elementary types generally occupy one slot
	if variable.IsElementaryType() {
		return 1
	}

	// For arrays, the size depends on whether it's a fixed-size or dynamic array
	if variable.IsArrayType() {
		// For simplicity, let's assume IsArrayType and GetArrayLength are implemented
		// Fixed-size arrays occupy a number of slots equal to their length
		// Dynamic arrays occupy a single slot for the length, data is stored separately
		if length := variable.GetArrayLength(); length > 0 {
			return length // For fixed-size arrays
		} else {
			return 1 // For dynamic arrays
		}
	}

	// For mappings, they don't occupy slots in a linear fashion like other types
	if variable.IsMappingType() {
		return 1 // Mappings use a single slot for the hash table pointer
	}

	// For structs, you must recursively calculate the size of each member
	if variable.IsStructType() {
		size := int64(0)
		for _, member := range variable.GetStructMembers() {
			size += calculateSize(member)
		}
		return size
	}

	// Default to 1 slot if the type is not recognized
	// This is a simplification, in practice, you'll want to handle all types explicitly
	return 1
}

func calculateSlot(variable *Variable, currentSlot int64) int64 {
	if variable.IsElementaryType() {
		// Elementary types are stored sequentially
		return currentSlot
	}

	if variable.IsArrayType() {
		// For dynamic arrays, the slot only stores the length or a pointer
		// The actual content starts at a keccak256 hash of the slot number
		/* 		if variable.IsDynamicArray() {
			return currentSlot
		} */
		// For fixed-size arrays, the array elements occupy sequential slots
		// starting from currentSlot
		return currentSlot // Assumes the caller will handle incrementing currentSlot for each element
	}

	if variable.IsMappingType() {
		// Mappings use one slot for the pointer, but the actual data
		// is stored at a location derived from a keccak256 hash
		return currentSlot
	}

	if variable.IsStructType() {
		// Structs occupy sequential slots starting from currentSlot
		// The actual slot for each member needs to be calculated based on the size and order of previous members
		return currentSlot // Assumes the caller will handle incrementing currentSlot for each struct member
	}

	// Fallback to current slot for unrecognized types
	// This is a simplification, and you should ideally handle all types explicitly
	return currentSlot
}

func calculateOffset(variable *Variable) int64 {
	if variable.IsPackedType() {
		// Assuming you have a method to get the size of the variable in bytes
		variableSize := getSizeInBytes(variable)
		_ = variableSize

		// Assuming you have a method to get previous packed variables in the same slot
		previousVariables := getPreviousPackedVariables(variable)

		offset := int64(0)
		for _, prevVar := range previousVariables {
			prevVarSize := getSizeInBytes(prevVar)
			offset += prevVarSize // Increment offset by the size of each previous variable
		}

		return offset
	}

	// Non-packed types do not have an offset within a slot
	return 0
}

func getSizeInBytes(variable *Variable) int64 {
	variableType := variable.StateVariable.GetType() // Assuming GetType returns the type as a string

	// Mapping of Solidity types to their sizes in bytes
	typeSizes := map[string]int64{
		"bool": 1,
		"int8": 1, "uint8": 1,
		"int16": 2, "uint16": 2,
		"int32": 4, "uint32": 4,
		"int64": 8, "uint64": 8,
		"int128": 16, "uint128": 16,
		"int256": 32, "uint256": 32,
		// Add other types as needed
	}

	if size, ok := typeSizes[variableType]; ok {
		return size
	}

	// Default to 0 for unrecognized types
	return 0
}

func getPreviousPackedVariables(variable *Variable) []*Variable {
	// Assuming you have access to the contract's variables and their order
	contractVariables := getContractVariables(variable.Contract)

	var previousPackedVariables []*Variable
	for _, v := range contractVariables {
		if v == variable {
			break
		}
		if v.IsPackedType() {
			previousPackedVariables = append(previousPackedVariables, v)
		}
	}

	return previousPackedVariables
}

func getContractVariables(contract *ir.Contract) []*Variable {
	// Initialize a slice to hold the variables
	var variables []*Variable

	// Assuming contract has a method or field to get its variables
	// For example, contract.GetStateVariables() might return a slice of *ir.StateVariable
	stateVariables := contract.GetStateVariables()

	// Iterate over the state variables and wrap them in the Variable struct
	for _, stateVar := range stateVariables {
		variable := &Variable{
			Contract:      contract,
			StateVariable: stateVar,
			// Set other fields of Variable struct as necessary
		}
		variables = append(variables, variable)
	}

	return variables
}
