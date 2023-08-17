package solc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/txpull/solgo"
)

// Compiler represents a Solidity compiler instance.
type Compiler struct {
	ctx             context.Context // The context for the compiler.
	solc            *Select         // The solc selector.
	sources         *solgo.Sources  // The Solidity sources to compile.
	config          *Config         // The configuration for the compiler.
	compilerVersion string          // The version of the compiler to use.
}

// NewCompiler creates a new Compiler instance with the given context, configuration, and sources.
func NewCompiler(ctx context.Context, config *Config, sources *solgo.Sources) (*Compiler, error) {
	solc, err := NewSelect()
	if err != nil {
		return nil, err
	}

	// Ensure that the sources are prepared for future consumption in case they are not already.
	if !sources.ArePrepared() {
		if err := sources.Prepare(); err != nil {
			return nil, err
		}
	}

	return &Compiler{
		ctx:     ctx,
		solc:    solc,
		sources: sources,
		config:  config,
	}, nil
}

// SetCompilerVersion sets the version of the solc compiler to use.
func (v *Compiler) SetCompilerVersion(version string) {
	v.compilerVersion = version
}

// GetCompilerVersion returns the currently set version of the solc compiler.
func (v *Compiler) GetCompilerVersion() string {
	return v.compilerVersion
}

// GetContext returns the context associated with the compiler.
func (v *Compiler) GetContext() context.Context {
	return v.ctx
}

// GetSources returns the Solidity sources associated with the compiler.
func (v *Compiler) GetSources() *solgo.Sources {
	return v.sources
}

// GetSolc returns the solc selector associated with the compiler.
func (v *Compiler) GetSolc() *Select {
	return v.solc
}

// Compile compiles the Solidity sources using the configured compiler version and arguments.
func (v *Compiler) Compile() (*CompilerResults, error) {
	combinedSource := v.sources.GetCombinedSource()

	if v.compilerVersion == "" {
		sv, err := v.sources.GetSolidityVersion()
		if err != nil {
			return nil, err
		}
		v.compilerVersion = sv
	}

	if v.compilerVersion == "" {
		return nil, fmt.Errorf("no compiler version specified")
	}

	if _, _, _, err := v.solc.Use(v.compilerVersion); err != nil {
		return nil, err
	}

	args := []string{}
	sanitizedArgs, err := v.config.SanitizeArguments(v.config.Arguments)
	if err != nil {
		return nil, err
	}
	args = append(args, sanitizedArgs...)

	if err := v.config.Validate(); err != nil {
		return nil, err
	}

	// Prepare the command
	cmd := exec.Command("solc", args...)

	// Set the combined source as input
	cmd.Stdin = strings.NewReader(combinedSource)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		var errors []string
		var warnings []string

		// Parsing the error message to extract line and column information.
		errorMessage := stderr.String()
		if strings.Contains(errorMessage, "Error:") {
			errors = append(errors, errorMessage)
		} else if strings.HasPrefix(errorMessage, "Warning:") {
			warnings = append(warnings, errorMessage)
		}

		// Construct the CompilerResults structure with errors and warnings.
		results := &CompilerResults{
			RequestedVersion: v.compilerVersion,
			Errors:           errors,
			Warnings:         warnings,
		}
		return results, err
	}

	// Parse the output
	var compilationOutput struct {
		Contracts map[string]struct {
			Bin string      `json:"bin"`
			Abi interface{} `json:"abi"`
		} `json:"contracts"`
		Errors  []string `json:"errors"`
		Version string   `json:"version"`
	}

	err = json.Unmarshal(out.Bytes(), &compilationOutput)
	if err != nil {
		return nil, err
	}

	// Extract the first contract's results (assuming one contract for simplicity)
	var firstContractKey string
	for key := range compilationOutput.Contracts {
		firstContractKey = key
		break
	}

	if firstContractKey == "" {
		return nil, fmt.Errorf("no contracts found")
	}

	// Separate errors and warnings
	var errors, warnings []string
	for _, msg := range compilationOutput.Errors {
		if strings.Contains(msg, "Warning:") {
			warnings = append(warnings, msg)
		} else {
			errors = append(errors, msg)
		}
	}

	abi, err := json.Marshal(compilationOutput.Contracts[firstContractKey].Abi)
	if err != nil {
		return nil, err
	}

	results := &CompilerResults{
		RequestedVersion: v.compilerVersion,
		CompilerVersion:  compilationOutput.Version,
		Bytecode:         compilationOutput.Contracts[firstContractKey].Bin,
		ABI:              string(abi),
		ContractName:     strings.ReplaceAll(firstContractKey, "<stdin>:", ""),
		Errors:           errors,
		Warnings:         warnings,
	}

	return results, nil
}

// CompilerResults represents the results of a solc compilation.
type CompilerResults struct {
	RequestedVersion string   `json:"requested_version"`
	CompilerVersion  string   `json:"compiler_version"`
	Bytecode         string   `json:"bytecode"`
	ABI              string   `json:"abi"`
	ContractName     string   `json:"contract_name"`
	Errors           []string `json:"errors"`
	Warnings         []string `json:"warnings"`
}

// HasErrors returns true if there are compilation errors.
func (v *CompilerResults) HasErrors() bool {
	if v == nil {
		return false
	}

	return len(v.Errors) > 0
}

// HasWarnings returns true if there are compilation warnings.
func (v *CompilerResults) HasWarnings() bool {
	if v == nil {
		return false
	}

	return len(v.Warnings) > 0
}

// GetErrors returns the compilation errors.
func (v *CompilerResults) GetErrors() []string {
	return v.Errors
}

// GetWarnings returns the compilation warnings.
func (v *CompilerResults) GetWarnings() []string {
	return v.Warnings
}

// GetABI returns the compiled contract's ABI (Application Binary Interface) in JSON format.
func (v *CompilerResults) GetABI() string {
	return v.ABI
}

// GetBytecode returns the compiled contract's bytecode.
func (v *CompilerResults) GetBytecode() string {
	return v.Bytecode
}

// GetContractName returns the name of the compiled contract.
func (v *CompilerResults) GetContractName() string {
	return v.ContractName
}

// GetRequestedVersion returns the requested compiler version used for compilation.
func (v *CompilerResults) GetRequestedVersion() string {
	return v.RequestedVersion
}

// GetCompilerVersion returns the actual compiler version used for compilation.
func (v *CompilerResults) GetCompilerVersion() string {
	return v.CompilerVersion
}
