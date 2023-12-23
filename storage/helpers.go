package storage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// calculateSlot determines the appropriate storage slot and offset for a given variable.
// It handles different cases like mapping, dynamic arrays, and variable packing within a slot.
// Returns the slot number, offset within the slot, and an updated slice of variables.
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

// calculateOffset computes the total bit offset for a set of variables.
// Used to determine the offset within a storage slot for variable packing.
func calculateOffset(previousVars []*Variable) int64 {
	totalUsedBits := int64(0)
	for _, prevVar := range previousVars {
		bitSize, _ := prevVar.GetAST().GetTypeName().StorageSize()
		totalUsedBits += bitSize
	}

	return totalUsedBits
}

// canBePacked checks if a given variable can be packed into the same storage slot
// as previous variables. Considers the total bit size and special cases like boolean variables.
func canBePacked(variable *Variable, previousVars []*Variable) bool {
	totalUsedBits := int64(0)
	for _, prevVar := range previousVars {
		bitSize, _ := prevVar.GetAST().GetTypeName().StorageSize()
		totalUsedBits += bitSize
	}

	bitSize, _ := variable.GetAST().GetTypeName().StorageSize()

	// Check if the total size exceeds the 256-bit boundary of a single slot.
	if totalUsedBits+bitSize > 256 {
		return false
	}

	// Special handling for bool variables
	if variable.Type == "bool" {
		return totalUsedBits%256+bitSize <= 256
	}

	// General case for other types: Allow packing different-sized variables
	// as long as they fit within the 256-bit boundary of a single slot
	return totalUsedBits%256+bitSize <= 256
}

// convertStorageToValue converts raw storage bytes into a meaningful value based on the slot's type.
// Handles various Ethereum data types like integers, booleans, addresses, etc.
// Returns an error if the conversion is not possible or if the data format is not as expected.
func convertStorageToValue(storage *Storage, contractAddress common.Address, slot *SlotDescriptor, storageValue []byte) error {
	if len(storageValue) == 0 {
		return errors.New("storage value is empty")
	}

	switch {
	case strings.HasPrefix(slot.Type, "uint") || strings.HasPrefix(slot.Type, "int"):
		slot.Value = new(big.Int).SetBytes(storageValue)
		return nil

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
		if len(storageValue) < 32 {
			return errors.New("storage value too short for an Ethereum address")
		}

		var addressBytes []byte
		if slot.Offset == 0 {
			// If the address starts at the beginning of the slot, use the last 20 bytes
			addressBytes = storageValue[len(storageValue)-20:]
		} else {
			// Calculate the start index based on the offset (in bits)
			startIndex := slot.Offset / 8

			// Adjust the endIndex to extract 20 bytes after the offset
			endIndex := startIndex + 20
			if endIndex > int64(len(storageValue)) {
				return errors.New("storage value too short for an Ethereum address with given offset")
			}
			addressBytes = storageValue[startIndex:endIndex]
		}

		slot.Value = common.BytesToAddress(addressBytes)

	case strings.HasPrefix(slot.Type, "bytes"):
		slot.Value = storageValue

	case strings.HasPrefix(slot.Type, "string"):
		decodedString, err := decodeSolidityString(storage, contractAddress, slot.Slot, storageValue, slot.BlockNumber)
		if err != nil {
			return fmt.Errorf("error decoding string: %v", err)
		}
		slot.Value = decodedString

	case strings.HasPrefix(slot.Type, "struct"):
		slot.Value = struct{}{}

	case strings.HasPrefix(slot.Type, "mapping"):
		slot.Value = struct{}{}

	default:
		// Fuck this shit, will figure out later on how to deal with it properly...
		if common.BytesToAddress(storageValue) == utils.ZeroAddress {
			slot.Value = big.NewInt(0)
			return nil
		}

		slot.Value = storageValue
	}

	return nil
}

// decodeSolidityString decodes a string stored in Ethereum's Solidity format from raw storage bytes.
// It handles strings that span multiple slots.
func decodeSolidityString(storage *Storage, contractAddress common.Address, startSlot int64, storageValue []byte, blockNumber *big.Int) (string, error) {
	if len(storageValue) != 32 {
		return "", errors.New("initial storage value is not 32 bytes long")
	}

	// Length is read from the last 8 bytes of the storageValue
	length := binary.BigEndian.Uint64(storageValue[24:32])

	// Guard against excessively large length values
	/* 	const maxLength = 10 * 1024 // For example, 10 KB
	   	if length > maxLength {
	   		return "", fmt.Errorf("string length %d exceeds maximum allowed length of %d", length, maxLength)
	   	} */

	if length <= 31 { // Fits in a single slot
		return string(storageValue[:length]), nil
	}

	// For now.... Don't have time fixing up the multi strings...
	return "", nil
}
