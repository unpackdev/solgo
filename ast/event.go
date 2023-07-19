package ast

import (
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseEventDefinition(sourceUnit *ast_pb.SourceUnit, identifierNode *ast_pb.Node, eventDefinitionCtx *parser.EventDefinitionContext) *ast_pb.Node {
	nodeCtx := &ast_pb.Node{
		Id: identifierNode.Id,
		Src: &ast_pb.Src{
			Line:        int64(eventDefinitionCtx.GetStart().GetLine()),
			Start:       int64(eventDefinitionCtx.GetStart().GetStart()),
			End:         int64(eventDefinitionCtx.GetStop().GetStop()),
			Length:      int64(eventDefinitionCtx.GetStop().GetStop() - eventDefinitionCtx.GetStart().GetStart() + 1),
			ParentIndex: identifierNode.Id,
		},
		NodeType:  ast_pb.NodeType_EVENT_DEFINITION,
		Name:      eventDefinitionCtx.Identifier().GetText(),
		Anonymous: eventDefinitionCtx.Anonymous() != nil,
	}

	var parametersList *ast_pb.ParametersList

	if len(eventDefinitionCtx.AllEventParameter()) > 0 {
		eventParametersCtx := eventDefinitionCtx.AllEventParameter()
		parametersList = &ast_pb.ParametersList{
			Id:       atomic.AddInt64(&b.nextID, 1) - 1,
			NodeType: ast_pb.NodeType_PARAMETER_LIST,
			Src: &ast_pb.Src{
				Line:        int64(eventParametersCtx[0].GetStart().GetLine()),
				Start:       int64(eventParametersCtx[0].GetStart().GetStart()),
				End:         int64(eventParametersCtx[0].GetStop().GetStop()),
				Length:      int64(eventParametersCtx[0].GetStop().GetStop() - eventDefinitionCtx.GetStart().GetStart() + 1),
				ParentIndex: nodeCtx.Id,
			},
		}

		parametersList.Parameters = make([]*ast_pb.Parameter, 0)

		for _, paramCtx := range eventDefinitionCtx.AllEventParameter() {
			paramNode := &ast_pb.Parameter{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(paramCtx.GetStart().GetLine()),
					Column:      int64(paramCtx.GetStart().GetColumn()),
					Start:       int64(paramCtx.GetStart().GetStart()),
					End:         int64(paramCtx.GetStop().GetStop()),
					Length:      int64(paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1),
					ParentIndex: parametersList.Id,
				},
				Scope: parametersList.Id,
				Name: func() string {
					if paramCtx.Identifier() != nil {
						return paramCtx.Identifier().GetText()
					}
					return ""
				}(),
				NodeType: ast_pb.NodeType_VARIABLE_DECLARATION,
				Indexed:  paramCtx.Indexed() != nil,
				// Just hardcoding it here to internal as I am not sure how
				// could it be possible to be anything else.
				// @TODO: Check if it is possible to be anything else.
				Visibility: ast_pb.Visibility_INTERNAL,
				// Mutable is the default state for parameter declarations.
				Mutability: ast_pb.Mutability_MUTABLE,
			}

			if paramCtx.GetType_().ElementaryTypeName() != nil {
				paramNode.NodeType = ast_pb.NodeType_ELEMENTARY_TYPE_NAME
				typeCtx := paramCtx.GetType_().ElementaryTypeName()
				normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
					typeCtx.GetText(),
				)
				paramNode.TypeName = &ast_pb.TypeName{
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
					NodeType: paramNode.NodeType,
					TypeDescriptions: &ast_pb.TypeDescriptions{
						TypeIdentifier: normalizedTypeIdentifier,
						TypeString:     normalizedTypeName,
					},
				}
			}

			parametersList.Parameters = append(parametersList.Parameters, paramNode)
		}
	}

	nodeCtx.Parameters = parametersList
	b.currentEvents = append(b.currentEvents, nodeCtx)

	return nodeCtx
}
