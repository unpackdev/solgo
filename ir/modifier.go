package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Modifier struct {
	unit          *ast.ModifierInvocation `json:"-"`
	Id            int64                   `json:"id"`
	NodeType      ast_pb.NodeType         `json:"node_type"`
	Name          string                  `json:"name"`
	ArgumentTypes []*ast.TypeDescription  `json:"argument_types"`
}

func (m *Modifier) GetAST() *ast.ModifierInvocation {
	return m.unit
}

func (m *Modifier) GetId() int64 {
	return m.Id
}

func (m *Modifier) GetName() string {
	return m.Name
}

func (m *Modifier) GetNodeType() ast_pb.NodeType {
	return m.NodeType
}

func (m *Modifier) GetArgumentTypes() []*ast.TypeDescription {
	return m.ArgumentTypes
}

func (m *Modifier) ToProto() *ir_pb.Modifier {
	proto := &ir_pb.Modifier{
		Id:       m.GetId(),
		NodeType: m.GetNodeType(),
		Name:     m.GetName(),
	}

	for _, typeArgument := range m.GetArgumentTypes() {
		proto.ArgumentTypes = append(proto.ArgumentTypes, typeArgument.ToProto())
	}

	return proto
}
