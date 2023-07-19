package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseUsingForDeclaration(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, ctx *parser.UsingDirectiveContext) *ast_pb.Node {
	node.NodeType = ast_pb.NodeType_USING_FOR_DIRECTIVE

	for _, identifierCtx := range ctx.AllIdentifierPath() {
		node.LibraryName = &ast_pb.LibraryName{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(identifierCtx.GetStart().GetLine()),
				Start:       int64(identifierCtx.GetStart().GetStart()),
				End:         int64(identifierCtx.GetStop().GetStop()),
				Length:      int64(identifierCtx.GetStop().GetStop() - identifierCtx.GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			Name:     identifierCtx.GetText(),
			NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
			ReferencedDeclaration: func() int64 {
				for _, unit := range b.sourceUnits {
					for _, symbol := range unit.ExportedSymbols {
						if symbol.Name == identifierCtx.GetText() {
							return symbol.Id
						}
					}
				}
				return 0
			}(),
		}
	}

	if typeNameCtx := ctx.TypeName(); typeNameCtx != nil {
		normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
			typeNameCtx.GetText(),
		)

		node.TypeName = &ast_pb.TypeName{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(typeNameCtx.GetStart().GetLine()),
				Start:       int64(typeNameCtx.GetStart().GetStart()),
				End:         int64(typeNameCtx.GetStop().GetStop()),
				Length:      int64(typeNameCtx.GetStop().GetStop() - typeNameCtx.GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			Name: typeNameCtx.GetText(),
			NodeType: func() ast_pb.NodeType {
				if typeNameCtx.ElementaryTypeName() != nil {
					return ast_pb.NodeType_ELEMENTARY_TYPE_NAME
				} else if typeNameCtx.MappingType() != nil {
					return ast_pb.NodeType_MAPPING_TYPE_NAME
				} else if typeNameCtx.FunctionTypeName() != nil {
					return ast_pb.NodeType_FUNCTION_TYPE_NAME
				}
				return ast_pb.NodeType_UNKNOWN_TYPE_NAME
			}(),
			TypeDescriptions: &ast_pb.TypeDescriptions{
				TypeIdentifier: normalizedTypeIdentifier,
				TypeString:     normalizedTypeName,
			},
		}

	}

	return node
}
