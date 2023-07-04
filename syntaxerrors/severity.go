package syntaxerrors

import "strings"

// SeverityLevel represents the severity of a syntax error.
type SeverityLevel int

const (
	// SeverityInfo represents a syntax error of informational level.
	SeverityInfo SeverityLevel = iota

	// SeverityWarning represents a syntax error of warning level.
	SeverityWarning

	// SeverityError represents a syntax error of error level.
	SeverityError
)

// String returns a string representation of the SeverityLevel.
func (s SeverityLevel) String() string {
	switch s {
	case SeverityInfo:
		return "info"
	case SeverityWarning:
		return "warning"
	case SeverityError:
		return "error"
	default:
		return "unknown"
	}
}

var (
	// severityMap maps error messages to their severity level.
	severityMap = map[string]SeverityLevel{
		"missing":               SeverityError,
		"mismatched":            SeverityError,
		"no viable alternative": SeverityError,
		"extraneous input":      SeverityError,
		"cannot find symbol":    SeverityWarning,
		"method not found":      SeverityWarning,
	}
)

// DetermineSeverity determines the severity of a syntax error.
// It returns the severity of the error.
func (l *SyntaxErrorListener) DetermineSeverity(msg, context string) SeverityLevel {
	// Replace the error message if needed
	msg = ReplaceErrorMessage(msg, "missing Semicolon at '}'", "missing ';' at '}'")

	// Check if the error message is in the map
	for err, severity := range severityMap {
		if strings.Contains(msg, err) {
			return severity
		}
	}

	// Context-specific rules
	if context == "FunctionDeclaration" && strings.Contains(msg, "expected") {
		return SeverityError
	}

	// Default to low severity
	return SeverityInfo
}
