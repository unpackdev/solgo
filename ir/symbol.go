package ir

import (
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Symbol represents a symbol in the Intermediate Representation (IR) of Solidity contracts' Abstract Syntax Tree (AST).
type Symbol struct {
	Id           int64  `json:"id"`            // Id is the unique identifier of the symbol.
	Name         string `json:"name"`          // Name is the name of the symbol.
	AbsolutePath string `json:"absolute_path"` // AbsolutePath is the absolute path of the symbol.
}

// GetId returns the unique identifier of the symbol.
func (s *Symbol) GetId() int64 {
	return s.Id
}

// GetName returns the name of the symbol.
func (s *Symbol) GetName() string {
	return s.Name
}

// GetAbsolutePath returns the absolute path of the symbol.
func (s *Symbol) GetAbsolutePath() string {
	return s.AbsolutePath
}

// ToProto is a function that converts the Symbol to a protobuf message.
func (s *Symbol) ToProto() *ir_pb.Symbol {
	return &ir_pb.Symbol{
		Id:           s.GetId(),
		Name:         s.GetName(),
		AbsolutePath: s.GetAbsolutePath(),
	}
}

// processSymbol is a function that processes the given symbol and returns a Symbol.
func (b *Builder) processSymbol(unit ast.Symbol) *Symbol {
	return &Symbol{
		Id:           unit.Id,
		Name:         unit.Name,
		AbsolutePath: unit.AbsolutePath,
	}
}
