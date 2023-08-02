package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

// PragmaNode represents a pragma directive in a Solidity source file.
// A pragma directive provides instructions to the compiler about how to treat the source code (e.g., compiler version).
type PragmaNode struct {
	// Id is the unique identifier of the pragma directive.
	Id int64 `json:"id"`
	// NodeType is the type of the node.
	// For a PragmaNode, this is always NodeType_PRAGMA_DIRECTIVE.
	NodeType ast_pb.NodeType `json:"node_type"`
	// SrcNode contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// Literals is a slice of strings that represent the literals of the pragma directive.
	// For example, for the pragma directive "pragma solidity ^0.5.0;", the literals would
	// be ["solidity", "^", "0", ".", "5", ".", "0"].
	Literals []string `json:"literals"`
}

// SetReferenceDescriptor sets the reference descriptions of the PragmaNode node.
func (p PragmaNode) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the pragma directive.
func (p PragmaNode) GetId() int64 {
	return p.Id
}

// GetType returns the type of the node. For a PragmaNode, this is always NodeType_PRAGMA_DIRECTIVE.
func (p PragmaNode) GetType() ast_pb.NodeType {
	return p.NodeType
}

// GetSrc returns the source information about the node, such as its line and column numbers in the source file.
func (p PragmaNode) GetSrc() SrcNode {
	return p.Src
}

// GetTypeDescription returns the type description of the node. For a PragmaNode, this is always nil.
func (p PragmaNode) GetTypeDescription() *TypeDescription {
	return nil
}

// GetLiterals returns a slice of strings that represent the literals of the pragma directive.
func (p PragmaNode) GetLiterals() []string {
	return p.Literals
}

// GetNodes returns the child nodes of the node. For a PragmaNode, this is always nil.
func (p PragmaNode) GetNodes() []Node[NodeType] {
	return nil
}

// ToProto returns the protobuf representation of the node.
func (p PragmaNode) ToProto() NodeType {
	proto := ast_pb.Pragma{
		Id:       p.Id,
		NodeType: p.NodeType,
		Src:      p.Src.ToProto(),
		Literals: p.Literals,
	}

	// Marshal the Pragma into JSON
	jsonBytes, err := protojson.Marshal(&proto)
	if err != nil {
		panic(err)
	}

	s := &structpb.Struct{}
	if err := protojson.Unmarshal(jsonBytes, s); err != nil {
		panic(err)
	}

	return &v3.TypedStruct{
		TypeUrl: "github.com/txpull/protos/txpull.v1.ast.Pragma",
		Value:   s,
	}
}

// CreatePragmaNodeFromCtx creates a new PragmaNode from the provided pragma context.
// It sets the ID of the new node to the next available ID from the provided ASTBuilder,
// and sets the source information of the node based on the provided pragma context.
// The NodeType of the new node is set to NodeType_PRAGMA_DIRECTIVE, and the literals
// of the node are set to the literals of the pragma context.
//
// The function takes the following parameters:
//   - b: The ASTBuilder from which to get the next available ID.
//   - unit: The SourceUnit to which the new node will belong. The ID of the unit is set
//     as the ParentIndex of the new node.
//   - pragmaCtx: The pragma context from which to create the new node. The source information
//     and literals of the new node are set based on this context.
//
// The function returns a pointer to the newly created PragmaNode.
func CreatePragmaNodeFromCtx(b *ASTBuilder, unit *SourceUnit[Node[ast_pb.SourceUnit]], pragmaCtx *parser.PragmaDirectiveContext) *PragmaNode {
	return &PragmaNode{
		Id: b.GetNextID(),
		Src: SrcNode{
			Id:          b.GetNextID(),
			Line:        int64(pragmaCtx.GetStart().GetLine()),
			Column:      int64(pragmaCtx.GetStart().GetColumn()),
			Start:       int64(pragmaCtx.GetStart().GetStart()),
			End:         int64(pragmaCtx.GetStop().GetStop()),
			Length:      int64(pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1),
			ParentIndex: unit.Id,
		},
		NodeType: ast_pb.NodeType_PRAGMA_DIRECTIVE,
		Literals: getLiterals(pragmaCtx.GetText()),
	}
}

