package ast

import (
	"encoding/json"
	"io/ioutil"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/parser"
)

// ASTBuilder is a structure that helps in building and manipulating an Abstract Syntax Tree (AST).
// It contains a reference to a Solidity parser, the source code of the Solidity files, and various nodes of the AST.
type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	resolver              *Resolver
	tree                  *Tree
	sources               solgo.Sources          // sources is the source code of the Solidity files.
	parser                *parser.SolidityParser // parser is the Solidity parser instance.
	nextID                int64                  // nextID is the next ID to assign to a node.
	comments              []*CommentNode
	commentsParsed        bool
	entrySourceUnit       *SourceUnit[Node[ast_pb.SourceUnit]]
	sourceUnits           []*SourceUnit[Node[ast_pb.SourceUnit]]
	currentStateVariables []*StateVariableDeclaration
	currentEvents         []Node[NodeType]
	currentEnums          []Node[NodeType]
	currentStructs        []Node[NodeType]
	currentErrors         []Node[NodeType]
	currentModifiers      []Node[NodeType]
	currentFunctions      []Node[NodeType]
	currentVariables      []Node[NodeType]
}

// NewAstBuilder creates a new ASTBuilder with the provided Solidity parser and source code.
func NewAstBuilder(parser *parser.SolidityParser, sources solgo.Sources) *ASTBuilder {
	builder := &ASTBuilder{
		parser:      parser,
		sources:     sources,
		comments:    make([]*CommentNode, 0),
		sourceUnits: make([]*SourceUnit[Node[ast_pb.SourceUnit]], 0),
		nextID:      1,
	}

	// Used for resolving references.
	builder.resolver = NewResolver(builder)

	// Used for traversing the AST.
	builder.tree = NewTree(builder)

	return builder
}

// GetResolver returns the Resolver of the ASTBuilder.
func (b *ASTBuilder) GetResolver() *Resolver {
	return b.resolver
}

// GetRoot returns the root node of the AST from the Tree of the ASTBuilder.
func (b *ASTBuilder) GetRoot() *RootNode {
	return b.tree.GetRoot()
}

// ToJSON converts the root node of the AST to a JSON byte array.
func (b *ASTBuilder) ToJSON() ([]byte, error) {
	return json.Marshal(b.tree.GetRoot())
}

// ToPrettyJSON converts the provided data to a pretty (indented) JSON byte array.
func (b *ASTBuilder) ToPrettyJSON(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

// WriteJSONToFile writes the root node of the AST as a JSON byte array to a file at the provided path.
func (b *ASTBuilder) WriteJSONToFile(path string) error {
	bts, err := b.ToJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bts, 0644)
}

// WriteToFile writes the provided data byte array to a file at the provided path.
func (b *ASTBuilder) WriteToFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

// ResolveReferences resolves the references in the AST using the Resolver of the ASTBuilder.
func (b *ASTBuilder) ResolveReferences() error {
	if err := b.resolver.Resolve(); err != nil {
		return err
	}
	return nil
}
