package inspector

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type Report struct {
	Address       common.Address       `json:"address"`
	Network       utils.Network        `json:"network"`
	UsesTransfers bool                 `json:"uses_transfers"`
	Detectors     map[DetectorType]any `json:"detectors"`
}

func (r *Report) HasDetector(detectorType DetectorType) bool {
	_, ok := r.Detectors[detectorType]
	return ok
}

func (r *Report) GetDetector(detectorType DetectorType) any {
	return r.Detectors[detectorType]
}

func (r *Report) AddDetector(detectorType DetectorType, detector any) {
	r.Detectors[detectorType] = detector
}
