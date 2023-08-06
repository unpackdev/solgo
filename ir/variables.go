package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type StateVariable struct {
	unit *ast.StateVariableDeclaration `json:"-"`

	Id              int64                  `json:"id"`
	ContractId      int64                  `json:"contract_id"`
	Name            string                 `json:"name"`
	NodeType        ast_pb.NodeType        `json:"node_type"`
	Visibility      ast_pb.Visibility      `json:"visibility"`
	Constant        bool                   `json:"is_constant"`
	StorageLocation ast_pb.StorageLocation `json:"storage_location"`
	StateMutability ast_pb.Mutability      `json:"state_mutability"`
	Type            string                 `json:"type"`
	TypeDescription *ast.TypeDescription   `json:"type_description"`
}

func (v *StateVariable) GetAST() *ast.StateVariableDeclaration {
	return v.unit
}

func (v *StateVariable) GetId() int64 {
	return v.Id
}

func (v *StateVariable) GetContractId() int64 {
	return v.ContractId
}

func (v *StateVariable) GetName() string {
	return v.Name
}

func (v *StateVariable) GetNodeType() ast_pb.NodeType {
	return v.NodeType
}

func (v *StateVariable) GetVisibility() ast_pb.Visibility {
	return v.Visibility
}

func (v *StateVariable) IsConstant() bool {
	return v.Constant
}

func (v *StateVariable) GetStorageLocation() ast_pb.StorageLocation {
	return v.StorageLocation
}

func (v *StateVariable) GetStateMutability() ast_pb.Mutability {
	return v.StateMutability
}

func (v *StateVariable) GetType() string {
	return v.Type
}

func (v *StateVariable) GetTypeDescription() *ast.TypeDescription {
	return v.TypeDescription
}

func (v *StateVariable) GetSrc() ast.SrcNode {
	return v.unit.GetSrc()
}

func (v *StateVariable) ToProto() *ir_pb.StateVariable {
	proto := &ir_pb.StateVariable{
		Id:              v.GetId(),
		ContractId:      v.GetContractId(),
		Name:            v.GetName(),
		NodeType:        v.GetNodeType(),
		Visibility:      v.GetVisibility(),
		IsConstant:      v.IsConstant(),
		StorageLocation: v.GetStorageLocation(),
		StateMutability: v.GetStateMutability(),
		Type:            v.GetType(),
		TypeDescription: v.GetTypeDescription().ToProto(),
	}

	return proto
}

func (b *Builder) processStateVariables(unit *ast.StateVariableDeclaration) *StateVariable {
	variableNode := &StateVariable{
		unit:            unit,
		Id:              unit.GetId(),
		ContractId:      unit.GetScope(),
		Name:            unit.GetName(),
		NodeType:        unit.GetType(),
		Visibility:      unit.GetVisibility(),
		Constant:        unit.IsConstant(),
		StorageLocation: unit.GetStorageLocation(),
		StateMutability: unit.GetStateMutability(),
		Type:            unit.GetTypeName().GetName(),
		TypeDescription: unit.GetTypeName().GetTypeDescription(),
	}

	// It could be that name of the type name node is not set, but the type description string is.
	if variableNode.Type == "" {
		variableNode.Type = variableNode.TypeDescription.TypeString
	}

	return variableNode
}
