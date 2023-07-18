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
			panic("Interface definitions are not supported")
			_ = interfaceCtx
		}

		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			b.currentSourceUnit = b.parseContractDefinition(ctx, contractCtx)
			b.sourceUnits = append(b.sourceUnits, b.currentSourceUnit)
		}
	}
}

func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &ast_pb.RootSourceUnit{SourceUnits: b.sourceUnits}
	b.currentSourceUnit = nil
}
