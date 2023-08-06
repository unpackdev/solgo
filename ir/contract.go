package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
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
	GetStructs() []*ast.StructDefinition
	GetConstructor() *ast.Constructor
	GetFunctions() []*ast.Function
	GetFallback() *ast.Fallback
	GetReceive() *ast.Receive
	GetEnums() []*ast.EnumDefinition
	GetEvents() []*ast.EventDefinition
	GetErrors() []*ast.ErrorDefinition
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
	Symbols        []*Symbol                                    `json:"symbols"`
	Imports        []*Import                                    `json:"imports"`
	Pragmas        []*Pragma                                    `json:"pragmas"`
	StateVariables []*StateVariable                             `json:"state_variables"`
	Structs        []*Struct                                    `json:"structs"`
	Enums          []*Enum                                      `json:"enums"`
	Events         []*Event                                     `json:"events"`
	Errors         []*Error                                     `json:"errors"`
	Constructor    *Constructor                                 `json:"constructor,omitempty"`
	Functions      []*Function                                  `json:"functions"`
	Fallback       *Fallback                                    `json:"fallback,omitempty"`
	Receive        *Receive                                     `json:"receive,omitempty"`
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

func (c *Contract) GetKind() ast_pb.NodeType {
	return c.Kind
}

func (c *Contract) GetImports() []*Import {
	return c.Imports
}

func (c *Contract) GetPragmas() []*Pragma {
	return c.Pragmas
}

func (c *Contract) GetStructs() []*Struct {
	return c.Structs
}

func (c *Contract) GetEnums() []*Enum {
	return c.Enums
}

func (c *Contract) GetEvents() []*Event {
	return c.Events
}

func (c *Contract) GetErrors() []*Error {
	return c.Errors
}

func (c *Contract) GetConstructor() *Constructor {
	return c.Constructor
}

func (c *Contract) GetFunctions() []*Function {
	return c.Functions
}

func (c *Contract) GetFallback() *Fallback {
	return c.Fallback
}

func (c *Contract) GetReceive() *Receive {
	return c.Receive
}

func (c *Contract) GetSymbols() []*Symbol {
	return c.Symbols
}

func (c *Contract) GetLanguage() Language {
	return c.Language
}

func (c *Contract) ToProto() *ir_pb.Contract {
	proto := &ir_pb.Contract{
		Id:             c.GetId(),
		NodeType:       c.GetNodeType(),
		Kind:           c.GetKind(),
		Name:           c.GetName(),
		License:        c.GetLicense(),
		Language:       c.GetLanguage().String(),
		AbsolutePath:   c.GetAbsolutePath(),
		SourceUnitId:   c.GetSourceUnitId(),
		Symbols:        make([]*ir_pb.Symbol, 0),
		Imports:        make([]*ir_pb.Import, 0),
		Pragmas:        make([]*ir_pb.Pragma, 0),
		StateVariables: make([]*ir_pb.StateVariable, 0),
		Structs:        make([]*ir_pb.Struct, 0),
		Enums:          make([]*ir_pb.Enum, 0),
		Events:         make([]*ir_pb.Event, 0),
		Errors:         make([]*ir_pb.Error, 0),
	}

	for _, symbol := range c.GetSymbols() {
		proto.Symbols = append(proto.Symbols, symbol.ToProto())
	}

	for _, imp := range c.GetImports() {
		proto.Imports = append(proto.Imports, imp.ToProto())
	}

	for _, pragma := range c.GetPragmas() {
		proto.Pragmas = append(proto.Pragmas, pragma.ToProto())
	}

	for _, stateVar := range c.GetStateVariables() {
		proto.StateVariables = append(proto.StateVariables, stateVar.ToProto())
	}

	for _, strct := range c.GetStructs() {
		proto.Structs = append(proto.Structs, strct.ToProto())
	}

	for _, enum := range c.GetEnums() {
		proto.Enums = append(proto.Enums, enum.ToProto())
	}

	for _, event := range c.GetEvents() {
		proto.Events = append(proto.Events, event.ToProto())
	}

	for _, err := range c.GetErrors() {
		proto.Errors = append(proto.Errors, err.ToProto())
	}

	if c.GetConstructor() != nil {
		proto.Constructor = c.GetConstructor().ToProto()
	}

	for _, fn := range c.GetFunctions() {
		proto.Functions = append(proto.Functions, fn.ToProto())
	}

	if c.GetFallback() != nil {
		proto.Fallback = c.GetFallback().ToProto()
	}

	if c.GetReceive() != nil {
		proto.Receive = c.GetReceive().ToProto()
	}

	return proto
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
		Symbols:        make([]*Symbol, 0),
		StateVariables: make([]*StateVariable, 0),
		Structs:        make([]*Struct, 0),
		Enums:          make([]*Enum, 0),
		Events:         make([]*Event, 0),
		Errors:         make([]*Error, 0),
		Functions:      make([]*Function, 0),
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

	// Process symbols of the contract.
	for _, symbol := range unit.GetExportedSymbols() {
		contractNode.Symbols = append(
			contractNode.Symbols,
			b.processSymbol(symbol),
		)
	}

	// Process state variables of the contract.
	for _, stateVariable := range contract.GetStateVariables() {
		contractNode.StateVariables = append(
			contractNode.StateVariables,
			b.processStateVariables(stateVariable),
		)
	}

	// Process structs of the contract.
	for _, structNode := range contract.GetStructs() {
		contractNode.Structs = append(
			contractNode.Structs,
			b.processStruct(structNode),
		)
	}

	// Process enums of the contract.
	for _, enum := range contract.GetEnums() {
		contractNode.Enums = append(
			contractNode.Enums,
			b.processEnum(enum),
		)
	}

	// Process events of the contract.
	for _, event := range contract.GetEvents() {
		contractNode.Events = append(
			contractNode.Events,
			b.processEvent(event),
		)
	}

	// Process errors of the contract.
	for _, errorNode := range contract.GetErrors() {
		contractNode.Errors = append(
			contractNode.Errors,
			b.processError(errorNode),
		)
	}

	// Process constructor of the contract.
	if contract.GetConstructor() != nil {
		contractNode.Constructor = b.processConstructor(contract.GetConstructor())
	}

	// Process functions of the contract.
	for _, function := range contract.GetFunctions() {
		contractNode.Functions = append(
			contractNode.Functions,
			b.processFunction(function),
		)
	}

	// Process fallback of the contract.
	if contract.GetFallback() != nil {
		contractNode.Fallback = b.processFallback(contract.GetFallback())
	}

	// Process receive of the contract.
	if contract.GetReceive() != nil {
		contractNode.Receive = b.processReceive(contract.GetReceive())
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
