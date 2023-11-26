package inspector

import (
	"github.com/ethereum/go-ethereum/common"
)

type Report struct {
	Addresses     []common.Address     `json:"addresses"`
	UsesTransfers bool                 `json:"uses_transfers"`
	Detectors     map[DetectorType]any `json:"detectors"`
}
