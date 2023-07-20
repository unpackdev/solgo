package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseContractDefinition(sourceUnitCtx *parser.SourceUnitContext, ctx *parser.ContractDefinitionContext) *ast_pb.SourceUnit {
	sourceUnit := &ast_pb.SourceUnit{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		AbsolutePath: func() string {
			for _, unit := range b.sources.SourceUnits {
				if unit.Name == ctx.Identifier().GetText() {
					return unit.Path
				}
			}
			return ""
		}(),
		ExportedSymbols: make([]*ast_pb.ExportedSymbol, 0),
		NodeType:        ast_pb.NodeType_SOURCE_UNIT,
		Root:            &ast_pb.RootNode{},
		Src: &ast_pb.Src{
			Line:   int64(ctx.GetStart().GetLine()),
			Column: int64(ctx.GetStart().GetColumn()),
			Start:  int64(ctx.GetStart().GetStart()),
		},
		Comments: b.comments,
	}

	// Alright lets get the license of the contract...
	sourceUnit.License = b.GetLicense(b.comments)

	// Alright lets extract bloody pragmas...
	sourceUnit.Root.Nodes = append(
		sourceUnit.Root.Nodes,
		b.findPragmasForSourceUnit(sourceUnitCtx, sourceUnit, nil, ctx, nil)...,
	)

	// Now extraction of import paths...
	sourceUnit.Root.Nodes = append(
		sourceUnit.Root.Nodes,
		b.findImportPathsForSourceUnit(sourceUnitCtx, sourceUnit, nil, ctx, nil)...,
	)

	id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	sourceUnit.ExportedSymbols = append(
		sourceUnit.ExportedSymbols,
		&ast_pb.ExportedSymbol{
			Id:   id,
			Name: identifierName,
			AbsolutePath: func() string {
				for _, unit := range b.sources.SourceUnits {
					if unit.Name == identifierName {
						return unit.Path
					}
				}
				return ""
			}(),
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
			ParentIndex: sourceUnit.Id,
		},
		Abstract:             false,
		NodeType:             ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:                 ast_pb.NodeType_KIND_CONTRACT,
		BaseContracts:        make([]*ast_pb.BaseContract, 0),
		ContractDependencies: make([]int64, 0),
	}

	// Discover linearized base contracts...
	// The linearizedBaseContracts field contains an array of IDs that represent the
	// contracts in the inheritance hierarchy, starting from the most derived contract
	// (the contract itself) and ending with the most base contract.
	// The IDs correspond to the id fields of the ContractDefinition nodes in the AST.
	identifierNode.LinearizedBaseContracts = []int64{id}

	if ctx.InheritanceSpecifierList() != nil {
		inheritanceCtx := ctx.InheritanceSpecifierList()

		for _, inheritanceSpecifierCtx := range inheritanceCtx.AllInheritanceSpecifier() {
			baseContract := &ast_pb.BaseContract{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(inheritanceSpecifierCtx.GetStart().GetLine()),
					Column:      int64(inheritanceSpecifierCtx.GetStart().GetColumn()),
					Start:       int64(inheritanceSpecifierCtx.GetStart().GetStart()),
					End:         int64(inheritanceSpecifierCtx.GetStop().GetStop()),
					Length:      int64(inheritanceSpecifierCtx.GetStop().GetStop() - inheritanceSpecifierCtx.GetStart().GetStart() + 1),
					ParentIndex: identifierNode.Id,
				},
				NodeType: ast_pb.NodeType_INHERITANCE_SPECIFIER,
				BaseName: &ast_pb.BaseContractName{
					Id: atomic.AddInt64(&b.nextID, 1) - 1,
					Src: &ast_pb.Src{
						Line:        int64(inheritanceSpecifierCtx.GetStart().GetLine()),
						Column:      int64(inheritanceSpecifierCtx.GetStart().GetColumn()),
						Start:       int64(inheritanceSpecifierCtx.GetStart().GetStart()),
						End:         int64(inheritanceSpecifierCtx.GetStop().GetStop()),
						Length:      int64(inheritanceSpecifierCtx.GetStop().GetStop() - inheritanceSpecifierCtx.GetStart().GetStart() + 1),
						ParentIndex: identifierNode.Id,
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
					Name:     inheritanceSpecifierCtx.IdentifierPath().GetText(),
				},
			}

			for _, unitNode := range b.sourceUnits {
				if unitNode.GetName() == inheritanceSpecifierCtx.IdentifierPath().GetText() {
					baseContract.BaseName.ReferencedDeclaration = unitNode.GetId()
					identifierNode.LinearizedBaseContracts = append(
						identifierNode.LinearizedBaseContracts, unitNode.GetId(),
					)
					identifierNode.ContractDependencies = append(
						identifierNode.ContractDependencies, unitNode.GetId(),
					)
				}
			}

			identifierNode.BaseContracts = append(identifierNode.BaseContracts, baseContract)
		}
	}

	// Allright now the fun part begins. We need to traverse through the body of the library
	// and extract all of the nodes...

	// First lets define nodes...
	identifierNode.Nodes = make([]*ast_pb.Node, 0)

	fullyImplemented := true

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			continue
		}

		bodyNode := b.parseBodyElement(sourceUnit, identifierNode, bodyElement)
		identifierNode.Nodes = append(identifierNode.Nodes, bodyNode)

		// Lets check if we have any functions in the body...
		if bodyNode.NodeType == ast_pb.NodeType_FUNCTION_DEFINITION {
			if !bodyNode.Implemented {
				fullyImplemented = false
			}
		}
	}

	identifierNode.FullyImplemented = fullyImplemented
	sourceUnit.Root.Nodes = append(sourceUnit.Root.Nodes, identifierNode)

	return sourceUnit
}
