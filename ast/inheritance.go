package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

func parseInheritanceFromCtx(b *ASTBuilder, unit *SourceUnit[Node[ast_pb.SourceUnit]], contractNode Node[NodeType], ctx parser.IInheritanceSpecifierListContext) []*BaseContract {
	toReturn := []*BaseContract{}

	if ctx == nil {
		return toReturn
	}

	for _, specifierCtx := range ctx.AllInheritanceSpecifier() {
		baseContract := &BaseContract{
			Id: b.GetNextID(),
			Src: SrcNode{
				Id:          b.GetNextID(),
				Line:        int64(specifierCtx.GetStart().GetLine()),
				Column:      int64(specifierCtx.GetStart().GetColumn()),
				Start:       int64(specifierCtx.GetStart().GetStart()),
				End:         int64(specifierCtx.GetStop().GetStop()),
				Length:      int64(specifierCtx.GetStop().GetStop() - specifierCtx.GetStart().GetStart() + 1),
				ParentIndex: contractNode.GetId(),
			},
			NodeType: ast_pb.NodeType_INHERITANCE_SPECIFIER,
			BaseName: &BaseContractName{
				Id: b.GetNextID(),
				Src: SrcNode{
					Id:          b.GetNextID(),
					Line:        int64(specifierCtx.GetStart().GetLine()),
					Column:      int64(specifierCtx.GetStart().GetColumn()),
					Start:       int64(specifierCtx.GetStart().GetStart()),
					End:         int64(specifierCtx.GetStop().GetStop()),
					Length:      int64(specifierCtx.GetStop().GetStop() - specifierCtx.GetStart().GetStart() + 1),
					ParentIndex: contractNode.GetId(),
				},
				NodeType: ast_pb.NodeType_IDENTIFIER_PATH,
				Name:     specifierCtx.IdentifierPath().GetText(),
			},
		}

		for _, unitNode := range b.sourceUnits {
			if unitNode.GetName() == specifierCtx.IdentifierPath().GetText() {
				baseContract.BaseName.ReferencedDeclaration = unitNode.GetId()

				switch nodeCtx := contractNode.(type) {
				case *ContractNode[ast_pb.Contract]:
					nodeCtx.LinearizedBaseContracts = append(
						nodeCtx.LinearizedBaseContracts, unitNode.GetId(),
					)
					nodeCtx.ContractDependencies = append(
						nodeCtx.ContractDependencies, unitNode.GetId(),
					)
				case *InterfaceNode[ast_pb.Contract]:
					nodeCtx.LinearizedBaseContracts = append(
						nodeCtx.LinearizedBaseContracts, unitNode.GetId(),
					)
					nodeCtx.ContractDependencies = append(
						nodeCtx.ContractDependencies, unitNode.GetId(),
					)
				}

				symbolFound := false
				for _, symbol := range unitNode.GetExportedSymbols() {
					if symbol.GetName() == unitNode.GetName() {
						symbolFound = true
					}
				}

				if !symbolFound {
					unit.ExportedSymbols = append(
						unit.ExportedSymbols,
						Symbol{
							Id:   unitNode.GetId(),
							Name: unitNode.GetName(),
							AbsolutePath: func() string {
								for _, unit := range b.sources.SourceUnits {
									if unit.Name == unitNode.GetName() {
										return unit.Path
									}
								}
								return ""
							}(),
						},
					)
				}
			}
		}

		toReturn = append(toReturn, baseContract)
	}

	return toReturn
}
