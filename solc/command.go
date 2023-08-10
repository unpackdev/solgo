package solc

// Commander interface defines the methods for executing solc-select commands.
type Commander interface {
	Current() string
	Install(version string) (bool, []string, error)
	Use(version string) (bool, []string, []string, error)
	Versions() ([]Version, error)
	Upgrade() (bool, error)
}
