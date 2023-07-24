package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
)

type BaseContract struct {
	*ASTBuilder

	Id       int64             `json:"id"`
	NodeType ast_pb.NodeType   `json:"node_type"`
	Src      SrcNode           `json:"src"`
	BaseName *BaseContractName `json:"base_name"`
}

func NewBaseContract(b *ASTBuilder) *BaseContract {
	return &BaseContract{
		ASTBuilder: b,
	}
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

type BaseContractName struct {
	*ASTBuilder

	Id                    int64           `json:"id"`
	NodeType              ast_pb.NodeType `json:"node_type"`
	Src                   SrcNode         `json:"src"`
	Name                  string          `json:"name"`
	ReferencedDeclaration int64           `json:"referenced_declaration"`
}

func NewBaseContractName(b *ASTBuilder) *BaseContractName {
	return &BaseContractName{
		ASTBuilder: b,
	}
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
