package contracts

import (
	"context"

	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

func (c *Contract) Audit(ctx context.Context) error {
	// We are going to process the contract auditing only if we have the detector and contracts.
	// It's obvious tho...
	if c.descriptor.HasDetector() && c.descriptor.HasContracts() {
		// Make sure that auditor uses the same version of the compiler as the contract was compiled with.
		//descriptor.GetDetector().GetAuditor().GetConfig().SetCompilerVersion(dbMetadata.CompilerVersion)
		detector := c.descriptor.Detector

		semVer := utils.ParseSemanticVersion(c.descriptor.CompilerVersion)
		detector.GetAuditor().GetConfig().SetCompilerVersion(semVer.String())

		audit, err := c.descriptor.Detector.Analyze()
		if err != nil {
			zap.L().Error(
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
