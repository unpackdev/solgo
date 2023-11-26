package inspector

var registry = map[DetectorType]Detector{}

func DetectorExists(detectorType DetectorType) bool {
	_, ok := registry[detectorType]
	return ok
}

func GetDetector(detectorType DetectorType) Detector {
	return registry[detectorType]
}

func RegisterDetector(detectorType DetectorType, detector Detector) bool {
	if !DetectorExists(detectorType) {
		registry[detectorType] = detector
		return true
	}

	return false
}

func IsDetectorType(detectorType DetectorType, detectorTypes ...DetectorType) bool {
	for _, dt := range detectorTypes {
		if dt == detectorType {
			return true
		}
	}

	return false
}

func (i *Inspector) RegisterDetectors() {
	RegisterDetector(StateVariableDetectorType, NewStateVariableDetector(i.ctx, i))
	RegisterDetector(OwnershipDetectorType, NewOwnershipDetector(i.ctx, i))
	RegisterDetector(TransferDetectorType, NewTransferDetector(i.ctx, i))
	RegisterDetector(ProxyDetectorType, NewProxyDetector(i.ctx, i))
	RegisterDetector(MintDetectorType, NewMintDetector(i.ctx, i))
	RegisterDetector(BurnDetectorType, NewBurnDetector(i.ctx, i))
}
