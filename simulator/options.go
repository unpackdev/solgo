package simulator

type Options struct {
	NumberOfFaucets int    // Number of faucets to initialize
	GasLimit        uint64 // Gas limit for the simulated backend
	// Add other configuration options as needed
}

// NewDefaultOptions creates an Options instance with default settings.
func NewDefaultOptions() *Options {
	return &Options{
		NumberOfFaucets: 5,       // Default number of faucets
		GasLimit:        9000000, // Default gas limit
		// Set other defaults as necessary
	}
}
