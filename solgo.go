package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

type SolGo struct {
	input io.Reader
}

type TreeShapeListener struct {
	*parser.BaseSolidityParserListener
}

func NewSolGo(input io.Reader) *SolGo {
	return &SolGo{input: input}
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *SolGo) Parse() error {
	// Read the input
	inputBytes, err := ioutil.ReadAll(s.input)
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	// Create an Antlr InputStream
	is := antlr.NewInputStream(string(inputBytes))

	// Create the Lexer
	lexer := parser.NewSolidityLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Read all tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSolidityParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally parse the expression (by calling the start rule)
	tree := p.SourceUnit()

	// Create a new listener
	listener := NewTreeShapeListener()

	// Walk the tree with the listener
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	zap.L().Info("Parsing completed successfully")

	return nil
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	// Initialize zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)

	input := strings.NewReader("contract HelloWorld { function sayHello() public pure returns (string memory) { return 'Hello, World!'; } }")

	solGo := NewSolGo(input)
	if err := solGo.Parse(); err != nil {
		zap.L().Fatal("Error parsing input", zap.Error(err))
	}
}
