package storage

import "math/big"

type SlotDescriptor struct {
	Variable        *Variable   `json:"-"`
	BlockNumber     *big.Int    `json:"block_number"`
	Name            string      `json:"name"`
	Type            string      `json:"type"`
	DeclarationLine int64       `json:"declaration_line"`
	Slot            int64       `json:"slot"`
	Size            int64       `json:"size"`
	Offset          int64       `json:"offset"`
	RawValue        []byte      `json:"raw_value"`
	Value           interface{} `json:"value"`
}
