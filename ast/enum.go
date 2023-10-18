package ast

import (
	"encoding/json"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// EnumDefinition represents an enumeration definition in the Solidity abstract syntax tree (AST).
type EnumDefinition struct {
	*ASTBuilder                      // Embedding the ASTBuilder for common functionality
	SourceUnitName  string           `json:"-"`
	Id              int64            `json:"id"`               // Unique identifier for the enumeration definition
	NodeType        ast_pb.NodeType  `json:"node_type"`        // Type of the node (ENUM_DEFINITION for enumeration definition)
	Src             SrcNode          `json:"src"`              // Source information about the enumeration definition
	NameLocation    SrcNode          `json:"name_location"`    // Source information about the name of the enumeration
	Name            string           `json:"name"`             // Name of the enumeration
	CanonicalName   string           `json:"canonical_name"`   // Canonical name of the enumeration
	TypeDescription *TypeDescription `json:"type_description"` // Type description of the enumeration
	Members         []Node[NodeType] `json:"members"`          // Members of the enumeration
}

// NewEnumDefinition creates a new EnumDefinition instance.
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

// GetId returns the unique identifier of the enumeration definition.
func (e *EnumDefinition) GetId() int64 {
	return e.Id
}

// GetType returns the type of the node, which is 'ENUM_DEFINITION' for an enumeration definition.
func (e *EnumDefinition) GetType() ast_pb.NodeType {
	return e.NodeType
}

// GetSrc returns the source information about the enumeration definition.
func (e *EnumDefinition) GetSrc() SrcNode {
	return e.Src
}

// GetNameLocation returns the source information about the name of the enumeration.
func (e *EnumDefinition) GetNameLocation() SrcNode {
	return e.NameLocation
}

// GetName returns the name of the enumeration.
func (e *EnumDefinition) GetName() string {
	return e.Name
}

// GetTypeDescription returns the type description of the enumeration.
func (e *EnumDefinition) GetTypeDescription() *TypeDescription {
	return e.TypeDescription
}

// GetCanonicalName returns the canonical name of the enumeration.
func (e *EnumDefinition) GetCanonicalName() string {
	return e.CanonicalName
}

// GetMembers returns the members of the enumeration.
func (e *EnumDefinition) GetMembers() []*Parameter {
	toReturn := make([]*Parameter, 0)

	for _, member := range e.Members {
		toReturn = append(toReturn, member.(*Parameter))
	}

	return toReturn
}

// GetSourceUnitName returns the name of the source unit containing the enumeration.
func (e *EnumDefinition) GetSourceUnitName() string {
	return e.SourceUnitName
}

// ToProto returns the protobuf representation of the enumeration definition.
func (e *EnumDefinition) ToProto() NodeType {
	proto := ast_pb.Enum{
		Id:              e.GetId(),
		Name:            e.GetName(),
		CanonicalName:   e.GetCanonicalName(),
		NodeType:        e.GetType(),
		Src:             e.GetSrc().ToProto(),
		NameLocation:    e.GetNameLocation().ToProto(),
		Members:         make([]*ast_pb.Parameter, 0),
		TypeDescription: e.GetTypeDescription().ToProto(),
	}

	for _, member := range e.GetMembers() {
		proto.Members = append(
			proto.Members,
			member.ToProto().(*ast_pb.Parameter),
		)
	}

	return NewTypedStruct(&proto, "Enum")
}

// GetNodes returns the members of the enumeration.
func (e *EnumDefinition) GetNodes() []Node[NodeType] {
	return e.Members
}

func (e *EnumDefinition) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["id"], &e.Id); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["node_type"], &e.NodeType); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["name"], &e.Name); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["canonical_name"], &e.CanonicalName); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["type_description"], &e.TypeDescription); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["src"], &e.Src); err != nil {
		return err
	}

	if err := json.Unmarshal(tempMap["name_location"], &e.NameLocation); err != nil {
		return err
	}

	var tempMembers []json.RawMessage
	if err := json.Unmarshal(tempMap["members"], &tempMembers); err != nil {
		return err
	}

	for _, tempMember := range tempMembers {
		var tempMemberMap map[string]json.RawMessage
		if err := json.Unmarshal(tempMember, &tempMemberMap); err != nil {
			return err
		}

		var tempMemberType ast_pb.NodeType
		if err := json.Unmarshal(tempMemberMap["node_type"], &tempMemberType); err != nil {
			return err
		}

		node, err := unmarshalNode(tempMember, tempMemberType)
		if err != nil {
			return err
		}

		e.Members = append(e.Members, node)
	}

	return nil
}

