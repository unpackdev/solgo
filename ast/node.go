package ast

import ast_pb "github.com/txpull/protos/dist/go/ast"

type Node[T NodeType] interface {
	GetId() int64
	GetType() ast_pb.NodeType
	GetSrc() SrcNode
	ToProto() T
}

type NodeType interface {
	ast_pb.Pragma | ast_pb.Import | ast_pb.Modifier | ast_pb.SourceUnit |
		ast_pb.Function | ast_pb.Contract | ast_pb.Statement | ast_pb.Body |
		any
}
