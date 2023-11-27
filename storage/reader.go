package storage

import (
	"context"
	"fmt"
	"sort"
	"strings"
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
	var previousVars []*Variable

	sortedSlots := []*SlotDescriptor{}

	for _, variables := range r.descriptor.GetTargetVariables() {
		for _, variable := range variables {
			storageSize, found := variable.GetAST().GetTypeName().StorageSize()
			if !found {
				return fmt.Errorf("error calculating storage size for variable %s", variable.GetName())
			}

			typeName := variable.GetType()
			if strings.HasPrefix(typeName, "contract") {
				typeName = "address"
			}

			sortedSlots = append(sortedSlots, &SlotDescriptor{
				DeclarationId:   variable.StateVariable.GetId(),
				Variable:        variable,
				Contract:        variable.Contract,
				Name:            variable.GetName(),
				Type:            typeName,
				TypeDescription: variable.StateVariable.GetTypeDescription(),
				Size:            storageSize,
			})
		}
	}

	sort.Slice(sortedSlots, func(i, j int) bool {
		return sortedSlots[i].DeclarationId < sortedSlots[j].DeclarationId
	})

	for i, variable := range sortedSlots {
		slot, offset, updatedPreviousVars := calculateSlot(variable.Variable, currentSlot, previousVars)
		previousVars = updatedPreviousVars
		sortedSlots[i].Slot = slot
		sortedSlots[i].Offset = offset

		if slot != currentSlot {
			currentSlot = slot
		}
	}

	r.descriptor.StorageLayout = &StorageLayout{
		Slots: sortedSlots,
	}

	return nil
}
