package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseEnumDefinition(sourceUnit *ast_pb.SourceUnit, enumNode *ast_pb.Node, ctx *parser.EnumDefinitionContext) *ast_pb.Node {
	enumNode.NodeType = ast_pb.NodeType_ENUM_DEFINITION
	enumNode.Name = ctx.GetName().GetText()
	enumNode.CanonicalName = fmt.Sprintf("%s.%s", sourceUnit.Name, enumNode.Name)

	for _, enumCtx := range ctx.GetEnumValues() {
		enumNode.Members = append(
			enumNode.Members,
			&ast_pb.Parameter{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(enumCtx.GetStart().GetLine()),
					Column:      int64(enumCtx.GetStart().GetColumn()),
					Start:       int64(enumCtx.GetStart().GetStart()),
					End:         int64(enumCtx.GetStop().GetStop()),
					Length:      int64(enumCtx.GetStop().GetStop() - enumCtx.GetStart().GetStart()),
					ParentIndex: enumNode.Id,
				},
				Name:     enumCtx.GetText(),
				NodeType: ast_pb.NodeType_ENUM_VALUE,
			},
		)
	}

	return enumNode
}
