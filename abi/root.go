package abi

import (
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ir"
)

// Root represents the root of a Solidity contract's ...
type Root struct {
	unit              *ir.RootSourceUnit   `json:"-"`
	EntryContractId   int64                `json:"entry_contract_id"`
	EntryContractName string               `json:"entry_contract_name"`
	ContractsCount    int32                `json:"contracts_count"`
	Contracts         map[string]*Contract `json:"contracts"`
}

// GetAST returns the underlying IR node of the RootSourceUnit.
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

// GetContracts returns the list of contracts in the ABI.
func (r *Root) GetContracts() map[string]*Contract {
	return r.Contracts
}

// GetContractByName returns the contract with the given name from the ABI.
// If no contract with the given name is found, it returns nil.
func (r *Root) GetContractByName(name string) *Contract {
	for cName, contract := range r.Contracts {
		if cName == name {
			return contract
		}
	}
	return nil
}

// GetEntryContract returns the entry contract from the ABI.
func (r *Root) GetEntryContract() *Contract {
	return r.GetContractByName(r.EntryContractName)
}

// HasContracts returns true if the AST has one or more contracts, false otherwise.
func (r *Root) HasContracts() bool {
	return len(r.Contracts) > 0
}

// GetContractsCount returns the count of contracts in the ABI.
func (r *Root) GetContractsCount() int32 {
	return r.ContractsCount
}

// ToProto is a placeholder function for converting the RootSourceUnit to a protobuf message.
func (r *Root) ToProto() *ir_pb.Root {
	proto := &ir_pb.Root{
		/* 		Id:                0,
		   		NodeType:          r.GetNodeType(),
		   		EntryContractId:   r.GetEntryId(),
		   		EntryContractName: r.GetEntryName(),
		   		ContractsCount:    r.GetContractsCount(),
		   		Contracts:         make([]*ir_pb.Contract, 0), */
	}

	/* 	for _, c := range r.GetContracts() {
		proto.Contracts = append(proto.Contracts, c.ToProto())
	} */

	return proto
}

func (b *Builder) processRoot(root *ir.RootSourceUnit) *Root {
	rootNode := &Root{
		unit:           root,
		ContractsCount: int32(root.GetContractsCount()),
		Contracts:      make(map[string]*Contract, 0),
	}

	// No source units to process, so we're going to stop processing the root from here...
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
