package validation

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/0x19/solc-switch"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/utils"
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
	if !solc.IsSynced() {
		if err := solc.Sync(); err != nil {
			return nil, err
		}
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

func (v *Verifier) Compile(ctx context.Context, config *solc.CompilerConfig) (*solc.CompilerResults, error) {
	source, err := config.GetJsonConfig().ToJSON()
	if err != nil {
		return nil, err
	}

	results, err := v.solc.Compile(ctx, string(source), config)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// VerifyFromResults compiles the sources using the solc compiler and then verifies the bytecode.
// If the bytecode does not match the compiled result, it returns a diff of the two.
// Returns true if the bytecode matches, otherwise returns false.
// Also returns an error if there's any issue in the compilation or verification process.
func (v *Verifier) VerifyFromResults(bytecode []byte, results *solc.CompilerResults) (*VerifyResult, error) {
	result := results.GetEntryContract()

	if result == nil {
		return nil, errors.New("no entry contract found in the compilation results")
	}

	encoded := hex.EncodeToString(bytecode)
	/* 	spew.Dump(encoded[20:70])
	   	spew.Dump(result.GetDeployedBytecode()[20:70])
	   	fmt.Println(len(encoded), len(result.GetDeployedBytecode()))

	   	os.WriteFile("from.txt", []byte(encoded), 0777)
	   	os.WriteFile("to.txt", []byte(result.GetDeployedBytecode()), 0777) */

	if !strings.Contains(result.GetDeployedBytecode(), encoded) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(encoded, result.GetDeployedBytecode(), false)
		toReturn := &VerifyResult{
			Verified:            false,
			CompilerResult:      result,
			ExpectedBytecode:    encoded,
			Diffs:               diffs,
			DiffPretty:          dmp.DiffPrettyText(diffs),
			LevenshteinDistance: dmp.DiffLevenshtein(diffs),
		}

		return toReturn, errors.New("bytecode missmatch, failed to verify")
	}

	toReturn := &VerifyResult{
		Verified:         true,
		ExpectedBytecode: encoded,
		CompilerResult:   result,
		Diffs:            make([]diffmatchpatch.Diff, 0),
	}

	return toReturn, nil
}

// Verify compiles the sources using the solc compiler and then verifies the bytecode.
// If the bytecode does not match the compiled result, it returns a diff of the two.
// Returns true if the bytecode matches, otherwise returns false.
// Also returns an error if there's any issue in the compilation or verification process.
func (v *Verifier) Verify(ctx context.Context, bytecode []byte, config *solc.CompilerConfig) (*VerifyResult, error) {
	var source string

	if config.GetJsonConfig() != nil {
		sourceBytes, err := config.GetJsonConfig().ToJSON()
		if err != nil {
			return nil, err
		}
		source = string(sourceBytes)
	} else {
		source = utils.StripExtraSPDXLines(utils.SimplifyImportPaths(
			v.GetSources().GetCombinedSource(),
		))
	}

	results, err := v.solc.Compile(ctx, source, config)
	if err != nil {
		return nil, err
	}

	for _, result := range results.GetResults() {
		if result.IsEntry() {
			encoded := hex.EncodeToString(bytecode)
			var retBytecode string
			if result.GetDeployedBytecode() == "" {
				retBytecode = result.GetBytecode()
			} else {
				retBytecode = result.GetDeployedBytecode()
			}

			if encoded != retBytecode {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(encoded, retBytecode, false)
				toReturn := &VerifyResult{
					Verified:            false,
					CompilerResult:      result,
					ExpectedBytecode:    encoded,
					Diffs:               diffs,
					DiffPretty:          dmp.DiffPrettyText(diffs),
					LevenshteinDistance: dmp.DiffLevenshtein(diffs),
				}

				return toReturn, errors.New("bytecode missmatch, failed to verify")
			}

			toReturn := &VerifyResult{
				Verified:         true,
				ExpectedBytecode: encoded,
				CompilerResult:   result,
				Diffs:            make([]diffmatchpatch.Diff, 0),
			}

			return toReturn, nil
		}
	}

	for _, result := range results.GetResults() {
		if result.HasErrors() {
			return nil, fmt.Errorf("compilation failed with errors: %v", result.GetErrors())
		}
	}

	return nil, fmt.Errorf("compilation did not contain entry contract results")
}

// VerifyResult represents the result of the verification process.
type VerifyResult struct {
	Verified            bool                  `json:"verified"`             // Whether the verification was successful or not.
	CompilerResult      *solc.CompilerResult  `json:"compiler_results"`     // The results from the solc compiler.
	ExpectedBytecode    string                `json:"expected_bytecode"`    // The expected bytecode.
	Diffs               []diffmatchpatch.Diff `json:"diffs"`                // The diffs between the provided bytecode and the compiled bytecode.
	DiffPretty          string                `json:"diffs_pretty"`         // The pretty printed diff between the provided bytecode and the compiled bytecode.
	LevenshteinDistance int                   `json:"levenshtein_distance"` // The levenshtein distance between the provided bytecode and the compiled bytecode.
}

// IsVerified returns whether the verification was successful or not.
func (vr *VerifyResult) IsVerified() bool {
	return vr.Verified
}

// GetCompilerResults returns the results from the solc compiler.
func (vr *VerifyResult) GetCompilerResult() *solc.CompilerResult {
	return vr.CompilerResult
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

// GetLevenshteinDistance returns the levenshtein distance between the provided bytecode and the compiled bytecode.
func (vr *VerifyResult) GetLevenshteinDistance() int {
	return vr.LevenshteinDistance
}
