package standards

import (
	"fmt"
	"sort"
)

// storage is a map that holds registered Ethereum standards.
var storage map[Standard]EIP

// RegisterStandard registers a new Ethereum standard to the storage.
// If the standard already exists, it returns an error.
//
// Parameters:
// - s: The Ethereum standard type.
// - cs: The details of the Ethereum standard.
//
// Returns:
// - error: An error if the standard already exists, otherwise nil.
func RegisterStandard(s Standard, cs EIP) error {
	if Exists(s) {
		return fmt.Errorf("standard %s already exists", s)
	}

	storage[s] = cs
	return nil
}

// GetStandard retrieves the details of a registered Ethereum standard.
//
// Parameters:
// - s: The Ethereum standard type.
//
// Returns:
// - ContractStandard: The details of the Ethereum standard if it exists.
// - bool: A boolean indicating if the standard exists in the storage.
func GetStandard(s Standard) (EIP, bool) {
	cs, exists := storage[s]
	return cs, exists
}

// Exists checks if a given Ethereum standard is registered in the storage.
//
// Parameters:
// - s: The Ethereum standard type.
//
// Returns:
// - bool: A boolean indicating if the standard exists in the storage.
func Exists(s Standard) bool {
	_, exists := storage[s]
	return exists
}

// GetRegisteredStandards retrieves all the registered Ethereum standards from the storage.
//
// Returns:
// - map[Standard]ContractStandard: A map of all registered Ethereum standards.
func GetRegisteredStandards() map[Standard]EIP {
	return storage
}

// GetSortedRegisteredStandards retrieves all the registered Ethereum standards from the storage in a sorted order.
func GetSortedRegisteredStandards() []EIP {
	// Create a slice to hold all EIP values
	var eips []EIP
	for _, eip := range storage {
		eips = append(eips, eip)
	}

	// Sort the slice based on the Name field of EIP
	sort.Slice(eips, func(i, j int) bool {
		return eips[i].GetType().String() < eips[j].GetType().String()
	})

	return eips
}

// StandardsLoaded returns a boolean indicating whether the storage has any registered Ethereum standards.
func StandardsLoaded() bool {
	return len(storage) > 1
}
