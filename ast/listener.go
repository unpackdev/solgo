package ast

import (
	"github.com/txpull/solgo/common"
	"github.com/txpull/solgo/parser"
)

type AstListener struct {
	*parser.BaseSolidityParserListener
	parser *AstParser
}

func NewAstListener() *AstListener {
	return &AstListener{parser: &AstParser{
		ast:            common.AST{},
		definedStructs: make(map[string]common.MethodIO),
		definedEnums:   make(map[string]bool),
	}}
}

func (l *AstListener) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	l.parser.contractName = ctx.Identifier().GetText()
}

// ... other methods similar to AbiListener

func (l *AstListener) GetParser() *AstParser {
	return l.parser
}
