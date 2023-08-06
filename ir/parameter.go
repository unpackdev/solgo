package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Parameter struct {
	unit            *ast.Parameter       `json:"-"`
	Id              int64                `json:"id"`
	NodeType        ast_pb.NodeType      `json:"node_type"`
	Name            string               `json:"name"`
	Type            string               `json:"type"`
	TypeDescription *ast.TypeDescription `json:"type_description"`
}

func (p *Parameter) GetAST() *ast.Parameter {
	return p.unit
}

func (p *Parameter) GetId() int64 {
	return p.Id
}

func (p *Parameter) GetName() string {
	return p.Name
}

func (p *Parameter) GetNodeType() ast_pb.NodeType {
	return p.NodeType
}

func (p *Parameter) GetType() string {
	return p.Type
}

func (p *Parameter) GetTypeDescription() *ast.TypeDescription {
	return p.TypeDescription
}

func (p *Parameter) ToProto() *ir_pb.Parameter {
	proto := &ir_pb.Parameter{
		Id:              p.GetId(),
		NodeType:        p.GetNodeType(),
		Name:            p.GetName(),
		Type:            p.GetType(),
		TypeDescription: p.GetTypeDescription().ToProto(),
	}

	return proto
}
