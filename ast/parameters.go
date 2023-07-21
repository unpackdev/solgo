package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

func (b *ASTBuilder) traverseParameterList(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, parameterCtx parser.IParameterListContext) *ast_pb.ParametersList {
	if parameterCtx == nil || parameterCtx.IsEmpty() {
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

		if paramCtx.DataLocation() != nil {
			switch paramCtx.DataLocation().GetText() {
			case "memory":
				pNode.StorageLocation = ast_pb.StorageLocation_MEMORY
			case "storage":
				pNode.StorageLocation = ast_pb.StorageLocation_STORAGE
			case "calldata":
				pNode.StorageLocation = ast_pb.StorageLocation_CALLDATA
			}
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
					ParentIndex: parametersList.Id,
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
					ParentIndex: parametersList.Id,
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
			fmt.Println("Expression type is not supported yet")
		}

		parametersList.Parameters = append(parametersList.Parameters, pNode)
	}

	return parametersList
}
