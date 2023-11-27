package storage

import (
	"context"
	"fmt"
	"sort"

	"github.com/unpackdev/solgo/ir"
)

type Reader struct {
	ctx        context.Context
	storage    *Storage
	descriptor *Descriptor
}

func NewReader(ctx context.Context, s *Storage, d *Descriptor) (*Reader, error) {
	return &Reader{
		ctx:        ctx,
		storage:    s,
		descriptor: d,
	}, nil
}

func (r *Reader) GetDescriptor() *Descriptor {
	return r.descriptor
}

func (r *Reader) DiscoverStorageVariables() error {
	contract := r.descriptor.Contract

	if contract == nil {
		return fmt.Errorf("failed to get storage variables as contract is not set")
	}

	descriptor := contract.GetDescriptor()
	if descriptor == nil {
		return fmt.Errorf("failed to get storage variables as contract descriptor is not set (parsing did not occur or failed)")
	}

	detector := descriptor.Detector
	if detector == nil {
		return fmt.Errorf("failed to get storage variables as contract detector is not set (parsing did not occur or failed)")
	}

	ir := detector.GetIR()
	if ir == nil {
		return fmt.Errorf("failed to get storage variables as contract IR is not set (parsing did not occur or failed)")
	}

	for _, subContract := range ir.GetRoot().GetContracts() {
		var stateVariables []*Variable
		for _, stateVariable := range subContract.GetStateVariables() {
			stateVariables = append(stateVariables, &Variable{
				Contract:      subContract,
				EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
				StateVariable: stateVariable,
			})
		}

		r.descriptor.StateVariables[subContract.Name] = stateVariables

		// Append the sorted variables to the appropriate descriptor slices
		for _, variable := range stateVariables {
			if !variable.StateVariable.IsConstant() {
				r.descriptor.TargetVariables[subContract.Name] = append(
					r.descriptor.TargetVariables[subContract.Name],
					variable,
				)
			} else {
				r.descriptor.ConstanVariables[subContract.Name] = append(
					r.descriptor.ConstanVariables[subContract.Name],
					variable,
				)
			}
		}
	}

	return nil
}

func (r *Reader) CalculateStorageLayout() error {
	currentSlot := int64(0)

	for contractName, variables := range r.descriptor.GetTargetVariables() {
		var slots []*SlotDescriptor
		var previousVars []*Variable

		// Sort the variables by their declaration line number
		sort.Slice(variables, func(i, j int) bool {
			return variables[i].GetId() < variables[j].GetId()
		})

		for _, variable := range variables {
			slot, offset, updatedPreviousVars := calculateSlot(variable, currentSlot, previousVars)
			previousVars = updatedPreviousVars
			storageSize, found := variable.GetAST().GetTypeName().StorageSize()
			if !found {
				return fmt.Errorf("error calculating storage size for variable %s", variable.GetName())
			}

			slots = append(slots, &SlotDescriptor{
				DeclarationId: variable.StateVariable.GetId(),
				Variable:      variable,
				Name:          variable.GetName(),
				Type:          variable.GetType(),
				Slot:          slot,
				Size:          storageSize,
				Offset:        offset,
			})

			if slot != currentSlot {
				currentSlot = slot
			}
		}

		r.descriptor.StorageLayout[contractName] = &StorageLayout{
			Contract: func() *ir.Contract {
				if len(slots) != 0 {
					return slots[0].Variable.Contract
				}
				return nil
			}(),
			Slots: slots,
		}
	}

	return nil
}
