package solc

type Select struct {
	current   string
	commander Commander
}

// Get the current version of solc in use.
func (s *Select) Current() string {
	return s.current
}

// NewSelect initializes and returns a new instance of the Select struct.
// The function first checks if solc-select is installed on the system.
// If solc-select is not installed, it returns an error.
// If solc-select is installed, the function fetches the list of available solc versions.
// It then identifies the currently active solc version and sets it in the returned Select struct.
// If any step fails, an error is returned.
func NewSelect() (*Select, error) {
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
