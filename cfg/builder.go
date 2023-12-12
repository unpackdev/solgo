package cfg

import (
	"context"
	"errors"

	"github.com/goccy/go-graphviz"
	"github.com/unpackdev/solgo/ir"
)

// Builder is responsible for constructing the control flow graph (CFG) of Solidity contracts.
// It utilizes the Intermediate Representation (IR) provided by solgo and Graphviz for graph operations.
type Builder struct {
	ctx     context.Context    // Context for the builder operations.
	builder *ir.Builder        // IR builder from solgo, used for generating the IR of the contracts.
	viz     *graphviz.Graphviz // Graphviz instance for visualizing the CFG.
	graph   *Graph             // Internal representation of the CFG.
}

// NewBuilder initializes a new CFG builder with the given context and IR builder.
// Returns an error if the provided IR builder is nil or if it does not have a root contract set.
func NewBuilder(ctx context.Context, builder *ir.Builder) (*Builder, error) {
	if builder == nil || builder.GetRoot() == nil {
		return nil, errors.New("builder is not set")
	}

	return &Builder{
		ctx:     ctx,
		builder: builder,
		viz:     graphviz.New(),
		graph:   NewGraph(),
	}, nil
}

// GetGraph returns the internal Graph instance of the CFG.
func (b *Builder) GetGraph() *Graph {
	return b.graph
}

// Build processes the Solidity contracts using the IR builder to construct the CFG.
// It identifies the entry contract and explores all dependencies and inherited contracts.
// Returns an error if the root node or entry contract is not set in the IR builder.
func (b *Builder) Build() error {
	root := b.builder.GetRoot()
	if root == nil {
		return errors.New("root node is not set in IR builder")
	}

	entryContract := root.GetEntryContract()
	if entryContract == nil {
		return errors.New("no entry contract found")
	}

	if b.graph == nil {
		b.graph = NewGraph()
	}

	var dfs func(contract *ir.Contract, isEntryContract bool)
	dfs = func(contract *ir.Contract, isEntryContract bool) {
		if !b.graph.NodeExists(contract.GetName()) {
			b.graph.AddNode(contract.GetName(), contract, isEntryContract)
			allRelatedContracts := make([]*ir.Contract, 0)
			for _, importStmt := range contract.GetImports() {
				importedContract := root.GetContractById(importStmt.GetContractId())
				if importedContract != nil {
					b.graph.AddDependency(contract.GetName(), importStmt)
					allRelatedContracts = append(allRelatedContracts, importedContract)
				}
			}
			for _, baseContract := range contract.GetBaseContracts() {
				baseContractId := baseContract.GetBaseName().GetReferencedDeclaration()
				baseContractObj := root.GetContractBySourceUnitId(baseContractId)
				if baseContractObj != nil {
					b.graph.AddInheritance(contract.GetName(), baseContract)
					allRelatedContracts = append(allRelatedContracts, baseContractObj)
				}
			}
			for _, relatedContract := range allRelatedContracts {
				dfs(relatedContract, false)
			}
		}
	}

	dfs(entryContract, true)
	return nil
}
