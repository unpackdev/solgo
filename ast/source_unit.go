package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	for _, child := range ctx.GetChildren() {
		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			b.currentSourceUnit = b.parseLibraryDefinition(ctx, libraryCtx)
			b.sourceUnits = append(b.sourceUnits, b.currentSourceUnit)
		}

		if interfaceCtx, ok := child.(*parser.InterfaceDefinitionContext); ok {
			b.currentSourceUnit = b.parseInterfaceDefinition(ctx, interfaceCtx)
			b.sourceUnits = append(b.sourceUnits, b.currentSourceUnit)
		}

		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			b.currentSourceUnit = b.parseContractDefinition(ctx, contractCtx)
			b.sourceUnits = append(b.sourceUnits, b.currentSourceUnit)
		}
	}
}

func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &ast_pb.RootSourceUnit{SourceUnits: b.sourceUnits}

	// We should now discover the highest source unit that has the most of the
	// inheritance chain and set it as the execution source unit.
	// This is a bit of a hack, but it works for now.
	high := int64(0)

	for _, sourceUnit := range b.sourceUnits {
		for _, node := range sourceUnit.GetRoot().GetNodes() {
			if node.GetNodeType() == ast_pb.NodeType_CONTRACT_DEFINITION {
				for _, c := range node.GetLinearizedBaseContracts() {
					if c > high {
						high = c
					}
				}
			} else if node.GetNodeType() == ast_pb.NodeType_INTERFACE_DEFINITION {
				for _, c := range node.GetLinearizedBaseContracts() {
					if c > high {
						high = c
					}
				}
			} else if node.GetNodeType() == ast_pb.NodeType_LIBRARY_DEFINITION {
				for _, c := range node.GetLinearizedBaseContracts() {
					if c > high {
						high = c
					}
				}
			}
		}
	}

	if high > 0 {
		b.astRoot.EntrySourceUnit = high
		if node, ok := b.FindNodeById(high); ok {
			b.entrySourceUnit = node
		}
	}

	b.currentSourceUnit = nil
	b.currentStateVariables = nil
	b.currentEvents = nil
}
