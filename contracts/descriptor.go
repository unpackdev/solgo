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
	"github.com/unpackdev/solgo/inspector"
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
	ContractUUID    *uuid.UUID `json:"contract_uuid,omitempty"`
	TokenUUID       *uuid.UUID `json:"token_uuid,omitempty"`

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

	// Identity related fields
	Owner common.Address `json:"owner,omitempty"`

	// SourcesRaw is the raw sources from Etherscan|BscScan|etc. Should not be used anywhere except in
	// the contract discovery process.
	SourcesRaw      *etherscan.Contract `json:"-"`
	Sources         *solgo.Sources      `json:"sources,omitempty"`
	SourcesProvider string              `json:"sources_provider,omitempty"`

	// Source detection related fields.
	Detector    *detector.Detector    `json:"-"`
	IRRoot      *ir.RootSourceUnit    `json:"ir,omitempty"`
	Constructor *bytecode.Constructor `json:"constructor,omitempty"`
	Metadata    *bytecode.Metadata    `json:"metadata,omitempty"`

	// Auditing related fields.
	Verified             bool              `json:"verified,omitempty"`
	VerificationProvider string            `json:"verification_provider,omitempty"`
	Safe                 bool              `json:"safe,omitempty"`
	Audit                *audit.Report     `json:"audit,omitempty"`
	Introspection        *inspector.Report `json:"introspection,omitempty"`

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

func (d *Descriptor) HasNetworkUUID() bool {
	return d.NetworkUUID != nil
}

func (d *Descriptor) HasMetadata() bool {
	return d.Metadata != nil
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

func (d *Descriptor) SetCompilerVersion(ver string) {
	d.CompilerVersion = ver
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

func (d *Descriptor) GetAudit() *audit.Report {
	return d.Audit
}

func (d *Descriptor) GetIntrospection() *inspector.Report {
	return d.Introspection
}

func (d *Descriptor) GetMetadata() *bytecode.Metadata {
	return d.Metadata
}

func (d *Descriptor) GetConstructor() *bytecode.Constructor {
	return d.Constructor
}

func (d *Descriptor) GetSources() *solgo.Sources {
	return d.Sources
}

func (d *Descriptor) GetSourcesRaw() *etherscan.Contract {
	return d.SourcesRaw
}

func (d *Descriptor) GetSourcesProvider() string {
	return d.SourcesProvider
}

func (d *Descriptor) IsOptimized() bool {
	return d.Optimized
}

func (d *Descriptor) GetOptimizationRuns() uint64 {
	return d.OptimizationRuns
}

func (d *Descriptor) GetEVMVersion() string {
	return d.EVMVersion
}

func (d *Descriptor) GetABI() string {
	return d.ABI
}

func (d *Descriptor) GetLicense() string {
	return d.License
}

func (d *Descriptor) GetName() string {
	return d.Name
}

func (d *Descriptor) GetCompilerVersion() string {
	return d.CompilerVersion
}

func (d *Descriptor) GetSolgoVersion() string {
	return d.SolgoVersion
}

func (d *Descriptor) GetNetwork() utils.Network {
	return d.Network
}

func (d *Descriptor) GetNetworkID() utils.NetworkID {
	return d.NetworkID
}

func (d *Descriptor) GetAddress() common.Address {
	return d.Address
}

func (d *Descriptor) GetRuntimeBytecode() []byte {
	return d.RuntimeBytecode
}

func (d *Descriptor) GetDeployedBytecode() []byte {
	return d.DeployedBytecode
}

func (d *Descriptor) GetBlock() *types.Block {
	return d.Block
}

func (d *Descriptor) GetTransaction() *types.Transaction {
	return d.Transaction
}

func (d *Descriptor) GetReceipt() *types.Receipt {
	return d.Receipt
}

func (d *Descriptor) GetLiquidityPairs() map[utils.ExchangeType]common.Address {
	return d.LiquidityPairs
}

func (d *Descriptor) GetDetector() *detector.Detector {
	return d.Detector
}

func (d *Descriptor) GetUUID() *uuid.UUID {
	return d.UUID
}

func (d *Descriptor) GetBlockUUID() *uuid.UUID {
	return d.BlockUUID
}

func (d *Descriptor) GetTransactionUUID() *uuid.UUID {
	return d.TransactionUUID
}

func (d *Descriptor) GetNetworkUUID() *uuid.UUID {
	return d.NetworkUUID
}

func (d *Descriptor) IsVerified() bool {
	return d.Verified
}

func (d *Descriptor) GetVerificationProvider() string {
	return d.VerificationProvider
}

func (d *Descriptor) IsSafe() bool {
	return d.Safe
}

func (d *Descriptor) HasIntropection() bool {
	return d.Introspection != nil
}

func (d *Descriptor) HasAuditReport() bool {
	return d.Audit != nil
}

func (d *Descriptor) GetOwner() common.Address {
	return d.Owner
}
