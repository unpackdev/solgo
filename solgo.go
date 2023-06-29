package solgo

import (
	"context"
	"fmt"
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

type SolGo struct {
	ctx                   context.Context
	inputRaw              io.Reader
	inputStream           *antlr.InputStream
	lexer                 *parser.SolidityLexer
	tokenStream           *antlr.CommonTokenStream
	solidityParser        *parser.SolidityParser
	listeners             listeners
	errListener           *SyntaxErrorListener
	diagnosticErrListener *antlr.DiagnosticErrorListener
}

func New(ctx context.Context, input io.Reader) (*SolGo, error) {
	ib, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	inputStream := antlr.NewInputStream(string(ib))

	//diagnosticErrorListener := antlr.NewDiagnosticErrorListener(true)

	errListener := NewSyntaxErrorListener()

	lexer := parser.NewSolidityLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	solidityParser := parser.NewSolidityParser(stream)
	solidityParser.RemoveErrorListeners()
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

func (s *SolGo) GetInput() io.Reader {
	return s.inputRaw
}

func (s *SolGo) GetInputStream() *antlr.InputStream {
	return s.inputStream
}

func (s *SolGo) GetLexer() *parser.SolidityLexer {
	return s.lexer
}

func (s *SolGo) GetTokenStream() *antlr.CommonTokenStream {
	return s.tokenStream
}

func (s *SolGo) GetParser() *parser.SolidityParser {
	return s.solidityParser
}

func (s *SolGo) GetTree() antlr.ParseTree {
	return s.solidityParser.SourceUnit()
}

func (s *SolGo) Parse() error {
	tree := s.GetTree()

	for _, listener := range s.GetAllListeners() {
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	}

	if s.solidityParser.HasError() {
		// Print syntax errors
		err := s.solidityParser.GetError()

		zap.L().Error("Syntax error", zap.String("message", err.GetMessage()))

	}

	if len(s.errListener.Errors) > 0 {
		// Handle syntax errors here
		// For example, you could print them to the console:
		for _, err := range s.errListener.Errors {
			fmt.Printf("Syntax error at line %d, column %d: %s\n", err.Line, err.Column, err.Message)
			fmt.Printf("Context: %s, Severity: %d\n", err.Context, err.Severity)
		}
		return fmt.Errorf("parsing failed with %d syntax errors", len(s.errListener.Errors))
	}

	zap.L().Info("Parsing completed successfully")
	return nil
}
