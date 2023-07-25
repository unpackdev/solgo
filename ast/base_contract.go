package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
)

type BaseContract struct {
	Id       int64             `json:"id"`
	NodeType ast_pb.NodeType   `json:"node_type"`
	Src      SrcNode           `json:"src"`
	BaseName *BaseContractName `json:"base_name"`
}

func (b *BaseContract) GetId() int64 {
	return b.Id
}

func (b *BaseContract) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *BaseContract) GetSrc() SrcNode {
	return b.Src
}

func (b *BaseContract) ToProto() *ast_pb.BaseContract {
	return &ast_pb.BaseContract{
		Id:       b.Id,
		NodeType: b.NodeType,
		Src:      b.Src.ToProto(),
		BaseName: b.BaseName.ToProto(),
	}
}

type BaseContractName struct {
	Id                    int64           `json:"id"`
	NodeType              ast_pb.NodeType `json:"node_type"`
	Src                   SrcNode         `json:"src"`
	Name                  string          `json:"name"`
	ReferencedDeclaration int64           `json:"referenced_declaration"`
}

func (b *BaseContractName) GetId() int64 {
	return b.Id
}

func (b *BaseContractName) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *BaseContractName) GetSrc() SrcNode {
	return b.Src
}

func (b *BaseContractName) ToProto() *ast_pb.BaseContractName {
	return &ast_pb.BaseContractName{
		Id:                    b.Id,
		NodeType:              b.NodeType,
		Src:                   b.Src.ToProto(),
		Name:                  b.Name,
		ReferencedDeclaration: b.ReferencedDeclaration,
	}
}
