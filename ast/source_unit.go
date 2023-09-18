package ast

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/parser"
)

// SourceUnit represents a source unit in the abstract syntax tree.
// It includes various attributes like id, license, exported symbols, absolute path, name, node type, nodes, and source node.
type SourceUnit[T NodeType] struct {
	Id              int64            `json:"id"`               // Id is the unique identifier of the source unit.
	Contract        Node[NodeType]   `json:"-"`                // Contract is the contract associated with the source unit.
	BaseContracts   []*BaseContract  `json:"base_contracts"`   // BaseContracts are the base contracts of the source unit.
	License         string           `json:"license"`          // License is the license of the source unit.
	ExportedSymbols []Symbol         `json:"exported_symbols"` // ExportedSymbols is the list of source units, including its names and node tree ids used by current source unit.
	AbsolutePath    string           `json:"absolute_path"`    // AbsolutePath is the absolute path of the source unit.
	Name            string           `json:"name"`             // Name is the name of the source unit. This is going to be one of the following: contract, interface or library name. It's here for convenience.
	NodeType        ast_pb.NodeType  `json:"node_type"`        // NodeType is the type of the AST node.
	Nodes           []Node[NodeType] `json:"nodes"`            // Nodes is the list of AST nodes.
	Src             SrcNode          `json:"src"`              // Src is the source code location.
}

// NewSourceUnit creates a new SourceUnit with the provided ASTBuilder, name, and license.
// It returns a pointer to the created SourceUnit.
func NewSourceUnit[T any](builder *ASTBuilder, name string, license string) *SourceUnit[T] {
	return &SourceUnit[T]{
		Id:              builder.GetNextID(),
		Name:            name,
		License:         license,
		Nodes:           make([]Node[NodeType], 0),
		NodeType:        ast_pb.NodeType_SOURCE_UNIT,
		ExportedSymbols: make([]Symbol, 0),
		BaseContracts:   make([]*BaseContract, 0),
	}
}

// SetAbsolutePathFromSources sets the absolute path of the source unit from the provided sources.
func (s *SourceUnit[T]) SetAbsolutePathFromSources(sources *solgo.Sources) {
	for _, unit := range sources.SourceUnits {
		if unit.Name == s.Name {
			s.AbsolutePath = unit.Path
		}
	}
}

// SetReferenceDescriptor sets the reference descriptions of the SourceUnit node.
func (s *SourceUnit[T]) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetLicense returns the license of the source unit.
func (s *SourceUnit[T]) GetLicense() string {
	return s.License
}

// GetNodes returns the nodes associated with the source unit.
func (s *SourceUnit[T]) GetNodes() []Node[NodeType] {
	return s.Nodes
}

// GetName returns the name of the source unit.
func (s *SourceUnit[T]) GetName() string {
	return s.Name
}

// GetId returns the unique identifier of the source unit.
func (s *SourceUnit[T]) GetId() int64 {
	return s.Id
}

// GetType returns the type of the source unit.
func (s *SourceUnit[T]) GetType() ast_pb.NodeType {
	return s.NodeType
}

// GetSrc returns the source code location of the source unit.
func (s *SourceUnit[T]) GetSrc() SrcNode {
	return s.Src
}

// GetExportedSymbols returns the exported symbols of the source unit.
func (s *SourceUnit[T]) GetExportedSymbols() []Symbol {
	return s.ExportedSymbols
}

// GetAbsolutePath returns the absolute path of the source unit.
func (s *SourceUnit[T]) GetAbsolutePath() string {
	return s.AbsolutePath
}

// GetContract returns the contract associated with the source unit.
func (s *SourceUnit[T]) GetContract() Node[NodeType] {
	return s.Contract
}

// GetBaseContracts returns the base contracts of the source unit.
func (s *SourceUnit[T]) GetBaseContracts() []*BaseContract {
	return s.BaseContracts
}

