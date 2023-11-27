package storage

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func calculateSlot(variable *Variable, currentSlot int64, previousVars []*Variable) (int64, int64, []*Variable) {
	if variable.IsMappingType() || variable.IsDynamicArray() {
		return currentSlot + 1, 0, append(previousVars, variable)
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
	// Check if storageValue has at least 32 bytes (minimum size for dynamic types)
	if len(storageValue) < 32 {
		return "", errors.New("storage value too short for a dynamic string")
	}

	// The first 32 bytes represent the length of the string in bytes
	length := new(big.Int).SetBytes(storageValue[:32]).Uint64()

	// Check if the total length makes sense given the storageValue size
	if uint64(len(storageValue)) < 32+length {
		return "", errors.New("storage value length mismatch")
	}

	// Extract the string data
	strBytes := storageValue[32 : 32+length]

	// Convert bytes to string, ensuring valid UTF-8 encoding
	if !isValidUTF8(strBytes) {
		return "", errors.New("invalid UTF-8 encoding in string data")
	}

	return string(strBytes), nil
}

// isValidUTF8 checks if the provided byte slice is valid UTF-8.
func isValidUTF8(data []byte) bool {
	return len(data) == len(string(data))
}
