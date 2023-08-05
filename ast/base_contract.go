package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
)

// BaseContract represents a base contract in a Solidity source file.
// A base contract is a contract that is inherited by another contract.
type BaseContract struct {
	// Id is the unique identifier of the base contract.
	Id int64 `json:"id"`
	// NodeType is the type of the node.
	// For a BaseContract, this is always NodeType_BASE_CONTRACT.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Src contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// BaseName is the name of the base contract.
	BaseName *BaseContractName `json:"base_name"`
}

// GetId returns the unique identifier of the base contract.
func (b *BaseContract) GetId() int64 {
	return b.Id
}

// GetType returns the type of the node.
// For a BaseContract, this is always NodeType_BASE_CONTRACT.
func (b *BaseContract) GetType() ast_pb.NodeType {
	return b.NodeType
}

// GetSrc returns source information about the node, such as its line and column numbers in the source file.
func (b *BaseContract) GetSrc() SrcNode {
	return b.Src
}

// GetBaseName returns the name of the base contract.
func (b *BaseContract) GetBaseName() *BaseContractName {
	return b.BaseName
}

// ToProto returns the protobuf representation of the base contract.
func (b *BaseContract) ToProto() *ast_pb.BaseContract {
	return &ast_pb.BaseContract{
		Id:       b.Id,
		NodeType: b.NodeType,
		Src:      b.Src.ToProto(),
		BaseName: b.BaseName.ToProto(),
	}
}

// BaseContractName represents the name of a base contract in a Solidity source file.
type BaseContractName struct {
	// Id is the unique identifier of the base contract name.
	Id int64 `json:"id"`
	// NodeType is the type of the node.
	// For a BaseContractName, this is always NodeType_BASE_CONTRACT_NAME.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Src contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// Name is the name of the base contract.
	Name string `json:"name"`
	// ReferencedDeclaration is the unique identifier of the contract declaration that this name references.
	ReferencedDeclaration int64 `json:"referenced_declaration"`
}

// GetId returns the unique identifier of the base contract name.
func (b *BaseContractName) GetId() int64 {
	return b.Id
}

// GetType returns the type of the node.
// For a BaseContractName, this is always NodeType_BASE_CONTRACT_NAME.
func (b *BaseContractName) GetType() ast_pb.NodeType {
	return b.NodeType
}

// GetSrc returns source information about the node, such as its line and column numbers in the source file.
func (b *BaseContractName) GetSrc() SrcNode {
	return b.Src
}

// GetName returns the name of the base contract name.
func (b *BaseContractName) GetName() string {
	return b.Name
}

// GetReferencedDeclaration returns the unique identifier of the source unit contract declaration that this name references.
func (b *BaseContractName) GetReferencedDeclaration() int64 {
	return b.ReferencedDeclaration
}

// ToProto returns the protobuf representation of the base contract name.
func (b *BaseContractName) ToProto() *ast_pb.BaseContractName {
	return &ast_pb.BaseContractName{
		Id:                    b.Id,
		NodeType:              b.NodeType,
		Src:                   b.Src.ToProto(),
		Name:                  b.Name,
		ReferencedDeclaration: b.ReferencedDeclaration,
	}
}
