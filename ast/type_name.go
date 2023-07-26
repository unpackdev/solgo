package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

type TypeName struct {
	*ASTBuilder

	Id                    int64             `json:"id"`
	NodeType              ast_pb.NodeType   `json:"node_type"`
	Src                   SrcNode           `json:"src"`
	Name                  string            `json:"name,omitempty"`
	TypeDescription       *TypeDescription  `json:"type_descriptions,omitempty"`
	KeyType               *TypeName         `json:"key_type,omitempty"`
	ValueType             *TypeName         `json:"value_type,omitempty"`
	PathNode              *PathNode         `json:"path_node,omitempty"`
	StateMutability       ast_pb.Mutability `json:"state_mutability,omitempty"`
	ReferencedDeclaration int64             `json:"referenced_declaration"`
}

func NewTypeName(b *ASTBuilder) *TypeName {
	return &TypeName{
		ASTBuilder: b,
	}
}

func (t *TypeName) GetId() int64 {
	return t.Id
}

func (t *TypeName) GetType() ast_pb.NodeType {
	return t.NodeType
}

func (t *TypeName) GetSrc() SrcNode {
	return t.Src
}

func (t *TypeName) GetName() string {
	return t.Name
}

func (t *TypeName) GetTypeDescription() *TypeDescription {
	return t.TypeDescription
}

func (t *TypeName) GetPathNode() *PathNode {
	return t.PathNode
}

func (t *TypeName) GetReferencedDeclaration() int64 {
	return t.ReferencedDeclaration
}

func (t *TypeName) GetKeyType() *TypeName {
	return t.KeyType
}

func (t *TypeName) GetValueType() *TypeName {
	return t.ValueType
}

func (t *TypeName) parseElementaryTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.ElementaryTypeNameContext) {
	t.Name = ctx.GetText()
	t.NodeType = ast_pb.NodeType_ELEMENTARY_TYPE_NAME

	normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
		ctx.GetText(),
	)

	switch normalizedTypeIdentifier {
	case "t_address":
		t.StateMutability = ast_pb.Mutability_NONPAYABLE
	case "t_address_payable":
		t.StateMutability = ast_pb.Mutability_PAYABLE
	}

	t.TypeDescription = &TypeDescription{
		TypeIdentifier: normalizedTypeIdentifier,
		TypeString:     normalizedTypeName,
	}
}

func (t *TypeName) parseIdentifierPath(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.IdentifierPathContext) {
	t.NodeType = ast_pb.NodeType_USER_DEFINED_PATH_NAME
	if len(ctx.AllIdentifier()) > 0 {
		identifierCtx := ctx.Identifier(0)
		t.PathNode = &PathNode{
			Id:   t.GetNextID(),
			Name: identifierCtx.GetText(),
			Src: SrcNode{
				Id:          t.GetNextID(),
				Line:        int64(identifierCtx.GetStart().GetLine()),
				Column:      int64(identifierCtx.GetStart().GetColumn()),
				Start:       int64(identifierCtx.GetStart().GetStart()),
				End:         int64(identifierCtx.GetStop().GetStop()),
				Length:      int64(identifierCtx.GetStop().GetStop() - identifierCtx.GetStart().GetStart() + 1),
				ParentIndex: t.Id,
			},
			NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
		}

		if ref, refTypeDescription := discoverReferenceByCtxName(t.ASTBuilder, identifierCtx.GetText()); ref != nil {
			t.PathNode.ReferencedDeclaration = ref.GetId()
			t.ReferencedDeclaration = ref.GetId()
			t.TypeDescription = refTypeDescription
		}
	}
}

func (t *TypeName) parseMappingTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.MappingTypeContext) {
	keyCtx := ctx.GetKey()
	valueCtx := ctx.GetValue()

	t.KeyType = t.generateTypeName(unit, keyCtx, t, t)
	t.ValueType = t.generateTypeName(unit, valueCtx, t, t)
}

