package inspector

// DetectorType represents the type of detector used in analyzing smart contracts.
type DetectorType string

// String returns the string representation of the detector type.
func (dt DetectorType) String() string {
	return string(dt)
}

const (
	// StateVariableDetectorType represents a detector type focused on state variables in smart contracts.
	StateVariableDetectorType DetectorType = "state_variable"

	// TransferDetectorType represents a detector type focused on transfer operations in smart contracts.
	TransferDetectorType DetectorType = "transfer"

	// MintDetectorType represents a detector type focused on minting operations in smart contracts.
	MintDetectorType DetectorType = "mint"

	// BurnDetectorType represents a detector type focused on burning operations in smart contracts.
	BurnDetectorType DetectorType = "burn"

	// ProxyDetectorType represents a detector type focused on proxy patterns in smart contracts.
	ProxyDetectorType DetectorType = "proxy"

	// OwnershipDetectorType represents a detector type focused on ownership patterns in smart contracts.
	OwnershipDetectorType DetectorType = "ownership"

	// StorageDetectorType represents a detector type focused on storage patterns in smart contracts.
	StorageDetectorType DetectorType = "storage"

	// StandardsDetectorType represents a detector type focused on standards patterns in smart contracts.
	StandardsDetectorType DetectorType = "standards"

	// TokenDetectorType represents a detector type focused on token (ERC20) patterns in smart contracts.
	TokenDetectorType DetectorType = "token"

	// AuditDetectorType represents a detector type focused on simulation of smart contract behaviors.
	AuditDetectorType DetectorType = "audit"

	// PausableDetectorType represents a detector type focused on pausable patterns in smart contracts.
	PausableDetectorType DetectorType = "pausable"

	// ExternalCallsDetectorType represents a detector type focused on external calls in smart contracts.
	ExternalCallsDetectorType DetectorType = "external_calls"

	// AntiWhaleDetectorType represents a detector type focused on anti-whale patterns in smart contracts.
	AntiWhaleDetectorType DetectorType = "anti_whale"
)

// SeverityType represents the severity level of a detection.
type SeverityType string

// String returns the string representation of the severity type.
func (st SeverityType) String() string {
	return string(st)
}

const (
	// SeverityInfo represents an informational severity level of a detection.
	SeverityInfo SeverityType = "informatinal"

	// SeverityLow represents a low severity level of a detection.
	SeverityLow SeverityType = "low"

	// SeverityMedium represents a medium severity level of a detection.
	SeverityMedium SeverityType = "medium"

	// SeverityHigh represents a high severity level of a detection.
	SeverityHigh SeverityType = "high"

	// SeverityCritical represents a critical severity level of a detection.
	SeverityCritical SeverityType = "critical"

	// SeverityOk represents an okay or non-issue severity level of a detection.
	SeverityOk SeverityType = "ok"
)

// ConfidenceLevelType represents the confidence level of a detection result.
type ConfidenceLevelType string

// String returns the string representation of the confidence level type.
func (cl ConfidenceLevelType) String() string {
	return string(cl)
}

const (
	// ConfidenceLevelLow represents a low confidence level in a detection result.
	ConfidenceLevelLow ConfidenceLevelType = "low"

	// ConfidenceLevelMedium represents a medium confidence level in a detection result.
	ConfidenceLevelMedium ConfidenceLevelType = "medium"

	// ConfidenceLevelHigh represents a high confidence level in a detection result.
	ConfidenceLevelHigh ConfidenceLevelType = "high"
)

// DetectionType represents the type of detection made by the inspector.
type DetectionType string

// String returns the string representation of the detection type.
func (dt DetectionType) String() string {
	return string(dt)
}
