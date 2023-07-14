package ast

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/txpull/solgo/parser"
)

type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	parser            *parser.SolidityParser // parser is the Solidity parser instance.
	nextID            int64                  // nextID is the next ID to assign to a node.
	comments          []*CommentNode
	commentsParsed    bool
	currentSourceUnit *SourceUnit
	astRoot           *RootSourceUnit
}

func NewAstBuilder(parser *parser.SolidityParser) *ASTBuilder {
	return &ASTBuilder{
		parser:   parser,
		comments: make([]*CommentNode, 0),
	}
}

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	id := atomic.AddInt64(&b.nextID, 1) - 1

	b.currentSourceUnit = &SourceUnit{
		ID:              id,
		AbsolutePath:    ctx.GetStart().GetInputStream().GetSourceName(),
		ExportedSymbols: make([]ExportedSymbols, 0),
		NodeType:        NodeTypeSourceUnit.String(),
		Nodes:           []Node{},
		Src: Src{
			Line:  ctx.GetStart().GetLine(),
			Start: ctx.GetStart().GetStart(),
			// @TODO: GetStop() is always nil due to some reason so we cannot get lenght
			// just yet. We need to figure out why.
			//Length: ctx.GetStop() - ctx.GetStart() + 1,
			Index: int64(ctx.GetStart().GetTokenIndex()),
		},
		Comments: b.comments,
	}

	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			_ = contractCtx
		}

		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			b.currentSourceUnit.License = b.GetLicense()

			// Alright lets extract bloody pragmas...
			pragmas := b.findPragmasForLibrary(ctx, libraryCtx)
			b.currentSourceUnit.Nodes = append(
				b.currentSourceUnit.Nodes,
				pragmas...,
			)
		}
	}
}

// ExitSourceUnit is called when production sourceUnit is exited.
func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &RootSourceUnit{
		SourceUnits: []SourceUnit{*b.currentSourceUnit},
	}
	b.currentSourceUnit = nil
}

func (b *ASTBuilder) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	//id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	fmt.Println("EnterContractDefinition AAA", identifierName)
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

func (b *ASTBuilder) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	if ctx.IsEmpty() {
		return
	}

	id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	b.currentSourceUnit.ExportedSymbols = append(
		b.currentSourceUnit.ExportedSymbols,
		ExportedSymbols{
			ID:   id,
			Name: identifierName,
		},
	)

	identifierNode := Node{
		ID:   id,
		Name: identifierName,
		Src: Src{
			Line:   ctx.GetStart().GetLine(),
			Start:  ctx.GetStart().GetStart(),
			End:    ctx.GetStop().GetStop(),
			Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
			Index:  b.currentSourceUnit.Src.Index,
		},
		Abstract: false,
		NodeType: NodeTypeContractDefinition,
		Kind:     NodeTypeKindLibrary.String(),
	}

	// Check if all of the functions discovered in the library are fully implemented...
	// @TODO: Implement this.
	identifierNode.FullyImplemented = false

	// Discover linearized base contracts...
	// The linearizedBaseContracts field contains an array of IDs that represent the
	// contracts in the inheritance hierarchy, starting from the most derived contract
	// (the contract itself) and ending with the most base contract.
	// The IDs correspond to the id fields of the ContractDefinition nodes in the AST.
	identifierNode.LinearizedBaseContracts = []int64{id}

	// Allright now the fun part begins. We need to traverse through the body of the library
	// and extract all of the nodes...

	// First lets define nodes...
	identifierNode.Nodes = make([]Node, 0)

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			continue
		}

		bodyNode := b.traverseBodyElement(identifierNode, bodyElement)
		identifierNode.Nodes = append(identifierNode.Nodes, bodyNode)
	}

	b.currentSourceUnit.Nodes = append(b.currentSourceUnit.Nodes, identifierNode)

}

func (b *ASTBuilder) traverseBodyElement(identifierNode Node, bodyElement parser.IContractBodyElementContext) Node {
	id := atomic.AddInt64(&b.nextID, 1) - 1
	toReturn := Node{
		ID: id,
		Src: Src{
			Line:   bodyElement.GetStart().GetLine(),
			Start:  bodyElement.GetStart().GetStart(),
			End:    bodyElement.GetStop().GetStop(),
			Length: bodyElement.GetStop().GetStop() - bodyElement.GetStart().GetStart() + 1,
			Index:  identifierNode.Src.Index,
		},
	}

	if functionDefinition := bodyElement.FunctionDefinition(); functionDefinition != nil {
		toReturn = b.traverseFunctionDefinition(
			toReturn,
			functionDefinition.(*parser.FunctionDefinitionContext),
		)
	}

	return toReturn
}

