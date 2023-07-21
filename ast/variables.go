package ast

import (
	"fmt"
	"sync/atomic"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func (b *ASTBuilder) parseStateVariableDeclaration(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, ctx *parser.StateVariableDeclarationContext) *ast_pb.Node {
	nodeCtx := &ast_pb.Node{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		IsStateVariable: true,
		Name:            ctx.Identifier().GetText(),
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		Scope:           node.Id,
		StateMutability: ast_pb.Mutability_MUTABLE,
		Visibility: func() ast_pb.Visibility {
			if len(ctx.AllPublic()) > 0 {
				return ast_pb.Visibility_PUBLIC
			} else if len(ctx.AllPrivate()) > 0 {
				return ast_pb.Visibility_PRIVATE
			} else if len(ctx.AllInternal()) > 0 {
				return ast_pb.Visibility_INTERNAL
			} else {
				return ast_pb.Visibility_INTERNAL
			}
		}(),
		StorageLocation: ast_pb.StorageLocation_DEFAULT,
	}

	for _, constantCtx := range ctx.AllConstant() {
		nodeCtx.IsConstant = constantCtx != nil
	}

	for _, immutableCtx := range ctx.AllImmutable() {
		if immutableCtx != nil {
			nodeCtx.StateMutability = ast_pb.Mutability_IMMUTABLE
		}
	}

	typeNameCtx := ctx.GetType_()
	normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
		typeNameCtx.GetText(),
	)

	nodeCtx.TypeDescriptions = &ast_pb.TypeDescriptions{
		TypeString:     normalizedTypeName,
		TypeIdentifier: normalizedTypeIdentifier,
	}

	nodeCtx.TypeName = &ast_pb.TypeName{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(typeNameCtx.GetStart().GetLine()),
			Column:      int64(typeNameCtx.GetStart().GetColumn()),
			Start:       int64(typeNameCtx.GetStart().GetStart()),
			End:         int64(typeNameCtx.GetStop().GetStop()),
			Length:      int64(typeNameCtx.GetStop().GetStop() - typeNameCtx.GetStart().GetStart() + 1),
			ParentIndex: nodeCtx.Id,
		},
		Name:             typeNameCtx.GetText(),
		TypeDescriptions: nodeCtx.TypeDescriptions,
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
	}

	node.TypeName = b.traverseTypeName(sourceUnit, typeNameCtx, nodeCtx, nil)

	b.currentStateVariables = append(b.currentStateVariables, nodeCtx)
	return nodeCtx
}

func (b *ASTBuilder) traverseTypeName(sourceUnit *ast_pb.SourceUnit, typeNameCtx parser.ITypeNameContext, node *ast_pb.Node, typeNameNode *ast_pb.TypeName) *ast_pb.TypeName {
	if mappingCtx := typeNameCtx.MappingType(); mappingCtx != nil {
		keyCtx := mappingCtx.GetKey()
		valueCtx := mappingCtx.GetValue()

		node.TypeName.KeyType = b.generateTypeName(sourceUnit, keyCtx, node, typeNameNode)
		node.TypeName.ValueType = b.generateTypeName(sourceUnit, valueCtx, node, typeNameNode)

		if node.TypeName.KeyType != nil &&
			node.TypeName.ValueType != nil &&
			node.TypeName.KeyType.TypeDescriptions != nil &&
			node.TypeName.ValueType.TypeDescriptions != nil {
			node.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeString: fmt.Sprintf("mapping(%s => %s)", node.TypeName.KeyType.Name, node.TypeName.ValueType.Name),
				TypeIdentifier: fmt.Sprintf(
					"t_mapping_$t_%s_$t_%s$",
					node.TypeName.KeyType.TypeDescriptions.TypeString,
					node.TypeName.ValueType.TypeDescriptions.TypeString,
				),
			}
		}
	} else if elementaryCtx := typeNameCtx.ElementaryTypeName(); elementaryCtx != nil {
		normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
			elementaryCtx.GetText(),
		)

		node.TypeDescriptions = &ast_pb.TypeDescriptions{
			TypeString:     normalizedTypeName,
			TypeIdentifier: normalizedTypeIdentifier,
		}
	}

	return node.TypeName
}

