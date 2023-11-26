package inspector

type DetectorType string

type SeverityType string

type ConfidenceLevelType string

type DetectionType string

const (
	StateVariableDetectorType DetectorType = "state_variable"
	TransferDetectorType      DetectorType = "transfer"
	MintDetectorType          DetectorType = "mint"
	BurnDetectorType          DetectorType = "burn"
	ProxyDetectorType         DetectorType = "proxy"

	SeverityInfo     SeverityType = "informatinal"
	SeverityLow      SeverityType = "low"
	SeverityMedium   SeverityType = "medium"
	SeverityHigh     SeverityType = "high"
	SeverityCritical SeverityType = "critical"

	ConfidenceLevelLow    ConfidenceLevelType = "low"
	ConfidenceLevelMedium ConfidenceLevelType = "medium"
	ConfidenceLevelHigh   ConfidenceLevelType = "high"
)
