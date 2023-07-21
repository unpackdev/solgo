package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

func (b *ASTBuilder) parseErrorDefinition(
	sourceUnit *ast_pb.SourceUnit,
	defNode *ast_pb.Node,
	ctx *parser.ErrorDefinitionContext,
) *ast_pb.Node {
	node := &ast_pb.Node{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: defNode.Id,
		},
		NodeType: ast_pb.NodeType_ERROR_DEFINITION,
		Name:     ctx.GetName().GetText(),
	}

	parameters := &ast_pb.ParametersList{
		Id:         atomic.AddInt64(&b.nextID, 1) - 1,
		Parameters: make([]*ast_pb.Parameter, 0),
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType: ast_pb.NodeType_PARAMETER_LIST,
	}

	for _, paramCtx := range ctx.AllErrorParameter() {
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
				ParentIndex: parameters.Id,
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

		typeCtx := paramCtx.TypeName()
		if typeCtx.ElementaryTypeName() != nil {
			pNode.NodeType = ast_pb.NodeType_ELEMENTARY_TYPE_NAME
			typeCtx := typeCtx.ElementaryTypeName()
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
					ParentIndex: parameters.Id,
				},
				NodeType: pNode.NodeType,
				TypeDescriptions: &ast_pb.TypeDescriptions{
					TypeIdentifier: normalizedTypeIdentifier,
					TypeString:     normalizedTypeName,
				},
			}
		} else if typeCtx.MappingType() != nil {
			zap.L().Warn("Mapping type is not supported yet")
		} else if typeCtx.FunctionTypeName() != nil {
			zap.L().Warn("Function type is not supported yet")
		} else {
			// It seems to be a user defined type but that does not exist as type in parser...
			pNode.TypeName = &ast_pb.TypeName{
				Id:   atomic.AddInt64(&b.nextID, 1) - 1,
				Name: typeCtx.GetText(),
				Src: &ast_pb.Src{
					Line:        int64(paramCtx.GetStart().GetLine()),
					Column:      int64(paramCtx.GetStart().GetColumn()),
					Start:       int64(paramCtx.GetStart().GetStart()),
					End:         int64(paramCtx.GetStop().GetStop()),
					Length:      int64(paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1),
					ParentIndex: parameters.Id,
				},
				NodeType: ast_pb.NodeType_USER_DEFINED_PATH_NAME,
			}

			pathCtx := typeCtx.IdentifierPath()
			if pathCtx != nil {
				pNode.TypeName.PathNode = &ast_pb.PathNode{
					Id:   atomic.AddInt64(&b.nextID, 1) - 1,
					Name: pathCtx.GetText(),
					Src: &ast_pb.Src{
						Line:        int64(pathCtx.GetStart().GetLine()),
						Column:      int64(pathCtx.GetStart().GetColumn()),
						Start:       int64(pathCtx.GetStart().GetStart()),
						End:         int64(pathCtx.GetStop().GetStop()),
						Length:      int64(pathCtx.GetStop().GetStop() - pathCtx.GetStart().GetStart() + 1),
						ParentIndex: pNode.TypeName.Id,
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
				}
			}

			// Lets figure out type...
			// Search for argument reference in state variable declarations.
			referenceFound := false

			for _, node := range b.currentStateVariables {
				if node.GetName() == pathCtx.GetText() {
					referenceFound = true
					pNode.TypeName.PathNode.ReferencedDeclaration = node.Id
					pNode.TypeName.ReferencedDeclaration = node.Id
					pNode.TypeName.TypeDescriptions = node.TypeDescriptions
					pNode.TypeDescriptions = node.TypeDescriptions
				}
			}

			if !referenceFound {
				for _, node := range b.currentEnums {
					if node.GetName() == pathCtx.GetText() {
						referenceFound = true
						pNode.TypeName.PathNode.ReferencedDeclaration = node.Id
						pNode.TypeName.ReferencedDeclaration = node.Id

						typeDescription := &ast_pb.TypeDescriptions{
							TypeIdentifier: func() string {
								return fmt.Sprintf(
									"t_enum_$_%s_$%d",
									pathCtx.GetText(),
									node.Id,
								)
							}(),
							TypeString: func() string {
								return fmt.Sprintf(
									"enum %s.%s",
									sourceUnit.GetName(),
									pathCtx.GetText(),
								)
							}(),
						}

						pNode.TypeName.TypeDescriptions = typeDescription
						pNode.TypeDescriptions = typeDescription
					}
				}
			}
		}

		if typeCtx.Expression() != nil {
			zap.L().Warn(
				"Expression type is not supported yet for error definition",
				zap.String("expression", typeCtx.Expression().GetText()),
				zap.String("source_unit_name", sourceUnit.GetName()),
				zap.String("error_name", node.Name),
			)
		}
	}

	node.Parameters = parameters
	b.currentErrors = append(b.currentErrors, node)

	return node
}
