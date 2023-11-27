package storage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func calculateSlot(variable *Variable, currentSlot int64, previousVars []*Variable) (int64, int64, []*Variable) {
	isNewSlotNeeded := variable.IsMappingType() || variable.IsDynamicArray()

	// If this is a mapping or dynamic array, or no previous vars, assign a new slot
	if isNewSlotNeeded && len(previousVars) == 0 {
		return currentSlot, 0, []*Variable{variable}
	}

	if canBePacked(variable, previousVars) {
		offset := calculateOffset(previousVars)
		return currentSlot, offset, append(previousVars, variable)
	}

	return currentSlot + 1, 0, []*Variable{variable}
}

func calculateOffset(previousVars []*Variable) int64 {
	totalUsedBits := int64(0)
	for _, prevVar := range previousVars {
		bitSize, _ := prevVar.GetAST().GetTypeName().StorageSize()
		totalUsedBits += bitSize
	}

	return totalUsedBits
}

func canBePacked(variable *Variable, previousVars []*Variable) bool {
	totalUsedBits := int64(0)
	for _, prevVar := range previousVars {
		bitSize, _ := prevVar.GetAST().GetTypeName().StorageSize()
		totalUsedBits += bitSize
	}

	bitSize, _ := variable.GetAST().GetTypeName().StorageSize()
	return totalUsedBits+bitSize <= 256
}

func convertStorageToValue(slot *SlotDescriptor, storageValue []byte) error {
	if len(storageValue) == 0 {
		return errors.New("storage value is empty")
	}

	switch {
	case strings.HasPrefix(slot.Type, "uint") || strings.HasPrefix(slot.Type, "int"):
		slot.Value = new(big.Int).SetBytes(storageValue)

	case strings.HasPrefix(slot.Type, "bool"):
		if slot.Offset >= 8*int64(len(storageValue)) {
			return fmt.Errorf("offset %d out of range for storage value", slot.Offset)
		}

		// Convert from big-endian to little-endian
		littleEndianValue := make([]byte, len(storageValue))
		for i, b := range storageValue {
			littleEndianValue[len(storageValue)-1-i] = b
		}

		// Determine the byte index and bit position within the byte
		byteIndex := slot.Offset / 8
		bitPos := slot.Offset % 8

		// Extract the bit and assign the boolean value
		slot.Value = (littleEndianValue[byteIndex] & (1 << bitPos)) != 0

	case strings.HasPrefix(slot.Type, "address") || strings.HasPrefix(slot.Type, "contract"):
		if len(storageValue) < 20 {
			return errors.New("storage value too short for an Ethereum address")
		}
		slot.Value = common.BytesToAddress(storageValue)

	case strings.HasPrefix(slot.Type, "bytes"):
		slot.Value = storageValue

	case strings.HasPrefix(slot.Type, "string"):
		// Assume decodeSolidityString returns an error if decoding fails
		decodedString, err := decodeSolidityString(storageValue)
		if err != nil {
			return fmt.Errorf("error decoding string: %v", err)
		}
		slot.Value = decodedString

	default:
		slot.Value = storageValue
	}

	return nil
}

func decodeSolidityString(storageValue []byte) (string, error) {
	// Check if storageValue has exactly 32 bytes
	if len(storageValue) != 32 {
		return "", errors.New("storage value is not 32 bytes long")
	}

	length := binary.BigEndian.Uint32(storageValue[len(storageValue)-4:])
	if length/2 < 31 {
		return string(storageValue[length/2]), nil
	}

	return "", errors.New("string spans multiple slots, handling not implemented in this function")
}
