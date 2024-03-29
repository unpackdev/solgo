package utils

import (
	"strconv"
	"strings"
)

// SemanticVersion represents a semantic version, following the Major.Minor.Patch
// format. Optionally, a version can include a commit revision as a metadata string
// appended after a plus sign (+).
type SemanticVersion struct {
	Major  int    `json:"major"`    // Major version, incremented for incompatible API changes.
	Minor  int    `json:"minor"`    // Minor version, incremented for backwards-compatible enhancements.
	Patch  int    `json:"patch"`    // Patch version, incremented for backwards-compatible bug fixes.
	Commit string `json:"revision"` // Optional commit revision for tracking specific builds.
}

// String returns the string representation of the SemanticVersion, excluding the
// commit revision. It adheres to the "Major.Minor.Patch" format.
func (v SemanticVersion) String() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
}

// ParseSemanticVersion parses a version string into a SemanticVersion struct.
// The function supports version strings in the "Major.Minor.Patch" format, with
// an optional commit revision appended after a plus sign (+). The 'v' prefix
// in version strings is also supported and ignored during parsing.
func ParseSemanticVersion(version string) SemanticVersion {
	versions := strings.Split(version, "+")
	var commit string
	if len(versions) > 1 {
		version = versions[0]
		commit = versions[1]
	}

	version = strings.Replace(version, "v", "", 1)
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return SemanticVersion{}
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	return SemanticVersion{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Commit: commit,
	}
}

// IsSemanticVersionGreaterOrEqualTo checks if the version represented by a string
// is greater than or equal to the provided SemanticVersion. The comparison takes
// into account the Major, Minor, and Patch components of the semantic version.
func IsSemanticVersionGreaterOrEqualTo(versionStr string, version SemanticVersion) bool {
	parsedVersion := ParseSemanticVersion(versionStr)

	if parsedVersion.Major == version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch >= version.Patch {
		return true
	}

	if parsedVersion.Major == version.Major && parsedVersion.Minor > version.Minor {
		return true
	}

	if parsedVersion.Major > version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch == version.Patch {
		return true
	}

	return false
}

// IsSemanticVersionLowerOrEqualTo checks if the version represented by a string
// is lower than or equal to the provided SemanticVersion. The comparison considers
// the Major, Minor, and Patch components of the semantic version.
func IsSemanticVersionLowerOrEqualTo(versionStr string, version SemanticVersion) bool {
	parsedVersion := ParseSemanticVersion(versionStr)
	if parsedVersion.Major == version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch < version.Patch {
		return true
	}

	if parsedVersion.Major == version.Major && parsedVersion.Minor < version.Minor {
		return true
	}

	if parsedVersion.Major < version.Major && parsedVersion.Minor == version.Minor && parsedVersion.Patch == version.Patch {
		return true
	}

	return false
}
