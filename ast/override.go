package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// OverridePath represents an override path node in the abstract syntax tree.
type OverridePath struct {
	Id                    int64            `json:"id"`                     // Unique identifier of the modifier name node.
	Name                  string           `json:"name"`                   // Name of the modifier.
	NodeType              ast_pb.NodeType  `json:"node_type"`              // Type of the node.
	Src                   SrcNode          `json:"src"`                    // Source location information.
	ReferencedDeclaration int64            `json:"referenced_declaration"` // Referenced declaration identifier.
	TypeDescription       *TypeDescription `json:"type_description"`       // Type description of the override specifier.
}

// SetReferenceDescriptor sets the reference descriptor of the OverridePath.
func (op *OverridePath) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	op.ReferencedDeclaration = refId
	op.TypeDescription = refDesc
	return true
}

// GetId returns the unique identifier of the override path node.
func (m *OverridePath) GetId() int64 {
	return m.Id
}

// GetName returns the name of the override path node.
func (m *OverridePath) GetName() string {
	return m.Name
}

// GetType returns the type of the node.
func (m *OverridePath) GetType() ast_pb.NodeType {
	return m.NodeType
}

// GetSrc returns the source location information of the override path node.
func (m *OverridePath) GetSrc() SrcNode {
	return m.Src
}

// GetNodes returns an empty slice of nodes associated with the override path.
func (m *OverridePath) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetTypeDescription returns the type description of the override path.
func (m *OverridePath) GetTypeDescription() *TypeDescription {
	return m.TypeDescription
}

// GetReferencedDeclaration returns the referenced declaration identifier of the override path.
func (m *OverridePath) GetReferencedDeclaration() int64 {
	return m.ReferencedDeclaration
}

// ToProto converts the OverridePath node to its corresponding protobuf representation.
func (m *OverridePath) ToProto() *ast_pb.OverridePath {
	return &ast_pb.OverridePath{
		Id:                    m.Id,
		Name:                  m.Name,
		NodeType:              m.NodeType,
		Src:                   m.Src.ToProto(),
		ReferencedDeclaration: m.ReferencedDeclaration,
		TypeDescription:       m.TypeDescription.ToProto(),
	}
}

// OverrideSpecifier represents an override specifier node in the abstract syntax tree.
type OverrideSpecifier struct {
	*ASTBuilder

	Id                    int64            `json:"id"`                     // Unique identifier of the override specifier node.
	NodeType              ast_pb.NodeType  `json:"node_type"`              // Type of the node.
	Src                   SrcNode          `json:"src"`                    // Source location information.
	Overrides             []*OverridePath  `json:"overrides"`              // List of override paths.
	ReferencedDeclaration int64            `json:"referenced_declaration"` // Referenced declaration identifier.
	TypeDescription       *TypeDescription `json:"type_descriptions"`      // Type description of the override specifier.
}

// NewOverrideSpecifier creates a new instance of OverrideSpecifier with the provided ASTBuilder.
func NewOverrideSpecifier(b *ASTBuilder) *OverrideSpecifier {
	return &OverrideSpecifier{
		ASTBuilder: b,
		Overrides:  make([]*OverridePath, 0),
		TypeDescription: &TypeDescription{
			TypeString:     "override",
			TypeIdentifier: "$_t_override",
		},
	}
}

// SetReferenceDescriptor sets the reference descriptor of the OverrideSpecifier.
func (o *OverrideSpecifier) SetReferenceDescriptor(refId int64, refType *TypeDescription) bool {
	o.ReferencedDeclaration = refId
	o.TypeDescription = refType
	return true
}

// GetId returns the unique identifier of the override specifier node.
func (o *OverrideSpecifier) GetId() int64 {
	return o.Id
}

// GetType returns the type of the node.
func (o *OverrideSpecifier) GetType() ast_pb.NodeType {
	return o.NodeType
}

// GetSrc returns the source location information of the override specifier node.
func (o *OverrideSpecifier) GetSrc() SrcNode {
	return o.Src
}

// GetNodes returns an empty slice of nodes associated with the override specifier.
func (o *OverrideSpecifier) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	return toReturn
}

// GetTypeDescription returns the type description of the override specifier.
func (o *OverrideSpecifier) GetTypeDescription() *TypeDescription {
	return o.TypeDescription
}

// GetReferencedDeclaration returns the referenced declaration identifier of the override specifier.
func (o *OverrideSpecifier) GetReferencedDeclaration() int64 {
	return o.ReferencedDeclaration
}

// GetOverrides returns the list of override paths.
func (o *OverrideSpecifier) GetOverrides() []*OverridePath {
	return o.Overrides
}

// ToProto converts the OverrideSpecifier node to its corresponding protobuf representation.
func (o *OverrideSpecifier) ToProto() NodeType {
	proto := &ast_pb.OverrideSpecifier{
		Id:                    o.GetId(),
		NodeType:              o.GetType(),
		Src:                   o.GetSrc().ToProto(),
		ReferencedDeclaration: o.GetReferencedDeclaration(),
		TypeDescription:       o.GetTypeDescription().ToProto(),
	}

	for _, override := range o.GetOverrides() {
		proto.Overrides = append(proto.Overrides, override.ToProto())
	}

	return proto
}

// Parse parses the override specifier context and populates the OverrideSpecifier fields.
func (o *OverrideSpecifier) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], ctx parser.IOverrideSpecifierContext) {
	o.Id = o.GetNextID()
	o.Src = SrcNode{
		Id:          o.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}
	o.NodeType = ast_pb.NodeType_OVERRIDE_SPECIFIER

	// Not yet able to figure this part out entirely...
	// @TODO: Figure out how to parse this, as override can be specified, without any paths which
	// is what the hell to figure out what actual function name is...
	// Findings afer this note is that these are derived contracts.
	if ctx.GetOverrides() != nil {
		for _, override := range ctx.GetOverrides() {
			oId := o.GetNextID()
			for _, identifier := range override.AllIdentifier() {
				bO := &OverridePath{
					Id:       oId,
					Name:     identifier.GetText(),
					NodeType: ast_pb.NodeType_OVERRIDE_PATH,
					Src: SrcNode{
						Id:          o.GetNextID(),
						Line:        int64(override.GetStart().GetLine()),
						Column:      int64(override.GetStart().GetColumn()),
						Start:       int64(override.GetStart().GetStart()),
						End:         int64(override.GetStop().GetStop()),
						Length:      int64(override.GetStop().GetStop() - override.GetStart().GetStart() + 1),
						ParentIndex: o.Id,
					},
				}

				if refId, refTypeDescription := o.GetResolver().ResolveByNode(o, identifier.GetText()); refTypeDescription != nil {
					bO.ReferencedDeclaration = refId
					bO.TypeDescription = refTypeDescription
				}

				o.Overrides = append(o.Overrides, bO)
			}
		}
	}
}
