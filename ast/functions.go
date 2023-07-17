package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseFunctionDefinition(node *ast_pb.Node, fd *parser.FunctionDefinitionContext) *ast_pb.Node {
	// Extract the function name.
	node.Name = fd.Identifier().GetText()

	// Set the function type and its kind.
	node.NodeType = ast_pb.NodeType_FUNCTION_DEFINITION
	node.Kind = ast_pb.NodeType_FUNCTION_DEFINITION

	// If block is not empty we are going to assume that the function is implemented.
	// @TODO: Take assumption to the next level in the future.
	node.Implemented = !fd.Block().IsEmpty()

	// Get function visibility state.
	for _, visibility := range fd.AllVisibility() {
		node.Visibility = visibility.GetText()
	}

	// Get function state mutability.
	for _, stateMutability := range fd.AllStateMutability() {
		node.StateMutability = stateMutability.GetText()
	}

	// Get function modifiers.
	for _, modifier := range fd.AllModifierInvocation() {
		panic("Modifier here...")
		_ = modifier
		//node.Modifiers = append(node.Modifiers, modifier.GetText())
	}

	// Check if function is virtual.
	for _, virtual := range fd.AllVirtual() {
		node.Virtual = virtual.GetText() == "virtual"
	}

	// Check if function is override.
	// @TODO: Implement override specification.
	for _, override := range fd.AllOverrideSpecifier() {
		panic("Override here...")
		_ = override
	}

	// Extract function parameters.
	if len(fd.AllParameterList()) > 0 {
		node.Parameters = b.traverseParameterList(node, fd.AllParameterList()[0])
	}

	// Extract function return parameters.
	node.ReturnParameters = b.traverseParameterList(node, fd.GetReturnParameters())

	// And now we are going to the big league. We are going to traverse the function body.
	if !fd.Block().IsEmpty() {
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(fd.Block().GetStart().GetLine()),
				Column:      int64(fd.Block().GetStart().GetColumn()),
				Start:       int64(fd.Block().GetStart().GetStart()),
				End:         int64(fd.Block().GetStop().GetStop()),
				Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range fd.Block().AllStatement() {
			if statement.IsEmpty() {
				continue
			}

			// Parent index statement in this case is used only to be able provide
			// index to the parent node. It is not used for anything else.
			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
				node, bodyNode, parentIndexStmt, statement,
			))
		}

		node.Body = bodyNode
	}

	return node
}

func (b *ASTBuilder) parseFunctionCall(fnNode *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, fnCtx *parser.FunctionCallContext) *ast_pb.Statement {
	statementNode.Expression = &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(fnCtx.GetStart().GetLine()),
			Column:      int64(fnCtx.GetStart().GetColumn()),
			Start:       int64(fnCtx.GetStart().GetStart()),
			End:         int64(fnCtx.GetStop().GetStop()),
			Length:      int64(fnCtx.GetStop().GetStop() - fnCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		NodeType: ast_pb.NodeType_FUNCTION_CALL,
	}

	expressionCtx := fnCtx.Expression()
	nameExpression := &ast_pb.Expression{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(expressionCtx.GetStart().GetLine()),
			Column:      int64(expressionCtx.GetStart().GetColumn()),
			Start:       int64(expressionCtx.GetStart().GetStart()),
			End:         int64(expressionCtx.GetStop().GetStop()),
			Length:      int64(expressionCtx.GetStop().GetStop() - expressionCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Expression.Id,
		},
		NodeType: ast_pb.NodeType_IDENTIFIER,
		Name:     expressionCtx.GetText(),
	}

	if fnCtx.CallArgumentList() != nil {
		for _, expressionCtx := range fnCtx.CallArgumentList().AllExpression() {
			argument := b.parseExpression(
				fnNode, bodyNode, nil, nameExpression.Id, expressionCtx,
			)
			statementNode.Arguments = append(statementNode.Arguments, argument)
			nameExpression.ArgumentTypes = append(
				nameExpression.ArgumentTypes, argument.TypeDescriptions,
			)
		}

		for _, _ = range fnCtx.CallArgumentList().AllNamedArgument() {
			panic("Named argument for function call here (functions.go)...")
			//argument := b.parseNamedArgument(statementNode, argumentCtx.(*parser.NamedArgumentContext))
			//statementNode.Arguments = append(statementNode.Arguments, argument)
		}
	}

	statementNode.Expression.Expression = nameExpression

	return statementNode
}
