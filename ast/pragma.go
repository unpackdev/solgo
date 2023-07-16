package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) findPragmasForLibrary(sourceUnit *parser.SourceUnitContext, library *parser.LibraryDefinitionContext) []*ast_pb.Node {
	pragmas := make([]*ast_pb.Node, 0)
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
				pragmas = append(pragmas, &ast_pb.Node{
					Id: id,
					Src: &ast_pb.Src{
						Line:        int64(pragmaLine),
						Column:      int64(pragmaCtx.GetStart().GetColumn()),
						Start:       int64(pragmaCtx.GetStart().GetStart()),
						End:         int64(pragmaCtx.GetStop().GetStop()),
						Length:      int64(pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1),
						ParentIndex: int64(b.currentSourceUnit.Src.ParentIndex),
					},
					NodeType: ast_pb.NodeType_NODE_TYPE_PRAGMA_DIRECTIVE,
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
			pragmas = append(pragmas, &ast_pb.Node{
				Id: id,
				Src: &ast_pb.Src{
					Line:        int64(pragmaLine),
					Column:      int64(pragmaCtx.GetStart().GetColumn()),
					Start:       int64(pragmaCtx.GetStart().GetStart()),
					End:         int64(pragmaCtx.GetStop().GetStop()),
					Length:      int64(pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1),
					ParentIndex: int64(b.currentSourceUnit.Src.ParentIndex),
				},
				NodeType: ast_pb.NodeType_NODE_TYPE_PRAGMA_DIRECTIVE,
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
	filteredPragmas := make([]*ast_pb.Node, 0)
	maxLine := int64(-1)

	// Iterate through the collected pragmas in reverse order
	for i := len(pragmas) - 1; i >= 0; i-- {
		pragma := pragmas[i]
		if maxLine == -1 || (pragma.Src.Line-maxLine <= 10 && pragma.Src.Line-maxLine >= -1) {
			filteredPragmas = append([]*ast_pb.Node{pragma}, filteredPragmas...)
			maxLine = pragma.Src.Line
		}
	}

	return pragmas
}
