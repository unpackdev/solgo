package ast

import (
	"bytes"
	"context"
	"encoding/json"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/parser"
)

// ASTBuilder is a structure that helps in building and manipulating an Abstract Syntax Tree (AST).
// It contains a reference to a Solidity parser, the source code of the Solidity files, and various nodes of the AST.
type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	resolver              *Resolver
	tree                  *Tree
	sources               *solgo.Sources         // sources is the source code of the Solidity files.
	parser                *parser.SolidityParser // parser is the Solidity parser instance.
	nextID                int64                  // nextID is the next ID to assign to a node.
	comments              []*Comment
	commentsParsed        bool
	sourceUnits           []*SourceUnit[Node[ast_pb.SourceUnit]]
	currentStateVariables []*StateVariableDeclaration
	currentEvents         []Node[NodeType]
	currentEnums          []Node[NodeType]
	currentStructs        []Node[NodeType]
	currentErrors         []Node[NodeType]
	currentModifiers      []Node[NodeType]
	currentFunctions      []Node[NodeType]
	currentVariables      []Node[NodeType]
	globalDefinitions     []Node[NodeType]
	currentImports        []Node[NodeType]
}

// NewAstBuilder creates a new ASTBuilder with the provided Solidity parser and source code.
func NewAstBuilder(parser *parser.SolidityParser, sources *solgo.Sources) *ASTBuilder {
	builder := &ASTBuilder{
		parser:                parser,
		sources:               sources,
		comments:              make([]*Comment, 0),
		sourceUnits:           make([]*SourceUnit[Node[ast_pb.SourceUnit]], 0),
		currentImports:        make([]Node[NodeType], 0),
		currentStateVariables: make([]*StateVariableDeclaration, 0),
		currentEvents:         make([]Node[NodeType], 0),
		currentEnums:          make([]Node[NodeType], 0),
		currentStructs:        make([]Node[NodeType], 0),
		currentErrors:         make([]Node[NodeType], 0),
		currentModifiers:      make([]Node[NodeType], 0),
		currentFunctions:      make([]Node[NodeType], 0),
		currentVariables:      make([]Node[NodeType], 0),
		globalDefinitions:     make([]Node[NodeType], 0),
		nextID:                1,
	}

	// Used for resolving references.
	builder.resolver = NewResolver(builder)

	// Used for traversing the AST.
	builder.tree = NewTree(builder)

	return builder
}

// GetParser returns the Solidity parser of the ASTBuilder.
func (b *ASTBuilder) GetParser() *parser.SolidityParser {
	return b.parser
}

// GetResolver returns the Resolver of the ASTBuilder.
func (b *ASTBuilder) GetResolver() *Resolver {
	return b.resolver
}

// GetTree returns the Tree of the ASTBuilder.
func (b *ASTBuilder) GetTree() *Tree {
	return b.tree
}

// GetRoot returns the root node of the AST from the Tree of the ASTBuilder.
func (b *ASTBuilder) GetRoot() *RootNode {
	return b.tree.GetRoot()
}

func (b *ASTBuilder) ToProto() *ast_pb.RootSourceUnit {
	return b.tree.GetRoot().ToProto()
}

// ToJSON converts the root node of the AST to a JSON byte array.
func (b *ASTBuilder) ToJSON() ([]byte, error) {
	return b.InterfaceToJSON(b.tree.GetRoot())
}

// ToPrettyJSON converts the provided data to a JSON byte array.
func (b *ASTBuilder) InterfaceToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// ResolveReferences resolves the references in the AST using the Resolver of the ASTBuilder.
func (b *ASTBuilder) ResolveReferences() []error {
	if err := b.resolver.Resolve(); err != nil {
		return err
	}

	// Cleanup the builder so garbage collector and memory usage is minimized...
	b.GarbageCollect()

	return nil
}

// ImportFromJSON imports the AST from a JSON byte array.
// Note that parser content won't be imported. Only the results for future manipulation.
func (b *ASTBuilder) ImportFromJSON(ctx context.Context, jsonBytes []byte) (*RootNode, error) {
	var toReturn *RootNode
	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
	if err := decoder.Decode(&toReturn); err != nil {
		return nil, err
	}

	if b.tree == nil {
		b.tree = NewTree(b)
	}

	b.tree.SetRoot(toReturn)

	return toReturn, nil
}

// GarbageCollect cleans up the ASTBuilder after resolving references.
func (b *ASTBuilder) GarbageCollect() {
	b.currentEnums = nil
	b.currentErrors = nil
	b.currentEvents = nil
	b.currentFunctions = nil
	b.currentModifiers = nil
	b.currentStateVariables = nil
	b.currentStructs = nil
	b.currentVariables = nil
	b.globalDefinitions = nil
	b.currentImports = nil
}
