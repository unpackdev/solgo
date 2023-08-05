package ir

import "github.com/txpull/solgo/ast"

type Builder struct {
	*ast.ASTBuilder
}

func NewBuilder(astBuilder *ast.ASTBuilder) *Builder {
	return &Builder{
		ASTBuilder: astBuilder,
	}
}
