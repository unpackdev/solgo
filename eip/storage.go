package eip

import "fmt"

// storage is a map that holds registered Ethereum standards.
var storage map[Standard]ContractStandard

// RegisterStandard registers a new Ethereum standard to the storage.
// If the standard already exists, it returns an error.
//
// Parameters:
// - s: The Ethereum standard type.
// - cs: The details of the Ethereum standard.
//
// Returns:
// - error: An error if the standard already exists, otherwise nil.
func RegisterStandard(s Standard, cs ContractStandard) error {
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
func GetStandard(s Standard) (ContractStandard, bool) {
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
func GetRegisteredStandards() map[Standard]ContractStandard {
	return storage
}
