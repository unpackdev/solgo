package storage

// Options defines the configuration parameters for the storage system.
// It includes settings to determine the use of a simulator for storage operations.
type Options struct {
	// UseSimulator indicates whether to use the simulator for storage operations.
	// When set to true, storage operations are simulated. This is useful for testing
	// or development purposes where actual storage resources are not required.
	// The default value is false.
	UseSimulator bool `json:"use_simulator"`
}

// NewDefaultOptions creates and returns a new instance of Options with default settings.
// By default, the UseSimulator field is set to false, indicating that the simulator
// is not used, and actual storage operations are performed.
func NewDefaultOptions() *Options {
	return &Options{
		UseSimulator: false,
	}
}

// NewSimulatorOptions creates and returns a new instance of Options with the
// UseSimulator field set to true. This configuration is useful for scenarios
// where storage operations need to be simulated, such as during testing or
// when actual storage resources are not available or necessary.
func NewSimulatorOptions() *Options {
	return &Options{
		UseSimulator: true,
	}
}
