package audit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/txpull/solgo"
)

type Slither struct {
	ctx    context.Context
	config *Config
}

func NewSlither(ctx context.Context, config *Config) (*Slither, error) {
	toReturn := &Slither{
		ctx:    ctx,
		config: config,
	}

	if config.GetTempDir() == "" {
		return nil, ErrTempDirNotSet
	}

	if !toReturn.IsInstalled() {
		return nil, ErrSlitherNotInstalled
	}

	return toReturn, nil
}

// IsInstalled checks if slither is installed on the machine
func (s *Slither) IsInstalled() bool {
	cmd := exec.CommandContext(s.ctx, "slither", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// Version retrieves the installed version of slither
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

func (s *Slither) Analyze(sources *solgo.Sources) (*Response, []byte, error) {
	if sources == nil {
		return nil, nil, ErrSourcesNotSet
	}

	// Make sure that sources are prepared for future consumption
	if !sources.ArePrepared() {
		if err := sources.Prepare(); err != nil {
			return nil, nil, err
		}
	}

	// Write sources into the temporary directory for slither to consume.
	// Later on we are going to drop this directory.
	dirName := strings.ToLower(filepath.Base(sources.EntrySourceUnitName))
	dir := filepath.Clean(filepath.Join(s.config.GetTempDir(), dirName))
	if err := sources.WriteToDir(dir); err != nil {
		return nil, nil, err
	}

	args := []string{dir}
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
	cmd := exec.CommandContext(s.ctx, "slither", args...)

	output, err := cmd.CombinedOutput()

	// @WARN: Figure out better exist status error management!
	if err != nil && err.Error() != "exit status 255" {
		fmt.Println(err.Error())

		// Error itself usually gives exit status which is less then helpful to us
		// at this point so instead, output will be returned back as error.
		return nil, nil, errors.New(string(output))
	}

	// Cast the output into the response struct
	response, err := NewResponse(output)
	if err != nil {
		return nil, nil, err
	}

	// Lets first truncate the directory to make sure that we are not going to
	// have any leftovers from the previous runs.
	if err := sources.TruncateDir(dir); err != nil {
		return nil, nil, err
	}

	// Just in case let's remove the directory once we are done with it.
	// No need for it. TruncateDir() should stay the same as purpose of it is to truncate
	// the directory and not to remove it.
	if err := os.Remove(dir); err != nil {
		return nil, nil, err
	}

	return response, output, nil
}
