package ast

import ast_pb "github.com/unpackdev/protos/dist/go/ast"

type YulIdentifier struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
	Name     string          `json:"name"`
}

// SetReferenceDescriptor sets the reference descriptions of the YulIdentifier node.
func (y *YulIdentifier) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulIdentifier) GetId() int64 {
	return y.Id
}

func (y *YulIdentifier) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulIdentifier) GetSrc() SrcNode {
	return y.Src
}

func (y *YulIdentifier) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	return toReturn
}

func (y *YulIdentifier) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulIdentifier) GetName() string {
	return y.Name
}

func (y *YulIdentifier) ToProto() NodeType {
	toReturn := ast_pb.YulIdentifier{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
		Name:     y.GetName(),
	}

	return NewTypedStruct(&toReturn, "YulIdentifier")
}
