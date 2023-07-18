package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseBodyElement(identifierNode *ast_pb.Node, bodyElement parser.IContractBodyElementContext) *ast_pb.Node {
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
		fmt.Println("Found using for declaration", usingForDeclaration.GetText())
		toReturn = b.parseUsingForDeclaration(
			toReturn,
			usingForDeclaration.(*parser.UsingDirectiveContext),
		)
	} else if stateVariableDeclaration := bodyElement.StateVariableDeclaration(); stateVariableDeclaration != nil {
		fmt.Println("Found state variable declaration", stateVariableDeclaration.GetText())
		toReturn = b.parseStateVariableDeclaration(
			toReturn,
			stateVariableDeclaration.(*parser.StateVariableDeclarationContext),
		)
	} else if functionDefinition := bodyElement.FunctionDefinition(); functionDefinition != nil {
		toReturn = b.parseFunctionDefinition(
			toReturn,
			functionDefinition.(*parser.FunctionDefinitionContext),
		)
	} else {
		panic("Another type of body element that needs to be parsed...")
	}

	return toReturn
}
