package ast

import ast_pb "github.com/unpackdev/protos/dist/go/ast"

type YulIdentifier struct {
	*ASTBuilder

	Id           int64           `json:"id"`
	NodeType     ast_pb.NodeType `json:"node_type"`
	Src          SrcNode         `json:"src"`
	Name         string          `json:"name"`
	NameLocation SrcNode         `json:"name_location"`
}

type YulEVMBuiltin struct {
	*ASTBuilder

	Id           int64           `json:"id"`
	NodeType     ast_pb.NodeType `json:"node_type"`
	Src          SrcNode         `json:"src"`
	Name         string          `json:"name"`
	NameLocation SrcNode         `json:"name_location"`
}
