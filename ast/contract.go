package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ContractNode[T NodeType] struct {
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
}

func NewContractDefinition(b *ASTBuilder) *ContractNode[Node[ast_pb.Contract]] {
	return &ContractNode[Node[ast_pb.Contract]]{
		ASTBuilder: b,
	}
}

func (l ContractNode[T]) GetId() int64 {
	return l.Id
}

func (l ContractNode[T]) GetType() ast_pb.NodeType {
	return l.NodeType
}

func (l ContractNode[T]) GetSrc() SrcNode {
	return l.Src
}

func (l ContractNode[T]) GetName() string {
	return l.Name
}

func (l ContractNode[T]) IsAbstract() bool {
	return l.Abstract
}

func (l ContractNode[T]) GetKind() ast_pb.NodeType {
	return l.Kind
}

func (l ContractNode[T]) IsFullyImplemented() bool {
	return l.FullyImplemented
}

func (l ContractNode[T]) GetNodes() []Node[NodeType] {
	return l.Nodes
}

func (l ContractNode[T]) GetLinearizedBaseContracts() []int64 {
	return l.LinearizedBaseContracts
}

func (l ContractNode[T]) GetBaseContracts() []*BaseContract {
	return l.BaseContracts
}

func (l ContractNode[T]) GetContractDependencies() []int64 {
	return l.ContractDependencies
}

func (l ContractNode[T]) GetTypeDescription() *TypeDescription {
	return nil
}

func (l ContractNode[T]) ToProto() NodeType {
	return ast_pb.Contract{}
}

func (l ContractNode[T]) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.ContractDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
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
		parsePragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, ctx, nil)...,
	)

	// Now we are going to resolve import paths for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		parseImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, ctx, nil)...,
	)

	contractNode := &ContractNode[ast_pb.Contract]{
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
		Abstract:                false,
		NodeType:                ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:                    ast_pb.NodeType_KIND_CONTRACT,
		LinearizedBaseContracts: make([]int64, 0),
		ContractDependencies:    make([]int64, 0),
		BaseContracts:           make([]*BaseContract, 0),
		Nodes:                   make([]Node[NodeType], 0),
		FullyImplemented:        true,
	}

	contractNode.BaseContracts = append(
		contractNode.BaseContracts,
		parseInheritanceFromCtx(
			l.ASTBuilder, unit, contractNode, ctx.InheritanceSpecifierList(),
		)...,
	)

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			contractNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode(l.ASTBuilder)
		childNode := bodyNode.Parse(unit, contractNode, bodyElement)
		if childNode != nil {
			contractNode.Nodes = append(
				contractNode.Nodes,
				childNode,
			)

			if bodyNode.NodeType == ast_pb.NodeType_FUNCTION_DEFINITION {
				if !bodyNode.Implemented {
					contractNode.FullyImplemented = false
				}
			}

			//l.dumpNode(subBodyNode)
		} else {
			contractNode.FullyImplemented = false
		}
	}

	unit.Nodes = append(unit.Nodes, contractNode)
}
