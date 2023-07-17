package ast

import (
	"fmt"

	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	//id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	fmt.Println("EnterContractDefinition AAA", identifierName)
}
