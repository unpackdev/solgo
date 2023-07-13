package ast

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/davecgh/go-spew/spew"
	"github.com/txpull/solgo/parser"
)

// StatementNode represents a statement in Solidity.
type StatementNode struct {
	Expression string      `json:"expression"`
	Line       int         `json:"line"`
	Type       string      `json:"type"`
	Tokens     []TokenNode `json:"tokens"`
}

func (s *StatementNode) Children() []Node {
	// This will depend on the specific kind of statement.
	return nil
}

func (b *ASTBuilder) traverseStatements(body []*StatementNode, fnArguments []*VariableNode, node antlr.Tree) []*StatementNode {
	statements := []*StatementNode{}

	switch node := node.(type) {

	case *parser.ReturnStatementContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokenNode := TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			}

			if token.GetTokenType() == parser.SolidityParserFalse ||
				token.GetTokenType() == parser.SolidityParserTrue ||
				token.GetTokenType() == parser.SolidityParserBool {
				tokenNode.Type = "bool"
			}

			for _, fnArg := range fnArguments {
				if token.GetText() == fnArg.Name {
					tokenNode.IsFunctionArgument = true
					tokenNode.Type = fnArg.Type
				}
			}

			for _, stateVar := range b.currentContract.StateVariables {
				if token.GetText() == stateVar.Name {
					tokenNode.IsStateVariable = true
					tokenNode.Type = stateVar.Type
				}
			}

			tokensNode = append(tokensNode, tokenNode)
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				toReturn := []string{}

				for _, token := range tokens {
					if token.GetTokenType() != parser.SolidityParserSemicolon {
						toReturn = append(toReturn, token.GetText())
					}
				}

				return strings.TrimSpace(strings.Join(toReturn, " ")) + ";"
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "return",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

		spew.Dump(statementNode)

	case *parser.AssignmentContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokenNode := TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			}

			for _, fnArg := range fnArguments {
				if token.GetText() == fnArg.Name {
					tokenNode.IsFunctionArgument = true
					tokenNode.Type = fnArg.Type
				}
			}

			for _, stateVar := range b.currentContract.StateVariables {
				if token.GetText() == stateVar.Name {
					tokenNode.IsStateVariable = true
					tokenNode.Type = stateVar.Type
				}
			}

			tokensNode = append(tokensNode, tokenNode)
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				toReturn := []string{}

				for _, token := range tokens {
					if token.GetTokenType() != parser.SolidityParserSemicolon {
						toReturn = append(toReturn, token.GetText())
					}
				}

				return strings.TrimSpace(strings.Join(toReturn, " ")) + ";"
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "assignment",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	case *parser.VariableDeclarationStatementContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokenNode := TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			}

			if tokenNode.LexerTypeIndex == parser.SolidityParserIdentifier {
				tokenNode.IsFunctionArgument = true
			}

			tokensNode = append(tokensNode, tokenNode)
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				toReturn := []string{}

				for _, token := range tokens {
					if token.GetTokenType() != parser.SolidityParserSemicolon {
						toReturn = append(toReturn, token.GetText())
					}
				}

				return strings.TrimSpace(strings.Join(toReturn, " ")) + ";"
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "variable_declaration",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	case *parser.FunctionCallContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokensNode = append(tokensNode, TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			})
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				var toReturn string

				for _, token := range tokensNode {
					if token.LexerTypeIndex != parser.SolidityParserLParen {
						toReturn += token.Name + ""
					} else if token.LexerTypeIndex != parser.SolidityParserIdentifier {
						toReturn += token.Name + ""
					} else if token.LexerTypeIndex != parser.SolidityParserComma {
						toReturn += token.Name + " "
					} else if token.LexerTypeIndex != parser.SolidityParserSemicolon {
						toReturn += ""
					}
				}

				return toReturn
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "function_call",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	default:
		// This node is not a statement, so we recurse on its children.
		for i := 0; i < node.GetChildCount(); i++ {
			childStatements := b.traverseStatements(body, fnArguments, node.GetChild(i))
			if childStatements != nil {
				statements = append(statements, childStatements...)
			}
		}
	}

	return statements
}
