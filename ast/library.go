package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	if ctx.IsEmpty() {
		return
	}

	id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	b.currentSourceUnit.ExportedSymbols = append(
		b.currentSourceUnit.ExportedSymbols,
		&ast_pb.ExportedSymbols{
			Id:   id,
			Name: identifierName,
		},
	)

	identifierNode := &ast_pb.Node{
		Id:   id,
		Name: identifierName,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: b.currentSourceUnit.Id,
		},
		Abstract: false,
		NodeType: ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:     ast_pb.NodeType_KIND_LIBRARY,
	}

	// Check if all of the functions discovered in the library are fully implemented...
	// @TODO: Implement this.
	identifierNode.FullyImplemented = false

	// Discover linearized base contracts...
	// The linearizedBaseContracts field contains an array of IDs that represent the
	// contracts in the inheritance hierarchy, starting from the most derived contract
	// (the contract itself) and ending with the most base contract.
	// The IDs correspond to the id fields of the ContractDefinition nodes in the AST.
	identifierNode.LinearizedBaseContracts = []int64{id}

	// Allright now the fun part begins. We need to traverse through the body of the library
	// and extract all of the nodes...

	// First lets define nodes...
	identifierNode.Nodes = make([]*ast_pb.Node, 0)

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			continue
		}

		bodyNode := b.parseBodyElement(identifierNode, bodyElement)
		identifierNode.Nodes = append(identifierNode.Nodes, bodyNode)
	}

	b.currentSourceUnit.Nodes.Nodes = append(b.currentSourceUnit.Nodes.Nodes, identifierNode)

}
