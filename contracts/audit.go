package contracts

import (
	"context"

	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

// Audit performs a security analysis of the contract using its associated detector,
// if available. It updates the contract descriptor with the audit results.
func (c *Contract) Audit(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		if c.descriptor.HasDetector() && c.descriptor.HasContracts() {
			detector := c.descriptor.Detector

			semVer := utils.ParseSemanticVersion(c.descriptor.CompilerVersion)
			detector.GetAuditor().GetConfig().SetCompilerVersion(semVer.String())

			audit, err := c.descriptor.Detector.Analyze()
			if err != nil {
				zap.L().Debug(
					"failed to analyze contract",
					zap.Error(err),
					zap.String("contract_address", c.descriptor.Address.Hex()),
				)
				return err
			}
			c.descriptor.Audit = audit
		}

		return nil
	}
}
