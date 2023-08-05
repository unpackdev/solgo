package ast

import (
	"reflect"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// Library represents a library node in the abstract syntax tree.
// It includes various attributes like id, name, type, source node, abstract status, kind, implementation status, nodes, base contracts, contract dependencies and scope.
type Library struct {
	*ASTBuilder

	Id                      int64            `json:"id"`                        // Id is the unique identifier of the library node.
	Name                    string           `json:"name"`                      // Name is the name of the library.
	NodeType                ast_pb.NodeType  `json:"node_type"`                 // NodeType is the type of the node.
	Src                     SrcNode          `json:"src"`                       // Src is the source node associated with the library node.
	Abstract                bool             `json:"abstract"`                  // Abstract indicates if the library is abstract.
	Kind                    ast_pb.NodeType  `json:"kind"`                      // Kind is the kind of the node.
	FullyImplemented        bool             `json:"fully_implemented"`         // FullyImplemented indicates if the library is fully implemented.
	Nodes                   []Node[NodeType] `json:"nodes"`                     // Nodes are the nodes associated with the library.
	LinearizedBaseContracts []int64          `json:"linearized_base_contracts"` // LinearizedBaseContracts are the linearized base contracts of the library.
	BaseContracts           []*BaseContract  `json:"base_contracts"`            // BaseContracts are the base contracts of the library.
	ContractDependencies    []int64          `json:"contract_dependencies"`     // ContractDependencies are the contract dependencies of the library.
	Scope                   int64            `json:"scope"`                     // Scope is the scope of the library.
}

// NewLibraryDefinition creates a new Library with the provided ASTBuilder.
// It returns a pointer to the created Library.
func NewLibraryDefinition(b *ASTBuilder) *Library {
	return &Library{
		ASTBuilder: b,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Library node.
func (l *Library) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the library node.
func (l *Library) GetId() int64 {
	return l.Id
}

// GetType returns the type of the library node.
func (l *Library) GetType() ast_pb.NodeType {
	return l.NodeType
}

// GetSrc returns the source node associated with the library node.
func (l *Library) GetSrc() SrcNode {
	return l.Src
}

// GetTypeDescription returns the type description of the library node.
// Currently, it returns nil and needs to be implemented.
func (l *Library) GetTypeDescription() *TypeDescription {
	return nil
}

// GetName returns the name of the library.
func (l *Library) GetName() string {
	return l.Name
}

// IsAbstract returns a boolean indicating whether the library is abstract.
func (l *Library) IsAbstract() bool {
	return l.Abstract
}

// GetKind returns the kind of the library node.
func (l *Library) GetKind() ast_pb.NodeType {
	return l.Kind
}

// IsFullyImplemented returns a boolean indicating whether the library is fully implemented.
func (l *Library) IsFullyImplemented() bool {
	return l.FullyImplemented
}

// GetNodes returns the nodes associated with the library.
func (l *Library) GetNodes() []Node[NodeType] {
	return l.Nodes
}

// GetScope returns the scope of the library.
func (l *Library) GetScope() int64 {
	return l.Scope
}

// GetLinearizedBaseContracts returns the linearized base contracts of the library.
func (l *Library) GetLinearizedBaseContracts() []int64 {
	return l.LinearizedBaseContracts
}

// GetBaseContracts returns the base contracts of the library.
func (l *Library) GetBaseContracts() []*BaseContract {
	return l.BaseContracts
}

// GetContractDependencies returns the contract dependencies of the library.
func (l *Library) GetContractDependencies() []int64 {
	return l.ContractDependencies
}

func (l *Library) GetStateVariables() []*StateVariableDeclaration {
	toReturn := make([]*StateVariableDeclaration, 0)

	for _, node := range l.GetNodes() {
		if stateVariable, ok := node.(*StateVariableDeclaration); ok {
			toReturn = append(toReturn, stateVariable)
		}
	}

	return toReturn
}

func (l *Library) GetConstructor() *Constructor {
	for _, node := range l.GetNodes() {
		if constructor, ok := node.(*Constructor); ok {
			return constructor
		}
	}

	return nil
}

func (l *Library) GetFunctions() []*Function {
	toReturn := make([]*Function, 0)

	for _, node := range l.GetNodes() {
		if function, ok := node.(*Function); ok {
			toReturn = append(toReturn, function)
		}
	}

	return toReturn
}

// ToProto converts the Library to a protocol buffer representation.
// Currently, it returns an empty Contract and needs to be implemented.
func (l *Library) ToProto() NodeType {
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

// Parse parses the source unit context and library definition context to populate the library node.
// It takes a SourceUnitContext, a LibraryDefinitionContext, a RootNode and a SourceUnit as arguments.
// It does not return anything.
func (l *Library) Parse(unitCtx *parser.SourceUnitContext, ctx *parser.LibraryDefinitionContext, rootNode *RootNode, unit *SourceUnit[Node[ast_pb.SourceUnit]]) {
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
	libraryNode := &Library{
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
		childNode := bodyNode.ParseDefinitions(unit, libraryNode, bodyElement)
		if childNode != nil {
			libraryNode.Nodes = append(libraryNode.Nodes, childNode)
			if bodyNode.NodeType == ast_pb.NodeType_FUNCTION_DEFINITION && !bodyNode.Implemented {
				libraryNode.FullyImplemented = false
			}
		} else {
			libraryNode.FullyImplemented = false
			zap.L().Warn(
				"Discovered empty body node. Checkout why this is happening.",
				zap.String("contract", libraryNode.Name),
				zap.String("contract_kind", libraryNode.Kind.String()),
				zap.String("body_element_type", reflect.TypeOf(bodyElement).String()),
			)
		}
	}

	unit.Nodes = append(unit.Nodes, libraryNode)
	unit.Contract = libraryNode
}
