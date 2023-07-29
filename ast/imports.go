package ast

import (
	"path/filepath"
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ImportNode struct {
	// Id is the unique identifier of the import node.
	Id int64 `json:"id"`

	NodeType     ast_pb.NodeType `json:"node_type"`
	Src          SrcNode         `json:"src"`
	AbsolutePath string          `json:"absolute_path"`
	File         string          `json:"file"`
	Scope        int64           `json:"scope"`
	UnitAlias    string          `json:"unit_alias"`
	SourceUnit   int64           `json:"source_unit"`
}

// SetReferenceDescriptor sets the reference descriptions of the ImportNode node.
func (i ImportNode) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (i ImportNode) GetId() int64 {
	return i.Id
}

func (i ImportNode) GetType() ast_pb.NodeType {
	return i.NodeType
}

func (i ImportNode) GetSrc() SrcNode {
	return i.Src
}

func (i *ImportNode) GetTypeDescription() *TypeDescription {
	return nil
}

func (i ImportNode) GetAbsolutePath() string {
	return i.AbsolutePath
}

func (i ImportNode) GetFile() string {
	return i.File
}

func (i ImportNode) GetScope() int64 {
	return i.Scope
}

func (i ImportNode) GetUnitAlias() string {
	return i.UnitAlias
}

func (i ImportNode) GetSourceUnit() int64 {
	return i.SourceUnit
}

func (i ImportNode) GetNodes() []Node[NodeType] {
	return nil
}

func (i ImportNode) ToProto() NodeType {
	return ast_pb.Import{}
}

func parseImportPathsForSourceUnit(
	b *ASTBuilder,
	unitCtx *parser.SourceUnitContext,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	libraryCtx *parser.LibraryDefinitionContext,
	contractCtx *parser.ContractDefinitionContext,
	interfaceCtx *parser.InterfaceDefinitionContext,
) []Node[NodeType] {
	imports := make([]*ImportNode, 0)

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
	for _, child := range unitCtx.GetChildren() {
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
			importNode := &ImportNode{
				Id:       b.GetNextID(),
				NodeType: ast_pb.NodeType_IMPORT_DIRECTIVE,
				Src: SrcNode{
					Line:        int64(importCtx.GetStart().GetLine()),
					Column:      int64(importCtx.GetStart().GetColumn()),
					Start:       int64(importCtx.GetStart().GetStart()),
					End:         int64(importCtx.GetStop().GetStop()),
					Length:      int64(importCtx.GetStop().GetStop() - importCtx.GetStart().GetStart() + 1),
					ParentIndex: unit.Id,
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
				Scope: unit.Id,
			}

			if importCtx.GetUnitAlias() != nil {
				importNode.UnitAlias = importCtx.GetUnitAlias().GetText()
			}

			// Find the source unit that corresponds to the import path
			// and add the exported symbols to the current source unit exported symbols.
			// @TODO: Perhaps too much of iterations?
			for _, unitCtx := range b.sourceUnits {
				for _, source := range b.sources.SourceUnits {
					absolutePath := filepath.Base(source.Path)
					if importNode.AbsolutePath == absolutePath {
						importNode.SourceUnit = unitCtx.Id
					}
				}
			}

			imports = append(imports, importNode)
		}
	}

	filteredImports := make([]Node[NodeType], 0)

	for i := len(imports) - 1; i >= 0; i-- {
		importNode := imports[i]
		if int64(contractLine)-importNode.Src.Line <= 20 && int64(contractLine)-importNode.Src.Line >= -1 {
			importNode.Src.ParentIndex = unit.Id
			for _, unitCtx := range b.sourceUnits {
				for _, symbol := range unitCtx.ExportedSymbols {
					if symbol.AbsolutePath == importNode.AbsolutePath {
						unit.ExportedSymbols = append(
							unit.ExportedSymbols,
							Symbol{
								Id:           symbol.Id,
								Name:         symbol.Name,
								AbsolutePath: symbol.AbsolutePath,
							},
						)
					}
				}
			}
			filteredImports = append([]Node[NodeType]{importNode}, filteredImports...)
		}
	}

	return filteredImports
}
