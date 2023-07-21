package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseModifierDefinition(sourceUnit *ast_pb.SourceUnit, modifierNode *ast_pb.Node, ctx *parser.ModifierDefinitionContext) *ast_pb.Node {
	modifierNode.NodeType = ast_pb.NodeType_MODIFIER_DEFINITION
	modifierNode.Name = ctx.GetName().GetText()
	modifierNode.Visibility = ast_pb.Visibility_INTERNAL

	if ctx.AllVirtual() != nil {
		for _, virtualCtx := range ctx.AllVirtual() {
			if virtualCtx.GetText() == "virtual" {
				modifierNode.Virtual = true
			}
		}
	}

	modifierNode.Parameters = b.traverseParameterList(
		sourceUnit, modifierNode, ctx.ParameterList(),
	)

	if ctx.Block() != nil {
		blockCtx := ctx.Block()
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(blockCtx.GetStart().GetLine()),
				Column:      int64(blockCtx.GetStart().GetColumn()),
				Start:       int64(blockCtx.GetStart().GetStart()),
				End:         int64(blockCtx.GetStop().GetStop()),
				Length:      int64(blockCtx.GetStop().GetStop() - blockCtx.GetStart().GetStart() + 1),
				ParentIndex: modifierNode.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range blockCtx.AllStatement() {
			if statement.IsEmpty() {
				continue
			}

			// Parent index statement in this case is used only to be able provide
			// index to the parent node. It is not used for anything else.
			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
				sourceUnit, modifierNode, bodyNode, parentIndexStmt, statement,
			))
		}

		modifierNode.Body = bodyNode
	}

	b.dumpNode("")
	return modifierNode
}
