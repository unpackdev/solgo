package storage

import "math/big"

type SlotDescriptor struct {
	DeclarationId int64       `json:"declaration_id"`
	Variable      *Variable   `json:"-"`
	BlockNumber   *big.Int    `json:"block_number"`
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	Slot          int64       `json:"slot"`
	Size          int64       `json:"size"`
	Offset        int64       `json:"offset"`
	RawValue      []byte      `json:"raw_value"`
	Value         interface{} `json:"value"`
}
