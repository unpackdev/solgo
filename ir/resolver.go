package ir

import (
	"github.com/txpull/solgo/ast"
)

func (r *Builder) byFunction(name string) *Function {
	for _, unit := range r.astBuilder.GetRoot().GetSourceUnits() {
		contract := unit.GetContract()
		for _, node := range contract.GetNodes() {
			if function, ok := node.(*ast.Function); ok {
				if function.GetName() == name {
					return r.processFunction(function, false)
				}
			}
		}
	}

	return nil
}
