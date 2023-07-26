package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type LibraryNode struct {
	*ASTBuilder

	Id                      int64            `json:"id"`
	Name                    string           `json:"name"`
	NodeType                ast_pb.NodeType  `json:"node_type"`
	Src                     SrcNode          `json:"src"`
	Abstract                bool             `json:"abstract"`
	Kind                    ast_pb.NodeType  `json:"kind"`
	FullyImplemented        bool             `json:"fully_implemented"`
	Nodes                   []Node[NodeType] `json:"nodes"`
	LinearizedBaseContracts []int64          `json:"linearized_base_contracts"`
	BaseContracts           []*BaseContract  `json:"base_contracts"`
	ContractDependencies    []int64          `json:"contract_dependencies"`
	Scope                   int64            `json:"scope"`
}

func NewLibraryDefinition(b *ASTBuilder) *LibraryNode {
	return &LibraryNode{
		ASTBuilder: b,
	}
}

func (l LibraryNode) GetId() int64 {
	return l.Id
}

func (l LibraryNode) GetType() ast_pb.NodeType {
	return l.NodeType
}

func (l LibraryNode) GetSrc() SrcNode {
	return l.Src
}

func (l LibraryNode) GetTypeDescription() *TypeDescription {
	return nil
}

func (l LibraryNode) GetName() string {
	return l.Name
}

func (l LibraryNode) IsAbstract() bool {
	return l.Abstract
}

func (l LibraryNode) GetKind() ast_pb.NodeType {
	return l.Kind
}

func (l LibraryNode) IsFullyImplemented() bool {
	return l.FullyImplemented
}

func (l LibraryNode) GetNodes() []Node[NodeType] {
	return l.Nodes
}

func (l LibraryNode) GetScope() int64 {
	return l.Scope
}

func (l LibraryNode) GetLinearizedBaseContracts() []int64 {
	return l.LinearizedBaseContracts
}

func (l LibraryNode) GetBaseContracts() []*BaseContract {
	return l.BaseContracts
}

func (l LibraryNode) GetContractDependencies() []int64 {
	return l.ContractDependencies
}

func (l LibraryNode) ToProto() NodeType {
	return ast_pb.Contract{}
}

func (l LibraryNode) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.LibraryDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
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
		parsePragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, ctx, nil, nil)...,
	)

	// Now we are going to resolve import paths for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		parseImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, ctx, nil, nil)...,
	)

	libraryId := l.GetNextID()
	libraryNode := &LibraryNode{
		Id:   libraryId,
		Name: ctx.Identifier().GetText(),
		Src: SrcNode{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: unit.Id,
		},
		Abstract:                false,
		NodeType:                ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:                    ast_pb.NodeType_KIND_LIBRARY,
		Nodes:                   make([]Node[NodeType], 0),
		FullyImplemented:        true,
		LinearizedBaseContracts: []int64{libraryId},
		ContractDependencies:    make([]int64, 0),
		BaseContracts:           make([]*BaseContract, 0),
		Scope:                   unit.Id,
	}

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			libraryNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode(l.ASTBuilder)
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
		} else {
			libraryNode.FullyImplemented = false
		}
	}

	unit.Nodes = append(unit.Nodes, libraryNode)
}
