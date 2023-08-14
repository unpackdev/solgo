package audit

import (
	"fmt"
	"os"
)

// Define a list of allowed arguments for slither
var allowedArgs = map[string]bool{
	"--json": true,
	"-":      true,
}

// Define a list of required arguments for slither
var requiredArgs = map[string]bool{
	"--json": true,
	"-":      true,
}

type Config struct {
	tempDir   string   // temp directory to store temporary contract files
	Arguments []string // arguments to pass to slither
}

// NewDefaultConfig returns a default configuration for slither
func NewDefaultConfig(tempDir string) (*Config, error) {
	if _, err := os.Stat(tempDir); err != nil {
		return nil, err
	}

	toReturn := &Config{
		tempDir: tempDir,
		Arguments: []string{
			"--json", "-", // output to stdout
		},
	}

	if _, err := toReturn.SanitizeArguments(toReturn.Arguments); err != nil {
		return nil, err
	}

	if err := toReturn.Validate(); err != nil {
		return nil, err
	}

	return toReturn, nil
}

// SanitizeArguments sanitizes the provided arguments against a list of allowed arguments.
// It returns an error if any of the provided arguments are not in the allowed list.
func (c *Config) SanitizeArguments(args []string) ([]string, error) {
	var sanitizedArgs []string
	for _, arg := range args {
		if _, ok := allowedArgs[arg]; !ok {
			return nil, fmt.Errorf("invalid argument: %s", arg)
		}
		sanitizedArgs = append(sanitizedArgs, arg)
	}
	return sanitizedArgs, nil
}

func (c *Config) Validate() error {
	sanitized, err := c.SanitizeArguments(c.Arguments)
	if err != nil {
		return err
	}

	// Convert the sanitized slice into a map for easier lookup
	sanitizedMap := make(map[string]bool)
	for _, arg := range sanitized {
		sanitizedMap[arg] = true
	}

	for arg, _ := range requiredArgs {
		if _, ok := sanitizedMap[arg]; !ok {
			return fmt.Errorf("missing required argument: %s", arg)
		}
	}

	return nil
}

// GetTempDir returns the temp directory to store temporary contract files
func (c *Config) GetTempDir() string {
	return c.tempDir
}

// SetArguments sets the arguments to pass to slither
func (c *Config) SetArguments(args []string) {
	c.Arguments = args
}

// AppendArguments appends new arguments to already existing arguments
func (c *Config) AppendArguments(args []string) {
	c.Arguments = append(c.Arguments, args...)
}

// GetArguments returns the arguments to pass to slither
func (c *Config) GetArguments() []string {
	return c.Arguments
}
