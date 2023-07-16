package ast

import (
	"encoding/json"
	"fmt"
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

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	id := atomic.AddInt64(&b.nextID, 1) - 1

	b.currentSourceUnit = &ast_pb.SourceUnit{
		Id:              id,
		AbsolutePath:    ctx.GetStart().GetInputStream().GetSourceName(),
		ExportedSymbols: make([]*ast_pb.ExportedSymbols, 0),
		NodeType:        ast_pb.NodeType_SOURCE_UNIT,
		Nodes:           &ast_pb.RootNode{},
		Src: &ast_pb.Src{
			Line:   int64(ctx.GetStart().GetLine()),
			Column: int64(ctx.GetStart().GetColumn()),
			Start:  int64(ctx.GetStart().GetStart()),
			// @TODO: GetStop() is always nil due to some reason so we cannot get lenght
			// just yet. We need to figure out why.
			//Length: ctx.GetStop() - ctx.GetStart() + 1,
			ParentIndex: int64(ctx.GetStart().GetTokenIndex()),
		},
		Comments: b.comments,
	}

	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			_ = contractCtx
		}

		if libraryCtx, ok := child.(*parser.LibraryDefinitionContext); ok {
			b.currentSourceUnit.License = b.GetLicense()

			// Alright lets extract bloody pragmas...
			pragmas := b.findPragmasForLibrary(ctx, libraryCtx)
			b.currentSourceUnit.Nodes.Nodes = append(
				b.currentSourceUnit.Nodes.Nodes,
				pragmas...,
			)
		}
	}

	// Just temporary...
	b.currentSourceUnit.Comments = nil
}

// ExitSourceUnit is called when production sourceUnit is exited.
func (b *ASTBuilder) ExitSourceUnit(ctx *parser.SourceUnitContext) {
	b.astRoot = &ast_pb.RootSourceUnit{
		SourceUnits: []*ast_pb.SourceUnit{b.currentSourceUnit},
	}
	b.currentSourceUnit = nil
}

func (b *ASTBuilder) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	//id := atomic.AddInt64(&b.nextID, 1) - 1
	identifierName := ctx.Identifier().GetText()

	fmt.Println("EnterContractDefinition AAA", identifierName)
}

// EnterPragmaDirective is called when production pragmaDirective is entered.
// However, it won't return pragma directives properly. For example, if we have
// experimental pragma, it won't return it. It will return only the pragma.
// Because of it, we are parsing pragmas in EnterSourceUnit to be able capture all of the
// pragmas and assign them based on the contract they belong to. Source file can have multiple
// contracts and multiple files and therefore we need to be able to assign pragmas to the
// correct contract.
// @WARN: DO NOT USE THIS METHOD.
func (b *ASTBuilder) EnterPragmaDirective(ctx *parser.PragmaDirectiveContext) {}

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

func (b *ASTBuilder) traverseReturnStatement(node *ast_pb.Node, bodyNode *ast_pb.Body, returnStatement *parser.ReturnStatementContext) *ast_pb.Statement {
	toReturn := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(returnStatement.GetStart().GetLine()),
			Column:      int64(returnStatement.GetStart().GetColumn()),
			Start:       int64(returnStatement.GetStart().GetStart()),
			End:         int64(returnStatement.GetStop().GetStop()),
			Length:      int64(returnStatement.GetStop().GetStop() - returnStatement.GetStart().GetStart() + 1),
			ParentIndex: bodyNode.Id,
		},
		NodeType: ast_pb.NodeType_RETURN_STATEMENT,
	}

	if expression := returnStatement.Expression(); expression != nil {
		toReturn.Expression = &ast_pb.Expression{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(expression.GetStart().GetLine()),
				Column:      int64(expression.GetStart().GetColumn()),
				Start:       int64(expression.GetStart().GetStart()),
				End:         int64(expression.GetStop().GetStop()),
				Length:      int64(expression.GetStop().GetStop() - expression.GetStart().GetStart() + 1),
				ParentIndex: toReturn.Id,
			},
			NodeType: ast_pb.NodeType_IDENTIFIER,
			Name:     expression.GetText(),
		}
	}

	// @TODO: Need to parse whole structure prior return types can be properly addressed.
	// It can be type that is in the body and not the type that is in arguments of the function.
	if node.ReturnParameters != nil {
		toReturn.FunctionReturnParameters = node.ReturnParameters.Id
	}

	return toReturn
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
