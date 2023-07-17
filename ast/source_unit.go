package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	id := atomic.AddInt64(&b.nextID, 1) - 1

	b.currentSourceUnit = &ast_pb.SourceUnit{
		Id:              id,
		AbsolutePath:    ctx.GetStart().GetInputStream().GetSourceName(),
		ExportedSymbols: make([]*ast_pb.ExportedSymbols, 0),
		NodeType:        ast_pb.NodeType_SOURCE_UNIT,
		Nodes:           &ast_pb.RootNode{},
		Src: &ast_pb.Src{
			Line:   int64(ctx.GetStart().GetLine()),
			Column: int64(ctx.GetStart().GetColumn()),
			Start:  int64(ctx.GetStart().GetStart()),
			// @TODO: GetStop() is always nil due to some reason so we cannot get lenght
			// just yet. We need to figure out why.
			//Length: ctx.GetStop() - ctx.GetStart() + 1,
			ParentIndex: int64(ctx.GetStart().GetTokenIndex()),
		},
		Comments: b.comments,
	}

	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			_ = contractCtx
		}

		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			b.currentSourceUnit.License = b.GetLicense()

			// Alright lets extract bloody pragmas...
			pragmas := b.findPragmasForLibrary(ctx, libraryCtx)
			b.currentSourceUnit.Nodes.Nodes = append(
				b.currentSourceUnit.Nodes.Nodes,
				pragmas...,
			)
		}
	}

	// Just temporary...
	b.currentSourceUnit.Comments = nil
}

// ExitSourceUnit is called when production sourceUnit is exited.
func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &ast_pb.RootSourceUnit{
		SourceUnits: []*ast_pb.SourceUnit{b.currentSourceUnit},
	}
	b.currentSourceUnit = nil
}
