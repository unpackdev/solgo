package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
	"go.uber.org/zap"
)

// OverrideSpecifier represents an override specifier node in the abstract syntax tree.
type OverrideSpecifier struct {
	*ASTBuilder

	Id                    int64            `json:"id"`                     // Unique identifier of the override specifier node.
	NodeType              ast_pb.NodeType  `json:"node_type"`              // Type of the node.
	Name                  string           `json:"name"`                   // Name of the overridden identifier.
	Src                   SrcNode          `json:"src"`                    // Source location information.
	ReferencedDeclaration int64            `json:"referenced_declaration"` // Referenced declaration identifier.
	TypeDescription       *TypeDescription `json:"type_descriptions"`      // Type description of the override specifier.
}

// NewOverrideSpecifier creates a new instance of OverrideSpecifier with the provided ASTBuilder.
func NewOverrideSpecifier(b *ASTBuilder) *OverrideSpecifier {
	return &OverrideSpecifier{
		ASTBuilder: b,
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
	return []Node[NodeType]{}
}

// GetTypeDescription returns the type description of the override specifier.
func (o *OverrideSpecifier) GetTypeDescription() *TypeDescription {
	return o.TypeDescription
}

// GetReferencedDeclaration returns the referenced declaration identifier of the override specifier.
func (o *OverrideSpecifier) GetReferencedDeclaration() int64 {
	return o.ReferencedDeclaration
}

// GetName returns the name of the identifier that is being overridden.
func (o *OverrideSpecifier) GetName() string {
	return o.Name
}

// ToProto converts the OverrideSpecifier node to its corresponding protobuf representation.
func (o *OverrideSpecifier) ToProto() NodeType {
	return &ast_pb.OverrideSpecifier{
		Id:                    o.GetId(),
		Name:                  o.GetName(),
		NodeType:              o.GetType(),
		Src:                   o.GetSrc().ToProto(),
		ReferencedDeclaration: o.GetReferencedDeclaration(),
		TypeDescription:       o.GetTypeDescription().ToProto(),
	}
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
	for _, overrides := range ctx.GetOverrides() {
		zap.L().Warn(
			"Override specifier overrides not implemented",
			zap.String("identifier_text", overrides.GetText()),
		)
	}

	if ctx.AllIdentifierPath() != nil {
		for _, pathCtx := range ctx.AllIdentifierPath() {
			for _, identifierCtx := range pathCtx.AllIdentifier() {
				zap.L().Warn(
					"Override specifier identifier path not implemented",
					zap.String("identifier_text", identifierCtx.GetText()),
				)
			}
		}
	}

	// Figure out what is the function name if no derived contracts are set and forward search resolution of
	// referenced declaration id and its type.
	if o.ReferencedDeclaration == 0 {
		for _, child := range ctx.GetRuleContext().GetParent().GetChildren() {
			switch childCtx := child.(type) {
			case *parser.IdentifierContext:
				o.Name = childCtx.GetText()
				if refId, refTypeDescription := o.GetResolver().ResolveByNode(o, childCtx.GetText()); refTypeDescription != nil {
					o.ReferencedDeclaration = refId
					o.TypeDescription = refTypeDescription
				}
			}
		}
	}
}
