// Package solc provides functionality related to the Solidity compiler (solc).
package solc

import "errors"

// MockCommand provides a mock implementation of the Commander interface.
// It simulates the behavior of solc-select commands for testing purposes.
type MockCommand struct {
	// current represents the currently selected version of solc in the mock environment.
	current string
}

// Current returns the current version of solc in use in the mock environment.
func (mc *MockCommand) Current() string {
	return mc.current
}

// Install simulates the installation of a specific version of solc.
// If the version is "0.8.19", it simulates a successful installation.
// Otherwise, it simulates an installation error.
func (mc *MockCommand) Install(version string) (bool, []string, error) {
	if version == "0.8.19" {
		return true, []string{"Installing solc '0.8.19'...", "Version '0.8.19' installed."}, nil
	}
	return false, []string{"Error installing version."}, errors.New("installation error")
}

// Use simulates the process of switching to a specific version of solc.
// If the version is "0.8.19", it simulates a successful switch.
// Otherwise, it simulates an error in switching.
func (mc *MockCommand) Use(version string) (bool, []string, []string, error) {
	if version == "0.8.19" {
		return true, []string{}, []string{"Switched global version to 0.8.19"}, nil
	}
	return false, []string{}, []string{"Error switching version."}, errors.New("switch error")
}

// Versions simulates the retrieval of available solc versions.
// It returns a predefined list of versions for testing purposes.
func (mc *MockCommand) Versions() ([]Version, error) {
	return []Version{
		{Release: "0.8.19", Current: true},
		{Release: "0.8.18", Current: false},
	}, nil
}

// Upgrade simulates the process of upgrading to the latest version of solc.
// In this mock implementation, it always simulates a successful upgrade.
func (mc *MockCommand) Upgrade() (bool, error) {
	return true, nil
}
