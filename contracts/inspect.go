package contracts

import (
	"context"
	"strings"

	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/inspector"
	"github.com/unpackdev/solgo/standards"
	"go.uber.org/zap"
)

/* var mintFunctions = []string{
	"mint", "mintWithTokenURI", "mintBatch", "mintBatchWithTokenURI",
	"_mint", "_mintWithTokenURI", "_mintBatch", "_mintBatchWithTokenURI",
}

var burnFunctions = []string{
	"burn", "_burn", "burnFrom", "_burnFrom",
} */

func (c *Contract) Inspect(ctx context.Context) (*SafetyDescriptor, error) {
	descriptor := c.GetDescriptor()
	detector := descriptor.Detector

	inspector, err := inspector.NewInspector(detector, c.GetAddress())
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isRenounceOwnershipFunction(functionNode *ast.Function) bool {
	// Check the function name
	if functionNode.GetName() != "renounceOwnership" {
		return false
	}

	// Check if the function body contains logic to set the owner to zero address
	// This is simplified; you would need to parse and understand the function body
	if strings.Contains(functionNode.GetBody().ToString(), "owner = address(0)") {
		return true
	}

	return false
}
