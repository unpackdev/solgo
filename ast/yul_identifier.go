package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// YulIdentifier represents a YUL identifier in the abstract syntax tree.
type YulIdentifier struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL identifier.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL identifier node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL identifier.
	Src SrcNode `json:"src"`

	// Name is the name of the YUL identifier.
	Name string `json:"name"`
}

// SetReferenceDescriptor sets the reference descriptions of the YulIdentifier node.
func (y *YulIdentifier) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulIdentifier.
func (y *YulIdentifier) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulIdentifier.
func (y *YulIdentifier) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulIdentifier.
func (y *YulIdentifier) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns an empty list of nodes.
func (y *YulIdentifier) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{}
}

// GetTypeDescription returns the type description of the YulIdentifier.
func (y *YulIdentifier) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetName returns the name of the YulIdentifier.
func (y *YulIdentifier) GetName() string {
	return y.Name
}

// ToProto converts the YulIdentifier to its protocol buffer representation.
func (y *YulIdentifier) ToProto() NodeType {
	toReturn := ast_pb.YulIdentifier{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
		Name:     y.GetName(),
	}

	return NewTypedStruct(&toReturn, "YulIdentifier")
}