// EnterPragmaDirective is called when production pragmaDirective is entered.
// However, it won't return pragma directives properly. For example, if we have
// experimental pragma, it won't return it. It will return only the pragma.
// Because of it, we are parsing pragmas in EnterSourceUnit to be able capture all of the
// pragmas and assign them based on the contract they belong to. Source file can have multiple
// contracts and multiple files and therefore we need to be able to assign pragmas to the
// correct contract.
// @WARN: DO NOT USE THIS METHOD.
func (b *ASTBuilder) EnterPragmaDirective(ctx *parser.PragmaDirectiveContext) {}

// FindPragmasForSourceUnit traverses the children of the provided source unit until
// it finds the library, contract, or interface definition. It collects all pragma
// directives encountered along the way and returns them as a slice of Node.
//
// The function takes the following parameters:
//   - sourceUnit: The source unit context to traverse.
//   - unit: The SourceUnit instance to which the pragmas belong.
//   - libraryCtx: The library definition context. If provided, the function stops traversing
//     the source unit once it encounters this context.
//   - contractCtx: The contract definition context. If provided, the function stops traversing
//     the source unit once it encounters this context.
//   - interfaceCtx: The interface definition context. If provided, the function stops traversing
//     the source unit once it encounters this context.
//
// The function returns a slice of Node, each representing a pragma directive. The pragmas
// are filtered such that only those that are within 10-20 lines of the contract definition
// are kept. The returned pragmas are ordered by their line number in ascending order.
func parsePragmasForSourceUnit(
	b *ASTBuilder,
	sourceUnit *parser.SourceUnitContext,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	libraryCtx *parser.LibraryDefinitionContext,
	contractCtx *parser.ContractDefinitionContext,
	interfaceCtx *parser.InterfaceDefinitionContext) []Node[NodeType] {
	pragmas := make([]*PragmaNode, 0)

	contractLine := func() int64 {
		switch {
		case libraryCtx != nil:
			return int64(libraryCtx.GetStart().GetLine())
		case contractCtx != nil:
			return int64(contractCtx.GetStart().GetLine())
		case interfaceCtx != nil:
			return int64(interfaceCtx.GetStart().GetLine())
		default:
			return 0
		}
	}()

	prevLine := int64(-1)

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if libraryCtx != nil && child == libraryCtx {
			// Found the library definition, stop traversing
			break
		}

		if contractCtx != nil && child == contractCtx {
			// Found the contract definition, stop traversing
			break
		}

		if interfaceCtx != nil && child == interfaceCtx {
			// Found the interface definition, stop traversing
			break
		}

		if pragmaCtx, ok := child.(*parser.PragmaDirectiveContext); ok {
			pragmaLine := int64(pragmaCtx.GetStart().GetLine())

			// First pragma encountered, add it to the result
			if prevLine == -1 {
				pragma := CreatePragmaNodeFromCtx(b, unit, pragmaCtx)
				pragmas = append(pragmas, pragma)
				prevLine = int64(pragmaLine)

				continue
			}

			// Add the pragma to the result
			pragmas = append(pragmas, CreatePragmaNodeFromCtx(b, unit, pragmaCtx))

			// Update the previous line number
			prevLine = pragmaLine
		}
	}

	// Post pragma cleanup...
	// Remove pragmas that have large gaps between the lines, keep only higher lines
	filteredPragmas := make([]Node[NodeType], 0)
	maxLine := int64(-1)

	// Iterate through the collected pragmas in reverse order and ensure only
	// pragmas that are within 10-20 lines of the contract definition are kept
	for i := len(pragmas) - 1; i >= 0; i-- {
		pragma := pragmas[i]
		if maxLine == -1 || (int64(contractLine)-pragma.Src.Line <= 10 && pragma.Src.Line-maxLine >= -1) {
			pragma.Src.ParentIndex = unit.Id
			filteredPragmas = append([]Node[NodeType]{Node[NodeType](pragma)}, filteredPragmas...)
			maxLine = pragma.Src.Line
		}
	}

	return filteredPragmas
}
