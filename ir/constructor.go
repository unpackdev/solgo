package ir

import (
	"github.com/txpull/solgo/ast"
)

type Constructor struct {
	unit *ast.Constructor
}

func (c *Constructor) GetAST() *ast.Constructor {
	return c.unit
}

func (b *Builder) processConstructor(unit *ast.Constructor) *Constructor {
	toReturn := &Constructor{
		unit: unit,
	}

	return toReturn
}
