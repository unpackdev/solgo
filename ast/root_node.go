package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// RootNode is the root node of the AST.
type RootNode struct {
	// Id is the unique identifier of the root node.
	Id int64 `json:"id"`

	// NodeType is the type of the AST node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// EntrySourceUnit is the entry source unit of the root node.
	EntrySourceUnit int64 `json:"entry_source_unit"`

	// SourceUnits is the list of source units.
	SourceUnits []*SourceUnit[Node[ast_pb.SourceUnit]] `json:"root"`

	// Comments is the list of comments.
	Comments []*Comment `json:"comments"`
}

// NewRootNode creates a new RootNode with the provided ASTBuilder, entry source unit, source units, and comments.
func NewRootNode(builder *ASTBuilder, entrySourceUnit int64, sourceUnits []*SourceUnit[Node[ast_pb.SourceUnit]], comments []*Comment) *RootNode {
	return &RootNode{
		Id:              builder.GetNextID(),
		EntrySourceUnit: entrySourceUnit,
		NodeType:        ast_pb.NodeType_ROOT_SOURCE_UNIT,
		Comments:        comments,
		SourceUnits:     sourceUnits,
	}
}

// GetId returns the id of the RootNode node.
func (r *RootNode) GetId() int64 {
	return r.Id
}

// GetType returns the type of the RootNode node.
func (r *RootNode) GetType() ast_pb.NodeType {
	return r.NodeType
}

// GetSrc returns the source code location of the RootNode node.
func (r *RootNode) GetSrc() SrcNode {
	return SrcNode{}
}

// GetTypeDescription returns the type description of the RootNode node.
// RootNode nodes do not have type descriptions.
func (r *RootNode) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "root",
		TypeIdentifier: "t_root",
	}
}

// SetReferenceDescriptor sets the reference descriptions of the RootNode node.
func (r *RootNode) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetSourceUnits returns the source units of the root node.
func (r *RootNode) GetSourceUnits() []*SourceUnit[Node[ast_pb.SourceUnit]] {
	return r.SourceUnits
}

// GetSourceUnitByName returns the source unit with the provided name.
func (r *RootNode) GetSourceUnitByName(name string) *SourceUnit[Node[ast_pb.SourceUnit]] {
	for _, sourceUnit := range r.SourceUnits {
		if sourceUnit.Name == name {
			return sourceUnit
		}
	}

	return nil
}

// GetSourceUnitById returns the source unit with the provided id.
func (r *RootNode) GetSourceUnitById(id int64) *SourceUnit[Node[ast_pb.SourceUnit]] {
	for _, sourceUnit := range r.SourceUnits {
		if sourceUnit.Id == id {
			return sourceUnit
		}
	}

	return nil
}

// HasSourceUnits returns true if the root node has source units.
func (r *RootNode) HasSourceUnits() bool {
	return len(r.SourceUnits) > 0
}

// GetSourceUnitCount returns the number of source units of the root node.
func (r *RootNode) GetSourceUnitCount() int32 {
	return int32(len(r.SourceUnits))
}

// GetEntrySourceUnit returns the entry source unit of the root node.
func (r *RootNode) GetEntrySourceUnit() int64 {
	return r.EntrySourceUnit
}

// SetEntrySourceUnit sets the entry source unit of the root node.
func (r *RootNode) SetEntrySourceUnit(entrySourceUnit int64) {
	r.EntrySourceUnit = entrySourceUnit
}

// GetComments returns the comments of the root node.
func (r *RootNode) GetComments() []*Comment {
	return r.Comments
}

// GetNodes returns the nodes of the root node.
func (r *RootNode) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	for _, sourceUnit := range r.SourceUnits {
		toReturn = append(toReturn, sourceUnit)
	}
	return toReturn
}

// ToProto returns the protobuf representation of the root node.
func (r *RootNode) ToProto() *ast_pb.RootSourceUnit {
	sourceUnits := []*ast_pb.SourceUnit{}
	for _, sourceUnit := range r.SourceUnits {
		sourceUnits = append(sourceUnits, sourceUnit.ToProto().(*ast_pb.SourceUnit))
	}

	comments := []*ast_pb.Comment{}
	for _, comment := range r.Comments {
		comments = append(comments, comment.ToProto())
	}

	return &ast_pb.RootSourceUnit{
		Id:              r.GetId(),
		NodeType:        r.GetType(),
		EntrySourceUnit: r.GetEntrySourceUnit(),
		SourceUnits:     sourceUnits,
		Comments:        comments,
	}
}
