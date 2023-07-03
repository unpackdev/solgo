package ast

import (
	"github.com/antlr4-go/antlr/v4"
)

// StatementNode represents a statement in Solidity.
type StatementNode struct {
	Raw     []string `json:"raw"`
	TextRaw string   `json:"text_raw"`
}

func (s *StatementNode) Children() []Node {
	// This will depend on the specific kind of statement.
	return nil
}

/* // EnterStatement is called when the parser enters a statement.
func (b *ASTBuilder) EnterStatement(ctx *parser.StatementContext) {
	// Get the start and stop tokens of the statement.
	start := ctx.GetStart()
	stop := ctx.GetStop()

	// Get the text of the statement, preserving the original formatting.
	text := b.getTextWithOriginalFormatting(start, stop)

	// Create a new StatementNode with the text of the statement.
	statement := &StatementNode{
		Text: text,
	}

	// Add the statement to the current function.
	b.currentFunction.Body = append(b.currentFunction.Body, statement)
} */

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
