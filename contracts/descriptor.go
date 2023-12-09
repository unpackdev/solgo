package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/audit"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/tokens"
	"github.com/unpackdev/solgo/utils"
)

type TokenDescriptor struct {
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"total_supply"`
}

type SafetyDescriptor struct {
	Mintable             bool `json:"mintable"`
	Burnable             bool `json:"burnable"`
	CanRenounceOwnership bool `json:"can_renounce_ownership"`
}

type Descriptor struct {

	// Helpers, usually for database....
	UUID            *uuid.UUID `json:"uuid,omitempty"`
	NetworkUUID     *uuid.UUID `json:"network_uuid,omitempty"`
	BlockUUID       *uuid.UUID `json:"block_uuid,omitempty"`
	TransactionUUID *uuid.UUID `json:"transaction_uuid,omitempty"`

	// Contract related fields.
	Network          utils.Network                         `json:"network"`
	NetworkID        utils.NetworkID                       `json:"network_id"`
	Address          common.Address                        `json:"address"`
	RuntimeBytecode  []byte                                `json:"runtime_bytecode"`
	DeployedBytecode []byte                                `json:"deployed_bytecode"`
	Block            *types.Block                          `json:"block,omitempty"`
	Transaction      *types.Transaction                    `json:"transaction,omitempty"`
	Receipt          *types.Receipt                        `json:"receipt,omitempty"`
	ABI              string                                `json:"abi,omitempty"`
	Name             string                                `json:"name,omitempty"`
	License          string                                `json:"license,omitempty"`
	SolgoVersion     string                                `json:"solgo_version,omitempty"`
	CompilerVersion  string                                `json:"compiler_version,omitempty"`
	Optimized        bool                                  `json:"optimized,omitempty"`
	OptimizationRuns uint64                                `json:"optimization_runs,omitempty"`
	EVMVersion       string                                `json:"evm_version,omitempty"`
	LiquidityPairs   map[utils.ExchangeType]common.Address `json:"liquidity_pairs,omitempty"`

	// SourcesRaw is the raw sources from Etherscan|BscScan|etc. Should not be used anywhere except in
	// the contract discovery process.
	SourcesRaw *etherscan.Contract `json:"-"`
	Sources    *solgo.Sources      `json:"sources,omitempty"`

	// Source detection related fields.
	Detector    *detector.Detector    `json:"-"`
	IRRoot      *ir.RootSourceUnit    `json:"ir,omitempty"`
	Constructor *bytecode.Constructor `json:"constructor,omitempty"`

	// Auditing related fields.
	Audit  *audit.Report     `json:"audit,omitempty"`
	Safety *SafetyDescriptor `json:"safety,omitempty"`

	// Token related fields.
	Token *tokens.Descriptor `json:"token,omitempty"`
}

func (d *Descriptor) HasToken() bool {
	return d.Token != nil
}

func (d *Descriptor) HasConstructor() bool {
	return d.Constructor != nil
}

func (d *Descriptor) HasSources() bool {
	return d.Sources != nil
}

func (d *Descriptor) HasAudit() bool {
	return d.Audit != nil
}

func (d *Descriptor) HasLiquidityPairs() bool {
	return len(d.LiquidityPairs) > 0
}

func (d *Descriptor) HasUUID() bool {
	return d.UUID != nil
}

func (d *Descriptor) SetUUID(uuid *uuid.UUID) {
	d.UUID = uuid
}

func (d *Descriptor) HasBlockUUID() bool {
	return d.BlockUUID != nil
}

func (d *Descriptor) SetBlockUUID(uuid *uuid.UUID) {
	d.BlockUUID = uuid
}

func (d *Descriptor) HasTransactionUUID() bool {
	return d.TransactionUUID != nil
}

func (d *Descriptor) SetTransactionUUID(uuid *uuid.UUID) {
	d.TransactionUUID = uuid
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

func (d *Descriptor) IsERC20() bool {
	if !d.HasDetector() {
		return false
	}

	return d.Detector.GetIR().GetRoot().HasHighConfidenceStandard(standards.ERC20)
}

func (d *Descriptor) IsERC721() bool {
	if !d.HasDetector() {
		return false
	}

	return d.Detector.GetIR().GetRoot().HasHighConfidenceStandard(standards.ERC721)
}

func (d *Descriptor) IsERC1155() bool {
	if !d.HasDetector() {
		return false
	}

	return d.Detector.GetIR().GetRoot().HasHighConfidenceStandard(standards.ERC1155)
}

func (d *Descriptor) IsERC165() bool {
	if !d.HasDetector() {
		return false
	}

	return d.Detector.GetIR().GetRoot().HasHighConfidenceStandard(standards.ERC165)
}

func (d *Descriptor) GetIrRoot() *ir.RootSourceUnit {
	return d.IRRoot
}

func (d *Descriptor) GetToken() *tokens.Descriptor {
	return d.Token
}
