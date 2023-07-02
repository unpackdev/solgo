package ast

import (
	"github.com/txpull/solgo/common"
)

type AstParser struct {
	ast            common.AST
	contractName   string
	definedStructs map[string]common.MethodIO
	definedEnums   map[string]bool
}
