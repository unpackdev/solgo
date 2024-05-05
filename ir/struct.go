package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Struct represents a Solidity struct definition as an IR node.
type Struct struct {
	Unit                    *ast.StructDefinition  `json:"ast"`
	Id                      int64                  `json:"id"`
	NodeType                ast_pb.NodeType        `json:"nodeType"`
	Name                    string                 `json:"name"`
	CanonicalName           string                 `json:"canonicalName"`
	ReferencedDeclarationId int64                  `json:"referencedDeclarationId"`
	Visibility              ast_pb.Visibility      `json:"visibility"`
	StorageLocation         ast_pb.StorageLocation `json:"storageLocation"`
	Members                 []*Parameter           `json:"members"`
	Type                    string                 `json:"type"`
	TypeDescription         *ast.TypeDescription   `json:"typeDescription"`
}

// GetAST returns the underlying AST node of the Struct.
func (f *Struct) GetAST() *ast.StructDefinition {
	return f.Unit
}

// GetId returns the unique identifier of the struct.
func (f *Struct) GetId() int64 {
	return f.Id
}

// GetName returns the name of the struct.
func (f *Struct) GetName() string {
	return f.Name
}

// GetNodeType returns the type of the node in the AST.
func (f *Struct) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

// GetCanonicalName returns the canonical name of the struct.
func (f *Struct) GetCanonicalName() string {
	return f.CanonicalName
}

// GetReferencedDeclarationId returns the referenced declaration ID of the struct.
func (f *Struct) GetReferencedDeclarationId() int64 {
	return f.ReferencedDeclarationId
}

// GetVisibility returns the visibility of the struct.
func (f *Struct) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStorageLocation returns the storage location of the struct.
func (f *Struct) GetStorageLocation() ast_pb.StorageLocation {
	return f.StorageLocation
}

// GetMembers returns the list of members (parameters) in the struct.
func (f *Struct) GetMembers() []*Parameter {
	return f.Members
}

// GetType returns the type of the struct.
func (f *Struct) GetType() string {
	return f.Type
}

// GetTypeDescription returns the type description of the struct.
func (f *Struct) GetTypeDescription() *ast.TypeDescription {
	return f.TypeDescription
}

// GetSrc returns the source node of the struct.
func (f *Struct) GetSrc() ast.SrcNode {
	return f.Unit.GetSrc()
}

// ToProto is a placeholder function for converting the Struct to a protobuf message.
func (f *Struct) ToProto() *ir_pb.Struct {
	proto := &ir_pb.Struct{
		Id:                      f.GetId(),
		NodeType:                f.GetNodeType(),
		Name:                    f.GetName(),
		CanonicalName:           f.GetCanonicalName(),
		ReferencedDeclarationId: f.GetReferencedDeclarationId(),
		Visibility:              f.GetVisibility(),
		StorageLocation:         f.GetStorageLocation(),
		Members:                 make([]*ir_pb.Parameter, 0),
		Type:                    f.GetType(),
		TypeDescription:         f.GetTypeDescription().ToProto(),
	}

	for _, member := range f.GetMembers() {
		proto.Members = append(proto.Members, member.ToProto())
	}

	return proto
}

// processStruct processes the given struct definition node of an AST and returns a Struct.
// It populates the Struct with the members (parameters) from the AST.
func (b *Builder) processStruct(unit *ast.StructDefinition) *Struct {
	toReturn := &Struct{
		Unit:                    unit,
		Id:                      unit.GetId(),
		NodeType:                unit.GetType(),
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
			Unit:            parameter,
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
