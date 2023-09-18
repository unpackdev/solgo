package ast

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Contract represents a Solidity contract in the abstract syntax tree.
type Contract struct {
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

// NewContractDefinition creates a new instance of Contract.
func NewContractDefinition(b *ASTBuilder) *Contract {
	return &Contract{
		ASTBuilder: b,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Contract node.
// This function always returns false for now.
func (c *Contract) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the Contract.
func (c *Contract) GetId() int64 {
	return c.Id
}

// GetType returns the NodeType of the Contract.
func (c *Contract) GetType() ast_pb.NodeType {
	return c.NodeType
}

// GetSrc returns the source information of the Contract.
func (c *Contract) GetSrc() SrcNode {
	return c.Src
}

// GetName returns the name of the Contract.
func (c *Contract) GetName() string {
	return c.Name
}

// IsAbstract returns whether the Contract is abstract.
func (c *Contract) IsAbstract() bool {
	return c.Abstract
}

// GetKind returns the kind of the Contract.
func (c *Contract) GetKind() ast_pb.NodeType {
	return c.Kind
}

// IsFullyImplemented returns whether the Contract is fully implemented.
func (c *Contract) IsFullyImplemented() bool {
	return c.FullyImplemented
}

// GetNodes returns the child nodes of the Contract.
func (c *Contract) GetNodes() []Node[NodeType] {
	return c.Nodes
}

// GetLinearizedBaseContracts returns the linearized base contracts of the Contract.
func (c *Contract) GetLinearizedBaseContracts() []int64 {
	return c.LinearizedBaseContracts
}

// GetBaseContracts returns the base contracts of the Contract.
func (c *Contract) GetBaseContracts() []*BaseContract {
	return c.BaseContracts
}

// GetContractDependencies returns the contract dependencies of the Contract.
func (c *Contract) GetContractDependencies() []int64 {
	return c.ContractDependencies
}

// GetTypeDescription returns the type description associated with the Contract.
func (c *Contract) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     fmt.Sprintf("contract %s", c.Name),
		TypeIdentifier: fmt.Sprintf("$_t_contract_%s_%d", c.GetName(), c.GetId()),
	}
}

// GetStateVariables returns the state variables defined in the Contract.
func (s *Contract) GetStateVariables() []*StateVariableDeclaration {
	toReturn := make([]*StateVariableDeclaration, 0)

	for _, node := range s.GetNodes() {
		if stateVariable, ok := node.(*StateVariableDeclaration); ok {
			toReturn = append(toReturn, stateVariable)
		}
	}

	return toReturn
}

// GetStructs returns the struct definitions defined in the Contract.
func (s *Contract) GetStructs() []*StructDefinition {
	toReturn := make([]*StructDefinition, 0)

	for _, node := range s.GetNodes() {
		if structNode, ok := node.(*StructDefinition); ok {
			toReturn = append(toReturn, structNode)
		}
	}

	return toReturn
}

// GetEnums returns the enum definitions defined in the Contract.
func (s *Contract) GetEnums() []*EnumDefinition {
	toReturn := make([]*EnumDefinition, 0)

	for _, node := range s.GetNodes() {
		if enum, ok := node.(*EnumDefinition); ok {
			toReturn = append(toReturn, enum)
		}
	}

	return toReturn
}

// GetErrors returns the error definitions defined in the Contract.
func (s *Contract) GetErrors() []*ErrorDefinition {
	toReturn := make([]*ErrorDefinition, 0)

	for _, node := range s.GetNodes() {
		if errorNode, ok := node.(*ErrorDefinition); ok {
			toReturn = append(toReturn, errorNode)
		}
	}

	return toReturn
}

// GetEvents returns the event definitions defined in the Contract.
func (s *Contract) GetEvents() []*EventDefinition {
	toReturn := make([]*EventDefinition, 0)

	for _, node := range s.GetNodes() {
		if event, ok := node.(*EventDefinition); ok {
			toReturn = append(toReturn, event)
		}
	}

	return toReturn
}

// GetConstructor returns the constructor definition of the Contract.
func (s *Contract) GetConstructor() *Constructor {
	for _, node := range s.GetNodes() {
		if constructor, ok := node.(*Constructor); ok {
			return constructor
		}
	}

	return nil
}

// GetFunctions returns the function definitions defined in the Contract.
func (s *Contract) GetFunctions() []*Function {
	toReturn := make([]*Function, 0)

	for _, node := range s.GetNodes() {
		if function, ok := node.(*Function); ok {
			toReturn = append(toReturn, function)
		}
	}

	return toReturn
}

// GetFallback returns the fallback definition of the Contract.
func (s *Contract) GetFallback() *Fallback {
	for _, node := range s.GetNodes() {
		if function, ok := node.(*Fallback); ok {
			return function
		}
	}

	return nil
}

// GetReceive returns the receive definition of the Contract.
func (s *Contract) GetReceive() *Receive {
	for _, node := range s.GetNodes() {
		if function, ok := node.(*Receive); ok {
			return function
		}
	}

	return nil
}

// ToProto converts the Contract to its corresponding protocol buffer representation.
func (c *Contract) ToProto() NodeType {
	proto := ast_pb.Contract{
		Id:                      c.GetId(),
		NodeType:                c.GetType(),
		Kind:                    c.GetKind(),
		Src:                     c.GetSrc().ToProto(),
		Name:                    c.GetName(),
		Abstract:                c.IsAbstract(),
		FullyImplemented:        c.IsFullyImplemented(),
		LinearizedBaseContracts: c.GetLinearizedBaseContracts(),
		ContractDependencies:    c.GetContractDependencies(),
		Nodes:                   make([]*v3.TypedStruct, 0),
		BaseContracts:           make([]*ast_pb.BaseContract, 0),
	}

	for _, baseContract := range c.GetBaseContracts() {
		proto.BaseContracts = append(proto.BaseContracts, baseContract.ToProto())
	}

	for _, node := range c.GetNodes() {
		proto.Nodes = append(proto.Nodes, node.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Contract")
}

// Parse parses the Contract node from the parsing context and associates it with other nodes.
func (c *Contract) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.ContractDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
	unit.Src = SrcNode{
		Id:          c.GetNextID(),
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
	unit.SetAbsolutePathFromSources(c.sources)
	unit.ExportedSymbols = append(unit.ExportedSymbols, Symbol{
		Id:           unit.Id,
		Name:         unit.Name,
		AbsolutePath: unit.AbsolutePath,
	})

	// Now we are going to resolve pragmas for current source unit...
	unit.Nodes = append(
		unit.Nodes,
		parsePragmasForSourceUnit(c.ASTBuilder, unitCtx, unit, nil, ctx, nil)...,
	)

	// Now we are going to resolve import paths for current source unit...
	nodeImports := parseImportPathsForSourceUnit(c.ASTBuilder, unitCtx, unit, nil, ctx, nil)
	unit.Nodes = append(
		unit.Nodes,
		nodeImports...,
	)

	contractNode := &Contract{
		Id:   c.GetNextID(),
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
			c.ASTBuilder, unit, contractNode, ctx.InheritanceSpecifierList(),
		)...,
	)
	unit.BaseContracts = contractNode.BaseContracts

	contractNode.LinearizedBaseContracts = append(
		contractNode.LinearizedBaseContracts,
		contractNode.GetId(),
	)

	for _, nodeImport := range nodeImports {
		contractNode.LinearizedBaseContracts = append(
			contractNode.LinearizedBaseContracts,
			nodeImport.GetId(),
		)

		contractNode.ContractDependencies = append(
			contractNode.ContractDependencies,
			nodeImport.GetId(),
		)
	}

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			contractNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode(c.ASTBuilder)
		childNode := bodyNode.ParseDefinitions(unit, contractNode, bodyElement)
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
		} else {
			contractNode.FullyImplemented = false
		}
	}

	unit.Nodes = append(unit.Nodes, contractNode)
	unit.Contract = contractNode
}
