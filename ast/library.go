package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type LibraryNode[T any] struct {
	*ASTBuilder

	Id               int64           `json:"id"`
	Name             string          `json:"name"`
	NodeType         ast_pb.NodeType `json:"node_type"`
	Src              SrcNode         `json:"src"`
	Abstract         bool            `json:"abstract"`
	Kind             ast_pb.NodeType `json:"kind"`
	FullyImplemented bool            `json:"fully_implemented"`
	Nodes            []T             `json:"nodes"`
}

func NewLibraryDefinition(b *ASTBuilder) *LibraryNode[Node] {
	return &LibraryNode[Node]{
		ASTBuilder: b,
	}
}

func (l LibraryNode[T]) GetId() int64 {
	return l.Id
}

func (l LibraryNode[T]) GetType() ast_pb.NodeType {
	return l.NodeType
}

func (l LibraryNode[T]) GetSrc() SrcNode {
	return l.Src
}

func (l LibraryNode[T]) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.LibraryDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node]) {
	unit.Src = SrcNode{
		Id:          l.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: rootNode.Id,
	}

	// Set the absolute path of the source unit from provided sources map.
	// We are not dynamically loading files like the solc compiler does so we need to
	// provide the absolute path of the source unit from the sources map.
	unit.SetAbsolutePathFromSources(l.sources)
	unit.ExportedSymbols = append(unit.ExportedSymbols, Symbol{
		Id:           unit.Id,
		Name:         unit.Name,
		AbsolutePath: unit.AbsolutePath,
	})

	// Now we are going to resolve pragmas for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		FindPragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, ctx, nil, nil)...,
	)

	// Now we are going to resolve import paths for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		FindImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, ctx, nil, nil)...,
	)

	libraryNode := &LibraryNode[Node]{
		Id:   l.GetNextID(),
		Name: ctx.Identifier().GetText(),
		Src: SrcNode{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: unit.Id,
		},
		Abstract:         false,
		NodeType:         ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:             ast_pb.NodeType_KIND_LIBRARY,
		Nodes:            make([]Node, 0),
		FullyImplemented: true,
	}

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			libraryNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode[Node](l.ASTBuilder)
		subBodyNode := bodyNode.Parse(unit, libraryNode, bodyElement)
		if subBodyNode != nil {
			libraryNode.Nodes = append(
				libraryNode.Nodes,
				subBodyNode,
			)

			if bodyNode.NodeType == ast_pb.NodeType_FUNCTION_DEFINITION {
				if !bodyNode.Implemented {
					libraryNode.FullyImplemented = false
				}
			}

			//l.dumpNode(subBodyNode)
		} else {
			libraryNode.FullyImplemented = false
		}
	}

	unit.Nodes = append(unit.Nodes, libraryNode)
}
