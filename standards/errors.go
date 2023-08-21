package standards

import "errors"

var (

	// ErrStandardNotFound is returned when a standard is not found.
	ErrStandardNotFound = errors.New("standard not found")
)
