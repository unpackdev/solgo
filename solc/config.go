package solc

import (
	"fmt"
	"strings"
)

// allowedArgs defines a list of allowed arguments for solc.
var allowedArgs = map[string]bool{
	"--combined-json": true,
	"-":               true,
	"--optimize":      true,
	"--optimize-runs": true,
	"--evm-version":   true,
	"--overwrite":     true,
	"--libraries":     true,
	"--output-dir":    true,
}

// requiredArgs defines a list of required arguments for solc.
var requiredArgs = map[string]bool{
	"--overwrite":     true,
	"--combined-json": true,
	"-":               true,
}

// Config represents the configuration for the solc tool.
type Config struct {
	Arguments []string // Arguments to pass to the solc tool.
}

// NewDefaultConfig creates and returns a default configuration for solc.
func NewDefaultConfig() (*Config, error) {
	toReturn := &Config{
		Arguments: []string{
			"--overwrite", "--combined-json", "bin,abi,ast", "-", // Output to stdout.
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
// Returns an error if any of the provided arguments are not in the allowed list.
func (c *Config) SanitizeArguments(args []string) ([]string, error) {
	var sanitizedArgs []string
	for _, arg := range args {
		if strings.Contains(arg, "-") {
			if _, ok := allowedArgs[arg]; !ok {
				return nil, fmt.Errorf("invalid argument: %s", arg)
			}
		}
		sanitizedArgs = append(sanitizedArgs, arg)
	}
	return sanitizedArgs, nil
}

// Validate checks if the current configuration's arguments are valid.
// It ensures that all required arguments are present.
func (c *Config) Validate() error {
	sanitized, err := c.SanitizeArguments(c.Arguments)
	if err != nil {
		return err
	}

	// Convert the sanitized slice into a map for easier lookup.
	sanitizedMap := make(map[string]bool)
	for _, arg := range sanitized {
		sanitizedMap[arg] = true
	}

	for arg := range requiredArgs {
		if _, ok := sanitizedMap[arg]; !ok {
			return fmt.Errorf("missing required argument: %s", arg)
		}
	}

	return nil
}

// SetArguments sets the arguments to be passed to the solc tool.
func (c *Config) SetArguments(args []string) {
	c.Arguments = args
}

// AppendArguments appends new arguments to the existing set of arguments.
func (c *Config) AppendArguments(args ...string) {
	c.Arguments = append(c.Arguments, args...)
}

// GetArguments returns the arguments to be passed to the solc tool.
func (c *Config) GetArguments() []string {
	return c.Arguments
}
