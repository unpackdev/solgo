package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
	"time"

	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/providers/etherscan"
	"go.uber.org/zap"
)

func (c *Contract) DiscoverSourceCode(ctx context.Context) error {
	var response *etherscan.Contract // Assuming ScanResponse is the type returned by ScanContract
	var err error

	// Retry mechanism
	const maxRetries = 5
	for i := 0; i < maxRetries; i++ {
		response, err = c.etherscan.ScanContract(c.addr)
		if err != nil {
			if strings.Contains(err.Error(), "Max rate limit reached") {
				// Wait for 100ms before retrying
				time.Sleep(100 * time.Millisecond)
				continue
			} else if !strings.Contains(err.Error(), "not found") &&
				!strings.Contains(err.Error(), "not verified") {
				zap.L().Error(
					"failed to scan contract source code",
					zap.Error(err),
					zap.String("network", c.network.String()),
					zap.String("contract_address", c.addr.String()),
				)
			}
			return fmt.Errorf("failed to scan contract source code from %s: %s", c.etherscan.ProviderName(), err)
		}
		break // Exit loop if ScanContract is successful
	}

	// Handle the case when all retries fail
	if err != nil {
		return fmt.Errorf("after %d retries, failed to scan contract source code from %s: %s", maxRetries, c.etherscan.ProviderName(), err)
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

	license := strings.ReplaceAll(c.descriptor.SourcesRaw.LicenseType, "\r", "")
	license = strings.ReplaceAll(license, "\n", "")
	license = strings.TrimSpace(c.descriptor.SourcesRaw.LicenseType)
	c.descriptor.License = license

	// Contract has no source code available. This is not critical error but annoyance that we can't decompile
	// contract's source code. @TODO: Figure out with external toolings how to decompile bytecode...
	// However we could potentially get information such as ABI from etherscan for future use...
	// We are setting it here, however we are going to replace it with the one from the sources if we have it.

	optimized, err := strconv.ParseBool(c.descriptor.SourcesRaw.OptimizationUsed)
	if err != nil {
		zap.L().Error(
			"failed to parse OptimizationUsed to bool",
			zap.Error(err),
			zap.String("OptimizationUsed", c.descriptor.SourcesRaw.OptimizationUsed),
		)
		return err
	}

	optimizationRuns, err := strconv.ParseUint(c.descriptor.SourcesRaw.Runs, 10, 64)
	if err != nil {
		zap.L().Error(
			"failed to parse OptimizationRuns to uint64",
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
	c.descriptor.SourcesProvider = c.etherscan.ProviderName()
	c.descriptor.Verified = true
	c.descriptor.VerificationProvider = c.etherscan.ProviderName()

	proxy, _ := strconv.ParseBool(response.Proxy)
	c.descriptor.Proxy = proxy

	if len(response.Implementation) > 10 {
		c.descriptor.Implementations = append(
			c.descriptor.Implementations,
			common.HexToAddress(response.Implementation),
		)
	}

	return nil
}
