// Package ast defines data structures and methods for abstract syntax tree nodes used in a specific programming language.
// The package contains definitions for various AST nodes that represent different elements of the programming language's syntax.
package ast

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Interface represents an interface definition node in the abstract syntax tree (AST).
// It encapsulates information about the characteristics and properties of an interface within the contract.
type Interface struct {
	*ASTBuilder                              // Embedded ASTBuilder for building the AST.
	Id                      int64            `json:"id"`                        // Unique identifier for the Interface node.
	Name                    string           `json:"name"`                      // Name of the interface.
	NodeType                ast_pb.NodeType  `json:"node_type"`                 // Type of the AST node.
	Src                     SrcNode          `json:"src"`                       // Source location information.
	Abstract                bool             `json:"abstract"`                  // Indicates whether the interface is abstract.
	Kind                    ast_pb.NodeType  `json:"kind"`                      // Kind of the interface.
	FullyImplemented        bool             `json:"fully_implemented"`         // Indicates whether the interface is fully implemented.
	Nodes                   []Node[NodeType] `json:"nodes"`                     // List of child nodes within the interface.
	LinearizedBaseContracts []int64          `json:"linearized_base_contracts"` // List of linearized base contract identifiers.
	BaseContracts           []*BaseContract  `json:"base_contracts"`            // List of base contracts.
	ContractDependencies    []int64          `json:"contract_dependencies"`     // List of contract dependency identifiers.
}

