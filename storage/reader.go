package storage

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"go.uber.org/zap"
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

func (r *Reader) GetStorageVariables() error {
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
		if subContract != nil {
			for _, stateVariable := range subContract.GetStateVariables() {
				if stateVariable == nil {
					continue
				}

				zap.L().Debug(
					"Found state variable",
					zap.String("contract_address", contract.GetAddress().Hex()),
					zap.String("contract_name", subContract.GetName()),
					zap.String("name", stateVariable.GetName()),
					zap.String("type", stateVariable.GetType()),
					zap.Any("visibility", stateVariable.GetVisibility()),
					zap.Any("storage", stateVariable.GetStateMutability()),
					zap.Any("constant", stateVariable.IsConstant()),
				)

				r.descriptor.StateVariables[subContract.Name] = append(
					r.descriptor.StateVariables[subContract.Name],
					&Variable{
						Contract:      subContract,
						EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
						StateVariable: stateVariable,
					},
				)

				if !stateVariable.IsConstant() && stateVariable.GetStateMutability() == ast_pb.Mutability_IMMUTABLE {
					r.descriptor.TargetVariables[subContract.Name] = append(
						r.descriptor.TargetVariables[subContract.Name],
						&Variable{
							Contract:      subContract,
							EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
							StateVariable: stateVariable,
						},
					)
				} else {
					r.descriptor.ConstantStorageSlotVariables[subContract.Name] = append(
						r.descriptor.ConstantStorageSlotVariables[subContract.Name],
						&Variable{
							Contract:      subContract,
							EntryContract: ir.GetRoot().GetEntryName() == subContract.GetName(),
							StateVariable: stateVariable,
						},
					)
				}
			}
		}
	}

	return nil
}

func (r *Reader) GetStorageLayout() error {
	storageLayout := make(map[string]*SlotInfo)

	for contractName, variables := range r.descriptor.GetTargetVariables() {
		for _, variable := range variables {
			_ = contractName
			fmt.Println(variable.GetType())
			fmt.Println(variable.StorageLocation)
			fmt.Println(variable.GetAST().GetTypeName())
			fmt.Println(variable.GetAST().GetTypeName().StorageSize())
		}
	}

	fmt.Println(
		"Found storage layout",
		zap.Any("storage_layout", storageLayout),
	)

	return nil
}
