package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Fallback struct {
	unit *ast.Fallback

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
}

func (f *Fallback) GetAST() *ast.Fallback {
	return f.unit
}

func (f *Fallback) GetId() int64 {
	return f.Id
}

func (f *Fallback) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

func (b *Builder) processFallback(unit *ast.Fallback) *Fallback {
	toReturn := &Fallback{
		unit:     unit,
		Id:       unit.GetId(),
		NodeType: unit.GetType(),
	}

	return toReturn
}
