package ast

import (
	"fmt"
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Contract represents a Solidity contract in the abstract syntax tree.
type Contract struct {
	*ASTBuilder

	Id                      int64            `json:"id"`
	Name                    string           `json:"name"`
	NodeType                ast_pb.NodeType  `json:"nodeType"`
	Src                     SrcNode          `json:"src"`
	NameLocation            SrcNode          `json:"nameLocation"`
	Abstract                bool             `json:"abstract"`
	Kind                    ast_pb.NodeType  `json:"kind"`
	FullyImplemented        bool             `json:"fullyImplemented"`
	Nodes                   []Node[NodeType] `json:"nodes"`
	LinearizedBaseContracts []int64          `json:"linearizedBaseContracts"`
	BaseContracts           []*BaseContract  `json:"baseContracts"`
	ContractDependencies    []int64          `json:"contractDependencies"`
}

// NewContractDefinition creates a new instance of Contract.
func NewContractDefinition(b *ASTBuilder) *Contract {
	return &Contract{
		ASTBuilder:              b,
		NodeType:                ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:                    ast_pb.NodeType_KIND_CONTRACT,
		LinearizedBaseContracts: make([]int64, 0),
		ContractDependencies:    make([]int64, 0),
		BaseContracts:           make([]*BaseContract, 0),
		Nodes:                   make([]Node[NodeType], 0),
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

// GetNameLocation returns the source information of the name of the Contract.
func (c *Contract) GetNameLocation() SrcNode {
	return c.NameLocation
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

// UnmarshalJSON parses the JSON-encoded data and stores the result in the Contract.
func (s *Contract) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &s.Id); err != nil {
			return err
		}

	}

	if name, ok := tempMap["name"]; ok {
		if err := json.Unmarshal(name, &s.Name); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["nodeType"]; ok {
		if err := json.Unmarshal(nodeType, &s.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &s.Src); err != nil {
			return err
		}
	}

	if nameLocation, ok := tempMap["nameLocation"]; ok {
		if err := json.Unmarshal(nameLocation, &s.NameLocation); err != nil {
			return err
		}
	}

	if abstract, ok := tempMap["abstract"]; ok {
		if err := json.Unmarshal(abstract, &s.Abstract); err != nil {
			return err
		}
	}

	if kind, ok := tempMap["kind"]; ok {
		if err := json.Unmarshal(kind, &s.Kind); err != nil {
			return err
		}
	}

	if fullyImplemented, ok := tempMap["fullyImplemented"]; ok {
		if err := json.Unmarshal(fullyImplemented, &s.FullyImplemented); err != nil {
			return err
		}
	}

	if baseContracts, ok := tempMap["baseContracts"]; ok {
		if err := json.Unmarshal(baseContracts, &s.BaseContracts); err != nil {
			return err
		}
	}

	if lbc, ok := tempMap["linearizedBaseContracts"]; ok {
		if err := json.Unmarshal(lbc, &s.LinearizedBaseContracts); err != nil {
			return err
		}
	}

	if cd, ok := tempMap["contractDependencies"]; ok {
		if err := json.Unmarshal(cd, &s.ContractDependencies); err != nil {
			return err
		}
	}

	if n, ok := tempMap["nodes"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(n, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(tempNode, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(tempNode, tempNodeType)
			if err != nil {
				return err
			}
			s.Nodes = append(s.Nodes, node)
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
		NameLocation:            c.GetNameLocation().ToProto(),
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

	contractId := c.GetNextID()

	contractNode := &Contract{
		Id:   contractId,
		Name: ctx.Identifier().GetText(),
		Src: SrcNode{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: unit.Id,
		},
		NameLocation: SrcNode{
			Line:        int64(ctx.Identifier().GetStart().GetLine()),
			Column:      int64(ctx.Identifier().GetStart().GetColumn()),
			Start:       int64(ctx.Identifier().GetStart().GetStart()),
			End:         int64(ctx.Identifier().GetStop().GetStop()),
			Length:      int64(ctx.Identifier().GetStop().GetStop() - ctx.Identifier().GetStart().GetStart() + 1),
			ParentIndex: contractId,
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

		bodyNode := NewBodyNode(c.ASTBuilder, false)
		if childNode := bodyNode.ParseDefinitions(unit, contractNode, bodyElement); childNode != nil {
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
