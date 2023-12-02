package inspector

import (
	"github.com/ethereum/go-ethereum/common"
)

type Report struct {
	Address       common.Address       `json:"address"`
	UsesTransfers bool                 `json:"uses_transfers"`
	Detectors     map[DetectorType]any `json:"detectors"`
}
