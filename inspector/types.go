package inspector

type DetectorType string

const (
	StateVariableDetectorType DetectorType = "state_variable"
	TransferDetectorType      DetectorType = "transfer"
	MintDetectorType          DetectorType = "mint"
	BurnDetectorType          DetectorType = "burn"
)
