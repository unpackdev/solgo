package ast

import (
	"fmt"
	"path/filepath"
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// Import represents an import node in the abstract syntax tree.
type Import struct {
	Id           int64           `json:"id"`            // Unique identifier of the import node.
	NodeType     ast_pb.NodeType `json:"node_type"`     // Type of the node.
	Src          SrcNode         `json:"src"`           // Source location information.
	AbsolutePath string          `json:"absolute_path"` // Absolute path of the imported file.
	File         string          `json:"file"`          // Filepath of the import statement.
	Scope        int64           `json:"scope"`         // Scope of the import.
	UnitAlias    string          `json:"unit_alias"`    // Alias of the imported unit.
	SourceUnit   int64           `json:"source_unit"`   // Source unit identifier.
}

// SetReferenceDescriptor sets the reference descriptions of the Import node.
func (i *Import) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	// This function sets the source unit of the import in the resolver as a forward declaration hack.
	if refId > 0 && refDesc == nil {
		i.SourceUnit = refId
		return true
	}
	return false
}

// GetId returns the unique identifier of the import node.
func (i *Import) GetId() int64 {
	return i.Id
}

// GetType returns the type of the node.
func (i *Import) GetType() ast_pb.NodeType {
	return i.NodeType
}

// GetSrc returns the source location information of the import node.
func (i *Import) GetSrc() SrcNode {
	return i.Src
}

// GetTypeDescription returns the type description of the import node.
func (i *Import) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "import",
		TypeIdentifier: fmt.Sprintf("$_t_import_%s_%d", i.AbsolutePath, i.Id),
	}
}

// GetAbsolutePath returns the absolute path of the imported file.
func (i *Import) GetAbsolutePath() string {
	return i.AbsolutePath
}

// GetFile returns the filepath of the import statement.
func (i *Import) GetFile() string {
	return i.File
}

// GetScope returns the scope of the import.
func (i *Import) GetScope() int64 {
	return i.Scope
}

// GetUnitAlias returns the alias of the imported unit.
func (i *Import) GetUnitAlias() string {
	return i.UnitAlias
}

// GetSourceUnit returns the source unit identifier of the import.
func (i *Import) GetSourceUnit() int64 {
	return i.SourceUnit
}

// GetNodes returns an empty slice of nodes associated with the import.
func (i *Import) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetName returns the name of the imported file (excluding extension).
func (i *Import) GetName() string {
	base := filepath.Base(i.AbsolutePath)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// ToProto converts the Import node to its corresponding protobuf representation.
func (i *Import) ToProto() NodeType {
	proto := ast_pb.Import{
		Id:           i.GetId(),
		NodeType:     i.GetType(),
		Src:          i.GetSrc().ToProto(),
		AbsolutePath: i.GetAbsolutePath(),
		File:         i.GetFile(),
		Scope:        i.GetScope(),
		UnitAlias:    i.GetUnitAlias(),
		SourceUnit:   i.GetSourceUnit(),
	}

	return NewTypedStruct(&proto, "Import")
}

// parseImportPathsForSourceUnit is a utility function for parsing import paths within a source unit.
// It returns a slice of Import nodes corresponding to the imported paths.
func parseImportPathsForSourceUnit(
	b *ASTBuilder,
	unitCtx *parser.SourceUnitContext,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	libraryCtx *parser.LibraryDefinitionContext,
	contractCtx *parser.ContractDefinitionContext,
	interfaceCtx *parser.InterfaceDefinitionContext,
) []Node[NodeType] {
	imports := make([]*Import, 0)

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
			importNode := &Import{
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
							NewSymbol(symbol.Id, symbol.Name, symbol.AbsolutePath),
						)

					}
				}
			}
			filteredImports = append([]Node[NodeType]{importNode}, filteredImports...)
		}
	}

	return filteredImports
}
