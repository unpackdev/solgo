package contracts

import (
	"context"

	"github.com/unpackdev/solgo/inspector"
	"github.com/unpackdev/solgo/standards"
	"go.uber.org/zap"
)

func (c *Contract) Inspect(ctx context.Context) (*SafetyDescriptor, error) {
	descriptor := c.GetDescriptor()
	detector := descriptor.Detector

	return nil, nil

	inspector, err := inspector.NewInspector(c.ctx, c.network, detector, c.sim, c.stor, c.bindings, c.GetAddress())
	if err != nil {
		zap.L().Error("Error creating inspector", zap.Error(err))
	}

	// If contract does not have any source code available we don't want to check it here.
	// In that case we will in the future go towards the opcodes...
	if !inspector.IsReady() {
		return nil, nil
	}

	// First we don't want to do any type of inspections if contract is not ERC20
	if !inspector.HasStandard(standards.ERC20) {
		return nil, nil
	} else {
		// It can be that we're not able to successfully get the standard but it is still doing trading...
		if !inspector.UsesTransfers() {
			return nil, nil
		}
	}

	inspector.RegisterDetectors()

	// Alright now we're at the point that we know contract should be checked for any type of malicious activity
	if err := inspector.Inspect(); err != nil {
		zap.L().Error("Error inspecting contract", zap.Error(err))
		return nil, err
	}

	//utils.DumpNodeNoExit(inspector.GetReport())
	return descriptor.Safety, nil
}
