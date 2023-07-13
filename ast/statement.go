package ast

import (
	"github.com/antlr4-go/antlr/v4"
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

func (b *ASTBuilder) getTextSliceWithOriginalFormatting(tree antlr.Tree) []string {
	switch node := tree.(type) {
	case *antlr.TerminalNodeImpl:
		// If the node is a terminal node, return its text in a slice.
		return []string{node.GetText()}
	case antlr.RuleContext:
		// If the node is a rule context, get the text of its children.
		var text []string
		for _, child := range node.GetChildren() {
			text = append(text, b.getTextSliceWithOriginalFormatting(child)...)
		}
		return text
	default:
		// If the node is neither a terminal node nor a rule context, return an empty slice.
		return []string{}
	}
}
