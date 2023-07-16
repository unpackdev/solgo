package ast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	parser            *parser.SolidityParser // parser is the Solidity parser instance.
	nextID            int64                  // nextID is the next ID to assign to a node.
	comments          []*ast_pb.Comment
	commentsParsed    bool
	currentSourceUnit *ast_pb.SourceUnit
	astRoot           *ast_pb.RootSourceUnit
}

func NewAstBuilder(parser *parser.SolidityParser) *ASTBuilder {
	return &ASTBuilder{
		parser:   parser,
		comments: make([]*ast_pb.Comment, 0),
	}
}

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	id := atomic.AddInt64(&b.nextID, 1) - 1

	b.currentSourceUnit = &ast_pb.SourceUnit{
		Id:              id,
		AbsolutePath:    ctx.GetStart().GetInputStream().GetSourceName(),
		ExportedSymbols: make([]*ast_pb.ExportedSymbols, 0),
		NodeType:        ast_pb.NodeType_NODE_TYPE_SOURCE_UNIT,
		Nodes:           &ast_pb.RootNode{},
		Src: &ast_pb.Src{
			Line:   int64(ctx.GetStart().GetLine()),
			Column: int64(ctx.GetStart().GetColumn()),
			Start:  int64(ctx.GetStart().GetStart()),
			// @TODO: GetStop() is always nil due to some reason so we cannot get lenght
			// just yet. We need to figure out why.
			//Length: ctx.GetStop() - ctx.GetStart() + 1,
			ParentIndex: int64(ctx.GetStart().GetTokenIndex()),
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
			b.currentSourceUnit.Nodes.Nodes = append(
				b.currentSourceUnit.Nodes.Nodes,
				pragmas...,
			)
		}
	}
}

// ExitSourceUnit is called when production sourceUnit is exited.
func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &ast_pb.RootSourceUnit{
		SourceUnits: []*ast_pb.SourceUnit{b.currentSourceUnit},
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
		&ast_pb.ExportedSymbols{
			Id:   id,
			Name: identifierName,
		},
	)

	identifierNode := &ast_pb.Node{
		Id:   id,
		Name: identifierName,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: b.currentSourceUnit.Src.ParentIndex,
		},
		Abstract: false,
		NodeType: ast_pb.NodeType_NODE_TYPE_CONTRACT_DEFINITION,
		Kind:     ast_pb.NodeType_NODE_TYPE_KIND_LIBRARY,
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
	identifierNode.Nodes = make([]*ast_pb.Node, 0)

	for _, bodyElement := range ctx.AllContractBodyElement() {
		if bodyElement.IsEmpty() {
			continue
		}

		bodyNode := b.traverseBodyElement(identifierNode, bodyElement)
		identifierNode.Nodes = append(identifierNode.Nodes, bodyNode)
	}

	b.currentSourceUnit.Nodes.Nodes = append(b.currentSourceUnit.Nodes.Nodes, identifierNode)

}

func (b *ASTBuilder) traverseBodyElement(identifierNode *ast_pb.Node, bodyElement parser.IContractBodyElementContext) *ast_pb.Node {
	id := atomic.AddInt64(&b.nextID, 1) - 1
	toReturn := &ast_pb.Node{
		Id: id,
		Src: &ast_pb.Src{
			Line:        int64(bodyElement.GetStart().GetLine()),
			Start:       int64(bodyElement.GetStart().GetStart()),
			End:         int64(bodyElement.GetStop().GetStop()),
			Length:      int64(bodyElement.GetStop().GetStop() - bodyElement.GetStart().GetStart() + 1),
			ParentIndex: identifierNode.Src.ParentIndex,
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

func (b *ASTBuilder) traverseFunctionDefinition(node *ast_pb.Node, fd *parser.FunctionDefinitionContext) *ast_pb.Node {
	// Extract the function name.
	node.Name = fd.Identifier().GetText()

	// Set the function type and its kind.
	node.NodeType = ast_pb.NodeType_NODE_TYPE_FUNCTION_DEFINITION
	node.Kind = ast_pb.NodeType_NODE_TYPE_FUNCTION_DEFINITION

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
		_ = modifier
		//node.Modifiers = append(node.Modifiers, modifier.GetText())
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
	//if len(fd.AllParameterList()) > 0 {
	//	node.Parameters = b.traverseParameterList(node, fd.AllParameterList()[0])
	//}

	// Extract function return parameters.
	//node.ReturnParameters = b.traverseParameterList(node, fd.GetReturnParameters())

	// And now we are going to the big league. We are going to traverse the function body.
	if !fd.Block().IsEmpty() {
		node.Nodes = make([]*ast_pb.Node, 0)
		for _, statement := range fd.Block().AllStatement() {
			if statement.IsEmpty() {
				continue
			}
			node.Nodes = append(node.Nodes, b.traverseStatement(node, statement))
		}

	}

	return node
}

func (b *ASTBuilder) traverseStatement(node *ast_pb.Node, statement parser.IStatementContext) *ast_pb.Node {
	toReturn := &ast_pb.Node{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(statement.GetStart().GetLine()),
			Column:      int64(statement.GetStart().GetColumn()),
			Start:       int64(statement.GetStart().GetStart()),
			End:         int64(statement.GetStop().GetStop()),
			Length:      int64(statement.GetStop().GetStop() - statement.GetStart().GetStart() + 1),
			ParentIndex: node.Src.ParentIndex,
		},
	}

	if block := statement.Block(); block != nil {
		toReturn = b.traverseBlock(toReturn, block.(*parser.BlockContext))
	}

	return toReturn
}

func (b *ASTBuilder) traverseBlock(node *ast_pb.Node, block *parser.BlockContext) *ast_pb.Node {
	node.NodeType = ast_pb.NodeType_NODE_TYPE_BLOCK
	node.Nodes = make([]*ast_pb.Node, 0)

	for _, statement := range block.AllStatement() {
		if statement.IsEmpty() {
			continue
		}
		node.Nodes = append(node.Nodes, b.traverseStatement(node, statement))
	}

	return node
}

func (b *ASTBuilder) GetRoot() *ast_pb.RootSourceUnit {
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

func (b *ASTBuilder) WriteJSONToFile(path string) error {
	bts, err := b.ToJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bts, 0644)
}