func (b *ASTBuilder) generateTypeName(sourceUnit *ast_pb.SourceUnit, ctx interface{}, node *ast_pb.Node, typeNameNode *ast_pb.TypeName) *ast_pb.TypeName {
	var typeName ast_pb.TypeName

	typeName.Id = atomic.AddInt64(&b.nextID, 1) - 1
	typeName.NodeType = ast_pb.NodeType_ELEMENTARY_TYPE_NAME

	switch specificCtx := ctx.(type) {
	case parser.IMappingKeyTypeContext:
		typeName.Name = specificCtx.GetText()
		typeName.Src = &ast_pb.Src{
			Line:        int64(specificCtx.GetStart().GetLine()),
			Column:      int64(specificCtx.GetStart().GetColumn()),
			Start:       int64(specificCtx.GetStart().GetStart()),
			End:         int64(specificCtx.GetStop().GetStop()),
			Length:      int64(specificCtx.GetStop().GetStop() - specificCtx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		}
		if specificCtx.ElementaryTypeName() != nil {
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				specificCtx.ElementaryTypeName().GetText(),
			)

			typeName.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeString:     normalizedTypeName,
				TypeIdentifier: normalizedTypeIdentifier,
			}
		}
	case parser.IMappingTypeContext:
		typeNameNode.NodeType = ast_pb.NodeType_MAPPING_TYPE_NAME
		keyCtx := specificCtx.GetKey()
		valueCtx := specificCtx.GetValue()

		typeNameNode.KeyType = b.generateTypeName(sourceUnit, keyCtx, node, typeNameNode)
		typeNameNode.ValueType = b.generateTypeName(sourceUnit, valueCtx, node, typeNameNode)

		if typeNameNode.KeyType != nil &&
			typeNameNode.ValueType != nil &&
			typeNameNode.KeyType.TypeDescriptions != nil &&
			typeNameNode.ValueType.TypeDescriptions != nil {
			node.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeString: fmt.Sprintf("mapping(%s => %s)", typeNameNode.KeyType.Name, typeNameNode.ValueType.Name),
				TypeIdentifier: fmt.Sprintf(
					"t_mapping_$t_%s_$t_%s$",
					typeNameNode.KeyType.TypeDescriptions.TypeString,
					typeNameNode.ValueType.TypeDescriptions.TypeString,
				),
			}
		}
	case parser.ITypeNameContext:
		typeName.Name = specificCtx.GetText()
		typeName.Src = &ast_pb.Src{
			Line:        int64(specificCtx.GetStart().GetLine()),
			Column:      int64(specificCtx.GetStart().GetColumn()),
			Start:       int64(specificCtx.GetStart().GetStart()),
			End:         int64(specificCtx.GetStop().GetStop()),
			Length:      int64(specificCtx.GetStop().GetStop() - specificCtx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		}

		if specificCtx.ElementaryTypeName() != nil {
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				specificCtx.ElementaryTypeName().GetText(),
			)

			typeName.TypeDescriptions = &ast_pb.TypeDescriptions{
				TypeString:     normalizedTypeName,
				TypeIdentifier: normalizedTypeIdentifier,
			}
		} else if specificCtx.MappingType() != nil {
			typeName.NodeType = ast_pb.NodeType_MAPPING_TYPE_NAME
			b.generateTypeName(sourceUnit, specificCtx.MappingType(), node, &typeName)
		} else {

			// It seems to be a user defined type but that does not exist as type in parser...
			typeName.NodeType = ast_pb.NodeType_USER_DEFINED_PATH_NAME

			pathCtx := specificCtx.IdentifierPath()
			if pathCtx != nil {
				typeName.PathNode = &ast_pb.PathNode{
					Id:   atomic.AddInt64(&b.nextID, 1) - 1,
					Name: pathCtx.GetText(),
					Src: &ast_pb.Src{
						Line:        int64(pathCtx.GetStart().GetLine()),
						Column:      int64(pathCtx.GetStart().GetColumn()),
						Start:       int64(pathCtx.GetStart().GetStart()),
						End:         int64(pathCtx.GetStop().GetStop()),
						Length:      int64(pathCtx.GetStop().GetStop() - pathCtx.GetStart().GetStart() + 1),
						ParentIndex: typeName.Id,
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
				}
			}

			// Search for argument reference in state variable declarations.
			referenceFound := false

			for _, node := range b.currentStateVariables {
				if node.GetName() == pathCtx.GetText() {
					referenceFound = true
					typeName.PathNode.ReferencedDeclaration = node.Id
					typeName.ReferencedDeclaration = node.Id
					typeName.TypeDescriptions = node.TypeDescriptions
				}
			}

			if !referenceFound {
				for _, node := range b.currentEnums {
					if node.GetName() == pathCtx.GetText() {
						referenceFound = true
						typeName.PathNode.ReferencedDeclaration = node.Id
						typeName.ReferencedDeclaration = node.Id

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

						typeName.TypeDescriptions = typeDescription
						typeName.TypeDescriptions = typeDescription
					}
				}
			}

			if !referenceFound {
				for _, node := range b.currentStructs {
					if node.GetName() == pathCtx.GetText() {
						referenceFound = true
						typeName.PathNode.ReferencedDeclaration = node.Id
						typeName.ReferencedDeclaration = node.Id

						typeDescription := &ast_pb.TypeDescriptions{
							TypeIdentifier: func() string {
								return fmt.Sprintf(
									"t_struct_$_%s_$%d",
									pathCtx.GetText(),
									node.Id,
								)
							}(),
							TypeString: func() string {
								return fmt.Sprintf(
									"struct %s.%s",
									sourceUnit.GetName(),
									pathCtx.GetText(),
								)
							}(),
						}

						typeName.TypeDescriptions = typeDescription
						typeName.TypeDescriptions = typeDescription
					}
				}
			}
		}
	}

	return &typeName
}

