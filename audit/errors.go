package audit

import "errors"

var (
	// ErrSlitherNotInstalled is returned when slither is not installed on the machine
	ErrSlitherNotInstalled = errors.New("slither is not installed - please install slither using `pip3 install slither-analyzer`")

	// ErrTempDirNotSet is returned when temp directory is not set
	ErrTempDirNotSet = errors.New("directory where contracts will be temporary stored is not set")

	// ErrSourcesNotSet is returned when sources are not set
	ErrSourcesNotSet = errors.New("sources are not set")
)
