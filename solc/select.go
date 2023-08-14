package solc

import (
	"os"

	"go.uber.org/zap"
)

// Select represents a utility structure that manages the version of solc in use.
// It encapsulates the current version of solc and provides functionalities to interact with solc-select.
type Select struct {
	current   string    // The current version of solc in use.
	commander Commander // Commander interface to execute shell commands.
}

// Get the current version of solc in use.
func (s *Select) Current() string {
	return s.current
}

// NewSelect initializes and returns a new instance of the Select struct.
// The function performs the following steps:
// 1. Checks if solc-select is installed on the system.
// 2. If solc-select is not installed, it returns an error.
// 3. If solc-select is installed, the function fetches the list of available solc versions.
// 4. Identifies the currently active solc version and sets it in the returned Select struct.
// 5. Checks if a Python virtual environment is set and initializes the current solc version.
// If any of the above steps fail, an error and/or warning is returned.
func NewSelect() (*Select, error) {
	if os.Getenv("VIRTUAL_ENV") == "" {
		zap.L().Warn(
			"Python virtual environment is not set. Security risk!",
			zap.String("message", "Please run 'python3 -m venv solgoenv' and 'source solgoenv/bin/activate' before running this command."),
		)
	}

	toReturn := &Select{
		commander: &RealCommand{},
	}

	versions, err := toReturn.Versions()
	if err != nil {
		return nil, err
	}

	for _, version := range versions {
		if version.Current {
			toReturn.current = version.Release
			break
		}
	}

	return toReturn, nil
}

// Install a specific version of solc.
// Returns a boolean indicating if the version was installed, a slice of output lines from solc-select, and an error if any occurred.
func (s *Select) Install(version string) (bool, []string, error) {
	return s.commander.Install(version)
}

// Use a specific version of solc. If the version does not exist, install it.
// Returns a boolean indicating success, the outputs from the install and use commands, and an error if any occurred.
func (s *Select) Use(version string) (bool, []string, []string, error) {
	return s.commander.Use(version)
}

// Versions lists all available versions of solc using solc-select.
func (s *Select) Versions() ([]Version, error) {
	return s.commander.Versions()
}

// Upgrade to the latest version of solc.
// Returns a bool indicating if solc-select is up to date, and an error if any occurred.
func (s *Select) Upgrade() (bool, error) {
	return s.commander.Upgrade()
}
