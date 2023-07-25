package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type TypeName struct {
	*ASTBuilder

	Id                    int64           `json:"id"`
	NodeType              ast_pb.NodeType `json:"node_type"`
	Src                   SrcNode         `json:"src"`
	Name                  string          `json:"name,omitempty"`
	TypeDescription       TypeDescription `json:"type_descriptions,omitempty"`
	KeyType               *TypeName       `json:"key_type,omitempty"`
	ValueType             *TypeName       `json:"value_type,omitempty"`
	PathNode              *PathNode       `json:"path_node,omitempty"`
	ReferencedDeclaration int64           `json:"referenced_declaration"`
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

func (t *TypeName) GetTypeDescriptions() TypeDescription {
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
	t.TypeDescription = TypeDescription{
		TypeIdentifier: normalizedTypeIdentifier,
		TypeString:     normalizedTypeName,
	}
}

func (t *TypeName) parseUserDefinedTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx antlr.Tree) {
	panic("User defined type name is not supported yet @ TypeName.Parse")
}

func (t *TypeName) parseMappingTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.MappingTypeContext) {
	panic("Mapping type name is not supported yet @ TypeName.Parse")
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
		default:
			t.parseUserDefinedTypeName(unit, parentNodeId, childCtx)
		}
	}

	/*
		if ctx.MappingType() != nil {
			zap.L().Warn(
				"Mapping type is not supported yet @ TypeName.Parse",
				zap.String("mapping_type", ctx.MappingType().GetText()),
			)
		} else if ctx.FunctionTypeName() != nil {
			zap.L().Warn(
				"Function type is not supported yet @ TypeName.Parse",
				zap.String("function_type", ctx.FunctionTypeName().GetText()),
			)
		} else {
			// It seems to be a user defined type but that does not exist as type in parser...
			t.NodeType = ast_pb.NodeType_USER_DEFINED_PATH_NAME

			pathCtx := ctx.IdentifierPath()
			if pathCtx != nil {
				pNode.TypeName.PathNode = &PathNode{
					Id:   t.GetNextID(),
					Name: pathCtx.GetText(),
					Src: SrcNode{
						Id:          t.GetNextID(),
						Line:        int64(pathCtx.GetStart().GetLine()),
						Column:      int64(pathCtx.GetStart().GetColumn()),
						Start:       int64(pathCtx.GetStart().GetStart()),
						End:         int64(pathCtx.GetStop().GetStop()),
						Length:      int64(pathCtx.GetStop().GetStop() - pathCtx.GetStart().GetStart() + 1),
						ParentIndex: t.Id,
					},
					NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
				}
			}

			// Lets figure out type...
			// Search for argument reference in state variable declarations.
			referenceFound := false

			for _, node := range t.currentStateVariables {
				if node.GetName() == pathCtx.GetText() {
					referenceFound = true
					t.PathNode.ReferencedDeclaration = node.Id
					t.ReferencedDeclaration = node.Id
					//t.TypeDescriptions = node.TypeDescriptions
					//t.TypeDescriptions = node.TypeDescriptions
				}
			}

			if !referenceFound {
				for _, node := range t.currentEnums {
					if node.GetName() == pathCtx.GetText() {
						referenceFound = true
						t.PathNode.ReferencedDeclaration = node.Id
						t.ReferencedDeclaration = node.Id

						typeDescription := &TypeDescriptions{
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
									unit.GetName(),
									pathCtx.GetText(),
								)
							}(),
						}

						t.TypeDescriptions = typeDescription
						t.TypeDescriptions = typeDescription
					}
				}
			}
		}

		if ctx.Expression() != nil {
			zap.L().Warn(
				"Expression type is not supported yet @ TypeName.Parse",
				zap.String("expression", ctx.Expression().GetText()),
			)
		} */
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
