package storage

type Options struct {
	UseSimulator bool `json:"use_simulator"`
}

// NewDefaultOptions creates an Options instance with default settings.
func NewDefaultOptions() *Options {
	return &Options{
		UseSimulator: false,
	}
}

func NewSimulatorOptions() *Options {
	return &Options{
		UseSimulator: true,
	}
}
