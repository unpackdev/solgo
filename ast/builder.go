package ast

import (
	"encoding/json"
	"os"

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
}

// NewAstBuilder creates a new ASTBuilder with the provided Solidity parser and source code.
func NewAstBuilder(parser *parser.SolidityParser, sources solgo.Sources) *ASTBuilder {
	builder := &ASTBuilder{
		parser:      parser,
		sources:     sources,
		comments:    make([]*Comment, 0),
		sourceUnits: make([]*SourceUnit[Node[ast_pb.SourceUnit]], 0),
		nextID:      1,
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
	return os.WriteFile(path, bts, 0600)
}

// WriteToFile writes the provided data byte array to a file at the provided path.
func (b *ASTBuilder) WriteToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
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
}
