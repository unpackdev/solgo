package ast

import (
	"sync/atomic"

	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) findPragmasForLibrary(sourceUnit *parser.SourceUnitContext, library *parser.LibraryDefinitionContext) []Node {
	pragmas := make([]Node, 0)
	contractLine := library.GetStart().GetLine()
	prevLine := -1
	pragmasBeforeLine := false

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if child == library {
			// Found the contract definition, stop traversing
			break
		}

		if pragmaCtx, ok := child.(*parser.PragmaDirectiveContext); ok {
			id := atomic.AddInt64(&b.nextID, 1) - 1

			pragmaLine := pragmaCtx.GetStart().GetLine()

			if prevLine == -1 {
				// First pragma encountered, add it to the result
				pragmas = append(pragmas, Node{
					ID: id,
					Src: Src{
						Line:   pragmaLine,
						Start:  pragmaCtx.GetStart().GetStart(),
						End:    pragmaCtx.GetStop().GetStop(),
						Length: pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1,
						Index:  b.currentSourceUnit.Src.Index,
					},
					NodeType: NodeTypePragmaDirective,
					Literals: getLiterals(pragmaCtx.GetText()),
				})
				prevLine = pragmaLine
				continue
			}

			// Check if there are pragmas before the current line
			if pragmaLine-prevLine > 10 && pragmasBeforeLine {
				break
			}

			// Add the pragma to the result
			pragmas = append(pragmas, Node{
				ID: id,
				Src: Src{
					Line:   pragmaLine,
					Start:  pragmaCtx.GetStart().GetStart(),
					End:    pragmaCtx.GetStop().GetStop(),
					Length: pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1,
					Index:  b.currentSourceUnit.Src.Index,
				},
				NodeType: NodeTypePragmaDirective,
				Literals: getLiterals(pragmaCtx.GetText()),
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
	filteredPragmas := make([]Node, 0)
	maxLine := -1

	// Iterate through the collected pragmas in reverse order
	for i := len(pragmas) - 1; i >= 0; i-- {
		pragma := pragmas[i]
		if maxLine == -1 || (pragma.Src.Line-maxLine <= 10 && pragma.Src.Line-maxLine >= -1) {
			filteredPragmas = append([]Node{pragma}, filteredPragmas...)
			maxLine = pragma.Src.Line
		}
	}

	return pragmas
}
