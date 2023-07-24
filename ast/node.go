package ast

import ast_pb "github.com/txpull/protos/dist/go/ast"

type Node interface {
	GetId() int64
	GetType() ast_pb.NodeType
	GetSrc() SrcNode
}
