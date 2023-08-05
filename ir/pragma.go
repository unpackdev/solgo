package ir

import (
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type Pragma struct {
	unit *ast.Pragma

	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Literals []string        `json:"literals"`
	Text     string          `json:"text"`
}

func (p *Pragma) GetAST() *ast.Pragma {
	return p.unit
}

func (p *Pragma) GetId() int64 {
	return p.Id
}

func (p *Pragma) GetNodeType() ast_pb.NodeType {
	return p.NodeType
}

func (p *Pragma) GetLiterals() []string {
	return p.Literals
}

func (p *Pragma) GetText() string {
	return p.Text
}

func (p *Pragma) GetVersion() string {
	parts := strings.Split(p.Text, " ")
	return strings.Replace(parts[len(parts)-1], ";", "", -1)
}

func (b *Builder) processPragma(unit *ast.Pragma) *Pragma {
	return &Pragma{
		unit:     unit,
		Id:       unit.GetId(),
		NodeType: unit.GetType(),
		Literals: unit.GetLiterals(),
		Text:     unit.GetText(),
	}
}
