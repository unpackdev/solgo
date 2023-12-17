package contracts

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/inspector"
	"go.uber.org/zap"
)

func (c *Contract) Inspect(ctx context.Context) (*inspector.Report, error) {
	descriptor := c.GetDescriptor()

	inspector, err := inspector.NewInspector(c.ctx, c.network, descriptor.Detector, c.sim, c.stor, c.bindings, c.ipfsProvider, c.GetAddress(), c.token)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create inspector: %s, network: %s, network_id: %d, contract: %s",
			err,
			c.network.String(),
			c.descriptor.NetworkID,
			c.GetAddress().Hex(),
		)
	}

	// If contract does not have any source code available we don't want to check it here.
	// In that case we will in the future go towards the opcodes...
	if !inspector.IsReady() {
		return nil, fmt.Errorf(
			"inspection rejected as contract '%s' on network '%s' does not have any source code available",
			c.GetAddress().Hex(),
			c.network.String(),
		)
	}

	inspector.RegisterDetectors()

	// Alright now we're at the point that we know contract should be checked for any type of malicious activity
	if err := inspector.Inspect(); err != nil {
		zap.L().Error(
			"failure while inspecting contract",
			zap.Error(err),
			zap.Any("network", c.network),
			zap.Any("network_id", c.descriptor.NetworkID),
			zap.String("contract", descriptor.Address.Hex()),
		)
		return nil, err
	}

	descriptor.Introspection = inspector.GetReport()
	return inspector.GetReport(), nil
}
