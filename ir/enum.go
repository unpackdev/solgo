package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

// Enum represents an enumeration in the IR.
type Enum struct {
	unit          *ast.EnumDefinition
	Id            int64           `json:"id"`
	NodeType      ast_pb.NodeType `json:"node_type"`
	Name          string          `json:"name"`
	CanonicalName string          `json:"canonical_name"`
	Members       []*Parameter    `json:"members"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the enum.
func (e *Enum) GetAST() *ast.EnumDefinition {
	return e.unit
}

// GetNodeType returns the NodeType of the enum.
func (e *Enum) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetId returns the ID of the enum.
func (e *Enum) GetId() int64 {
	return e.Id
}

// GetName returns the name of the enum.
func (e *Enum) GetName() string {
	return e.Name
}

// GetCanonicalName returns the canonical name of the enum.
func (e *Enum) GetCanonicalName() string {
	return e.CanonicalName
}

// GetMembers returns the members of the enum.
func (e *Enum) GetMembers() []*Parameter {
	return e.Members
}

// GetSrc returns the source location of the enum.
func (e *Enum) GetSrc() ast.SrcNode {
	return e.unit.GetSrc()
}

// ToProto converts the Enum to its protobuf representation.
func (e *Enum) ToProto() *ir_pb.Enum {
	proto := &ir_pb.Enum{
		Id:            e.GetId(),
		NodeType:      e.GetNodeType(),
		Name:          e.GetName(),
		CanonicalName: e.GetCanonicalName(),
		Members:       make([]*ir_pb.Parameter, 0),
	}

	for _, member := range e.GetMembers() {
		proto.Members = append(proto.Members, member.ToProto())
	}

	return proto
}

// processEnum processes the enum unit and returns the Enum.
func (b *Builder) processEnum(unit *ast.EnumDefinition) *Enum {
	toReturn := &Enum{
		unit:          unit,
		Id:            unit.GetId(),
		NodeType:      unit.GetType(),
		Name:          unit.GetName(),
		CanonicalName: unit.GetCanonicalName(),
		Members:       make([]*Parameter, 0),
	}

	for _, parameter := range unit.GetMembers() {
		param := &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            "enum",
			TypeDescription: parameter.GetTypeDescription(),
		}

		toReturn.Members = append(toReturn.Members, param)
	}

	return toReturn
}
