package contracts

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

func (c *Contract) Parse(ctx context.Context) error {
	// We are interested in attempt to decompile source code only if we actually have source code available.
	if c.descriptor.Sources != nil && c.descriptor.Sources.HasUnits() {
		parser, err := detector.NewDetectorFromSources(ctx, c.compiler, c.descriptor.Sources)
		if err != nil {
			zap.L().Error(
				"Failed to create detector from sources",
				zap.Error(err),
				zap.Any("network", c.network),
				zap.String("contract_address", c.addr.String()),
			)
			return err
		}
		c.descriptor.Detector = parser
		c.descriptor.SolgoVersion = utils.GetBuildVersionByModule("github.com/unpackdev/solgo")

		// Up until this point all is good for all of the contracts, however from this stage moving forward
		// we are getting into issues where we are not capable ATM to parse contracts with < 0.6.0 version.
		// Because of it, we are going to disable all of the functionality for this particular contract related to
		// source code parsing. :( In the future we should sort this out but right now, MVP is the most important thing.
		if utils.IsSemanticVersionLowerOrEqualTo(c.descriptor.CompilerVersion, utils.SemanticVersion{Major: 0, Minor: 5, Patch: 10}) {
			return fmt.Errorf("not supported compiler version (older version): %v", c.descriptor.CompilerVersion)
		}

		if errs := parser.Parse(); errs != nil {
			for _, err := range errs {
				zap.L().Debug(
					"failed to parse contract sources",
					zap.Error(err),
					zap.Any("network", c.network),
					zap.String("contract_address", c.addr.String()),
				)
			}
		}

		if err := parser.Build(); err != nil {
			zap.L().Error(
				"failed to build contract sources",
				zap.Error(err),
				zap.Any("network", c.network),
				zap.String("contract_address", c.addr.String()),
			)
			return err
		}
		c.descriptor.IRRoot = parser.GetIR().GetRoot()

		// What we should update here is get basically missing information from external sources corrected now...
		if c.descriptor.Name == "" {
			c.descriptor.Name = c.descriptor.IRRoot.GetEntryName()
		}

		if c.descriptor.License == "" || c.descriptor.License == "None" && c.descriptor.IRRoot.GetEntryContract() != nil {
			c.descriptor.License = c.descriptor.IRRoot.GetEntryContract().GetLicense()
		}
	}

	return nil
}
