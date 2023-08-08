package abi

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

type Builder struct {
	ctx        context.Context // Context for the builder operations.
	sources    solgo.Sources   // Source files to be processed.
	parser     *ir.Builder     // Parser for the source code.
	astBuilder *ast.ASTBuilder // AST Builder for generating AST from parsed source.
	root       *Root           // Root of the generated ABI.
	resolver   *TypeResolver   // Type resolver for the ABI.
}

// NewBuilderFromSources creates a new ABI builder from given sources. It initializes
// the necessary IR builder from the provided sources.
func NewBuilderFromSources(ctx context.Context, sources solgo.Sources) (*Builder, error) {
	parser, err := ir.NewBuilderFromSources(context.TODO(), sources)
	if err != nil {
		return nil, err
	}

	return &Builder{
		ctx:        ctx,
		sources:    sources,
		parser:     parser,
		astBuilder: parser.GetAstBuilder(),
		resolver: &TypeResolver{
			parser: parser,
		},
	}, nil
}

// GetSources returns the source files being processed.
func (b *Builder) GetSources() solgo.Sources {
	return b.sources
}

// GetAstBuilder returns the AST builder.
func (b *Builder) GetAstBuilder() *ast.ASTBuilder {
	return b.astBuilder
}

// GetTypeResolver returns the type resolver for the ABI.
func (b *Builder) GetTypeResolver() *TypeResolver {
	return b.resolver
}

// GetParser returns the underlying intermediate representation parser.
func (b *Builder) GetParser() *ir.Builder {
	return b.parser
}

// GetRoot retrieves the root of the ABI.
func (b *Builder) GetRoot() *Root {
	return b.root
}

// ToJSON returns the JSON representation of the ABI.
func (b *Builder) ToJSON(d any) ([]byte, error) {
	if d != nil {
		return json.Marshal(d)
	}

	return json.Marshal(b.root)
}

// ToProto converts the ABI to its protocol buffer representation.
func (b *Builder) ToProto() *ir_pb.Root {
	return b.root.ToProto()
}

// ToJSONPretty provides a prettified JSON representation of the ABI.
func (b *Builder) ToJSONPretty() ([]byte, error) {
	return json.MarshalIndent(b.GetRoot(), "", "\t")
}

// ToProtoPretty provides a prettified JSON representation of the protocol buffer version of the ABIs.
func (b *Builder) ToProtoPretty() ([]byte, error) {
	return json.MarshalIndent(b.ToProto(), "", "\t")
}

// ToABI converts the ABI object into an ethereum/go-ethereum ABI object.
func (p *Builder) ToABI(contract *Contract) (*abi.ABI, error) {
	jsonData, err := p.ToJSON(contract)
	if err != nil {
		return nil, err
	}

	toReturn, err := abi.JSON(bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	return &toReturn, nil
}

func (b *Builder) Parse() (errs []error) {
	if syntaxErrs := b.GetParser().Parse(); syntaxErrs != nil {
		errs = append(errs, syntaxErrs...)
	}

	if err := b.GetParser().Build(); err != nil {
		errs = append(errs, err)
	}

	return errs
}

// Build constructs the ABIs from the sources.
func (b *Builder) Build() error {
	if root := b.GetParser().GetRoot(); root != nil {
		b.root = b.processRoot(root)
	}
	return nil
}
