package storage

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/ir"
)

// Descriptor holds information about a smart contract's state at a specific block.
// It includes details like state variables, target variables, constant variables, and storage layouts.
type Descriptor struct {
	Contract         *contracts.Contract
	Block            *big.Int
	StateVariables   map[string][]*Variable
	TargetVariables  map[string][]*Variable
	ConstanVariables map[string][]*Variable
	StorageLayout    *StorageLayout
}

// GetDetector retrieves the contract's detector, which is essential for contract analysis.
// It returns an error if the contract or its detector is not properly set up.
func (s *Descriptor) GetDetector() (*detector.Detector, error) {
	if s.Contract == nil {
		return nil, fmt.Errorf("contract not set in Descriptor")
	}

	descriptor := s.Contract.GetDescriptor()
	if descriptor == nil {
		return nil, fmt.Errorf("contract descriptor not set (parsing did not occur or failed)")
	}

	return descriptor.Detector, nil
}

// GetAST retrieves the abstract syntax tree (AST) builder for the contract.
// It returns an error if the AST builder is not available due to parsing failures or initialization issues.
func (s *Descriptor) GetAST() (*ast.ASTBuilder, error) {
	detector, err := s.GetDetector()
	if err != nil {
		return nil, err
	}

	return detector.GetAST(), nil
}

// GetIR retrieves the intermediate representation (IR) builder of the contract.
// It returns an error if the IR builder is not set, indicating parsing or initialization issues.
func (s *Descriptor) GetIR() (*ir.Builder, error) {
	detector, err := s.GetDetector()
	if err != nil {
		return nil, err
	}

	return detector.GetIR(), nil
}

// GetStateVariables returns a map of state variables associated with the contract.
func (s *Descriptor) GetStateVariables() map[string][]*Variable {
	return s.StateVariables
}

// GetTargetVariables returns a map of target variables associated with the contract.
func (s *Descriptor) GetTargetVariables() map[string][]*Variable {
	return s.TargetVariables
}

// GetConstantStorageSlotVariables returns a map of constant variables associated with the contract.
func (s *Descriptor) GetConstantStorageSlotVariables() map[string][]*Variable {
	return s.ConstanVariables
}

// GetContract returns the contract associated with this descriptor.
func (s *Descriptor) GetContract() *contracts.Contract {
	return s.Contract
}

// GetBlock returns the block number at which the contract's state is described.
func (s *Descriptor) GetBlock() *big.Int {
	return s.Block
}

// GetStorageLayout returns all storage layouts associated with the contract.
func (s *Descriptor) GetStorageLayout() *StorageLayout {
	return s.StorageLayout
}

// GetSlots returns  slots descriptor by its declaration line.
func (s *Descriptor) GetSlots() []*SlotDescriptor {
	return s.StorageLayout.Slots
}

// GetSortedSlots returns a slice of slot descriptors sorted by their declaration line.
// It aggregates and sorts slot descriptors from all storage layouts.
func (s *Descriptor) GetSortedSlots() []*SlotDescriptor {
	var slots []*SlotDescriptor
	slots = append(slots, s.StorageLayout.Slots...)

	sort.Slice(slots, func(i, j int) bool {
		return slots[i].DeclarationId < slots[j].DeclarationId
	})

	return slots
}
