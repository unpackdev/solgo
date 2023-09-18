package abi

import (
	abi_pb "github.com/unpackdev/protos/dist/go/abi"
	"github.com/unpackdev/solgo/ir"
)

// Root represents the root of a Solidity contract's ABI structure.
type Root struct {
	unit              *ir.RootSourceUnit   `json:"-"`                   // Underlying IR node of the RootSourceUnit.
	EntryContractId   int64                `json:"entry_contract_id"`   // ID of the entry contract.
	EntryContractName string               `json:"entry_contract_name"` // Name of the entry contract.
	ContractsCount    int32                `json:"contracts_count"`     // Count of contracts in the ABI.
	Contracts         map[string]*Contract `json:"contracts"`           // Map of contract names to their respective Contract structures.
}

// GetIR returns the underlying IR node of the RootSourceUnit.
func (r *Root) GetIR() *ir.RootSourceUnit {
	return r.unit
}

// GetEntryId returns the entry contract ID.
func (r *Root) GetEntryId() int64 {
	return r.EntryContractId
}

// GetEntryName returns the entry contract name.
func (r *Root) GetEntryName() string {
	return r.EntryContractName
}

// GetContracts returns the map of contracts.
func (r *Root) GetContracts() map[string]*Contract {
	return r.Contracts
}

// GetContractsAsSlice returns the slice contracts.
func (r *Root) GetContractsAsSlice() []*Contract {
	toReturn := make([]*Contract, 0)
	for _, c := range r.Contracts {
		toReturn = append(toReturn, c)
	}
	return toReturn
}

// GetContractByName retrieves a contract by its name from the ABI.
// Returns nil if the contract is not found.
func (r *Root) GetContractByName(name string) *Contract {
	return r.Contracts[name]
}

// GetEntryContract retrieves the entry contract from the ABI.
func (r *Root) GetEntryContract() *Contract {
	return r.GetContractByName(r.EntryContractName)
}

// HasContracts checks if the ABI has any contracts.
// Returns true if there are one or more contracts, otherwise false.
func (r *Root) HasContracts() bool {
	return len(r.Contracts) > 0
}

// GetContractsCount returns the total number of contracts in the ABI.
func (r *Root) GetContractsCount() int32 {
	return r.ContractsCount
}

// ToProto converts the Root structure to its protobuf representation.
func (r *Root) ToProto() *abi_pb.Root {
	proto := &abi_pb.Root{
		EntryContractId:   r.GetEntryId(),
		EntryContractName: r.GetEntryName(),
		ContractsCount:    r.GetContractsCount(),
		Contracts:         make(map[string]*abi_pb.Contract),
	}

	for name, c := range r.GetContracts() {
		proto.Contracts[name] = c.ToProto()
	}

	return proto
}

// processRoot processes the provided RootSourceUnit from the IR and constructs a Root structure.
func (b *Builder) processRoot(root *ir.RootSourceUnit) *Root {
	rootNode := &Root{
		unit:           root,
		ContractsCount: int32(root.GetContractsCount()),
		Contracts:      make(map[string]*Contract),
	}

	if !root.HasContracts() {
		return rootNode
	}

	rootNode.EntryContractId = root.GetEntryId()
	rootNode.EntryContractName = root.GetEntryName()

	for _, contract := range root.GetContracts() {
		rootNode.Contracts[contract.Name] = b.processContract(contract)
	}

	return rootNode
}
