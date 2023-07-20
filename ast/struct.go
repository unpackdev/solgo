package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

func (b *ASTBuilder) parseStructDefinition(sourceUnit *ast_pb.SourceUnit, structNode *ast_pb.Node, ctx *parser.StructDefinitionContext) *ast_pb.Node {
	structNode.NodeType = ast_pb.NodeType_STRUCT_DEFINITION
	structNode.Name = ctx.GetName().GetText()
	structNode.CanonicalName = fmt.Sprintf("%s.%s", sourceUnit.Name, structNode.Name)
	structNode.Visibility = ast_pb.Visibility_PUBLIC
	structNode.StorageLocation = ast_pb.StorageLocation_DEFAULT

	for _, memberCtx := range ctx.AllStructMember() {
		parameter := &ast_pb.Parameter{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(memberCtx.GetStart().GetLine()),
				Column:      int64(memberCtx.GetStart().GetColumn()),
				Start:       int64(memberCtx.GetStart().GetStart()),
				End:         int64(memberCtx.GetStop().GetStop()),
				Length:      int64(memberCtx.GetStop().GetStop() - memberCtx.GetStart().GetStart()),
				ParentIndex: structNode.Id,
			},
			Name:       memberCtx.GetName().GetText(),
			Mutability: ast_pb.Mutability_MUTABLE,
			Scope:      structNode.Id,
			NodeType:   ast_pb.NodeType_VARIABLE_DECLARATION,
			Visibility: ast_pb.Visibility_INTERNAL,
		}

		typeCtx := memberCtx.TypeName()

		if typeCtx.ElementaryTypeName() != nil {
			typeCtx := typeCtx.ElementaryTypeName()
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				typeCtx.GetText(),
			)
			parameter.TypeName = &ast_pb.TypeName{
				Id:   atomic.AddInt64(&b.nextID, 1) - 1,
				Name: typeCtx.GetText(),
				Src: &ast_pb.Src{
					Line:        int64(typeCtx.GetStart().GetLine()),
					Column:      int64(typeCtx.GetStart().GetColumn()),
					Start:       int64(typeCtx.GetStart().GetStart()),
					End:         int64(typeCtx.GetStop().GetStop()),
					Length:      int64(typeCtx.GetStop().GetStop() - typeCtx.GetStart().GetStart() + 1),
					ParentIndex: parameter.Id,
				},
				NodeType: ast_pb.NodeType_ELEMENTARY_TYPE_NAME,
				TypeDescriptions: &ast_pb.TypeDescriptions{
					TypeIdentifier: normalizedTypeIdentifier,
					TypeString:     normalizedTypeName,
				},
			}
			parameter.TypeDescriptions = parameter.TypeName.TypeDescriptions
		}

		if typeCtx.MappingType() != nil {
			zap.L().Warn(
				"Mapping types are not supported yet for structs...",
				zap.String("type", typeCtx.GetText()),
				zap.String("name", parameter.Name),
				zap.Int64("line", parameter.Src.Line),
			)
		}

		if typeCtx.FunctionTypeName() != nil {
			zap.L().Warn(
				"Function types are not supported yet for structs...",
				zap.String("type", typeCtx.GetText()),
				zap.String("name", parameter.Name),
				zap.Int64("line", parameter.Src.Line),
			)
		}

		structNode.Members = append(structNode.Members, parameter)
	}

	return structNode
}
