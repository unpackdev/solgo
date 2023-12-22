package inspector

type DetectorEntry struct {
	detectorType DetectorType
	detector     Detector
}

var registry []DetectorEntry

func getDetectorIndex(detectorType DetectorType) int {
	for index, entry := range registry {
		if entry.detectorType == detectorType {
			return index
		}
	}
	return -1
}

// IsDetectorType checks if the provided detectorType is in the list of detectorTypes.
func IsDetectorType(detectorType DetectorType, detectorTypes ...DetectorType) bool {
	for _, dt := range detectorTypes {
		if dt == detectorType {
			return true
		}
	}
	return false
}

func DetectorExists(detectorType DetectorType) bool {
	return getDetectorIndex(detectorType) != -1
}

func GetDetector(detectorType DetectorType) Detector {
	index := getDetectorIndex(detectorType)
	if index != -1 {
		return registry[index].detector
	}
	return nil
}

func RegisterDetector(detectorType DetectorType, detector Detector) bool {
	if !DetectorExists(detectorType) {
		registry = append(registry, DetectorEntry{detectorType, detector})
		return true
	}
	return false
}

// InsertDetector allows inserting a detector before or after an existing detector.
// If `before` is true, the detector is inserted before the specified type, otherwise after.
// If the specified type is not found, the detector is added to the end of the list.
func InsertDetector(newType DetectorType, newDetector Detector, existingType DetectorType, before bool) {
	index := getDetectorIndex(existingType)

	if index == -1 || !before {
		RegisterDetector(newType, newDetector)
	} else {
		registry = append(registry[:index+1], registry[index:]...)
		registry[index] = DetectorEntry{newType, newDetector}
	}
}

func (i *Inspector) RegisterDetectors() {
	RegisterDetector(StorageDetectorType, NewStorageDetector(i.ctx, i))
	RegisterDetector(StateVariableDetectorType, NewStateVariableDetector(i.ctx, i))
	RegisterDetector(OwnershipDetectorType, NewOwnershipDetector(i.ctx, i))
	RegisterDetector(TransferDetectorType, NewTransferDetector(i.ctx, i))
	RegisterDetector(ProxyDetectorType, NewProxyDetector(i.ctx, i))
	RegisterDetector(MintDetectorType, NewMintDetector(i.ctx, i))
	RegisterDetector(BurnDetectorType, NewBurnDetector(i.ctx, i))
	RegisterDetector(PausableDetectorType, NewPausableDetector(i.ctx, i))
	RegisterDetector(ExternalCallsDetectorType, NewExternalCallsDetector(i.ctx, i))
	RegisterDetector(AntiWhaleDetectorType, NewAntiWhaleDetector(i.ctx, i))
	RegisterDetector(StandardsDetectorType, NewStandardsDetector(i.ctx, i))
	RegisterDetector(FeeDetectorType, NewFeesDetector(i.ctx, i))
}