// GetTypeDescription returns the type description of the source unit.
func (s *SourceUnit[T]) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_contract$_%s_$%d", s.Name, s.Id),
		TypeString:     fmt.Sprintf("contract %s", s.Name),
	}
}

func (s *SourceUnit[T]) GetImports() []*Import {
	toReturn := make([]*Import, 0)

	for _, node := range s.Nodes {
		if node.GetType() == ast_pb.NodeType_IMPORT_DIRECTIVE {
			toReturn = append(toReturn, node.(*Import))
		}
	}

	return toReturn
}

func (s *SourceUnit[T]) GetPragmas() []*Pragma {
	toReturn := make([]*Pragma, 0)

	for _, node := range s.Nodes {
		if node.GetType() == ast_pb.NodeType_PRAGMA_DIRECTIVE {
			toReturn = append(toReturn, node.(*Pragma))
		}
	}

	return toReturn
}

// ToProto converts the SourceUnit to a protocol buffer representation.
func (s *SourceUnit[T]) ToProto() NodeType {
	exportedSymbols := []*ast_pb.ExportedSymbol{}

	for _, symbol := range s.ExportedSymbols {
		exportedSymbols = append(
			exportedSymbols,
			&ast_pb.ExportedSymbol{
				Id:           symbol.GetId(),
				Name:         symbol.GetName(),
				AbsolutePath: symbol.GetAbsolutePath(),
			},
		)
	}

	nodes := []*v3.TypedStruct{}

	for _, node := range s.Nodes {
		nodes = append(nodes, node.ToProto().(*v3.TypedStruct))
	}

	return &ast_pb.SourceUnit{
		Id:              s.Id,
		License:         s.License,
		AbsolutePath:    s.AbsolutePath,
		Name:            s.Name,
		NodeType:        s.NodeType,
		Src:             s.GetSrc().ToProto(),
		ExportedSymbols: exportedSymbols,
		Root: &ast_pb.RootNode{
			Nodes: nodes,
		},
	}
}

// EnterSourceUnit is called when the ASTBuilder enters a source unit context.
// It initializes a new root node and source units based on the context.
func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {

	rootNode := NewRootNode(b, 0, b.sourceUnits, b.comments)
	b.tree.SetRoot(rootNode)

	for _, child := range ctx.GetChildren() {
		if interfaceCtx, ok := child.(*parser.InterfaceDefinitionContext); ok {
			license := getLicenseFromSources(b.sources, b.comments, interfaceCtx.Identifier().GetText())
			sourceUnit := NewSourceUnit[Node[ast_pb.SourceUnit]](b, interfaceCtx.Identifier().GetText(), license)
			interfaceNode := NewInterfaceDefinition(b)
			interfaceNode.Parse(ctx, interfaceCtx, rootNode, sourceUnit)
			b.sourceUnits = append(b.sourceUnits, sourceUnit)
		}

		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			license := getLicenseFromSources(b.sources, b.comments, libraryCtx.Identifier().GetText())
			sourceUnit := NewSourceUnit[Node[ast_pb.SourceUnit]](b, libraryCtx.Identifier().GetText(), license)
			libraryNode := NewLibraryDefinition(b)
			libraryNode.Parse(ctx, libraryCtx, rootNode, sourceUnit)
			b.sourceUnits = append(b.sourceUnits, sourceUnit)
		}

		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			license := getLicenseFromSources(b.sources, b.comments, contractCtx.Identifier().GetText())
			sourceUnit := NewSourceUnit[Node[ast_pb.SourceUnit]](b, contractCtx.Identifier().GetText(), license)
			contractNode := NewContractDefinition(b)
			contractNode.Parse(ctx, contractCtx, rootNode, sourceUnit)
			b.sourceUnits = append(b.sourceUnits, sourceUnit)
		}
	}
}

// ExitSourceUnit is called when the ASTBuilder exits a source unit context.
// It appends the source units to the root node.
func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.tree.AppendRootNodes(b.sourceUnits...)
}
