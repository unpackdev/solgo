package solgo

import (
	"context"
	"fmt"
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
)

// SolGo is a struct that encapsulates the functionality for parsing and analyzing Solidity contracts.
type SolGo struct {
	// ctx is the context in which SolGo operates. It can be used to control cancellation of parsing.
	ctx context.Context
	// inputRaw is the raw input reader from which the Solidity contract is read.
	inputRaw io.Reader
	// inputStream is the ANTLR input stream which is used by the lexer.
	inputStream *antlr.InputStream
	// lexer is the Solidity lexer which tokenizes the input stream.
	lexer *parser.SolidityLexer
	// tokenStream is the stream of tokens produced by the lexer.
	tokenStream *antlr.CommonTokenStream
	// solidityParser is the Solidity parser which parses the token stream.
	solidityParser *parser.SolidityParser
	// listeners is a map of listener names to ParseTreeListener instances.
	// These listeners are invoked as the parser walks the parse tree.
	listeners listeners
	// errListener is a SyntaxErrorListener which collects syntax errors encountered during parsing.
	errListener *SyntaxErrorListener
}

// New creates a new instance of SolGo.
// It takes a context and an io.Reader from which the Solidity contract is read.
// It initializes an input stream, lexer, token stream, and parser, and sets up error listeners.
func New(ctx context.Context, input io.Reader) (*SolGo, error) {
	ib, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	// Create an input stream from the input
	inputStream := antlr.NewInputStream(string(ib))

	// Create a new SyntaxErrorListener
	errListener := NewSyntaxErrorListener()

	// Create a new Solidity lexer with the input stream
	lexer := parser.NewSolidityLexer(inputStream)

	// Remove the default error listeners
	lexer.RemoveErrorListeners()

	// Add our SyntaxErrorListener
	lexer.AddErrorListener(errListener)

	// Create a new token stream from the lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new Solidity parser with the token stream
	solidityParser := parser.NewSolidityParser(stream)

	// Remove the default error listeners
	solidityParser.RemoveErrorListeners()

	// Add our SyntaxErrorListener
	solidityParser.AddErrorListener(errListener)

	return &SolGo{
		ctx:            ctx,
		inputRaw:       input,
		inputStream:    inputStream,
		lexer:          lexer,
		tokenStream:    stream,
		solidityParser: solidityParser,
		errListener:    errListener,
		listeners:      make(listeners),
	}, nil
}

// GetInput returns the raw input reader from which the Solidity contract is read.
func (s *SolGo) GetInput() io.Reader {
	return s.inputRaw
}

// GetInputStream returns the ANTLR input stream which is used by the lexer.
func (s *SolGo) GetInputStream() *antlr.InputStream {
	return s.inputStream
}

// GetLexer returns the Solidity lexer which tokenizes the input stream.
func (s *SolGo) GetLexer() *parser.SolidityLexer {
	return s.lexer
}

// GetTokenStream returns the stream of tokens produced by the lexer.
func (s *SolGo) GetTokenStream() *antlr.CommonTokenStream {
	return s.tokenStream
}

// GetParser returns the Solidity parser which parses the token stream.
func (s *SolGo) GetParser() *parser.SolidityParser {
	return s.solidityParser
}

// GetTree returns the root of the parse tree that results from parsing the Solidity contract.
func (s *SolGo) GetTree() antlr.ParseTree {
	return s.solidityParser.SourceUnit()
}

// Parse initiates the parsing process. It walks the parse tree with all registered listeners
// and returns any syntax errors that were encountered during parsing.
func (s *SolGo) Parse() []SyntaxError {
	tree := s.GetTree()

	// Walk the parse tree with all registered listeners
	for _, listener := range s.GetAllListeners() {
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	}

	// If there were syntax errors, return them
	if len(s.errListener.Errors) > 0 {
		return s.errListener.Errors
	}

	return nil
}