func (t *TypeName) generateTypeName(sourceUnit *SourceUnit[Node[ast_pb.SourceUnit]], ctx interface{}, parentNode *TypeName, typeNameNode *TypeName) *TypeName {
	typeName := &TypeName{
		ASTBuilder: t.ASTBuilder,
		Id:         t.GetNextID(),
		NodeType:   ast_pb.NodeType_ELEMENTARY_TYPE_NAME,
	}

	switch specificCtx := ctx.(type) {
	case parser.IMappingKeyTypeContext:
		typeName.Name = specificCtx.GetText()
		typeName.Src = SrcNode{
			Id:          t.GetNextID(),
			Line:        int64(specificCtx.GetStart().GetLine()),
			Column:      int64(specificCtx.GetStart().GetColumn()),
			Start:       int64(specificCtx.GetStart().GetStart()),
			End:         int64(specificCtx.GetStop().GetStop()),
			Length:      int64(specificCtx.GetStop().GetStop() - specificCtx.GetStart().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}
		if specificCtx.ElementaryTypeName() != nil {
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				specificCtx.ElementaryTypeName().GetText(),
			)

			typeName.TypeDescription = &TypeDescription{
				TypeString:     normalizedTypeName,
				TypeIdentifier: normalizedTypeIdentifier,
			}
		}
	case parser.IMappingTypeContext:
		typeNameNode.NodeType = ast_pb.NodeType_MAPPING_TYPE_NAME
		keyCtx := specificCtx.GetKey()
		valueCtx := specificCtx.GetValue()

		typeNameNode.KeyType = t.generateTypeName(sourceUnit, keyCtx, parentNode, typeNameNode)
		typeNameNode.ValueType = t.generateTypeName(sourceUnit, valueCtx, parentNode, typeNameNode)

		if typeNameNode.KeyType != nil &&
			typeNameNode.ValueType != nil &&
			typeNameNode.KeyType.TypeDescription != nil &&
			typeNameNode.ValueType.TypeDescription != nil {
			parentNode.TypeDescription = &TypeDescription{
				TypeString: fmt.Sprintf("mapping(%s => %s)", typeNameNode.KeyType.Name, typeNameNode.ValueType.Name),
				TypeIdentifier: fmt.Sprintf(
					"t_mapping_$t_%s_$t_%s$",
					typeNameNode.KeyType.TypeDescription.TypeString,
					typeNameNode.ValueType.TypeDescription.TypeString,
				),
			}
		}
	case parser.ITypeNameContext:
		typeName.Name = specificCtx.GetText()
		typeName.Src = SrcNode{
			Id:          t.GetNextID(),
			Line:        int64(specificCtx.GetStart().GetLine()),
			Column:      int64(specificCtx.GetStart().GetColumn()),
			Start:       int64(specificCtx.GetStart().GetStart()),
			End:         int64(specificCtx.GetStop().GetStop()),
			Length:      int64(specificCtx.GetStop().GetStop() - specificCtx.GetStart().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}

		if specificCtx.ElementaryTypeName() != nil {
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				specificCtx.ElementaryTypeName().GetText(),
			)

			typeName.TypeDescription = &TypeDescription{
				TypeString:     normalizedTypeName,
				TypeIdentifier: normalizedTypeIdentifier,
			}
		} else if specificCtx.MappingType() != nil {
			typeName.NodeType = ast_pb.NodeType_MAPPING_TYPE_NAME
			t.generateTypeName(sourceUnit, specificCtx.MappingType(), parentNode, typeName)
		} else {
			// It seems to be a user defined type but that does not exist as type in parser...
			typeName.NodeType = ast_pb.NodeType_USER_DEFINED_PATH_NAME

			pathCtx := specificCtx.IdentifierPath()
			if pathCtx != nil {
				typeName.PathNode = &PathNode{
					Id:   t.GetNextID(),
					Name: pathCtx.GetText(),
					Src: SrcNode{
						Id:          t.GetNextID(),
						Line:        int64(pathCtx.GetStart().GetLine()),
						Column:      int64(pathCtx.GetStart().GetColumn()),
						Start:       int64(pathCtx.GetStart().GetStart()),
						End:         int64(pathCtx.GetStop().GetStop()),
						Length:      int64(pathCtx.GetStop().GetStop() - pathCtx.GetStart().GetStart() + 1),
						ParentIndex: typeName.GetId(),
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
				}
			}

			if ref, refTypeDescription := discoverReferenceByCtxName(t.ASTBuilder, pathCtx.GetText()); ref != nil {
				t.PathNode.ReferencedDeclaration = ref.GetId()
				t.ReferencedDeclaration = ref.GetId()
				t.TypeDescription = refTypeDescription
			}

			// Search for argument reference in state variable declarations.
			/* 			referenceFound := false

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
			   			} */
		}
	}

	return typeName
}

func (t *TypeName) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], parentNodeId int64, ctx parser.ITypeNameContext) {
	t.Id = t.GetNextID()
	t.Src = SrcNode{
		Id:          t.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNodeId,
	}

	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.ElementaryTypeNameContext:
			t.parseElementaryTypeName(unit, parentNodeId, childCtx)
		case *parser.MappingTypeContext:
			t.parseMappingTypeName(unit, parentNodeId, childCtx)
		case *parser.IdentifierPathContext:
			t.parseIdentifierPath(unit, parentNodeId, childCtx)
		default:
			panic(fmt.Sprintf("Unknown type name @ TypeName.Parse: %T", childCtx))
			//t.parseUserDefinedTypeName(unit, parentNodeId, childCtx)
		}
	}

	if ctx.Expression() != nil {
		zap.L().Warn(
			"Expression type is not supported yet @ TypeName.Parse",
			zap.String("expression", ctx.Expression().GetText()),
		)
	}
}

type PathNode struct {
	Id                    int64           `json:"id"`
	Name                  string          `json:"name"`
	NodeType              ast_pb.NodeType `json:"node_type"`
	ReferencedDeclaration int64           `json:"referenced_declaration"`
	Src                   SrcNode         `json:"src"`
}

type TypeDescription struct {
	TypeIdentifier string `json:"type_identifier"`
	TypeString     string `json:"type_string"`
}
