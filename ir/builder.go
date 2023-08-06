package ir

import (
	"context"
	"encoding/json"

	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/ast"
)

type Builder struct {
	ctx        context.Context
	sources    solgo.Sources
	parser     *solgo.Parser
	astBuilder *ast.ASTBuilder
	root       *RootSourceUnit
}

func NewBuilder(ctx context.Context, parser *solgo.Parser, astBuilder *ast.ASTBuilder) *Builder {
	toReturn := &Builder{
		ctx:        ctx,
		parser:     parser,
		sources:    parser.GetSources(),
		astBuilder: astBuilder,
	}

	return toReturn
}

func NewBuilderFromSources(ctx context.Context, sources solgo.Sources) (*Builder, error) {
	parser, err := solgo.NewParserFromSources(ctx, sources)
	if err != nil {
		return nil, err
	}

	astBuilder := ast.NewAstBuilder(parser.GetParser(), parser.GetSources())

	if err := parser.RegisterListener(solgo.ListenerAst, astBuilder); err != nil {
		return nil, err
	}

	toReturn := &Builder{
		ctx:        ctx,
		sources:    sources,
		parser:     parser,
		astBuilder: astBuilder,
	}

	return toReturn, nil
}

// GetParser returns the main parser.
func (b *Builder) GetParser() *solgo.Parser {
	return b.parser
}

// GetAstBuilder returns the AST builder.
func (b *Builder) GetAstBuilder() *ast.ASTBuilder {
	return b.astBuilder
}

// GetSources returns the sources used to build the IR.
func (b *Builder) GetSources() solgo.Sources {
	return b.sources
}

// Parse parses the sources and returns a list of errors encountered during parsing.
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

// GetRoot returns the root source unit.
func (b *Builder) GetRoot() *RootSourceUnit {
	return b.root
}

// ToJSON returns the JSON representation of the IR.
func (b *Builder) ToJSON() ([]byte, error) {
	return json.Marshal(b.root)
}

func (b *Builder) ToProto() *ir_pb.Root {
	return b.root.ToProto()
}

// ToJSONPretty returns the pretty JSON representation of the IR.
func (b *Builder) ToJSONPretty() ([]byte, error) {
	return json.MarshalIndent(b.root, "", "\t")
}

// ToJSONPretty returns the pretty JSON representation of the IR.
func (b *Builder) ToProtoPretty() ([]byte, error) {
	return json.MarshalIndent(b.root.ToProto(), "", "\t")
}

// Build returns the IR built from the sources.
func (b *Builder) Build() error {
	if root := b.GetAstBuilder().GetRoot(); root != nil {
		b.root = b.processRoot(root)
	}
	return nil
}
