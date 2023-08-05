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
	TypeDescription       *TypeDescription  `json:"type_description,omitempty"`
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

// SetReferenceDescriptor sets the reference descriptions of the TypeName node.
func (t *TypeName) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	t.ReferencedDeclaration = refId
	t.TypeDescription = refDesc
	return true
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

func (t *TypeName) GetStateMutability() ast_pb.Mutability {
	return t.StateMutability
}

func (t *TypeName) GetNodes() []Node[NodeType] {
	return nil
}

func (t *TypeName) ToProto() NodeType {
	toReturn := &ast_pb.TypeName{
		Id:                    t.GetId(),
		Name:                  t.GetName(),
		NodeType:              t.GetType(),
		Src:                   t.GetSrc().ToProto(),
		ReferencedDeclaration: t.ReferencedDeclaration,
		StateMutability:       t.StateMutability,
	}

	// In case it's nil, we'll do our best to resolve it later on in
	// the final resolver pass.
	// This usually means that contract A is not yet processed while
	// contract B is referencing it.
	if t.GetTypeDescription() != nil {
		toReturn.TypeDescription = t.GetTypeDescription().ToProto()
	}

	if t.GetPathNode() != nil {
		toReturn.PathNode = t.GetPathNode().ToProto()
	}

	if t.GetKeyType() != nil {
		toReturn.KeyType = t.GetKeyType().ToProto().(*ast_pb.TypeName)
	}

	if t.GetValueType() != nil {
		toReturn.ValueType = t.GetValueType().ToProto().(*ast_pb.TypeName)
	}

	return toReturn
}

func (t *TypeName) parseTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.TypeNameContext) {
	t.Name = ctx.GetText()
	t.Src = SrcNode{
		Id:          t.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNodeId,
	}

	if ctx.ElementaryTypeName() != nil {
		normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
			ctx.ElementaryTypeName().GetText(),
		)

		t.TypeDescription = &TypeDescription{
			TypeString:     normalizedTypeName,
			TypeIdentifier: normalizedTypeIdentifier,
		}
	} else if ctx.MappingType() != nil {
		t.NodeType = ast_pb.NodeType_MAPPING_TYPE_NAME
		t.generateTypeName(unit, ctx.MappingType(), t, t)
	} else if ctx.FunctionTypeName() != nil {
		panic(fmt.Sprintf("Function type name is not supported yet @ TypeName.generateTypeName: %T", ctx))
	} else {
		// It seems to be a user defined type but that does not exist as type in parser...
		t.NodeType = ast_pb.NodeType_USER_DEFINED_PATH_NAME

		pathCtx := ctx.IdentifierPath()
		if pathCtx != nil {
			t.PathNode = &PathNode{
				Id:   t.GetNextID(),
				Name: pathCtx.GetText(),
				Src: SrcNode{
					Id:          t.GetNextID(),
					Line:        int64(pathCtx.GetStart().GetLine()),
					Column:      int64(pathCtx.GetStart().GetColumn()),
					Start:       int64(pathCtx.GetStart().GetStart()),
					End:         int64(pathCtx.GetStop().GetStop()),
					Length:      int64(pathCtx.GetStop().GetStop() - pathCtx.GetStart().GetStart() + 1),
					ParentIndex: t.GetId(),
				},
				NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
			}
		}

		if refId, refTypeDescription := t.GetResolver().ResolveByNode(t, pathCtx.GetText()); refTypeDescription != nil {
			if t.PathNode != nil {
				t.PathNode.ReferencedDeclaration = refId
			}
			t.ReferencedDeclaration = refId
			t.TypeDescription = refTypeDescription
		}
	}
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

		if refId, refTypeDescription := t.GetResolver().ResolveByNode(t, identifierCtx.GetText()); refTypeDescription != nil {
			t.PathNode.ReferencedDeclaration = refId
			t.ReferencedDeclaration = refId
			t.TypeDescription = refTypeDescription
		}
	}
}

func (t *TypeName) parseMappingTypeName(unit *SourceUnit[Node[ast_pb.SourceUnit]], parentNodeId int64, ctx *parser.MappingTypeContext) {
	keyCtx := ctx.GetKey()
	valueCtx := ctx.GetValue()

	t.KeyType = t.generateTypeName(unit, keyCtx, t, t)
	t.ValueType = t.generateTypeName(unit, valueCtx, t, t)

	t.TypeDescription = &TypeDescription{
		TypeString: fmt.Sprintf("mapping(%s=>%s)", t.KeyType.Name, t.ValueType.Name),
		TypeIdentifier: fmt.Sprintf(
			"t_mapping_$%s_$%s$",
			t.KeyType.TypeDescription.TypeIdentifier,
			t.ValueType.TypeDescription.TypeIdentifier,
		),
	}
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
		typeNameNode.TypeDescription = &TypeDescription{
			TypeString: fmt.Sprintf("mapping(%s=>%s)", typeNameNode.KeyType.Name, typeNameNode.ValueType.Name),
			TypeIdentifier: fmt.Sprintf(
				"t_mapping_$%s_$%s",
				typeNameNode.KeyType.TypeDescription.TypeIdentifier,
				typeNameNode.ValueType.TypeDescription.TypeIdentifier,
			),
		}
		parentNode.TypeDescription = t.TypeDescription
		typeName = typeNameNode
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
		} else if specificCtx.FunctionTypeName() != nil {
			panic(fmt.Sprintf("Function type name is not supported yet @ TypeName.generateTypeName: %T", specificCtx))
		} else {
			t.parseTypeName(sourceUnit, parentNode.GetId(), specificCtx.(*parser.TypeNameContext))
		}
	}

	// We're still not able to discover reference, so what we're going to do now is look for the references...
	if typeName.TypeDescription == nil {
		if refId, refTypeDescription := t.GetResolver().ResolveByNode(typeName, typeName.Name); refTypeDescription != nil {
			typeName.ReferencedDeclaration = refId
			typeName.TypeDescription = refTypeDescription
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
		case *parser.TypeNameContext:
			t.parseTypeName(unit, parentNodeId, childCtx)
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

func (pn *PathNode) ToProto() *ast_pb.PathNode {
	return &ast_pb.PathNode{
		Id:                    pn.Id,
		Name:                  pn.Name,
		NodeType:              pn.NodeType,
		ReferencedDeclaration: pn.ReferencedDeclaration,
		Src:                   pn.Src.ToProto(),
	}
}

type TypeDescription struct {
	TypeIdentifier string `json:"type_identifier"`
	TypeString     string `json:"type_string"`
}

func (td TypeDescription) ToProto() *ast_pb.TypeDescription {
	return &ast_pb.TypeDescription{
		TypeString:     td.TypeString,
		TypeIdentifier: td.TypeIdentifier,
	}
}
