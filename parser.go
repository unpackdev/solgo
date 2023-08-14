package solgo

import (
	"context"
	"fmt"
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
	"github.com/txpull/solgo/syntaxerrors"
)

// Parser is a struct that encapsulates the functionality for parsing and analyzing Solidity contracts.
type Parser struct {
	// ctx is the context in which SolGo operates. It can be used to control cancellation of parsing.
	ctx context.Context
	// sources is a struct that contains the sources of the Solidity contract.
	sources *Sources
	// inputRaw is the raw input reader from which the Solidity contract is read.
	inputRaw io.Reader
	// inputStream is the ANTLR input stream which is used by the lexer.
	inputStream *antlr.InputStream
	// lexer is the Solidity lexer which tokenizes the input stream.
	lexer *parser.SolidityLexer
	// tokenStream is the stream of tokens produced by the lexer.
	tokenStream *antlr.CommonTokenStream
	// solidityParser is the Solidity parser which parses the token stream.
	solidityParser *syntaxerrors.ContextualParser
	// listeners is a map of listener names to ParseTreeListener instances.
	// These listeners are invoked as the parser walks the parse tree.
	listeners listeners
	// errListener is a SyntaxErrorListener which collects syntax errors encountered during parsing.
	errListener *syntaxerrors.SyntaxErrorListener
}

// New creates a new instance of SolGo.
// It takes a context and an io.Reader from which the Solidity contract is read.
// It initializes an input stream, lexer, token stream, and parser, and sets up error listeners.
func NewParser(ctx context.Context, input io.Reader) (*Parser, error) {
	ib, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	// Create an input stream from the input
	inputStream := antlr.NewInputStream(string(ib))

	// Create a new SyntaxErrorListener
	errListener := syntaxerrors.NewSyntaxErrorListener()

	// Create a new Solidity lexer with the input stream
	lexer := parser.NewSolidityLexer(inputStream)

	// Remove the default error listeners
	lexer.RemoveErrorListeners()

	// Add our SyntaxErrorListener
	lexer.AddErrorListener(errListener)

	// Create a new token stream from the lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new ContextualParser with the token stream and listener
	contextualParser := syntaxerrors.NewContextualParser(stream, errListener)

	return &Parser{
		ctx:            ctx,
		inputRaw:       input,
		inputStream:    inputStream,
		lexer:          lexer,
		tokenStream:    stream,
		solidityParser: contextualParser,
		errListener:    errListener,
		listeners:      make(listeners),
	}, nil
}

// NewParserFromSources creates a new instance of parser from a reader.
// It takes a context and an io.Reader from which the Solidity contract is read.
// It initializes an input stream, lexer, token stream, and parser, and sets up error listeners.
func NewParserFromSources(ctx context.Context, sources *Sources) (*Parser, error) {
	if err := sources.Prepare(); err != nil {
		return nil, fmt.Errorf("error preparing sources: %w", err)
	}

	// Create an input stream from the input
	inputStream := antlr.NewInputStream(sources.GetCombinedSource())

	// Create a new SyntaxErrorListener
	errListener := syntaxerrors.NewSyntaxErrorListener()

	// Create a new Solidity lexer with the input stream
	lexer := parser.NewSolidityLexer(inputStream)

	// Remove the default error listeners
	lexer.RemoveErrorListeners()

	// Add our SyntaxErrorListener
	lexer.AddErrorListener(errListener)

	// Create a new token stream from the lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new ContextualParser with the token stream and listener
	contextualParser := syntaxerrors.NewContextualParser(stream, errListener)

	return &Parser{
		ctx:            ctx,
		sources:        sources,
		inputRaw:       nil,
		inputStream:    inputStream,
		lexer:          lexer,
		tokenStream:    stream,
		solidityParser: contextualParser,
		errListener:    errListener,
		listeners:      make(listeners),
	}, nil
}

// GetSources returns the sources of the Solidity contract.
func (s *Parser) GetSources() *Sources {
	return s.sources
}

// GetInput returns the raw input reader from which the Solidity contract is read.
func (s *Parser) GetInput() io.Reader {
	return s.inputRaw
}

// GetInputStream returns the ANTLR input stream which is used by the lexer.
func (s *Parser) GetInputStream() *antlr.InputStream {
	return s.inputStream
}

// GetLexer returns the Solidity lexer which tokenizes the input stream.
func (s *Parser) GetLexer() *parser.SolidityLexer {
	return s.lexer
}

// GetTokenStream returns the stream of tokens produced by the lexer.
func (s *Parser) GetTokenStream() *antlr.CommonTokenStream {
	return s.tokenStream
}

// GetParser returns the Solidity parser which parses the token stream.
func (s *Parser) GetParser() *parser.SolidityParser {
	return s.solidityParser.SolidityParser
}

// GetContextualParser returns the ContextualParser which wraps the Solidity parser.
func (s *Parser) GetContextualParser() *syntaxerrors.ContextualParser {
	return s.solidityParser
}

// GetTree returns the root of the parse tree that results from parsing the Solidity contract.
func (s *Parser) GetTree() antlr.ParseTree {
	return s.solidityParser.SourceUnit()
}

// Parse initiates the parsing process. It walks the parse tree with all registered listeners
// and returns any syntax errors that were encountered during parsing.
func (s *Parser) Parse() []syntaxerrors.SyntaxError {
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