// NewInterfaceDefinition creates a new Interface node with default values and returns it.
func NewInterfaceDefinition(b *ASTBuilder) *Interface {
	return &Interface{
		ASTBuilder:              b,
		LinearizedBaseContracts: make([]int64, 0),
		ContractDependencies:    make([]int64, 0),
		BaseContracts:           make([]*BaseContract, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Interface node.
// This function currently returns false, as no reference description updates are performed.
func (l *Interface) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the Interface node.
func (l *Interface) GetId() int64 {
	return l.Id
}

// GetType returns the type of the AST node, which is NodeType_CONTRACT_DEFINITION for an interface.
func (l *Interface) GetType() ast_pb.NodeType {
	return l.NodeType
}

// GetSrc returns the source location information of the Interface node.
func (l *Interface) GetSrc() SrcNode {
	return l.Src
}

// GetTypeDescription returns the type description associated with the Interface node.
func (l *Interface) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     fmt.Sprintf("contract %s", l.Name),
		TypeIdentifier: fmt.Sprintf("$_t_contract_%s_%d", l.GetName(), l.GetId()),
	}
}

// GetName returns the name of the interface.
func (l *Interface) GetName() string {
	return l.Name
}

// IsAbstract returns true if the Interface is abstract, false otherwise.
func (l *Interface) IsAbstract() bool {
	return l.Abstract
}

// GetKind returns the kind of the Interface node.
func (l *Interface) GetKind() ast_pb.NodeType {
	return l.Kind
}

// IsFullyImplemented returns true if the Interface is fully implemented, false otherwise.
func (l *Interface) IsFullyImplemented() bool {
	return l.FullyImplemented
}

// GetNodes returns a slice of child nodes within the interface.
func (l *Interface) GetNodes() []Node[NodeType] {
	return l.Nodes
}

// GetBaseContracts returns a list of base contracts associated with the Interface.
func (l *Interface) GetBaseContracts() []*BaseContract {
	return l.BaseContracts
}

// GetContractDependencies returns a list of contract dependency identifiers for the Interface.
func (l *Interface) GetContractDependencies() []int64 {
	return l.ContractDependencies
}

// GetLinearizedBaseContracts returns a list of linearized base contract identifiers for the Interface.
func (l *Interface) GetLinearizedBaseContracts() []int64 {
	return l.LinearizedBaseContracts
}

// GetStateVariables returns a list of state variable declarations within the Interface.
func (l *Interface) GetStateVariables() []*StateVariableDeclaration {
	toReturn := make([]*StateVariableDeclaration, 0)

	for _, node := range l.GetNodes() {
		if stateVariable, ok := node.(*StateVariableDeclaration); ok {
			toReturn = append(toReturn, stateVariable)
		}
	}

	return toReturn
}

// GetStructs returns a list of struct definitions within the Interface.
func (l *Interface) GetStructs() []*StructDefinition {
	toReturn := make([]*StructDefinition, 0)

	for _, node := range l.GetNodes() {
		if structNode, ok := node.(*StructDefinition); ok {
			toReturn = append(toReturn, structNode)
		}
	}

	return toReturn
}

// GetEnums returns a list of enum definitions within the Interface.
func (l *Interface) GetEnums() []*EnumDefinition {
	toReturn := make([]*EnumDefinition, 0)

	for _, node := range l.GetNodes() {
		if enum, ok := node.(*EnumDefinition); ok {
			toReturn = append(toReturn, enum)
		}
	}

	return toReturn
}

// GetErrors returns a list of error definitions within the Interface.
func (l *Interface) GetErrors() []*ErrorDefinition {
	toReturn := make([]*ErrorDefinition, 0)

	for _, node := range l.GetNodes() {
		if errorNode, ok := node.(*ErrorDefinition); ok {
			toReturn = append(toReturn, errorNode)
		}
	}

	return toReturn
}

// GetEvents returns a list of event definitions within the Interface.
func (l *Interface) GetEvents() []*EventDefinition {
	toReturn := make([]*EventDefinition, 0)

	for _, node := range l.GetNodes() {
		if event, ok := node.(*EventDefinition); ok {
			toReturn = append(toReturn, event)
		}
	}

	return toReturn
}

// GetConstructor returns the constructor node within the Interface, if present.
func (l *Interface) GetConstructor() *Constructor {
	for _, node := range l.GetNodes() {
		if constructor, ok := node.(*Constructor); ok {
			return constructor
		}
	}

	return nil
}

// GetFunctions returns a list of function definitions within the Interface.
func (l *Interface) GetFunctions() []*Function {
	toReturn := make([]*Function, 0)

	for _, node := range l.GetNodes() {
		if function, ok := node.(*Function); ok {
			toReturn = append(toReturn, function)
		}
	}

	return toReturn
}

// GetFallback returns the fallback function node within the Interface, if present.
func (l *Interface) GetFallback() *Fallback {
	for _, node := range l.GetNodes() {
		if function, ok := node.(*Fallback); ok {
			return function
		}
	}

	return nil
}

// GetReceive returns the receive function node within the Interface, if present.
func (l *Interface) GetReceive() *Receive {
	for _, node := range l.GetNodes() {
		if function, ok := node.(*Receive); ok {
			return function
		}
	}

	return nil
}

// ToProto converts the Interface node to its corresponding protocol buffer representation.
func (l *Interface) ToProto() NodeType {
	proto := ast_pb.Contract{
		Id:                      l.GetId(),
		NodeType:                l.GetType(),
		Kind:                    l.GetKind(),
		Src:                     l.GetSrc().ToProto(),
		Name:                    l.GetName(),
		Abstract:                l.IsAbstract(),
		FullyImplemented:        l.IsFullyImplemented(),
		LinearizedBaseContracts: l.GetLinearizedBaseContracts(),
		ContractDependencies:    l.GetContractDependencies(),
		Nodes:                   make([]*v3.TypedStruct, 0),
		BaseContracts:           make([]*ast_pb.BaseContract, 0),
	}

	for _, baseContract := range l.GetBaseContracts() {
		proto.BaseContracts = append(proto.BaseContracts, baseContract.ToProto())
	}

	for _, node := range l.GetNodes() {
		proto.Nodes = append(proto.Nodes, node.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Contract")
}

// Parse is responsible for parsing the interface definition from the source unit context and populating the Interface node.
func (l *Interface) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.InterfaceDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
	// Set the source location information for the source unit.
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
	// The absolute path is used to locate the source unit.
	unit.SetAbsolutePathFromSources(l.sources)
	// Add the exported symbol information for the source unit.
	unit.ExportedSymbols = append(unit.ExportedSymbols, Symbol{
		Id:           unit.Id,
		Name:         unit.Name,
		AbsolutePath: unit.AbsolutePath,
	})

	// Resolve pragmas for the source unit.
	unit.Nodes = append(unit.Nodes, parsePragmasForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, nil, ctx)...)
	// Resolve import paths for the source unit.
	nodeImports := parseImportPathsForSourceUnit(l.ASTBuilder, unitCtx, unit, nil, nil, ctx)
	unit.Nodes = append(unit.Nodes, nodeImports...)

	// Create a new Interface node.
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
		Abstract:                false,
		NodeType:                ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:                    ast_pb.NodeType_KIND_LIBRARY,
		Nodes:                   make([]Node[NodeType], 0),
		BaseContracts:           make([]*BaseContract, 0),
		ContractDependencies:    make([]int64, 0),
		LinearizedBaseContracts: make([]int64, 0),
		FullyImplemented:        true,
	}

	// Resolve and add base contracts.
	interfaceNode.BaseContracts = append(interfaceNode.BaseContracts,
		parseInheritanceFromCtx(l.ASTBuilder, unit, interfaceNode, ctx.InheritanceSpecifierList())...)
	unit.BaseContracts = interfaceNode.BaseContracts

	// Add linearized base contracts and contract dependencies.
	interfaceNode.LinearizedBaseContracts = append(interfaceNode.LinearizedBaseContracts, interfaceNode.GetId())
	for _, nodeImport := range nodeImports {
		interfaceNode.LinearizedBaseContracts = append(interfaceNode.LinearizedBaseContracts, nodeImport.GetId())
		interfaceNode.ContractDependencies = append(interfaceNode.ContractDependencies, nodeImport.GetId())
	}

	// Parse contract body elements.
	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			interfaceNode.FullyImplemented = false
			continue
		}

		bodyNode := NewBodyNode(l.ASTBuilder)
		childNode := bodyNode.ParseDefinitions(unit, interfaceNode, bodyElement)
		if childNode != nil {
			interfaceNode.Nodes = append(interfaceNode.Nodes, childNode)
			// Check if the body node is a function definition and if it's not implemented, mark the interface as not fully implemented.
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