// Parse parses an enumeration definition from the provided parser.EnumDefinitionContext and updates the current instance.
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
	e.NameLocation = SrcNode{
		Line:        int64(ctx.GetName().GetStart().GetLine()),
		Column:      int64(ctx.GetName().GetStart().GetColumn()),
		Start:       int64(ctx.GetName().GetStart().GetStart()),
		End:         int64(ctx.GetName().GetStop().GetStop()),
		Length:      int64(ctx.GetName().GetStop().GetStop() - ctx.GetName().GetStart().GetStart() + 1),
		ParentIndex: e.Id,
	}
	e.CanonicalName = fmt.Sprintf("%s.%s", unit.GetName(), e.Name)
	e.TypeDescription = &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_enum_$_%s_$%d", e.Name, e.Id),
		TypeString:     fmt.Sprintf("enum %s", e.CanonicalName),
	}

	for _, enumCtx := range ctx.GetEnumValues() {
		id := e.GetNextID()
		e.Members = append(
			e.Members,
			&Parameter{
				Id: id,
				Src: SrcNode{
					Id:          e.GetNextID(),
					Line:        int64(enumCtx.GetStart().GetLine()),
					Column:      int64(enumCtx.GetStart().GetColumn()),
					Start:       int64(enumCtx.GetStart().GetStart()),
					End:         int64(enumCtx.GetStop().GetStop()),
					Length:      int64(enumCtx.GetStop().GetStop() - enumCtx.GetStart().GetStart()),
					ParentIndex: e.Id,
				},
				Name: enumCtx.GetText(),
				NameLocation: &SrcNode{
					Id:          e.GetNextID(),
					Line:        int64(enumCtx.Identifier().GetSymbol().GetLine()),
					Column:      int64(enumCtx.Identifier().GetSymbol().GetColumn()),
					Start:       int64(enumCtx.Identifier().GetSymbol().GetStart()),
					End:         int64(enumCtx.Identifier().GetSymbol().GetStop()),
					Length:      int64(enumCtx.Identifier().GetSymbol().GetStop() - enumCtx.Identifier().GetSymbol().GetStart() + 1),
					ParentIndex: e.Id,
				},
				NodeType: ast_pb.NodeType_ENUM_VALUE,
				TypeDescription: &TypeDescription{
					TypeIdentifier: fmt.Sprintf("t_enum_$_%s$_%s_$%d", e.Name, enumCtx.GetText(), id),
					TypeString:     fmt.Sprintf("enum %s.%s", e.CanonicalName, enumCtx.GetText()),
				},
			},
		)
	}
	e.currentEnums = append(e.currentEnums, e)

	return e
}

// ParseGlobal parses an global enumeration definition from the provided parser.EnumDefinitionContext and updates the current instance.
func (e *EnumDefinition) ParseGlobal(
	ctx *parser.EnumDefinitionContext,
) Node[NodeType] {
	e.Src = SrcNode{
		Id:     e.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart()),
	}
	e.SourceUnitName = "Global"
	e.Name = ctx.GetName().GetText()
	e.NameLocation = SrcNode{
		Line:        int64(ctx.GetName().GetStart().GetLine()),
		Column:      int64(ctx.GetName().GetStart().GetColumn()),
		Start:       int64(ctx.GetName().GetStart().GetStart()),
		End:         int64(ctx.GetName().GetStop().GetStop()),
		Length:      int64(ctx.GetName().GetStop().GetStop() - ctx.GetName().GetStart().GetStart() + 1),
		ParentIndex: e.Id,
	}
	e.CanonicalName = fmt.Sprintf("%s.%s", "Global", e.Name)
	e.TypeDescription = &TypeDescription{
		TypeIdentifier: fmt.Sprintf("t_enum_$_%s_$%d", e.Name, e.Id),
		TypeString:     fmt.Sprintf("enum %s", e.CanonicalName),
	}

	for _, enumCtx := range ctx.GetEnumValues() {
		id := e.GetNextID()
		e.Members = append(
			e.Members,
			&Parameter{
				Id: id,
				Src: SrcNode{
					Line:        int64(enumCtx.GetStart().GetLine()),
					Column:      int64(enumCtx.GetStart().GetColumn()),
					Start:       int64(enumCtx.GetStart().GetStart()),
					End:         int64(enumCtx.GetStop().GetStop()),
					Length:      int64(enumCtx.GetStop().GetStop() - enumCtx.GetStart().GetStart()),
					ParentIndex: e.Id,
				},
				Name: enumCtx.GetText(),
				NameLocation: &SrcNode{
					Line:        int64(enumCtx.Identifier().GetSymbol().GetLine()),
					Column:      int64(enumCtx.Identifier().GetSymbol().GetColumn()),
					Start:       int64(enumCtx.Identifier().GetSymbol().GetStart()),
					End:         int64(enumCtx.Identifier().GetSymbol().GetStop()),
					Length:      int64(enumCtx.Identifier().GetSymbol().GetStop() - enumCtx.Identifier().GetSymbol().GetStart() + 1),
					ParentIndex: e.Id,
				},
				NodeType: ast_pb.NodeType_ENUM_VALUE,
				TypeDescription: &TypeDescription{
					TypeIdentifier: fmt.Sprintf("t_enum_$_%s$_%s_$%d", e.Name, enumCtx.GetText(), id),
					TypeString:     fmt.Sprintf("enum %s.%s", e.CanonicalName, enumCtx.GetText()),
				},
			},
		)
	}

	e.globalDefinitions = append(e.globalDefinitions, e)

	return e
}

// There can be global enums that are outside of the contract body, so we need to handle them here.
func (b *ASTBuilder) EnterEnumDefinition(ctx *parser.EnumDefinitionContext) {
	enumDef := NewEnumDefinition(b)
	enumDef.ParseGlobal(ctx)
}
