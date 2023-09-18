package ir

import (
	"context"
	"encoding/json"
	"errors"

	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/standards"
)

// Builder facilitates the creation of the IR from source code using solgo and AST tools.
type Builder struct {
	ctx        context.Context // Context for the builder operations.
	sources    *solgo.Sources  // Source files to be processed.
	parser     *solgo.Parser   // Parser for the source code.
	astBuilder *ast.ASTBuilder // AST Builder for generating AST from parsed source.
	root       *RootSourceUnit // Root of the generated IR.
}

// NewBuilderFromSources creates a new IR builder from given sources. It initializes
// the necessary parser and AST builder from the provided sources.
func NewBuilderFromSources(ctx context.Context, sources *solgo.Sources) (*Builder, error) {
	if sources == nil {
		return nil, errors.New("sources needed to initialize ir builder")
	}

	if !standards.StandardsLoaded() {
		if err := standards.LoadStandards(); err != nil {
			return nil, err
		}
	}

	parser, err := solgo.NewParserFromSources(ctx, sources)
	if err != nil {
		return nil, err
	}

	astBuilder := ast.NewAstBuilder(parser.GetParser(), parser.GetSources())

	if err := parser.RegisterListener(solgo.ListenerAst, astBuilder); err != nil {
		return nil, err
	}

	return &Builder{
		ctx:        ctx,
		sources:    sources,
		parser:     parser,
		astBuilder: astBuilder,
	}, nil
}

// GetParser returns the underlying solgo parser.
func (b *Builder) GetParser() *solgo.Parser {
	return b.parser
}

// GetAstBuilder returns the AST builder.
func (b *Builder) GetAstBuilder() *ast.ASTBuilder {
	return b.astBuilder
}

// GetSources returns the source files being processed.
func (b *Builder) GetSources() *solgo.Sources {
	return b.sources
}

// Parse processes the sources using the parser and the AST builder and returns
// any encountered errors.
func (b *Builder) Parse() (errs []error) {
	if syntaxErrs := b.parser.Parse(); syntaxErrs != nil {
		for _, syntaxErr := range syntaxErrs {
			errs = append(errs, syntaxErr.Error())
		}
	}

	if err := b.astBuilder.ResolveReferences(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

// GetRoot retrieves the root of the IR.
func (b *Builder) GetRoot() *RootSourceUnit {
	return b.root
}

// ToJSON returns the JSON representation of the IR.
func (b *Builder) ToJSON() ([]byte, error) {
	return json.Marshal(b.root)
}

// ToProto converts the IR to its protocol buffer representation.
func (b *Builder) ToProto() *ir_pb.Root {
	return b.root.ToProto()
}

// ToJSONPretty provides a prettified JSON representation of the IR.
func (b *Builder) ToJSONPretty() ([]byte, error) {
	return json.MarshalIndent(b.root, "", "\t")
}

// ToProtoPretty provides a prettified JSON representation of the protocol buffer version of the IR.
func (b *Builder) ToProtoPretty() ([]byte, error) {
	return json.MarshalIndent(b.root.ToProto(), "", "\t")
}

// Build constructs the IR from the sources.
func (b *Builder) Build() error {
	if root := b.GetAstBuilder().GetRoot(); root != nil {
		b.root = b.processRoot(root)
	}
	return nil
}
