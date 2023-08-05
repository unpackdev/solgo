package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Struct struct {
	unit                    *ast.StructDefinition
	Id                      int64                  `json:"id"`
	NodeType                ast_pb.NodeType        `json:"node_type"`
	Kind                    ast_pb.NodeType        `json:"kind,omitempty"`
	Name                    string                 `json:"name"`
	CanonicalName           string                 `json:"canonical_name"`
	ReferencedDeclarationId int64                  `json:"referenced_declaration_id"`
	Visibility              ast_pb.Visibility      `json:"visibility"`
	StorageLocation         ast_pb.StorageLocation `json:"storage_location"`
	Members                 []*Parameter           `json:"members"`
	Type                    string                 `json:"type"`
	TypeDescription         *ast.TypeDescription   `json:"type_description"`
}

func (f *Struct) GetAST() *ast.StructDefinition {
	return f.unit
}

func (f *Struct) GetId() int64 {
	return f.Id
}

func (f *Struct) GetName() string {
	return f.Name
}

func (f *Struct) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Struct) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Struct) GetCanonicalName() string {
	return f.CanonicalName
}

func (f *Struct) GetReferencedDeclarationId() int64 {
	return f.ReferencedDeclarationId
}

func (f *Struct) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Struct) GetStorageLocation() ast_pb.StorageLocation {
	return f.StorageLocation
}

func (f *Struct) GetMembers() []*Parameter {
	return f.Members
}

func (f *Struct) GetType() string {
	return f.Type
}

func (f *Struct) GetTypeDescription() *ast.TypeDescription {
	return f.TypeDescription
}

func (f *Struct) GetSrc() ast.SrcNode {
	return f.unit.GetSrc()
}

func (b *Builder) processStruct(unit *ast.StructDefinition) *Struct {
	toReturn := &Struct{
		unit:                    unit,
		Id:                      unit.GetId(),
		NodeType:                unit.GetType(),
		Kind:                    unit.GetKind(),
		Name:                    unit.GetName(),
		CanonicalName:           unit.GetCanonicalName(),
		ReferencedDeclarationId: unit.GetReferencedDeclaration(),
		Visibility:              unit.GetVisibility(),
		StorageLocation:         unit.GetStorageLocation(),
		Members:                 make([]*Parameter, 0),
		Type:                    "struct",
		TypeDescription:         unit.GetTypeDescription(),
	}

	for _, parameter := range unit.GetMembers() {
		param := &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		}

		toReturn.Members = append(toReturn.Members, param)
	}

	return toReturn
}
