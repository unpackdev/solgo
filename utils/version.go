package utils

import (
	"strconv"
	"strings"
)

// SemanticVersion represents a version in the format Major.Minor.Patch.
type SemanticVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

// ParseSemanticVersion converts a string representation of a version into a SemanticVersion struct.
// It expects the version string to be in the format "Major.Minor.Patch".
func ParseSemanticVersion(version string) SemanticVersion {
	version = strings.Replace(version, "v", "", 1)
	parts := strings.Split(version, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	return SemanticVersion{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}
