package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// EnterPragmaDirective is called when production pragmaDirective is entered.
// However, it won't return pragma directives properly. For example, if we have
// experimental pragma, it won't return it. It will return only the pragma.
// Because of it, we are parsing pragmas in EnterSourceUnit to be able capture all of the
// pragmas and assign them based on the contract they belong to. Source file can have multiple
// contracts and multiple files and therefore we need to be able to assign pragmas to the
// correct contract.
// @WARN: DO NOT USE THIS METHOD.
func (b *ASTBuilder) EnterPragmaDirective(ctx *parser.PragmaDirectiveContext) {}

func (b *ASTBuilder) findPragmasForSourceUnit(sourceUnit *parser.SourceUnitContext, currentSourceUnit *ast_pb.SourceUnit, library *parser.LibraryDefinitionContext, contract *parser.ContractDefinitionContext) []*ast_pb.Node {
	pragmas := make([]*ast_pb.Node, 0)

	contractLine := func() int64 {
		if library != nil {
			return int64(library.GetStart().GetLine())
		} else if contract != nil {
			return int64(contract.GetStart().GetLine())
		}
		return 0
	}()

	prevLine := int64(-1)

	// Traverse the children of the source unit until the contract definition is found
	for _, child := range sourceUnit.GetChildren() {
		if library != nil && child == library {
			// Found the library definition, stop traversing
			break
		}

		if contract != nil && child == contract {
			// Found the contract definition, stop traversing
			break
		}

		if pragmaCtx, ok := child.(*parser.PragmaDirectiveContext); ok {
			id := atomic.AddInt64(&b.nextID, 1) - 1

			pragmaLine := int64(pragmaCtx.GetStart().GetLine())

			if prevLine == -1 {
				// First pragma encountered, add it to the result
				pragma := &ast_pb.Node{
					Id: id,
					Src: &ast_pb.Src{
						Line:        int64(pragmaLine),
						Column:      int64(pragmaCtx.GetStart().GetColumn()),
						Start:       int64(pragmaCtx.GetStart().GetStart()),
						End:         int64(pragmaCtx.GetStop().GetStop()),
						Length:      int64(pragmaCtx.GetStop().GetStop() - pragmaCtx.GetStart().GetStart() + 1),
						ParentIndex: currentSourceUnit.Id,
					},
					NodeType: ast_pb.NodeType_PRAGMA_DIRECTIVE,
					Literals: getLiterals(pragmaCtx.GetText()),
				}
				pragmas = append(pragmas, pragma)
				prevLine = int64(pragmaLine)

				continue
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
					ParentIndex: currentSourceUnit.Id,
				},
				NodeType: ast_pb.NodeType_PRAGMA_DIRECTIVE,
				Literals: getLiterals(pragmaCtx.GetText()),
			})

			// Update the previous line number
			prevLine = pragmaLine
		}
	}

	// Post pragma cleanup...
	// Remove pragmas that have large gaps between the lines, keep only higher lines
	filteredPragmas := make([]*ast_pb.Node, 0)
	maxLine := int64(-1)

	// Iterate through the collected pragmas in reverse order
	for i := len(pragmas) - 1; i >= 0; i-- {
		pragma := pragmas[i]

		/* 		fmt.Printf(
			"pragma: %v, line: %d, maxLine: %d, contractLine: %d, diff: %d\n",
			pragma.Literals,
			pragma.Src.Line,
			maxLine,
			contractLine,
			int64(contractLine)-pragma.Src.Line,
		) */
		if maxLine == -1 || (int64(contractLine)-pragma.Src.Line <= 10 && pragma.Src.Line-maxLine >= -1) {
			pragma.Src.ParentIndex = currentSourceUnit.Id
			filteredPragmas = append([]*ast_pb.Node{pragma}, filteredPragmas...)
			maxLine = pragma.Src.Line
		}
	}

	return filteredPragmas
}
