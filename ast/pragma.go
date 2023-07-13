package ast

import (
	"github.com/txpull/solgo/parser"
)

// PragmaNode represents a pragma directive in Solidity.
type PragmaNode struct {
	Directive string `json:"name"`
	Line      int    `json:"line"`
}

// Helper function to find pragmas for a specific contract definition
func (b *ASTBuilder) findPragmasForContract(sourceUnit *parser.SourceUnitContext, contract *parser.ContractDefinitionContext) []*PragmaNode {
	pragmas := make([]*PragmaNode, 0)
	contractLine := contract.GetStart().GetLine()
	prevLine := -1
	pragmasBeforeLine := false

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if child == contract {
			// Found the contract definition, stop traversing
			break
		}

		if pragmaCtx, ok := child.(*parser.PragmaDirectiveContext); ok {
			pragmaLine := pragmaCtx.GetStart().GetLine()

			if prevLine == -1 {
				// First pragma encountered, add it to the result
				pragmas = append(pragmas, &PragmaNode{
					Directive: pragmaCtx.GetText(),
					Line:      pragmaLine,
				})
				prevLine = pragmaLine
				continue
			}

			// Check if there are pragmas before the current line
			if pragmaLine-prevLine > 10 && pragmasBeforeLine {
				break
			}

			// Add the pragma to the result
			pragmas = append(pragmas, &PragmaNode{
				Directive: pragmaCtx.GetText(),
				Line:      pragmaLine,
			})

			// Update the previous line number
			prevLine = pragmaLine

			if pragmaLine < contractLine {
				pragmasBeforeLine = true
			}
		}
	}

	// Post pragma cleanup...
	// Remove pragmas that have large gaps between the lines, keep only higher lines
	filteredPragmas := make([]*PragmaNode, 0)
	maxLine := -1

	// Iterate through the collected pragmas in reverse order
	for i := len(pragmas) - 1; i >= 0; i-- {
		pragma := pragmas[i]

		if maxLine == -1 || (pragma.Line-maxLine <= 10 && pragma.Line-maxLine >= -1) {
			filteredPragmas = append([]*PragmaNode{pragma}, filteredPragmas...)
			maxLine = pragma.Line
		}
	}

	return filteredPragmas
}
