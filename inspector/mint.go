package inspector

import (
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

func (i *Inspector) DetectMintable(nodeCtx *ast.Function) (Mintable, bool) {
	if utils.StringInSlice(nodeCtx.GetName(), mintableFunctions) {
		i.report.Mintable.Enabled = true
		//i.report.Mintable.Function = nodeCtx

		utils.DumpNodeWithExit(nodeCtx)
	}

	return i.report.Mintable, i.report.Mintable.Enabled
}
