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

	if len(parts) != 3 {
		return SemanticVersion{}
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	return SemanticVersion{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

// IsSemanticVersionGreaterOrEqualTo checks if the version represented by the string is greater than or equal to the provided SemanticVersion.
func IsSemanticVersionGreaterOrEqualTo(versionStr string, version SemanticVersion) bool {
	parsedVersion := ParseSemanticVersion(versionStr)

	if parsedVersion.Major == version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch >= version.Patch {
		return true
	}

	if parsedVersion.Major == version.Major && parsedVersion.Minor >= version.Minor {
		return true
	}

	if parsedVersion.Major >= version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch == version.Patch {
		return true
	}

	return false
}

// IsSemanticVersionLowerOrEqualTo checks if the version represented by the string is lower than or equal to the provided SemanticVersion.
func IsSemanticVersionLowerOrEqualTo(versionStr string, version SemanticVersion) bool {
	parsedVersion := ParseSemanticVersion(versionStr)

	if parsedVersion.Major == version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch >= version.Patch {
		return true
	}

	if parsedVersion.Major == version.Major && parsedVersion.Minor <= version.Minor {
		return true
	}

	if parsedVersion.Major <= version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch == version.Patch {
		return true
	}

	return false
}
