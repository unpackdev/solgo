package ast

import ast_pb "github.com/txpull/protos/dist/go/ast"

type Node[T NodeType] interface {
	GetId() int64
	GetType() ast_pb.NodeType
	GetSrc() SrcNode
	GetTypeDescription() *TypeDescription
	GetNodes() []Node[NodeType]
	ToProto() T
	SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool
}

type NodeType interface {
	ast_pb.Pragma | ast_pb.Import | ast_pb.Modifier | ast_pb.SourceUnit |
		ast_pb.Function | ast_pb.Contract | ast_pb.Statement | ast_pb.Body |
		ast_pb.Variable | ast_pb.PrimaryExpression | ast_pb.Expression | ast_pb.Using |
		ast_pb.Declaration | ast_pb.TypeName | ast_pb.BaseContract | ast_pb.TypeDescription |
		ast_pb.BinaryOperationExpression | ast_pb.Return | ast_pb.ParameterList | ast_pb.Parameter |
		ast_pb.StateVariable | ast_pb.Event | any
}
