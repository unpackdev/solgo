package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/parser"
)

type SourceUnit[T NodeType] struct {
	// Id is the unique identifier of the source unit.
	Id int64 `json:"id"`

	// License is the license of the source unit.
	License string `json:"license"`

	// ExportedSymbols is the list of source units, including its names
	// and node tree ids used by current source unit.
	ExportedSymbols []Symbol `json:"exported_symbols"`

	// AbsolutePath is the absolute path of the source unit.
	AbsolutePath string `json:"absolute_path"`

	// Name is the name of the source unit.
	// This is going to be one of the following: contract, interface or library name.
	// It's here for convenience.
	Name string `json:"name"`

	// NodeType is the type of the AST node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Nodes is the list of AST nodes.
	Nodes []Node[NodeType] `json:"nodes"`

	// Src is the source code location.
	Src SrcNode `json:"src"`
}

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

func (s *SourceUnit[T]) SetAbsolutePathFromSources(sources solgo.Sources) {
	for _, unit := range sources.SourceUnits {
		if unit.Name == s.Name {
			s.AbsolutePath = unit.Path
		}
	}
}

func (s *SourceUnit[T]) GetNodes() []Node[NodeType] {
	return s.Nodes
}

func (s *SourceUnit[T]) GetName() string {
	return s.Name
}

func (s *SourceUnit[T]) GetId() int64 {
	return s.Id
}

func (s *SourceUnit[T]) GetType() ast_pb.NodeType {
	return s.NodeType
}

func (s *SourceUnit[T]) GetSrc() SrcNode {
	return s.Src
}

func (s *SourceUnit[T]) GetExportedSymbols() []Symbol {
	return s.ExportedSymbols
}

func (s *SourceUnit[T]) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_contract$_%s_$%d", s.Name, s.Id),
		TypeString:     fmt.Sprintf("contract %s", s.Name),
	}
}

func (s *SourceUnit[T]) ToProto() NodeType {
	return &ast_pb.SourceUnit{
		Id:           s.Id,
		License:      s.License,
		AbsolutePath: s.AbsolutePath,
		Name:         s.Name,
		NodeType:     s.NodeType,
	}
}

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	license := GetLicense(b.comments)

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

func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot.SourceUnits = append(b.astRoot.SourceUnits, b.sourceUnits...)
}
