package ast

import (
	"encoding/json"
	"io/ioutil"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	parser            *parser.SolidityParser // parser is the Solidity parser instance.
	nextID            int64                  // nextID is the next ID to assign to a node.
	comments          []*ast_pb.Comment
	commentsParsed    bool
	currentSourceUnit *ast_pb.SourceUnit
	astRoot           *ast_pb.RootSourceUnit
}

func NewAstBuilder(parser *parser.SolidityParser) *ASTBuilder {
	return &ASTBuilder{
		parser:   parser,
		comments: make([]*ast_pb.Comment, 0),
	}
}

func (b *ASTBuilder) traverseBodyElement(identifierNode *ast_pb.Node, bodyElement parser.IContractBodyElementContext) *ast_pb.Node {
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

	if functionDefinition := bodyElement.FunctionDefinition(); functionDefinition != nil {
		toReturn = b.traverseFunctionDefinition(
			toReturn,
			functionDefinition.(*parser.FunctionDefinitionContext),
		)
	} else {
		panic("Another type of body element that needs to be parsed...")
	}

	return toReturn
}

func (b *ASTBuilder) traverseStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, statementCtx parser.IStatementContext) *ast_pb.Statement {
	if simpleStatement := statementCtx.SimpleStatement(); simpleStatement != nil {
		return b.traverseSimpleStatement(node, bodyNode, simpleStatement.(*parser.SimpleStatementContext))
	}

	if returnStatement := statementCtx.ReturnStatement(); returnStatement != nil {
		return b.traverseReturnStatement(node, bodyNode, returnStatement.(*parser.ReturnStatementContext))
	}

	//panic("There are statements that needs to be traversed...")
	return nil
}

func (b *ASTBuilder) traverseSimpleStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, statement *parser.SimpleStatementContext) *ast_pb.Statement {
	toReturn := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(statement.GetStart().GetLine()),
			Column:      int64(statement.GetStart().GetColumn()),
			Start:       int64(statement.GetStart().GetStart()),
			End:         int64(statement.GetStop().GetStop()),
			Length:      int64(statement.GetStop().GetStop() - statement.GetStart().GetStart() + 1),
			ParentIndex: bodyNode.Id,
		},
	}

	if variableDeclaration := statement.VariableDeclarationStatement(); variableDeclaration != nil {
		toReturn.NodeType = ast_pb.NodeType_VARIABLE_DECLARATION_STATEMENT

		toReturn = b.traverseVariableDeclaration(
			node, bodyNode, toReturn,
			variableDeclaration.(*parser.VariableDeclarationStatementContext),
		)

	} else if expressionStatement := statement.ExpressionStatement(); expressionStatement != nil {
		//toReturn = b.traverseExpressionStatement(toReturn, expressionStatement.(*parser.ExpressionStatementContext))
	} else {
		panic("Unknown simple statement type...")
	}

	return toReturn
}

func (b *ASTBuilder) GetRoot() *ast_pb.RootSourceUnit {
	return b.astRoot
}

func (b *ASTBuilder) ToJSON() ([]byte, error) {
	return json.Marshal(b.astRoot)
}

func (b *ASTBuilder) ToJSONString() (string, error) {
	bts, err := b.ToJSON()
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

func (b *ASTBuilder) ToPrettyJSON() ([]byte, error) {
	return json.MarshalIndent(b.astRoot, "", "  ")
}

func (b *ASTBuilder) WriteJSONToFile(path string) error {
	bts, err := b.ToJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bts, 0644)
}
