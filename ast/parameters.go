package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) traverseParameterList(node *ast_pb.Node, parameterCtx parser.IParameterListContext) *ast_pb.ParametersList {
	if parameterCtx.IsEmpty() {
		return nil
	}

	parametersList := &ast_pb.ParametersList{
		Id:         atomic.AddInt64(&b.nextID, 1) - 1,
		Parameters: make([]*ast_pb.Parameter, 0),
		Src: &ast_pb.Src{
			Line:        int64(parameterCtx.GetStart().GetLine()),
			Start:       int64(parameterCtx.GetStart().GetStart()),
			End:         int64(parameterCtx.GetStop().GetStop()),
			Length:      int64(parameterCtx.GetStop().GetStop() - parameterCtx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType: ast_pb.NodeType_PARAMETER_LIST,
	}

	for _, paramCtx := range parameterCtx.AllParameterDeclaration() {
		if paramCtx.IsEmpty() {
			continue
		}

		pNode := &ast_pb.Parameter{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(paramCtx.GetStart().GetLine()),
				Column:      int64(paramCtx.GetStart().GetColumn()),
				Start:       int64(paramCtx.GetStart().GetStart()),
				End:         int64(paramCtx.GetStop().GetStop()),
				Length:      int64(paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1),
				ParentIndex: parametersList.Id,
			},
			Scope: node.Id,
			Name: func() string {
				if paramCtx.Identifier() != nil {
					return paramCtx.Identifier().GetText()
				}
				return ""
			}(),
			NodeType: ast_pb.NodeType_VARIABLE_DECLARATION,
			// Just hardcoding it here to internal as I am not sure how
			// could it be possible to be anything else.
			// @TODO: Check if it is possible to be anything else.
			Visibility: ast_pb.Visibility_INTERNAL,
			// Mutable is the default state for parameter declarations.
			Mutability: ast_pb.Mutability_MUTABLE,
		}

		if paramCtx.GetType_().ElementaryTypeName() != nil {
			pNode.NodeType = ast_pb.NodeType_ELEMENTARY_TYPE_NAME
			typeCtx := paramCtx.GetType_().ElementaryTypeName()
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				typeCtx.GetText(),
			)
			pNode.TypeName = &ast_pb.TypeName{
				Id:   atomic.AddInt64(&b.nextID, 1) - 1,
				Name: typeCtx.GetText(),
				Src: &ast_pb.Src{
					Line:        int64(paramCtx.GetStart().GetLine()),
					Column:      int64(paramCtx.GetStart().GetColumn()),
					Start:       int64(paramCtx.GetStart().GetStart()),
					End:         int64(paramCtx.GetStop().GetStop()),
					Length:      int64(paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1),
					ParentIndex: parametersList.Id,
				},
				NodeType: pNode.NodeType,
				TypeDescriptions: &ast_pb.TypeDescriptions{
					TypeIdentifier: normalizedTypeIdentifier,
					TypeString:     normalizedTypeName,
				},
			}
		}

		parametersList.Parameters = append(parametersList.Parameters, pNode)
	}

	return parametersList
}
