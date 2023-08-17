package validation

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/txpull/solgo"
	"github.com/txpull/solgo/solc"
)

// Verifier is a utility that facilitates the verification of Ethereum smart contracts.
// It uses the solc compiler to compile the provided sources and then verifies the bytecode.
type Verifier struct {
	ctx      context.Context // The context for the verifier operations.
	compiler *solc.Compiler  // The solc compiler instance.
	config   *solc.Config    // The configuration for the solc compiler.
	sources  *solgo.Sources  // The sources of the Ethereum smart contracts to be verified.
}

// NewVerifier creates a new instance of Verifier.
// It initializes the solc compiler with the provided configuration and sources.
// If the sources are not prepared, it prepares them before initializing the compiler.
// Returns an error if there's any issue in preparing the sources or initializing the compiler.
func NewVerifier(ctx context.Context, config *solc.Config, sources *solgo.Sources) (*Verifier, error) {
	if config == nil {
		return nil, errors.New("config must be set")
	}

	if sources == nil {
		return nil, errors.New("sources must be set")
	}

	// Ensure that the sources are prepared for future consumption in case they are not already.
	if !sources.ArePrepared() {
		if err := sources.Prepare(); err != nil {
			return nil, err
		}
	}

	compiler, err := solc.NewCompiler(ctx, config, sources)
	if err != nil {
		return nil, err
	}

	return &Verifier{
		ctx:      ctx,
		compiler: compiler,
		sources:  sources,
		config:   config,
	}, nil
}

// GetContext returns the context associated with the verifier.
func (v *Verifier) GetContext() context.Context {
	return v.ctx
}

// GetSources returns the sources of the Ethereum smart contracts associated with the verifier.
func (v *Verifier) GetSources() *solgo.Sources {
	return v.sources
}

// GetCompiler returns the solc compiler instance associated with the verifier.
func (v *Verifier) GetCompiler() *solc.Compiler {
	return v.compiler
}

// Verify compiles the sources using the solc compiler and then verifies the bytecode.
// If the bytecode does not match the compiled result, it returns a diff of the two.
// Returns true if the bytecode matches, otherwise returns false.
// Also returns an error if there's any issue in the compilation or verification process.
func (v *Verifier) Verify(bytecode []byte) (bool, string, error) {
	results, err := v.compiler.Compile()
	if err != nil {
		return false, "", err
	}
	encoded := hex.EncodeToString(bytecode)
	if encoded != results.Bytecode {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(encoded, results.Bytecode, false)
		return false, dmp.DiffPrettyText(diffs), errors.New("bytecode missmatch, failed to verify")
	}

	return true, "", nil
}
