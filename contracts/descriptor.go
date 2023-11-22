package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/audit"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/utils"
)

type TokenDescriptor struct {
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"total_supply"`
}

type Descriptor struct {
	Network          utils.Network                             `json:"network"`
	NetworkID        utils.NetworkID                           `json:"network_id"`
	Address          common.Address                            `json:"address"`
	RuntimeBytecode  []byte                                    `json:"runtime_bytecode"`
	DeployedBytecode []byte                                    `json:"deployed_bytecode"`
	Block            *types.Block                              `json:"block,omitempty"`
	Transaction      *types.Transaction                        `json:"transaction,omitempty"`
	Receipt          *types.Receipt                            `json:"receipt,omitempty"`
	Token            *TokenDescriptor                          `json:"token,omitempty"`
	ABI              string                                    `json:"abi,omitempty"`
	License          string                                    `json:"license,omitempty"`
	SolgoVersion     string                                    `json:"solgo_version,omitempty"`
	CompilerVersion  string                                    `json:"compiler_version,omitempty"`
	Optimized        bool                                      `json:"optimized,omitempty"`
	OptimizationRuns uint64                                    `json:"optimization_runs,omitempty"`
	EVMVersion       string                                    `json:"evm_version,omitempty"`
	LiquidityPairs   map[exchanges.ExchangeType]common.Address `json:"liquidity_pairs,omitempty"`

	// SourcesRaw is the raw sources from Etherscan|BscScan|etc. Should not be used anywhere except in
	// the contract discovery process.
	SourcesRaw *etherscan.Contract `json:"-"`
	Sources    *solgo.Sources      `json:"sources,omitempty"`

	// Source detection related fields.
	Detector    *detector.Detector    `json:"-"`
	IRRoot      *ir.RootSourceUnit    `json:"ir,omitempty"`
	Constructor *bytecode.Constructor `json:"constructor,omitempty"`

	// Auditing related fields.
	Audit *audit.Report `json:"audit,omitempty"`
}

func (d *Descriptor) HasDetector() bool {
	if d.Detector != nil && d.Detector.GetIR() != nil && d.Detector.GetIR().GetRoot() != nil {
		return true
	}

	return false
}

func (d *Descriptor) HasContracts() bool {
	if d.HasDetector() && d.Detector.GetIR().GetRoot().HasContracts() {
		return true
	}

	return false
}
