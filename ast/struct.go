package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type StructDefinition struct {
	*ASTBuilder

	SourceUnitName        string                 `json:"-"`
	Id                    int64                  `json:"id"`
	NodeType              ast_pb.NodeType        `json:"node_type"`
	Src                   SrcNode                `json:"src"`
	Kind                  ast_pb.NodeType        `json:"kind,omitempty"`
	Name                  string                 `json:"name"`
	CanonicalName         string                 `json:"canonical_name"`
	ReferencedDeclaration int64                  `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription       `json:"type_description"`
	Members               []*Parameter           `json:"members"`
	Visibility            ast_pb.Visibility      `json:"visibility"`
	StorageLocation       ast_pb.StorageLocation `json:"storage_location"`
}

func NewStructDefinition(b *ASTBuilder) *StructDefinition {
	return &StructDefinition{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_STRUCT_DEFINITION,
		Visibility:      ast_pb.Visibility_PUBLIC,
		StorageLocation: ast_pb.StorageLocation_DEFAULT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the StructDefinition node.
func (s *StructDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	s.ReferencedDeclaration = refId
	s.TypeDescription = refDesc
	return false
}

func (s *StructDefinition) GetId() int64 {
	return s.Id
}

func (s *StructDefinition) GetType() ast_pb.NodeType {
	return s.NodeType
}

func (s *StructDefinition) GetSrc() SrcNode {
	return s.Src
}

func (s *StructDefinition) GetName() string {
	return s.Name
}

func (s *StructDefinition) GetTypeDescription() *TypeDescription {
	return s.TypeDescription
}

func (s *StructDefinition) GetCanonicalName() string {
	return s.CanonicalName
}

func (s *StructDefinition) GetMembers() []*Parameter {
	return s.Members
}

func (s *StructDefinition) GetSourceUnitName() string {
	return s.SourceUnitName
}

func (s *StructDefinition) GetKind() ast_pb.NodeType {
	return s.Kind
}

func (s *StructDefinition) GetVisibility() ast_pb.Visibility {
	return s.Visibility
}

func (s *StructDefinition) GetStorageLocation() ast_pb.StorageLocation {
	return s.StorageLocation
}

func (s *StructDefinition) GetNodes() []Node[NodeType] {
	return nil
}

func (s *StructDefinition) ToProto() NodeType {
	return ast_pb.Struct{}
}

func (s *StructDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.StructDefinitionContext,
) Node[NodeType] {
	s.Src = SrcNode{
		Id:          s.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: unit.GetId(),
	}
	s.SourceUnitName = unit.GetName()

	s.Name = ctx.GetName().GetText()
	s.CanonicalName = fmt.Sprintf("%s.%s", s.SourceUnitName, s.Name)

	s.TypeDescription = &TypeDescription{
		TypeIdentifier: fmt.Sprintf(
			"t_struct$_%s_%s_$%d", s.SourceUnitName, s.GetName(), s.GetId(),
		),
		TypeString: fmt.Sprintf(
			"struct %s.%s", s.SourceUnitName, s.GetName(),
		),
	}

	for _, memberCtx := range ctx.AllStructMember() {
		parameter := NewParameter(s.ASTBuilder)
		parameter.ParseStructParameter(unit, contractNode, s, memberCtx)
		s.Members = append(s.Members, parameter)
	}

	s.currentStructs = append(s.currentStructs, s)
	return s
}
