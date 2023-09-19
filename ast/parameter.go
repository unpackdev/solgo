package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Parameter represents a parameter node in the abstract syntax tree.
type Parameter struct {
	*ASTBuilder

	Id              int64                  `json:"id"`                         // Unique identifier of the parameter node.
	NodeType        ast_pb.NodeType        `json:"node_type"`                  // Type of the node.
	Src             SrcNode                `json:"src"`                        // Source location information.
	NameLocation    *SrcNode               `json:"name_location,omitempty"`    // Source location information of the name.
	Scope           int64                  `json:"scope,omitempty"`            // Scope of the parameter.
	Name            string                 `json:"name"`                       // Name of the parameter.
	TypeName        *TypeName              `json:"type_name,omitempty"`        // Type name of the parameter.
	StorageLocation ast_pb.StorageLocation `json:"storage_location,omitempty"` // Storage location of the parameter.
	Visibility      ast_pb.Visibility      `json:"visibility,omitempty"`       // Visibility of the parameter.
	StateMutability ast_pb.Mutability      `json:"state_mutability,omitempty"` // State mutability of the parameter.
	Constant        bool                   `json:"constant,omitempty"`         // Whether the parameter is constant.
	StateVariable   bool                   `json:"state_variable,omitempty"`   // Whether the parameter is a state variable.
	TypeDescription *TypeDescription       `json:"type_description,omitempty"` // Type description of the parameter.
	Indexed         bool                   `json:"indexed,omitempty"`          // Whether the parameter is indexed.
}

// NewParameter creates a new instance of Parameter with the provided ASTBuilder.
func NewParameter(b *ASTBuilder) *Parameter {
	return &Parameter{
		ASTBuilder:      b,
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		Visibility:      ast_pb.Visibility_INTERNAL,
		StateMutability: ast_pb.Mutability_MUTABLE,
	}
}

// SetReferenceDescriptor sets the reference descriptors of the Parameter node.
func (p *Parameter) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the parameter node.
func (p *Parameter) GetId() int64 {
	return p.Id
}

// GetType returns the type of the node.
func (p *Parameter) GetType() ast_pb.NodeType {
	return p.NodeType
}

// GetSrc returns the source location information of the parameter node.
func (p *Parameter) GetSrc() SrcNode {
	return p.Src
}

// GetNameLocation returns the source location information of the name of the parameter.
func (p *Parameter) GetNameLocation() *SrcNode {
	return p.NameLocation
}

// GetName returns the name of the parameter.
func (p *Parameter) GetName() string {
	return p.Name
}

// GetScope returns the scope of the parameter.
func (p *Parameter) GetScope() int64 {
	return p.Id
}

// GetTypeDescription returns the type description of the parameter.
func (p *Parameter) GetTypeDescription() *TypeDescription {
	if p.TypeName != nil {
		return p.TypeName.TypeDescription
	}
	return p.TypeDescription
}

// GetVisibility returns the visibility of the parameter.
func (p *Parameter) GetVisibility() ast_pb.Visibility {
	return p.Visibility
}

// GetStateMutability returns the state mutability of the parameter.
func (p *Parameter) GetStateMutability() ast_pb.Mutability {
	return p.StateMutability
}

// GetStorageLocation returns the storage location of the parameter.
func (p *Parameter) GetStorageLocation() ast_pb.StorageLocation {
	return p.StorageLocation
}

// IsConstant returns whether the parameter is constant.
func (p *Parameter) IsConstant() bool {
	return p.Constant
}

// IsStateVariable returns whether the parameter is a state variable.
func (p *Parameter) IsStateVariable() bool {
	return p.StateVariable
}

// GetTypeName returns the type name of the parameter.
func (p *Parameter) GetTypeName() *TypeName {
	return p.TypeName
}

// IsIndexed returns whether the parameter is indexed.
func (p *Parameter) IsIndexed() bool {
	return p.Indexed
}

// GetNodes returns a slice of nodes associated with the parameter.
func (p *Parameter) GetNodes() []Node[NodeType] {
	if p.TypeName != nil {
		return []Node[NodeType]{p.TypeName}
	}

	return []Node[NodeType]{}
}

