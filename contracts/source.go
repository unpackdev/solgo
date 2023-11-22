package contracts

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/contracts/storage"
	"go.uber.org/zap"
)

func (c *Contract) DiscoverSourceCode(ctx context.Context) error {
	if storage := storage.GetStorage(c.addr); storage != nil {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)
		sourcesDir := filepath.Clean(filepath.Join(dir, "..", "sources"))
		storage.Sources.LocalSourcesPath = sourcesDir
		c.descriptor.Name = storage.Sources.EntrySourceUnitName
		c.descriptor.Sources = storage.Sources
		c.descriptor.CompilerVersion = storage.CompilerVersion.String()
		c.descriptor.Optimized = storage.Optimized
		c.descriptor.OptimizationRuns = uint64(storage.OptimizationRuns)
		c.descriptor.EVMVersion = storage.EvmVersion
		c.descriptor.ABI = storage.ABI
		c.descriptor.License = storage.License
		return nil
	}

	response, err := c.etherscan.ScanContract(c.addr)
	if err != nil {
		if err.Error() != "contract not found" { // Do not print error if contract is not found. Just clusterfucks the logs...
			zap.L().Error(
				"failed to scan contract source code",
				zap.Error(err),
				zap.String("network", c.network.String()),
				zap.String("contract_address", c.addr.String()),
			)
		}
		return fmt.Errorf("failed to scan contract source code from %s: %s", c.etherscan.ProviderName(), err)
	}
	c.descriptor.SourcesRaw = response

	sources, err := solgo.NewSourcesFromEtherScan(response.Name, response.SourceCode)
	if err != nil {
		zap.L().Error(
			"failed to create new sources from etherscan response",
			zap.Error(err),
			zap.String("network", c.network.String()),
			zap.String("contract_address", c.addr.String()),
		)
		return fmt.Errorf("failed to create new sources from etherscan response: %s", err)
	}

	c.descriptor.Sources = sources
	c.descriptor.License = c.descriptor.SourcesRaw.LicenseType

	// Contract has no source code available. This is not critical error but annoyance that we can't decompile
	// contract's source code. @TODO: Figure out with external toolings how to decompile bytecode...
	// However we could potentially get information such as ABI from etherscan for future use...
	// We are setting it here, however we are going to replace it with the one from the sources if we have it.

	optimized, err := strconv.ParseBool(c.descriptor.SourcesRaw.OptimizationUsed)
	if err != nil {
		zap.L().Error(
			"Failed to parse OptimizationUsed to bool",
			zap.Error(err),
			zap.String("OptimizationUsed", c.descriptor.SourcesRaw.OptimizationUsed),
		)
		return err
	}

	optimizationRuns, err := strconv.ParseUint(c.descriptor.SourcesRaw.Runs, 10, 64)
	if err != nil {
		zap.L().Error(
			"Failed to parse OptimizationRuns to uint64",
			zap.Error(err),
			zap.String("OptimizationRuns", c.descriptor.SourcesRaw.Runs),
		)
		return err
	}

	c.descriptor.Name = response.Name
	c.descriptor.CompilerVersion = c.descriptor.SourcesRaw.CompilerVersion
	c.descriptor.Optimized = optimized
	c.descriptor.OptimizationRuns = optimizationRuns
	c.descriptor.EVMVersion = c.descriptor.SourcesRaw.EVMVersion
	c.descriptor.ABI = c.descriptor.SourcesRaw.ABI

	return nil
}