func (b *ASTBuilder) traverseFunctionDefinition(node Node, fd *parser.FunctionDefinitionContext) Node {
	// Extract the function name.
	node.Name = fd.Identifier().GetText()

	// Set the function type and its kind.
	node.NodeType = NodeTypeFunctionDefinition
	node.Kind = NodeTypeKindFunction.String()

	// If block is not empty we are going to assume that the function is implemented.
	// @TODO: Take assumption to the next level in the future.
	node.Implemented = !fd.Block().IsEmpty()

	// Get function visibility state.
	for _, visibility := range fd.AllVisibility() {
		node.Visibility = visibility.GetText()
	}

	// Get function state mutability.
	for _, stateMutability := range fd.AllStateMutability() {
		node.StateMutability = stateMutability.GetText()
	}

	// Get function modifiers.
	for _, modifier := range fd.AllModifierInvocation() {
		node.Modifiers = append(node.Modifiers, modifier.GetText())
	}

	// Check if function is virtual.
	for _, virtual := range fd.AllVirtual() {
		node.Virtual = virtual.GetText() == "virtual"
	}

	// Check if function is override.
	// @TODO: Implement override specification.
	for _, override := range fd.AllOverrideSpecifier() {
		_ = override
	}

	// Extract function parameters.
	if len(fd.AllParameterList()) > 0 {
		node.Parameters = b.traverseParameterList(node, fd.AllParameterList()[0])
	}

	// Extract function return parameters.
	node.ReturnParameters = b.traverseParameterList(node, fd.GetReturnParameters())

	return node
}

func (b *ASTBuilder) traverseParameterList(node Node, parameterCtx parser.IParameterListContext) *ParametersList {
	if parameterCtx.IsEmpty() {
		return nil
	}

	id := atomic.AddInt64(&b.nextID, 1) - 1

	parametersList := &ParametersList{
		ID:         id,
		Parameters: make([]Parameter, 0),
		Src: Src{
			Line:   parameterCtx.GetStart().GetLine(),
			Start:  parameterCtx.GetStart().GetStart(),
			End:    parameterCtx.GetStop().GetStop(),
			Length: parameterCtx.GetStop().GetStop() - parameterCtx.GetStart().GetStart() + 1,
			Index:  node.Src.Index,
		},
		NodeType: NodeTypeParameterList.String(),
	}

	for _, parameterCtx := range parameterCtx.AllParameterDeclaration() {
		if parameterCtx.IsEmpty() {
			continue
		}

		parameterID := atomic.AddInt64(&b.nextID, 1) - 1
		pNode := Parameter{
			ID: parameterID,
			Src: Src{
				Line:   parameterCtx.GetStart().GetLine(),
				Start:  parameterCtx.GetStart().GetStart(),
				End:    parameterCtx.GetStop().GetStop(),
				Length: parameterCtx.GetStop().GetStop() - parameterCtx.GetStart().GetStart() + 1,
				Index:  parametersList.ID,
			},
			Scope: node.ID,
			Name: func() string {
				if parameterCtx.Identifier() != nil {
					return parameterCtx.Identifier().GetText()
				}
				return ""
			}(),
			NodeType: NodeTypeVariableDeclaration.String(),
			// Just hardcoding it here to internal as I am not sure how
			// could it be possible to be anything else.
			// @TODO: Check if it is possible to be anything else.
			Visibility: NodeTypeVisibilityInternal.String(),
		}

		if parameterCtx.GetType_().ElementaryTypeName() != nil {
			typeNameID := atomic.AddInt64(&b.nextID, 1) - 1
			pNode.NodeType = NodeTypeElementaryTypeName.String()
			typeCtx := parameterCtx.GetType_().ElementaryTypeName()
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				typeCtx.GetText(),
			)
			pNode.TypeName = &TypeName{
				ID:   typeNameID,
				Name: typeCtx.GetText(),
				Src: Src{
					Line:   parameterCtx.GetStart().GetLine(),
					Start:  parameterCtx.GetStart().GetStart(),
					End:    parameterCtx.GetStop().GetStop(),
					Length: parameterCtx.GetStop().GetStop() - parameterCtx.GetStart().GetStart() + 1,
					Index:  pNode.ID,
				},
				NodeType: NodeTypeElementaryTypeName.String(),
				TypeDescriptions: &TypeDescriptions{
					TypeIdentifier: normalizedTypeIdentifier,
					TypeString:     normalizedTypeName,
				},
			}
		}

		parametersList.Parameters = append(parametersList.Parameters, pNode)
	}

	return parametersList
}

func (b *ASTBuilder) GetRoot() *RootSourceUnit {
	return b.astRoot
}

func (b *ASTBuilder) ToJSON() ([]byte, error) {
	return json.Marshal(b.astRoot)
}

func (b *ASTBuilder) ToJSONString() (string, error) {
	bts, err := b.ToJSON()
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

func (b *ASTBuilder) ToPrettyJSON() ([]byte, error) {
	return json.MarshalIndent(b.astRoot, "", "  ")
}
