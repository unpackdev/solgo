package audit

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/0x19/solc-switch"
	"github.com/google/uuid"
	"github.com/unpackdev/solgo"
)

// Slither represents a wrapper around the Slither static analysis tool.
type Slither struct {
	ctx      context.Context // Context for executing commands.
	config   *Config         // Configuration for the Slither tool.
	compiler *solc.Solc      // Instance of the solc compiler.
}

// NewSlither initializes a new Slither instance with the given context and configuration.
// It checks for the presence of Slither on the machine and returns an error if not found.
func NewSlither(ctx context.Context, compiler *solc.Solc, config *Config) (*Slither, error) {
	toReturn := &Slither{
		ctx:      ctx,
		config:   config,
		compiler: compiler,
	}

	if config.GetTempDir() == "" {
		return nil, ErrTempDirNotSet
	}

	return toReturn, nil
}

// IsInstalled checks if Slither is installed on the machine by querying its version.
// Returns true if installed, false otherwise.
func (s *Slither) IsInstalled() bool {
	cmd := exec.CommandContext(s.ctx, "slither", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// Version retrieves the installed version of Slither.
// Returns the version string or an error if unable to determine the version.
func (s *Slither) Version() (string, error) {
	cmd := exec.CommandContext(s.ctx, "slither", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	version := strings.TrimSpace(out.String())
	return version, nil
}

// Analyze performs a static analysis on the given sources using Slither.
// It writes the sources to a temporary directory, runs Slither, and then cleans up.
// Returns the analysis response, raw output, and any errors encountered.
func (s *Slither) Analyze(sources *solgo.Sources) (*Report, []byte, error) {
	if sources == nil {
		return nil, nil, ErrSourcesNotSet
	}

	// Ensure sources are prepared for analysis.
	if !sources.ArePrepared() {
		if err := sources.Prepare(); err != nil {
			return nil, nil, err
		}
	}

	// At this stage we should prepare compiler related settings.
	// First we are going to start with understanding which compiler version we are going to use.
	// To do so we'll extract it from parsed source.
	solVersion := s.config.GetCompilerVersion()

	if solVersion == "" {
		var err error
		solVersion, err = sources.GetSolidityVersion()
		if err != nil {
			return nil, nil, err
		}
	}

	solcPath, err := s.compiler.GetBinary(solVersion)
	if err != nil {
		return nil, nil, err
	}

	// Write sources to a temporary directory for Slither to analyze.
	dirName := strings.ToLower(filepath.Base(sources.EntrySourceUnitName))
	randomUuid, _ := uuid.NewRandom()
	dir := filepath.Clean(filepath.Join(s.config.GetTempDir(), randomUuid.String(), dirName))
	if err := sources.WriteToDir(dir); err != nil {
		return nil, nil, err
	}

	args := []string{dir, "--solc-solcs-bin", solcPath}
	sanitizedArgs, err := s.config.SanitizeArguments(s.config.Arguments)
	if err != nil {
		return nil, nil, err
	}
	args = append(args, sanitizedArgs...)

	if err := s.config.Validate(); err != nil {
		return nil, nil, err
	}

	// #nosec G204
	// G204 (CWE-78): Subprocess launched with variable (Confidence: HIGH, Severity: MEDIUM)
	// We did sanitization and verification of the arguments above, so we are safe to use them.
	cmd := exec.CommandContext(s.ctx, "slither", args...)

	// Errors from the output return only exit codes and therefore makes no sense at this moment
	// to deal with them. This should be changed in the future, when we have nothing else to
	// work on and we are bored or someone from community starts to complain about it :)
	output, _ := cmd.CombinedOutput()

	// Parse the output into the Response struct.
	response, err := NewResponse(output)
	if err != nil {
		return nil, nil, err
	}

	// Clean up the temporary directory.
	if err := sources.TruncateDir(dir); err != nil {
		return nil, nil, err
	}
	if err := os.Remove(dir); err != nil {
		return nil, nil, err
	}

	return response, output, nil
}
