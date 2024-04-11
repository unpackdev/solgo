package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printImport(node *ast.Import, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("import ")
	if node.UnitAlias != "" {
		writeStrings(sb, node.GetFile(), " as ", node.UnitAlias)
	} else if len(node.UnitAliases) > 0 {
		sb.WriteString("{")
		writeSeperatedList(sb, ", ", node.UnitAliases)
		writeStrings(sb, "} from ", node.GetFile())
	}
	writeStrings(sb, ";\n")
	return success
}
