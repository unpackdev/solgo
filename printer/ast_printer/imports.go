package ast_printer

import (
	"fmt"
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printImport(node *ast.Import, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("import ")
	file := fmt.Sprintf("'%s'", node.GetFile())
	if node.UnitAlias != "" {
		writeStrings(sb, file, " as ", node.UnitAlias)
	} else if len(node.UnitAliases) > 0 {
		sb.WriteString("{")
		writeSeperatedList(sb, ", ", node.UnitAliases)
		writeStrings(sb, "} from ", file)
	} else {
		writeStrings(sb, file)
	}
	writeStrings(sb, ";")
	return success
}
