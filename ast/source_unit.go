package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/parser"
)

// SourceUnit represents a source unit in the abstract syntax tree.
// It includes various attributes like id, license, exported symbols, absolute path, name, node type, nodes, and source node.
type SourceUnit[T NodeType] struct {
	Id              int64            `json:"id"`               // Id is the unique identifier of the source unit.
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
	}
}

// SetAbsolutePathFromSources sets the absolute path of the source unit from the provided sources.
func (s *SourceUnit[T]) SetAbsolutePathFromSources(sources solgo.Sources) {
	for _, unit := range sources.SourceUnits {
		if unit.Name == s.Name {
			s.AbsolutePath = unit.Path
		}
	}
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

// GetTypeDescription returns the type description of the source unit.
func (s *SourceUnit[T]) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_contract$_%s_$%d", s.Name, s.Id),
		TypeString:     fmt.Sprintf("contract %s", s.Name),
	}
}

// ToProto converts the SourceUnit to a protocol buffer representation.
func (s *SourceUnit[T]) ToProto() NodeType {
	return &ast_pb.SourceUnit{
		Id:           s.Id,
		License:      s.License,
		AbsolutePath: s.AbsolutePath,
		Name:         s.Name,
		NodeType:     s.NodeType,
	}
}

// EnterSourceUnit is called when the ASTBuilder enters a source unit context.
// It initializes a new root node and source units based on the context.
func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	license := getLicense(b.comments)

	rootNode := NewRootNode(b, 0, b.sourceUnits, b.comments).(*RootNode)
	b.astRoot = rootNode

	for _, child := range ctx.GetChildren() {
		if interfaceCtx, ok := child.(*parser.InterfaceDefinitionContext); ok {
			sourceUnit := NewSourceUnit[Node[ast_pb.SourceUnit]](b, interfaceCtx.Identifier().GetText(), license)
			interfaceNode := NewInterfaceDefinition(b)
			interfaceNode.Parse(ctx, interfaceCtx, rootNode, sourceUnit)
			b.sourceUnits = append(b.sourceUnits, sourceUnit)
		}

		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			sourceUnit := NewSourceUnit[Node[ast_pb.SourceUnit]](b, libraryCtx.Identifier().GetText(), license)
			libraryNode := NewLibraryDefinition(b)
			libraryNode.Parse(ctx, libraryCtx, rootNode, sourceUnit)
			b.sourceUnits = append(b.sourceUnits, sourceUnit)
		}

		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
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
	b.astRoot.SourceUnits = append(b.astRoot.SourceUnits, b.sourceUnits...)
}
