package storage

import (
	"math/big"

	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/ir"
)

type Variable struct {
	Contract      *ir.Contract
	EntryContract bool
	*ir.StateVariable
}

type Descriptor struct {
	Contract                     *contracts.Contract
	Block                        *big.Int
	StateVariables               map[string][]*Variable
	TargetVariables              map[string][]*Variable
	ConstantStorageSlotVariables map[string][]*Variable
}

func (s *Descriptor) GetStateVariables() map[string][]*Variable {
	return s.StateVariables
}

func (s *Descriptor) GetTargetVariables() map[string][]*Variable {
	return s.TargetVariables
}

func (s *Descriptor) GetConstantStorageSlotVariables() map[string][]*Variable {
	return s.ConstantStorageSlotVariables
}

func (s *Descriptor) GetContract() *contracts.Contract {
	return s.Contract
}

func (s *Descriptor) GetBlock() *big.Int {
	return s.Block
}
