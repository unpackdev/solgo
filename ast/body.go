package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseBodyElement(sourceUnit *ast_pb.SourceUnit, identifierNode *ast_pb.Node, bodyElement parser.IContractBodyElementContext) *ast_pb.Node {
	id := atomic.AddInt64(&b.nextID, 1) - 1
	toReturn := &ast_pb.Node{
		Id: id,
		Src: &ast_pb.Src{
			Line:        int64(bodyElement.GetStart().GetLine()),
			Start:       int64(bodyElement.GetStart().GetStart()),
			End:         int64(bodyElement.GetStop().GetStop()),
			Length:      int64(bodyElement.GetStop().GetStop() - bodyElement.GetStart().GetStart() + 1),
			ParentIndex: identifierNode.Id,
		},
	}

	if usingForDeclaration := bodyElement.UsingDirective(); usingForDeclaration != nil {
		toReturn = b.parseUsingForDeclaration(
			sourceUnit,
			toReturn,
			usingForDeclaration.(*parser.UsingDirectiveContext),
		)
	} else if stateVariableDeclaration := bodyElement.StateVariableDeclaration(); stateVariableDeclaration != nil {
		toReturn = b.parseStateVariableDeclaration(
			sourceUnit,
			toReturn,
			stateVariableDeclaration.(*parser.StateVariableDeclarationContext),
		)
	} else if functionDefinition := bodyElement.FunctionDefinition(); functionDefinition != nil {
		toReturn = b.parseFunctionDefinition(
			sourceUnit,
			toReturn,
			functionDefinition.(*parser.FunctionDefinitionContext),
		)
	} else {
		panic("Another type of body element that needs to be parsed...")
	}

	return toReturn
}
