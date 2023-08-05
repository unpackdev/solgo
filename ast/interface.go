package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Interface struct {
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

func NewInterfaceDefinition(b *ASTBuilder) *Interface {
	return &Interface{
		ASTBuilder:              b,
		LinearizedBaseContracts: make([]int64, 0),
		ContractDependencies:    make([]int64, 0),
		BaseContracts:           make([]*BaseContract, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Interface node.
func (l *Interface) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (l *Interface) GetId() int64 {
	return l.Id
}

func (l *Interface) GetType() ast_pb.NodeType {
	return l.NodeType
}

func (l *Interface) GetSrc() SrcNode {
	return l.Src
}

func (l *Interface) GetTypeDescription() *TypeDescription {
	return nil
}

func (l *Interface) GetName() string {
	return l.Name
}

func (l *Interface) IsAbstract() bool {
	return l.Abstract
}

func (l *Interface) GetKind() ast_pb.NodeType {
	return l.Kind
}

func (l *Interface) IsFullyImplemented() bool {
	return l.FullyImplemented
}

func (l *Interface) GetNodes() []Node[NodeType] {
	return l.Nodes
}

func (l *Interface) GetBaseContracts() []*BaseContract {
	return l.BaseContracts
}

func (l *Interface) GetContractDependencies() []int64 {
	return l.ContractDependencies
}

func (l *Interface) GetLinearizedBaseContracts() []int64 {
	return l.LinearizedBaseContracts
}

func (l *Interface) GetStateVariables() []*StateVariableDeclaration {
	toReturn := make([]*StateVariableDeclaration, 0)
	return toReturn
}

func (l *Interface) ToProto() NodeType {
	proto := ast_pb.Contract{
		Id:                      l.Id,
		NodeType:                l.NodeType,
		Kind:                    l.Kind,
		Src:                     l.Src.ToProto(),
		Name:                    l.Name,
		Abstract:                l.Abstract,
		FullyImplemented:        l.FullyImplemented,
		LinearizedBaseContracts: l.LinearizedBaseContracts,
		ContractDependencies:    l.ContractDependencies,
		Nodes:                   make([]*v3.TypedStruct, 0),
		BaseContracts:           make([]*ast_pb.BaseContract, 0),
	}

	for _, baseContract := range l.BaseContracts {
		proto.BaseContracts = append(proto.BaseContracts, baseContract.ToProto())
	}

	for _, node := range l.Nodes {
		proto.Nodes = append(proto.Nodes, node.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Contract")
}

func (l *Interface) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.InterfaceDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
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
		parsePragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, nil, ctx)...,
	)

	// Now we are going to resolve import paths for current source unit...
	nodeImports := parseImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, nil, ctx)
	unit.Nodes = append(
		unit.Nodes,
		nodeImports...,
	)

	interfaceNode := &Interface{
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
		Nodes:            make([]Node[NodeType], 0),
		FullyImplemented: true,
	}

	interfaceNode.BaseContracts = append(
		interfaceNode.BaseContracts,
		parseInheritanceFromCtx(
			l.ASTBuilder, unit, interfaceNode, ctx.InheritanceSpecifierList(),
		)...,
	)
	unit.BaseContracts = interfaceNode.BaseContracts

	interfaceNode.LinearizedBaseContracts = append(
		interfaceNode.LinearizedBaseContracts,
		interfaceNode.GetId(),
	)

	for _, nodeImport := range nodeImports {
		interfaceNode.LinearizedBaseContracts = append(
			interfaceNode.LinearizedBaseContracts,
			nodeImport.GetId(),
		)

		interfaceNode.ContractDependencies = append(
			interfaceNode.ContractDependencies,
			nodeImport.GetId(),
		)
	}

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			interfaceNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode(l.ASTBuilder)
		childNode := bodyNode.ParseDefinitions(unit, interfaceNode, bodyElement)
		if childNode != nil {
			interfaceNode.Nodes = append(
				interfaceNode.Nodes,
				childNode,
			)

			if bodyNode.NodeType == ast_pb.NodeType_FUNCTION_DEFINITION {
				if !bodyNode.Implemented {
					interfaceNode.FullyImplemented = false
				}
			}
		} else {
			interfaceNode.FullyImplemented = false
		}
	}

	unit.Nodes = append(unit.Nodes, interfaceNode)
	unit.Contract = interfaceNode
}
