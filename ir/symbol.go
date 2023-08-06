package ir

import (
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Symbol struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AbsolutePath string `json:"absolute_path"`
}

func (s *Symbol) GetId() int64 {
	return s.Id
}

func (s *Symbol) GetName() string {
	return s.Name
}

func (s *Symbol) GetAbsolutePath() string {
	return s.AbsolutePath
}

func (s *Symbol) ToProto() *ir_pb.Symbol {
	return &ir_pb.Symbol{
		Id:           s.GetId(),
		Name:         s.GetName(),
		AbsolutePath: s.GetAbsolutePath(),
	}
}

func (b *Builder) processSymbol(unit ast.Symbol) *Symbol {
	return &Symbol{
		Id:           unit.Id,
		Name:         unit.Name,
		AbsolutePath: unit.AbsolutePath,
	}
}
