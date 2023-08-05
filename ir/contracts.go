package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

type ContractNode interface {
	GetId() int64
	GetType() ast_pb.NodeType
	GetKind() ast_pb.NodeType
	GetSrc() ast.SrcNode
	GetTypeDescription() *ast.TypeDescription
	GetNodes() []ast.Node[ast.NodeType]
	ToProto() ast.NodeType
	SetReferenceDescriptor(refId int64, refDesc *ast.TypeDescription) bool
	GetStateVariables() []*ast.StateVariableDeclaration
}

type Contract struct {
	unit           *ast.SourceUnit[ast.Node[ast_pb.SourceUnit]] `json:"-"`
	Id             int64                                        `json:"id"`
	SourceUnitId   int64                                        `json:"source_unit_id"`
	NodeType       ast_pb.NodeType                              `json:"node_type"`
	Kind           ast_pb.NodeType                              `json:"kind"`
	Name           string                                       `json:"name"`
	License        string                                       `json:"license"`
	Language       Language                                     `json:"language"`
	AbsolutePath   string                                       `json:"absolute_path"`
	Symbols        []ast.Symbol                                 `json:"symbols"`
	Imports        []*Import                                    `json:"imports"`
	Pragmas        []*Pragma                                    `json:"pragmas"`
	StateVariables []*StateVariable                             `json:"state_variables"`
}

func (c *Contract) GetAST() *ast.SourceUnit[ast.Node[ast_pb.SourceUnit]] {
	return c.unit
}

func (c *Contract) GetId() int64 {
	return c.Id
}

func (c *Contract) GetNodeType() ast_pb.NodeType {
	return c.NodeType
}

func (c *Contract) GetName() string {
	return c.Name
}

func (c *Contract) GetLicense() string {
	return c.License
}

func (c *Contract) GetAbsolutePath() string {
	return c.AbsolutePath
}

func (c *Contract) GetStateVariables() []*StateVariable {
	return c.StateVariables
}

func (c *Contract) GetSourceUnitId() int64 {
	return c.SourceUnitId
}

func (c *Contract) GetUnitSrc() ast.SrcNode {
	return c.unit.GetSrc()
}

func (c *Contract) GetSrc() ast.SrcNode {
	return c.unit.GetContract().GetSrc()
}

func (b *Builder) processContract(unit *ast.SourceUnit[ast.Node[ast_pb.SourceUnit]]) *Contract {
	contract := getContractByNodeType(unit.GetContract())
	contractNode := &Contract{
		unit: unit,

		Id:             contract.GetId(),
		NodeType:       contract.GetType(),
		Kind:           contract.GetKind(),
		Name:           unit.GetName(),
		SourceUnitId:   unit.GetId(),
		License:        unit.GetLicense(),
		Language:       LanguageSolidity,
		AbsolutePath:   unit.GetAbsolutePath(),
		Pragmas:        make([]*Pragma, 0),
		Imports:        make([]*Import, 0),
		Symbols:        unit.GetExportedSymbols(),
		StateVariables: make([]*StateVariable, 0),
	}

	for _, pragma := range unit.GetPragmas() {
		contractNode.Pragmas = append(
			contractNode.Pragmas,
			b.processPragma(pragma),
		)
	}

	for _, importNode := range unit.GetImports() {
		contractNode.Imports = append(
			contractNode.Imports,
			b.processImport(importNode),
		)
	}

	// Process state variables of the contract.
	for _, stateVariable := range contract.GetStateVariables() {
		contractNode.StateVariables = append(
			contractNode.StateVariables,
			b.processStateVariables(stateVariable),
		)
	}

	return contractNode
}

func getContractByNodeType(c ast.Node[ast.NodeType]) ContractNode {
	switch contract := c.(type) {
	case *ast.Library:
		return contract
	case *ast.Interface:
		return contract
	case *ast.Contract:
		return contract
	}

	return nil
}
