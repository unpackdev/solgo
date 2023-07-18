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
	library *parser.LibraryDefinitionContext,
	contract *parser.ContractDefinitionContext,
) []*ast_pb.Node {
	imports := make([]*ast_pb.Node, 0)

	contractLine := func() int64 {
		if library != nil {
			return int64(library.GetStart().GetLine())
		} else if contract != nil {
			return int64(contract.GetStart().GetLine())
		}
		return 0
	}()

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if library != nil && child == library {
			// Found the library definition, stop traversing
			break
		}

		if contract != nil && child == contract {
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
						for _, symbol := range sourceUnit.ExportedSymbols {
							if symbol.AbsolutePath == absolutePath {
								currentSourceUnit.ExportedSymbols = append(
									currentSourceUnit.ExportedSymbols,
									&ast_pb.ExportedSymbol{
										Id:           symbol.Id,
										Name:         symbol.Name,
										AbsolutePath: absolutePath,
									},
								)
								break
							}
						}
					}
				}
			}

			imports = append(imports, importNode)
		}
	}

	filteredImports := make([]*ast_pb.Node, 0)
	maxLine := int64(-1)

	for i := len(imports) - 1; i >= 0; i-- {
		pragma := imports[i]
		if maxLine == -1 || (int64(contractLine)-pragma.Src.Line <= 20 && pragma.Src.Line-maxLine >= -1) {
			pragma.Src.ParentIndex = currentSourceUnit.Id
			filteredImports = append([]*ast_pb.Node{pragma}, filteredImports...)
			maxLine = pragma.Src.Line
		}
	}

	return filteredImports
}