// ToProto converts the Parameter node to its corresponding protobuf representation.
func (p *Parameter) ToProto() NodeType {
	toReturn := &ast_pb.Parameter{
		Id:              p.GetId(),
		Name:            p.GetName(),
		NodeType:        p.GetType(),
		Src:             p.GetSrc().ToProto(),
		Scope:           p.GetScope(),
		Constant:        p.IsConstant(),
		StateVariable:   p.IsStateVariable(),
		StateMutability: p.GetStateMutability(),
		Visibility:      p.GetVisibility(),
		StorageLocation: p.GetStorageLocation(),
		Indexed:         p.IsIndexed(),
	}

	if p.GetTypeName() != nil {
		toReturn.TypeName = p.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	if p.GetTypeDescription() != nil {
		toReturn.TypeDescription = p.GetTypeDescription().ToProto()
	}

	if p.GetNameLocation() != nil {
		toReturn.NameLocation = p.GetNameLocation().ToProto()
	}

	return toReturn
}

// Parse parses the parameter declaration context and populates the Parameter fields.
func (p *Parameter) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx *parser.ParameterDeclarationContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: plNode.GetId(),
	}
	p.Scope = fnNode.GetId()

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	p.StorageLocation = p.getStorageLocationFromCtx(ctx)

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, fnNode, p.GetId(), ctx.TypeName())

	if typeName.TypeDescription != nil {
		switch typeName.TypeDescription.TypeIdentifier {
		case "t_address":
			p.StateMutability = ast_pb.Mutability_NONPAYABLE
		case "t_address_payable":
			p.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	p.TypeName = typeName
	p.currentVariables = append(p.currentVariables, p)
}

// ParseEventParameter parses the event parameter context and populates the Parameter fields for event parameters.
func (p *Parameter) ParseEventParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx parser.IEventParameterContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: plNode.GetId(),
	}
	p.Scope = fnNode.GetId()
	p.Indexed = ctx.Indexed() != nil

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	p.StorageLocation = ast_pb.StorageLocation_MEMORY

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, fnNode, p.GetId(), ctx.TypeName())

	if typeName.TypeDescription != nil {
		switch typeName.TypeDescription.TypeIdentifier {
		case "t_address":
			p.StateMutability = ast_pb.Mutability_NONPAYABLE
		case "t_address_payable":
			p.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	p.TypeName = typeName
	p.currentVariables = append(p.currentVariables, p)
}

// ParseStructParameter parses the struct parameter context and populates the Parameter fields for struct members.
func (p *Parameter) ParseStructParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], contractNode Node[NodeType], structNode *StructDefinition, ctx parser.IStructMemberContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: structNode.GetId(),
	}
	p.Scope = contractNode.GetId()

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, contractNode, p.GetId(), ctx.TypeName())

	if typeName.TypeDescription != nil {
		switch typeName.TypeDescription.TypeIdentifier {
		case "t_address":
			p.StateMutability = ast_pb.Mutability_NONPAYABLE
		case "t_address_payable":
			p.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	p.TypeName = typeName
	p.currentVariables = append(p.currentVariables, p)
}

// ParseErrorParameter parses the error parameter context and populates the Parameter fields for error definitions.
func (p *Parameter) ParseErrorParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx parser.IErrorParameterContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: plNode.GetId(),
	}
	p.Scope = fnNode.GetId()

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	p.StorageLocation = ast_pb.StorageLocation_MEMORY

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, fnNode, p.GetId(), ctx.TypeName())

	if typeName.TypeDescription != nil {
		switch typeName.TypeDescription.TypeIdentifier {
		case "t_address":
			p.StateMutability = ast_pb.Mutability_NONPAYABLE
		case "t_address_payable":
			p.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	p.TypeName = typeName
	p.currentVariables = append(p.currentVariables, p)
}

// getStorageLocationFromCtx extracts the storage location information from the parameter declaration context.
func (p *Parameter) getStorageLocationFromCtx(ctx *parser.ParameterDeclarationContext) ast_pb.StorageLocation {
	storageLocationMap := map[string]ast_pb.StorageLocation{
		"memory":   ast_pb.StorageLocation_MEMORY,
		"storage":  ast_pb.StorageLocation_STORAGE,
		"calldata": ast_pb.StorageLocation_CALLDATA,
	}

	if ctx.DataLocation() != nil {
		if s, ok := storageLocationMap[ctx.DataLocation().GetText()]; ok {
			return s
		}
	}

	return ast_pb.StorageLocation_MEMORY
}
