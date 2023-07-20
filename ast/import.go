package ast

import (
	"path/filepath"
	"strings"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) findImportPathsForSourceUnit(
	sourceUnit *parser.SourceUnitContext,
	currentSourceUnit *ast_pb.SourceUnit,
	libraryCtx *parser.LibraryDefinitionContext,
	contractCtx *parser.ContractDefinitionContext,
	interfaceCtx *parser.InterfaceDefinitionContext,
) []*ast_pb.Node {
	imports := make([]*ast_pb.Node, 0)

	contractLine := func() int64 {
		if libraryCtx != nil {
			return int64(libraryCtx.GetStart().GetLine())
		} else if contractCtx != nil {
			return int64(contractCtx.GetStart().GetLine())
		} else if interfaceCtx != nil {
			return int64(interfaceCtx.GetStart().GetLine())
		}
		return 0
	}()

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if libraryCtx != nil && child == libraryCtx {
			// Found the library definition, stop traversing
			break
		}

		if contractCtx != nil && child == contractCtx {
			// Found the contract definition, stop traversing
			break
		}

		if importCtx, ok := child.(*parser.ImportDirectiveContext); ok {
			// First pragma encountered, add it to the result
			importNode := &ast_pb.Node{
				NodeType: ast_pb.NodeType_IMPORT_DIRECTIVE,
				Id:       atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(importCtx.GetStart().GetLine()),
					Column:      int64(importCtx.GetStart().GetColumn()),
					Start:       int64(importCtx.GetStart().GetStart()),
					End:         int64(importCtx.GetStop().GetStop()),
					Length:      int64(importCtx.GetStop().GetStop() - importCtx.GetStart().GetStart() + 1),
					ParentIndex: currentSourceUnit.Id,
				},
				AbsolutePath: func() string {
					toReturn := filepath.Base(importCtx.Path().GetText())
					toReturn = strings.ReplaceAll(toReturn, "\"", "")
					return toReturn
				}(),
				File: func() string {
					toReturn := importCtx.Path().GetText()
					toReturn = strings.ReplaceAll(toReturn, "\"", "")
					return toReturn
				}(),
				Scope: currentSourceUnit.Id,
			}

			if importCtx.GetUnitAlias() != nil {
				importNode.UnitAlias = importCtx.GetUnitAlias().GetText()
			}

			// Find the source unit that corresponds to the import path
			// and add the exported symbols to the current source unit exported symbols.
			// @TODO: Perhaps too much of iterations?
			for _, sourceUnit := range b.sourceUnits {
				for _, source := range b.sources.SourceUnits {
					absolutePath := filepath.Base(source.Path)
					if importNode.AbsolutePath == absolutePath {
						importNode.SourceUnit = sourceUnit.Id
					}
				}
			}

			imports = append(imports, importNode)
		}
	}

	filteredImports := make([]*ast_pb.Node, 0)

	for i := len(imports) - 1; i >= 0; i-- {
		importNode := imports[i]
		/* 		fmt.Println(
			importNode.Src.Line,
			contractLine,
			contractLine-importNode.Src.Line,
			importNode.AbsolutePath,
			importNode.Src.Line-int64(contractLine),
			(int64(contractLine)-importNode.Src.Line <= 20 && int64(contractLine)-importNode.Src.Line >= -1),
		) */
		if int64(contractLine)-importNode.Src.Line <= 20 && int64(contractLine)-importNode.Src.Line >= -1 {
			importNode.Src.ParentIndex = currentSourceUnit.Id
			filteredImports = append([]*ast_pb.Node{importNode}, filteredImports...)
		}
	}

	for _, importNode := range filteredImports {
		for _, sourceUnit := range b.sourceUnits {
			for _, symbol := range sourceUnit.ExportedSymbols {
				if symbol.AbsolutePath == importNode.AbsolutePath {
					currentSourceUnit.ExportedSymbols = append(
						currentSourceUnit.ExportedSymbols,
						&ast_pb.ExportedSymbol{
							Id:           symbol.Id,
							Name:         symbol.Name,
							AbsolutePath: symbol.AbsolutePath,
						},
					)
				}
			}
		}
	}

	return filteredImports
}