func (b *ASTBuilder) parseVariableDeclaration(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, statementNode *ast_pb.Statement, variableCtx *parser.VariableDeclarationStatementContext) *ast_pb.Statement {
	declarationCtx := variableCtx.VariableDeclaration()
	identifierCtx := declarationCtx.Identifier()

	declaration := &ast_pb.Declaration{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(declarationCtx.GetStart().GetLine()),
			Column:      int64(declarationCtx.GetStart().GetColumn()),
			Start:       int64(declarationCtx.GetStart().GetStart()),
			End:         int64(declarationCtx.GetStop().GetStop()),
			Length:      int64(declarationCtx.GetStop().GetStop() - declarationCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		Name:            identifierCtx.GetText(),
		Mutability:      ast_pb.Mutability_MUTABLE,
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		Scope:           bodyNode.Id,
		StorageLocation: ast_pb.StorageLocation_DEFAULT,
		Visibility:      ast_pb.Visibility_INTERNAL,
	}

	if declarationCtx.DataLocation() != nil {
		if declarationCtx.DataLocation().Memory() != nil {
			declaration.StorageLocation = ast_pb.StorageLocation_MEMORY
		} else if declarationCtx.DataLocation().Storage() != nil {
			declaration.StorageLocation = ast_pb.StorageLocation_STORAGE
		} else if declarationCtx.DataLocation().Calldata() != nil {
			declaration.StorageLocation = ast_pb.StorageLocation_CALLDATA
		}
	}

	typeCtx := declarationCtx.GetType_().ElementaryTypeName()
	normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
		typeCtx.GetText(),
	)

	declaration.TypeName = &ast_pb.TypeName{
		Id:   atomic.AddInt64(&b.nextID, 1) - 1,
		Name: typeCtx.GetText(),
		Src: &ast_pb.Src{
			Line:        int64(typeCtx.GetStart().GetLine()),
			Column:      int64(typeCtx.GetStart().GetColumn()),
			Start:       int64(typeCtx.GetStart().GetStart()),
			End:         int64(typeCtx.GetStop().GetStop()),
			Length:      int64(typeCtx.GetStop().GetStop() - typeCtx.GetStart().GetStart() + 1),
			ParentIndex: statementNode.Id,
		},
		NodeType: ast_pb.NodeType_ELEMENTARY_TYPE_NAME,
		TypeDescriptions: &ast_pb.TypeDescriptions{
			TypeIdentifier: normalizedTypeIdentifier,
			TypeString:     normalizedTypeName,
		},
	}

	declaration.TypeDescriptions = declaration.TypeName.TypeDescriptions

	statementNode.Declarations = append(statementNode.Declarations, declaration)
	statementNode.Assignments = append(statementNode.Assignments, declaration.Id)

	statementNode.InitialValue = b.parseExpression(
		sourceUnit, node, bodyNode, nil, statementNode.Id, variableCtx.Expression(),
	)

	return statementNode
}
