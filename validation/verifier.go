package validation

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/0x19/solc-switch"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/txpull/solgo"
)

// Verifier is a utility that facilitates the verification of Ethereum smart contracts.
// It uses the solc compiler to compile the provided sources and then verifies the bytecode.
type Verifier struct {
	ctx     context.Context // The context for the verifier operations.
	solc    *solc.Solc      // The solc compiler instance.
	config  *solc.Config    // The configuration for the solc compiler.
	sources *solgo.Sources  // The sources of the Ethereum smart contracts to be verified.
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

	solc, err := solc.New(ctx, config)
	if err != nil {
		return nil, err
	}

	// Do the releases synchronization in the background...
	if err := solc.Sync(); err != nil {
		return nil, err
	}

	return &Verifier{
		ctx:     ctx,
		solc:    solc,
		sources: sources,
		config:  config,
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
func (v *Verifier) GetCompiler() *solc.Solc {
	return v.solc
}

// Verify compiles the sources using the solc compiler and then verifies the bytecode.
// If the bytecode does not match the compiled result, it returns a diff of the two.
// Returns true if the bytecode matches, otherwise returns false.
// Also returns an error if there's any issue in the compilation or verification process.
func (v *Verifier) Verify(ctx context.Context, bytecode []byte, compilerConfig *solc.CompilerConfig) (*VerifyResult, error) {
	results, err := v.solc.Compile(ctx, v.sources.GetCombinedSource(), compilerConfig)
	if err != nil {
		return nil, err
	}

	encoded := hex.EncodeToString(bytecode)
	if encoded != results.Bytecode {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(encoded, results.Bytecode, false)

		toReturn := &VerifyResult{
			Verified:         false,
			CompilerResults:  results,
			ExpectedBytecode: encoded,
			Diffs:            diffs,
			DiffPretty:       dmp.DiffPrettyText(diffs),
		}

		return toReturn, errors.New("bytecode missmatch, failed to verify")
	}

	toReturn := &VerifyResult{
		Verified:         true,
		ExpectedBytecode: encoded,
		CompilerResults:  results,
		Diffs:            make([]diffmatchpatch.Diff, 0),
	}

	return toReturn, nil
}

// VerifyResult represents the result of the verification process.
type VerifyResult struct {
	Verified         bool                  // Whether the verification was successful or not.
	CompilerResults  *solc.CompilerResults // The results from the solc compiler.
	ExpectedBytecode string                // The expected bytecode.
	Diffs            []diffmatchpatch.Diff // The diffs between the provided bytecode and the compiled bytecode.
	DiffPretty       string                // The pretty printed diff between the provided bytecode and the compiled bytecode.
}

// IsVerified returns whether the verification was successful or not.
func (vr *VerifyResult) IsVerified() bool {
	return vr.Verified
}

// GetCompilerResults returns the results from the solc compiler.
func (vr *VerifyResult) GetCompilerResults() *solc.CompilerResults {
	return vr.CompilerResults
}

// GetExpectedBytecode returns the expected bytecode.
func (vr *VerifyResult) GetExpectedBytecode() string {
	return vr.ExpectedBytecode
}

// GetDiffs returns the diffs between the provided bytecode and the compiled bytecode.
func (vr *VerifyResult) GetDiffs() []diffmatchpatch.Diff {
	return vr.Diffs
}

// GetDiffPretty returns the pretty printed diff between the provided bytecode and the compiled bytecode.
func (vr *VerifyResult) GetDiffPretty() string {
	return vr.DiffPretty
}
