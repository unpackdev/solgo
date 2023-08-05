package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Receive struct {
	unit     *ast.Receive
	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
}

func (f *Receive) GetAST() *ast.Receive {
	return f.unit
}

func (f *Receive) GetId() int64 {
	return f.Id
}

func (f *Receive) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (b *Builder) processReceive(unit *ast.Receive) *Receive {
	toReturn := &Receive{
		unit:     unit,
		Id:       unit.GetId(),
		NodeType: unit.GetType(),
	}

	return toReturn
}
