package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ContractNode[T any] struct {
	*ASTBuilder

	Id                      int64           `json:"id"`
	Name                    string          `json:"name"`
	NodeType                ast_pb.NodeType `json:"node_type"`
	Src                     SrcNode         `json:"src"`
	Abstract                bool            `json:"abstract"`
	Kind                    ast_pb.NodeType `json:"kind"`
	FullyImplemented        bool            `json:"fully_implemented"`
	Nodes                   []T             `json:"nodes"`
	LinearizedBaseContracts []int64         `json:"linearized_base_contracts"`
	BaseContracts           []*BaseContract `json:"base_contracts"`
	ContractDependencies    []int64         `json:"contract_dependencies"`
}

func NewContractDefinition(b *ASTBuilder) *ContractNode[Node] {
	return &ContractNode[Node]{
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

func (l ContractNode[T]) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.ContractDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node]) {
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
		FindPragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, ctx, nil)...,
	)

	// Now we are going to resolve import paths for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		FindImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, ctx, nil)...,
	)

	contractNode := &ContractNode[Node]{
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
		Nodes:                   make([]Node, 0),
		FullyImplemented:        true,
	}

	if ctx.InheritanceSpecifierList() != nil {
		inheritanceCtx := ctx.InheritanceSpecifierList()

		for _, inheritanceSpecifierCtx := range inheritanceCtx.AllInheritanceSpecifier() {
			baseContract := &BaseContract{
				Id: l.GetNextID(),
				Src: SrcNode{
					Id:          l.GetNextID(),
					Line:        int64(inheritanceSpecifierCtx.GetStart().GetLine()),
					Column:      int64(inheritanceSpecifierCtx.GetStart().GetColumn()),
					Start:       int64(inheritanceSpecifierCtx.GetStart().GetStart()),
					End:         int64(inheritanceSpecifierCtx.GetStop().GetStop()),
					Length:      int64(inheritanceSpecifierCtx.GetStop().GetStop() - inheritanceSpecifierCtx.GetStart().GetStart() + 1),
					ParentIndex: contractNode.Id,
				},
				NodeType: ast_pb.NodeType_INHERITANCE_SPECIFIER,
				BaseName: &BaseContractName{
					Id: l.GetNextID(),
					Src: SrcNode{
						Id:          l.GetNextID(),
						Line:        int64(inheritanceSpecifierCtx.GetStart().GetLine()),
						Column:      int64(inheritanceSpecifierCtx.GetStart().GetColumn()),
						Start:       int64(inheritanceSpecifierCtx.GetStart().GetStart()),
						End:         int64(inheritanceSpecifierCtx.GetStop().GetStop()),
						Length:      int64(inheritanceSpecifierCtx.GetStop().GetStop() - inheritanceSpecifierCtx.GetStart().GetStart() + 1),
						ParentIndex: contractNode.Id,
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
					Name:     inheritanceSpecifierCtx.IdentifierPath().GetText(),
				},
			}

			for _, unitNode := range l.sourceUnits {
				if unitNode.GetName() == inheritanceSpecifierCtx.IdentifierPath().GetText() {
					baseContract.BaseName.ReferencedDeclaration = unitNode.GetId()
					contractNode.LinearizedBaseContracts = append(
						contractNode.LinearizedBaseContracts, unitNode.GetId(),
					)
					contractNode.ContractDependencies = append(
						contractNode.ContractDependencies, unitNode.GetId(),
					)

					symbolFound := false
					for _, symbol := range unitNode.GetExportedSymbols() {
						if symbol.GetName() == unitNode.GetName() {
							symbolFound = true
						}
					}

					if !symbolFound {
						unit.ExportedSymbols = append(
							unit.ExportedSymbols,
							Symbol{
								Id:   unitNode.GetId(),
								Name: unitNode.GetName(),
								AbsolutePath: func() string {
									for _, unit := range l.sources.SourceUnits {
										if unit.Name == unitNode.GetName() {
											return unit.Path
										}
									}
									return ""
								}(),
							},
						)
					}
				}
			}

			contractNode.BaseContracts = append(contractNode.BaseContracts, baseContract)
		}
	}

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			contractNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode[Node](l.ASTBuilder)
		subBodyNode := bodyNode.Parse(unit, contractNode, bodyElement)
		if subBodyNode != nil {
			contractNode.Nodes = append(
				contractNode.Nodes,
				subBodyNode,
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
