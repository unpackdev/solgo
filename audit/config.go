package audit

import "os"

type Config struct {
	tempDir   string   // temp directory to store temporary contract files
	Arguments []string // arguments to pass to slither
}

// NewDefaultConfig returns a default configuration for slither
func NewDefaultConfig(tempDir string) (*Config, error) {
	if _, err := os.Stat(tempDir); err != nil {
		return nil, err
	}

	return &Config{
		tempDir: tempDir,
		Arguments: []string{
			"--json", "-", // output to stdout
		},
	}, nil
}

// GetTempDir returns the temp directory to store temporary contract files
func (c *Config) GetTempDir() string {
	return c.tempDir
}

// GetArguments returns the arguments to pass to slither
func (c *Config) GetArguments() []string {
	return c.Arguments
}
