package ast

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type EnumDefinition struct {
	*ASTBuilder
	SourceUnitName  string           `json:"-"`
	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"node_type"`
	Src             SrcNode          `json:"src"`
	Name            string           `json:"name"`
	CanonicalName   string           `json:"canonical_name"`
	TypeDescription *TypeDescription `json:"type_description"`
	Members         []*Parameter     `json:"members"`
}

func NewEnumDefinition(b *ASTBuilder) *EnumDefinition {
	return &EnumDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_ENUM_DEFINITION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the EnumDefinition node.
// We don't need to do any reference description updates here, at least for now...
func (e *EnumDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (e *EnumDefinition) GetId() int64 {
	return e.Id
}

func (e *EnumDefinition) GetType() ast_pb.NodeType {
	return e.NodeType
}

func (e *EnumDefinition) GetSrc() SrcNode {
	return e.Src
}

func (e *EnumDefinition) GetName() string {
	return e.Name
}

func (e *EnumDefinition) GetTypeDescription() *TypeDescription {
	return e.TypeDescription
}

func (e *EnumDefinition) GetCanonicalName() string {
	return e.CanonicalName
}

func (e *EnumDefinition) GetMembers() []*Parameter {
	return e.Members
}

func (e *EnumDefinition) GetSourceUnitName() string {
	return e.SourceUnitName
}

func (e *EnumDefinition) ToProto() NodeType {
	proto := ast_pb.Enum{
		Id:              e.GetId(),
		Name:            e.GetName(),
		CanonicalName:   e.GetCanonicalName(),
		NodeType:        e.GetType(),
		Src:             e.GetSrc().ToProto(),
		Members:         make([]*ast_pb.Parameter, 0),
		TypeDescription: e.GetTypeDescription().ToProto(),
	}

	for _, member := range e.GetMembers() {
		proto.Members = append(proto.Members, member.ToProto())
	}

	// Marshal the Pragma into JSON
	jsonBytes, err := protojson.Marshal(&proto)
	if err != nil {
		panic(err)
	}

	s := &structpb.Struct{}
	if err := protojson.Unmarshal(jsonBytes, s); err != nil {
		panic(err)
	}

	return &v3.TypedStruct{
		TypeUrl: "github.com/txpull/protos/txpull.v1.ast.Enum",
		Value:   s,
	}
}

func (e *EnumDefinition) GetNodes() []Node[NodeType] {
	return nil
}

func (e *EnumDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.EnumDefinitionContext,
) Node[NodeType] {
	e.Src = SrcNode{
		Id:          e.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart()),
		ParentIndex: contractNode.GetId(),
	}
	e.SourceUnitName = unit.GetName()
	e.Name = ctx.GetName().GetText()
	e.CanonicalName = fmt.Sprintf("%s.%s", unit.GetName(), e.Name)
	e.TypeDescription = &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_enum_$_%s_$%d", e.Name, e.Id),
		TypeString:     fmt.Sprintf("enum %s", e.CanonicalName),
	}

	for _, enumCtx := range ctx.GetEnumValues() {
		e.Members = append(
			e.Members,
			&Parameter{
				Id: e.GetNextID(),
				Src: SrcNode{
					Line:        int64(enumCtx.GetStart().GetLine()),
					Column:      int64(enumCtx.GetStart().GetColumn()),
					Start:       int64(enumCtx.GetStart().GetStart()),
					End:         int64(enumCtx.GetStop().GetStop()),
					Length:      int64(enumCtx.GetStop().GetStop() - enumCtx.GetStart().GetStart()),
					ParentIndex: e.Id,
				},
				Name:     enumCtx.GetText(),
				NodeType: ast_pb.NodeType_ENUM_VALUE,
			},
		)
	}

	e.currentEnums = append(e.currentEnums, e)

	return e
}

/**
func (b *ASTBuilder) parseEnumDefinition(sourceUnit *ast_pb.SourceUnit, enumNode *ast_pb.Node, ctx *parser.EnumDefinitionContext) *ast_pb.Node {
	enumNode.NodeType = ast_pb.NodeType_ENUM_DEFINITION
	enumNode.Name = ctx.GetName().GetText()
	enumNode.CanonicalName = fmt.Sprintf("%s.%s", sourceUnit.Name, enumNode.Name)

	for _, enumCtx := range ctx.GetEnumValues() {
		enumNode.Members = append(
			enumNode.Members,
			&ast_pb.Parameter{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(enumCtx.GetStart().GetLine()),
					Column:      int64(enumCtx.GetStart().GetColumn()),
					Start:       int64(enumCtx.GetStart().GetStart()),
					End:         int64(enumCtx.GetStop().GetStop()),
					Length:      int64(enumCtx.GetStop().GetStop() - enumCtx.GetStart().GetStart()),
					ParentIndex: enumNode.Id,
				},
				Name:     enumCtx.GetText(),
				NodeType: ast_pb.NodeType_ENUM_VALUE,
			},
		)
	}

	enumNode.TypeDescriptions = &ast_pb.TypeDescriptions{
		TypeIdentifier: func() string {
			return fmt.Sprintf(
				"t_enum_$_%s_$%d",
				enumNode.Name,
				enumNode.Id,
			)
		}(),
		TypeString: func() string {
			return fmt.Sprintf(
				"enum %s.%s",
				sourceUnit.GetName(),
				enumNode.Name,
			)
		}(),
	}

	b.currentEnums = append(b.currentEnums, enumNode)

	return enumNode
}
**/
