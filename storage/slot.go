package storage

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

// SlotDescriptor provides detailed information about a storage slot in an Ethereum smart contract.
// It includes metadata about the variable stored in the slot, the contract it belongs to, and the slot's state at a specific block number.
type SlotDescriptor struct {
	// DeclarationId is a unique identifier for the variable declaration associated with this slot.
	DeclarationId int64 `json:"declaration_id"`

	// Variable represents the variable metadata and is not serialized into JSON.
	Variable *Variable `json:"-"`

	// Contract points to the IR representation of the contract containing this slot.
	// It is not serialized into JSON.
	Contract *ir.Contract `json:"-"`

	// BlockNumber specifies the block at which the slot's state is considered.
	BlockNumber *big.Int `json:"block_number"`

	// Name is the name of the variable in the smart contract.
	Name string `json:"name"`

	// Type is the high-level type of the variable (e.g., "uint256").
	Type string `json:"type"`

	// TypeDescription provides a detailed AST-based type description of the variable.
	TypeDescription *ast.TypeDescription `json:"type_description"`

	// Slot is the index of the storage slot in the contract.
	Slot int64 `json:"slot"`

	// Size indicates the size of the variable in bytes.
	Size int64 `json:"size"`

	// Offset represents the byte offset of the variable within the storage slot.
	Offset int64 `json:"offset"`

	// RawValue is the raw Ethereum storage slot value at the specified block number.
	RawValue common.Hash `json:"raw_value"`

	// Value is the interpreted value of the variable, with its type depending on the variable's type.
	Value interface{} `json:"value"`
}
