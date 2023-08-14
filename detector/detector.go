package detector

import (
	"context"
	"encoding/hex"

	"github.com/txpull/solgo"
	"github.com/txpull/solgo/abi"
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/eip"
	"github.com/txpull/solgo/ir"
	"github.com/txpull/solgo/opcode"
	"github.com/txpull/solgo/solc"
)

// Detector is a utility structure that provides functionalities to detect and analyze
// Solidity source code. It encapsulates the context, sources, ABI builder, and solc compiler selector.
type Detector struct {
	ctx     context.Context // Context for the builder operations.
	sources *solgo.Sources  // Source files to be processed.
	builder *abi.Builder    // ABI builder for the source code.
	solc    *solc.Select    // Solc selector for the solc compiler.
}

// NewDetectorFromSources initializes a new Detector instance using the provided sources.
// It sets up the ABI builder and solc compiler selector which provide access to Global parser, AST and IR.
func NewDetectorFromSources(ctx context.Context, sources *solgo.Sources) (*Detector, error) {
	if !eip.StandardsLoaded() {
		if err := eip.LoadStandards(); err != nil {
			return nil, err
		}
	}

	builder, err := abi.NewBuilderFromSources(ctx, sources)
	if err != nil {
		return nil, err
	}

	solc, err := solc.NewSelect()
	if err != nil {
		return nil, err
	}

	return &Detector{
		ctx:     ctx,
		sources: sources,
		builder: builder,
		solc:    solc,
	}, nil
}

// GetContext returns the context associated with the Detector.
func (d *Detector) GetContext() context.Context {
	return d.ctx
}

// GetSources returns the Solidity source files associated with the Detector.
func (d *Detector) GetSources() *solgo.Sources {
	return d.sources
}

// GetABI returns the ABI builder associated with the Detector.
func (d *Detector) GetABI() *abi.Builder {
	return d.builder
}

// GetIR returns the intermediate representation (IR) builder associated with the Detector.
func (d *Detector) GetIR() *ir.Builder {
	return d.builder.GetParser()
}

// GetAST returns the abstract syntax tree (AST) builder associated with the Detector.
func (d *Detector) GetAST() *ast.ASTBuilder {
	return d.builder.GetAstBuilder()
}

// GetParser returns the parser associated with the Detector.
func (d *Detector) GetParser() *solgo.Parser {
	return d.builder.GetParser().GetParser()
}

// GetSolc returns the solc compiler selector associated with the Detector.
func (d *Detector) GetSolc() *solc.Select {
	return d.solc
}

// GetOpcodes returns the opcodes decompiler from provided byte array.
// It decompiles the bytecode of the contract, transaction, log to EVM opcodes.
func (d *Detector) GetOpcodes(data []byte) (*opcode.Decompiler, error) {
	return opcode.NewDecompiler(d.ctx, data)
}

// GetOpcodesFromHex returns the opcodes decompiler from provided hexadecimal string.
// It decompiles the bytecode of the contract, transaction, log to EVM opcodes.
func (d *Detector) GetOpcodesFromHex(data string) (*opcode.Decompiler, error) {
	hexData, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return opcode.NewDecompiler(d.ctx, hexData)
}

// Parse parses the Solidity source code and returns the errors encountered during the process.
func (d *Detector) Parse() []error {
	return d.builder.Parse()
}

// Build builds the ABI from the parsed source code and returns the errors encountered during the process.
func (d *Detector) Build() error {
	return d.builder.Build()
}
