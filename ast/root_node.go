package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
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
	SourceUnits []*SourceUnit[Node] `json:"root"`

	// Comments is the list of comments.
	Comments []*CommentNode `json:"comments"`
}

// NewRootNode creates a new root node.
func NewRootNode(builder *ASTBuilder, entrySourceUnit int64, sourceUnits []*SourceUnit[Node], comments []*CommentNode) Node {
	return Node(&RootNode{
		Id:              builder.GetNextID(),
		EntrySourceUnit: entrySourceUnit,
		NodeType:        ast_pb.NodeType_ROOT_SOURCE_UNIT,
		Comments:        comments,
		SourceUnits:     sourceUnits,
	})
}

// GetSourceUnits returns the source units of the root node.
func (r *RootNode) GetSourceUnits() []*SourceUnit[Node] {
	return r.SourceUnits
}

// GetSourceUnitCount returns the number of source units of the root node.
func (r *RootNode) GetSourceUnitCount() int32 {
	return int32(len(r.SourceUnits))
}

// ToProto returns the protobuf representation of the root node.
func (r *RootNode) ToProto() *ast_pb.RootNode {
	return &ast_pb.RootNode{
		Nodes: make([]*ast_pb.Node, 0),
	}
}

func (r *RootNode) GetId() int64 {
	return r.Id
}

func (r *RootNode) GetType() ast_pb.NodeType {
	return r.NodeType
}

func (r *RootNode) GetSrc() SrcNode {
	return SrcNode{}
}
